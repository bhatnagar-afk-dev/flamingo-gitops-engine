package main

import (
	"fmt"
	"os"

	"github.com/bhatnagar-afk-dev/flamingo-gitops-engine/git"
	"github.com/bhatnagar-afk-dev/flamingo-gitops-engine/k8s"
	"github.com/bhatnagar-afk-dev/flamingo-gitops-engine/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <repo-url>")
		return
	}

	repoURL := os.Args[1]

	dir, err := git.CloneRepo(repoURL, "main")
	if err != nil {
		panic(err)
	}

	files, err := utils.ReadYAMLFiles(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		err := k8s.ApplyYAML(file)
		if err != nil {
			fmt.Println("Error applying:", err)
		}
	}
}
