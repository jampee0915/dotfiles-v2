package tools

import (
	"fmt"
	"os/exec"
)

func InstallBrew() error {
	// homebrewがインストールされているか確認
	_, err := exec.LookPath("brew")
	if err != nil {
		fmt.Println("Installing homebrew...")

		// インストールコマンド
		exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)").Run()
	}

	// 再度homebrewがインストールされているか確認
	_, err = exec.LookPath("brew")
	if err != nil {
		return err
	}

	brewPackages := []string{
		"ghq",
		"jq",
		"peco",
	}

	for _, pkg := range brewPackages {
		fmt.Println("Installing homebrew package: ", pkg)
		if err := exec.Command("brew", "install", pkg).Run(); err != nil {
			return err
		}
	}

	return nil
}