package git

import (
	"fmt"
	"os"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

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