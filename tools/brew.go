package tools

import (
	"fmt"
	"os/exec"
)

const homebrewInstallURL = "https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh"

func checkBrewInstalled() bool {
	_, err := exec.LookPath("brew")
	return err == nil
}

func installHomebrew() error {
	fmt.Println("Installing Homebrew...")
	cmd := exec.Command("/bin/bash", "-c", "$(curl -fsSL "+homebrewInstallURL+" )")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install Homebrew: %w", err)
	}
	return nil
}

func installBrewPackages() error {
	brewPackages := []string{
		"ghq", "jq", "peco", "tmux", "zsh", "zplug",
		"fzf", "awscli", "asdf", "kotlin", "gradle",
		"go", "bat", // cat command alternative
		"exa", // ls command alternative
	}

	for _, pkg := range brewPackages {
		fmt.Printf("Installing Homebrew package: %s\n", pkg)
		if err := exec.Command("brew", "install", pkg).Run(); err != nil {
			return fmt.Errorf("failed to install package %s: %w", pkg, err)
		}
	}

	return nil
}

func InstallBrew() error {
	if !checkBrewInstalled() {
		if err := installHomebrew(); err != nil {
			return err
		}
	}

	if !checkBrewInstalled() {
		return fmt.Errorf("failed to verify Homebrew installation")
	}

	if err := installBrewPackages(); err != nil {
		return err
	}

	return nil
}