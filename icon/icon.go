package icon

import (
	"path"
	"path/filepath"
	"runtime"

	aw "github.com/deanishe/awgo"
)

// IconType...
type Type string

const (
	TypeBranch    Type = "git-branch.png"
	TypeDefault   Type = "octoface.png"
	TypeIssue     Type = "issue-opened.png"
	TypeMilestone Type = "milestone.png"
	TypeProject   Type = "project.png"
	TypePull      Type = "git-pull-request.png"
	TypeRepo      Type = "repo.png"
)

func (t Type) ToString() string {
	return string(t)
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func GetIcon(t Type) (*aw.Icon, error) {
	iconPath := filepath.Join(getCurrentPath(), "../../assets/icons", t.ToString())

	absPath, err := filepath.Abs(iconPath)
	if err != nil {
		return nil, err
	}

	return &aw.Icon{
		Value: absPath,
		Type:  aw.IconTypeImage,
	}, nil
}
