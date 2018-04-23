package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/soupstore/git-pull/config"
	"github.com/soupstore/git-pull/log"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	// clone the repo if we dont already have a copy
	// do a pull if we do have the folder
	if _, err = os.Stat(conf.DataFolder); os.IsNotExist(err) {
		log.Logger().Info("Data repository not found")
		cloneRepo(conf.Repository, conf.Branch, conf.DataFolder)
	} else {
		log.Logger().Info("Data repository found")
		pullRepo(conf.DataFolder)
	}

	// pull the repository every minute
	ticker := time.NewTicker(time.Minute)
	for _ = range ticker.C {
		pullRepo(conf.DataFolder)
	}
}

// clones the repository and checks out the branch into the target directory
func cloneRepo(repository, branch, target string) {
	log.Logger().Info("Cloning repository...")

	cmd := exec.Command("git", "clone", repository, "--branch", branch, target)
	if err := cmd.Run(); err != nil {
		log.Logger().Fatal("Failed to clone repository")
		return
	}

	log.Logger().Info("Cloning repository [OK]")
}

// pulls from the upstream repository of the specified directory
func pullRepo(target string) {
	log.Logger().Info("Pulling repository...")

	// create git pull command and set working directory
	cmd := exec.Command("git", "pull")
	cmd.Dir = target
	if err := cmd.Run(); err != nil {
		log.Logger().Fatal("Failed to pull repository")
		return
	}

	log.Logger().Info("Pulling repository [OK]")
}
