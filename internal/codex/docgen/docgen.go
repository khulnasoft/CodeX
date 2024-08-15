package docgen

import (
	_ "embed"
	"os"
	"text/template"

	"github.com/khulnasoft/codex/internal/codex"
	"github.com/khulnasoft/codex/internal/fileutil"
)

//go:embed readme.tmpl
var defaultReadmeTemplate string

const (
	defaultName         = "README.md"
	defaultTemplateName = "readme.tmpl"
)

func GenerateReadme(
	codex *codex.Codex,
	outputPath, templatePath string,
) error {
	readmeTemplate := defaultReadmeTemplate
	if templatePath != "" {
		readmeTemplateBytes, err := os.ReadFile(templatePath)
		if err != nil {
			return err
		}
		readmeTemplate = string(readmeTemplateBytes)
	} else if fileutil.Exists(defaultTemplateName) {
		readmeTemplateBytes, err := os.ReadFile(defaultTemplateName)
		if err != nil {
			return err
		}
		readmeTemplate = string(readmeTemplateBytes)
	}

	tmpl, err := template.New("readme").Parse(readmeTemplate)
	if err != nil {
		return err
	}

	if outputPath == "" {
		outputPath = defaultName
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	return tmpl.Execute(f, map[string]any{
		"Name":        codex.Config().Root.Name,
		"Description": codex.Config().Root.Description,
		"Scripts": codex.Config().Scripts().
			WithRelativePaths(codex.ProjectDir()),
		"EnvVars":  codex.Config().Env(),
		"InitHook": codex.Config().InitHook(),
		"Packages": codex.TopLevelPackages(),
		// TODO add includes
	})
}

func SaveDefaultReadmeTemplate(outputPath string) error {
	if outputPath == "" {
		outputPath = defaultTemplateName
	}
	return os.WriteFile(outputPath, []byte(defaultReadmeTemplate), 0o644)
}
