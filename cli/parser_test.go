package cli

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	examples := []struct {
		args []string
		cmd  string
		pkg  string
	}{
		{
			args: []string{},
			cmd:  "HelpCommand",
			pkg:  "subcommand",
		},
		// owner command
		{
			args: []string{"owner"},
			cmd:  "ReposCommand",
			pkg:  "subcommand",
		},
	}

	for _, example := range examples {
		result := ParseArgs(example.args)

		structType := reflect.TypeOf(result)
		if structType.Name() != example.cmd {
			t.Errorf("Struct expected %s, but %s", example.cmd, structType.Name())
		}
		if structType.PkgPath() != "github.com/hirakiuc/alfred-github-workflow/"+example.pkg {
			t.Errorf("The package is not expected:%s", structType.PkgPath())
		}
	}
}

func TestParseRepoCommand(t *testing.T) {
	examples := []struct {
		args []string
		cmd  string
		pkg  string
	}{
		// owner/repo commands
		{
			args: []string{"owner/repo"},
			cmd:  "HelpCommand",
			pkg:  "subcommand/repo",
		},
		{
			args: []string{"owner/repo branches"},
			cmd:  "BranchesCommand",
			pkg:  "subcommand/repo",
		},
		{
			args: []string{"owner/repo issues"},
			cmd:  "IssueCommand",
			pkg:  "subcommand/repo",
		},
		{
			args: []string{"owner/repo milestones"},
			cmd:  "MilestonesCommand",
			pkg:  "subcommand/repo",
		},
		{
			args: []string{"owner/repo pulls"},
			cmd:  "PullsCommand",
			pkg:  "subcommand/repo",
		},
	}

	for _, example := range examples {
		result := ParseArgs(example.args)

		structType := reflect.TypeOf(result)
		if structType.Name() != example.cmd {
			t.Errorf("Struct expected %s, but %s", example.cmd, structType.Name())
		}
		if structType.PkgPath() != "github.com/hirakiuc/alfred-github-workflow/"+example.pkg {
			t.Errorf("The package is not expected:%s", structType.PkgPath())
		}
	}
}

func TestParseMyCommand(t *testing.T) {
	examples := []struct {
		args []string
		cmd  string
		pkg  string
	}{
		// my commands
		{
			args: []string{"my"},
			cmd:  "HelpCommand",
			pkg:  "subcommand/my",
		},
		{
			args: []string{"my pulls"},
			cmd:  "HelpCommand",
			pkg:  "subcommand/my/pulls",
		},
	}

	for _, example := range examples {
		result := ParseArgs(example.args)

		structType := reflect.TypeOf(result)
		if structType.Name() != example.cmd {
			t.Errorf("Struct expected %s, but %s", example.cmd, structType.Name())
		}
		if structType.PkgPath() != "github.com/hirakiuc/alfred-github-workflow/"+example.pkg {
			t.Errorf("The package is not expected:%s", structType.PkgPath())
		}
	}
}

func TestParseConfigCommand(t *testing.T) {
	examples := []struct {
		args []string
		cmd  string
		pkg  string
	}{
		// config commands
		{
			args: []string{">"},
			cmd:  "HelpCommand",
			pkg:  "subcommand/config",
		},
		{
			args: []string{"> token"},
			cmd:  "TokenCommand",
			pkg:  "subcommand/config",
		},
		{
			args: []string{"> clear-cache"},
			cmd:  "ClearCacheCommand",
			pkg:  "subcommand/config",
		},
	}

	for _, example := range examples {
		result := ParseArgs(example.args)

		structType := reflect.TypeOf(result)
		if structType.Name() != example.cmd {
			t.Errorf("Struct expected %s, but %s", example.cmd, structType.Name())
		}
		if structType.PkgPath() != "github.com/hirakiuc/alfred-github-workflow/"+example.pkg {
			t.Errorf("The package is not expected:%s", structType.PkgPath())
		}
	}
}
