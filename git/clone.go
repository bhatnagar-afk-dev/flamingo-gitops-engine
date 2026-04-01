// Package git provides utilities to interact with Git repositories for the GitOps Engine.
package git

import (
	"fmt"
	"os"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// CloneRepo clones a remote target git repository to a local temporary directory.
// It clones a single specific branch securely without fetching the complete history (depth=1).
// Returns the absolute path to the temporary directory, and an error if the cloning fails.
// The caller is potentially responsible for cleaning up the directory when no longer needed.
func CloneRepo(repoURL, branch string) (string, error) {
	dir, err := os.MkdirTemp("", "repo-*")
	if err != nil {
		return "", err
	}

	_, err = gogit.PlainClone(dir, false, &gogit.CloneOptions{
		URL:           repoURL,
		ReferenceName: plumbing.ReferenceName("refs/heads/" + branch),
		SingleBranch:  true,
		Depth:         1,
	})

	if err != nil {
		return "", err
	}

	fmt.Println("Cloned repo to:", dir)
	return dir, nil
}