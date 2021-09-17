package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
)

//go:embed tmpl/*
var f embed.FS

func main() {
	if len(os.Args) != 1 {
		log.Fatal("extension name required, got none")
	}
	extension := os.Args[1]

	cmd := exec.Command("gh", "extension", "create", extension)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	ghExtension := "gh-" + extension
	if err := os.Chdir(ghExtension); err != nil {
		log.Fatal(err)
	}

	// gh extension create foo
	// apply templates
	fmt.Println("")
}
