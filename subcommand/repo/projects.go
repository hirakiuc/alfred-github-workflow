package repo

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/icon"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// ProjectsCommand describe a subcommand to fetch projects
type ProjectsCommand struct {
	Owner string
	Repo  string

	Query string
	Limit int
}

// NewProjectsCommand return a ProjectsCommand instance
func NewProjectsCommand(owner string, repo string, args []string) ProjectsCommand {
	return ProjectsCommand{
		Owner: owner,
		Repo:  repo,
		Query: strings.Join(args, " "),
		Limit: 100,
	}
}

func (cmd ProjectsCommand) fetchProjects(ctx context.Context, wf *aw.Workflow) ([]model.Project, error) {
	store := cache.NewProjectsCache(wf)

	projects, err := store.GetCache(cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Project{}, err
	}

	if len(projects) != 0 {
		return projects, nil
	}

	client, err := api.NewClient(ctx, wf)
	if err != nil {
		return []model.Project{}, err
	}

	projects, err = client.FetchProjects(ctx, cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Project{}, err
	}

	return store.Store(cmd.Owner, cmd.Repo, projects)
}

// Run start this subcommand.
func (cmd ProjectsCommand) Run(ctx context.Context, wf *aw.Workflow) {
	projects, err := cmd.fetchProjects(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	icon, _ := icon.GetIcon(icon.TypeProject)

	// Add items
	for _, project := range projects {
		item := wf.NewItem(project.Name)

		if icon != nil {
			item.Icon(icon)
		}
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No projects found.", "")
}
