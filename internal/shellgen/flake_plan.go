package shellgen

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime/trace"
	"strings"

	"github.com/khulnasoft/codex/internal/devpkg"
	"github.com/khulnasoft/codex/internal/nix"
)

// flakePlan contains the data to populate the top level flake.nix file
// that builds the codex environment
type flakePlan struct {
	NixpkgsInfo *NixpkgsInfo
	Packages    []*devpkg.Package
	FlakeInputs []flakeInput
	System      string
}

func newFlakePlan(ctx context.Context, codex codexer) (*flakePlan, error) {
	ctx, task := trace.NewTask(ctx, "codexFlakePlan")
	defer task.End()

	for _, pluginConfig := range codex.Config().IncludedPluginConfigs() {
		if err := codex.PluginManager().CreateFilesForConfig(pluginConfig); err != nil {
			return nil, err
		}
	}

	packages := codex.InstallablePackages()

	// Fill the NarInfo Cache concurrently as a perf-optimization, prior to invoking
	// IsInBinaryCache in flakeInputs() and in the flake.nix template.
	if err := devpkg.FillNarInfoCache(ctx, packages...); err != nil {
		return nil, err
	}

	flakeInputs := flakeInputs(ctx, packages)
	nixpkgsInfo := getNixpkgsInfo(codex.Config().NixPkgsCommitHash())

	// This is an optimization. Try to reuse the nixpkgs info from the flake
	// inputs to avoid introducing a new one.
	for _, input := range flakeInputs {
		if input.IsNixpkgs() {
			nixpkgsInfo = getNixpkgsInfo(input.HashFromNixPkgsURL())
			break
		}
	}

	return &flakePlan{
		FlakeInputs: flakeInputs,
		NixpkgsInfo: nixpkgsInfo,
		Packages:    packages,
		System:      nix.System(),
	}, nil
}

func (f *flakePlan) needsGlibcPatch() bool {
	for _, in := range f.FlakeInputs {
		if in.URL == glibcPatchFlakeRef {
			return true
		}
	}
	return false
}

type glibcPatchFlake struct {
	// NixpkgsGlibcFlakeRef is a flake reference to the nixpkgs flake
	// containing the new glibc package.
	NixpkgsGlibcFlakeRef string

	// Inputs is the attribute set of flake inputs. The key is the input
	// name and the value is a flake reference.
	Inputs map[string]string

	// Outputs is the attribute set of flake outputs. It follows the
	// standard flake output schema of system.name = derivation. The
	// derivation can be any valid Nix expression.
	Outputs struct {
		Packages map[string]map[string]string
	}
}

func newGlibcPatchFlake(nixpkgsGlibcRev string, packages []*devpkg.Package) (glibcPatchFlake, error) {
	flake := glibcPatchFlake{NixpkgsGlibcFlakeRef: "flake:nixpkgs/" + nixpkgsGlibcRev}
	for _, pkg := range packages {
		if !pkg.PatchGlibc() {
			continue
		}

		err := flake.addPackageOutput(pkg)
		if err != nil {
			return glibcPatchFlake{}, err
		}
	}
	return flake, nil
}

func (g *glibcPatchFlake) addPackageOutput(pkg *devpkg.Package) error {
	if g.Inputs == nil {
		g.Inputs = make(map[string]string)
	}
	inputName := pkg.FlakeInputName()
	g.Inputs[inputName] = pkg.URLForFlakeInput()

	attrPath, err := pkg.FullPackageAttributePath()
	if err != nil {
		return err
	}
	// Remove the legacyPackages.<system> prefix.
	outputName := strings.SplitN(attrPath, ".", 3)[2]

	if g.Outputs.Packages == nil {
		g.Outputs.Packages = map[string]map[string]string{nix.System(): {}}
	}
	if cached, err := pkg.IsInBinaryCache(); err == nil && cached {
		if expr, err := g.fetchClosureExpr(pkg); err == nil {
			g.Outputs.Packages[nix.System()][outputName] = expr
			return nil
		}
	}
	g.Outputs.Packages[nix.System()][outputName] = strings.Join([]string{"pkgs", inputName, nix.System(), outputName}, ".")
	return nil
}

// TODO: this only handles the first store path, but we should handle all of them
func (g *glibcPatchFlake) fetchClosureExpr(pkg *devpkg.Package) (string, error) {
	storePaths, err := pkg.InputAddressedPaths()
	if err != nil {
		return "", err
	}
	if len(storePaths) == 0 {
		return "", fmt.Errorf("no store path for package %s", pkg.Raw)
	}
	return fmt.Sprintf(`builtins.fetchClosure {
  fromStore = "%s";
  fromPath = "%s";
  inputAddressed = true;
}`, "devpkg.BinaryCache", storePaths[0]), nil
}

func (g *glibcPatchFlake) writeTo(dir string) error {
	err := writeFromTemplate(dir, g, "glibc-patch.nix", "flake.nix")
	if err != nil {
		return err
	}
	return writeGlibcPatchScript(filepath.Join(dir, "glibc-patch.bash"))
}
