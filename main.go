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

	// register ssh key
	if err = registerSSH(conf.SSHPath); err != nil {
		log.Logger().Fatal("Failed to initialize SSH")
	}

	// clone the repo if we dont already have a copy
	// do a pull if we do have the folder
	if _, err = os.Stat(conf.DataFolder); os.IsNotExist(err) {
		log.Logger().Info("Data repository not found")
		if err = cloneRepo(conf.Repository, conf.Branch, conf.DataFolder); err != nil {
			log.Logger().Fatal(err.Error())
		}
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

// loads the ssh agent and loads the key for git
func registerSSH(sshKeyPath string) error {
	cmd := exec.Command("sh", "-c", "eval $(ssh-agent -s)", "&&", "ssh-add", sshKeyPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Logger().Info(string(out))
		log.Logger().Error(err.Error())
		log.Logger().Fatal("Failed to start SSH agent")
		return err
	}

	return nil
}

// clones the repository and checks out the branch into the target directory
func cloneRepo(repository, branch, target string) error {
	log.Logger().Info("Cloning repository...")

	cmd := exec.Command("git", "clone", repository, "--branch", branch, "--depth=1", target)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Logger().Info(string(out))
		log.Logger().Error(err.Error())
		log.Logger().Fatal("Failed to clone repository")
		return err
	}

	log.Logger().Info("Cloning repository [OK]")
	return nil
}

// pulls from the upstream repository of the specified directory
func pullRepo(target string) {
	log.Logger().Info("Pulling repository...")

	// create git pull command and set working directory
	cmd := exec.Command("git", "pull", "--depth=1")
	cmd.Dir = target
	if err := cmd.Run(); err != nil {
		log.Logger().Error(err.Error())
		log.Logger().Fatal("Failed to pull repository")
		return
	}

	log.Logger().Info("Pulling repository [OK]")
}
