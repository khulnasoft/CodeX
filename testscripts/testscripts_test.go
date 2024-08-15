package testrunner

import (
	"os"
	"strconv"
	"testing"

	"github.com/khulnasoft/codex/testscripts/testrunner"
)

// When true, tests that `codex run run_test` succeeds on every project (i.e. having codex.json)
// found in examples/.. and testscripts/..
const runProjectTests = "CODEX_RUN_PROJECT_TESTS"

func TestScripts(t *testing.T) {
	// To run a specific test, say, testscripts/foo/bar.test.text, then run
	// go test ./testscripts -run TestScripts/bar
	testrunner.RunTestscripts(t, ".")
}

func TestMain(m *testing.M) {
	os.Exit(testrunner.Main(m))
}

// TestExamples runs testscripts on the codex-projects in the examples folder.
func TestExamples(t *testing.T) {
	isOn, err := strconv.ParseBool(os.Getenv(runProjectTests))
	if err != nil || !isOn {
		t.Skipf("Skipping TestExamples. To enable, set %s=1.", runProjectTests)
	}

	// To run a specific test, say, examples/foo/bar, then run
	// go test ./testscripts -run TestExamples/foo_bar_run_test
	testrunner.RunCodexTestscripts(t, "../examples")
}

// TestScriptsWithProjects runs testscripts on the codex-projects in the testscripts folder.
func TestScriptsWithProjects(t *testing.T) {
	isOn, err := strconv.ParseBool(os.Getenv(runProjectTests))
	if err != nil || !isOn {
		t.Skipf("Skipping TestScriptsWithProjects. To enable, set %s=1.", runProjectTests)
	}

	testrunner.RunCodexTestscripts(t, ".")
}
