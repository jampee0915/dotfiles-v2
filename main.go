package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jampee0915/dotfiles-v2/tools"
)

func main() {
	brewInstall := flag.Bool("brew", false, "Install brew")
	allInstall := flag.Bool("all", false, "Install all tools")
	flag.Parse()

	if *brewInstall || *allInstall {
		if err := tools.InstallBrew(); err != nil {
            fmt.Println("Brew installation failed: ", err)    
            os.Exit(1)
        }
	}
}
