package storybranch

import (
	"os/exec"
	"strings"
)

type Repository struct{}

// GetBranchName gets the current branch name of the repository.
func (repo Repository) GetBranchName() string {
	output, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.TrimSpace(string(output))
}

// GitRepository comment
type GitRepository interface {
	GetBranchName() string
}

func NewRepository() *Repository {
	return &Repository{}
}
