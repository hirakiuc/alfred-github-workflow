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
		// Invalid command
		{
			args: []string{"/"},
			cmd:  "HelpCommand",
			pkg:  "subcommand",
		},
		// owner command
		{
			args: []string{"owner"},
			cmd:  "ReposCommand",
			pkg:  "subcommand",
		},
		// owner/repo command
		{
			args: []string{"owner/repo"},
			cmd:  "HelpCommand",
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
		{
			args: []string{"owner/repo new"},
			cmd:  "NewCommand",
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

func TestParseMyIssuesCommand(t *testing.T) {
	examples := []struct {
		args []string
		cmd  string
		pkg  string
	}{
		{
			args: []string{"my issues"},
			cmd:  "HelpCommand",
			pkg:  "subcommand/my/issues",
		},
		// my issues assigned
		{
			args: []string{"my issues assigned"},
			cmd:  "AssignedCommand",
			pkg:  "subcommand/my/issues",
		},
		// my issues created
		{
			args: []string{"my issues created"},
			cmd:  "CreatedCommand",
			pkg:  "subcommand/my/issues",
		},
		// my issues mentioned
		{
			args: []string{"my issues mentioned"},
			cmd:  "MentionedCommand",
			pkg:  "subcommand/my/issues",
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

func TestParseMyPullsCommand(t *testing.T) {
	examples := []struct {
		args []string
		cmd  string
		pkg  string
	}{
		{
			args: []string{"my pulls"},
			cmd:  "HelpCommand",
			pkg:  "subcommand/my/pulls",
		},
		// my pulls assigned
		{
			args: []string{"my pulls assigned"},
			cmd:  "AssignedCommand",
			pkg:  "subcommand/my/pulls",
		},
		// my pulls created
		{
			args: []string{"my pulls created"},
			cmd:  "CreatedCommand",
			pkg:  "subcommand/my/pulls",
		},
		// my pulls mentioned
		{
			args: []string{"my pulls mentioned"},
			cmd:  "MentionedCommand",
			pkg:  "subcommand/my/pulls",
		},
		// my pulls review-requests
		{
			args: []string{"my pulls review-requests"},
			cmd:  "ReviewRequestsCommand",
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
		opts string
	}{
		// config commands
		{
			args: []string{">"},
			cmd:  "HelpCommand",
			pkg:  "subcommand/config",
			opts: "",
		},
		{
			args: []string{"> token"},
			cmd:  "TokenCommand",
			pkg:  "subcommand/config",
			opts: "",
		},
		{
			args: []string{"> token api-token-string"},
			cmd:  "TokenCommand",
			pkg:  "subcommand/config",
			opts: "api-token-string",
		},
		{
			args: []string{"> clear-cache"},
			cmd:  "ClearCacheCommand",
			pkg:  "subcommand/config",
			opts: "",
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

func TestParseConfigTokenCommand(t *testing.T) {
	examples := []struct {
		args []string
		cmd  string
		pkg  string
		opts string
	}{
		// config commands
		{
			args: []string{"> token"},
			cmd:  "TokenCommand",
			pkg:  "subcommand/config",
			opts: "",
		},
		{
			args: []string{"> token api-token-string"},
			cmd:  "TokenCommand",
			pkg:  "subcommand/config",
			opts: "api-token-string",
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
