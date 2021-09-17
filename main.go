package main

import (
	"embed"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

//go:embed tmpl/*
var f embed.FS

type values struct {
	User      string
	Extension string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("extension name required, got none")
	}
	extension := os.Args[1]

	// gh extension create foo
	cmd := exec.Command("gh", "extension", "create", extension)
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to run gh command: %v\n", err)
	}

	ghExtension := "gh-" + extension
	if err := os.Chdir(ghExtension); err != nil {
		log.Fatalf("failed to Chdir: %v\n", err)
	}

	// apply templates
	tmpl := template.Must(template.ParseFS(f, "tmpl/*.tmpl"))
	vals := values{"n4to4", ghExtension}
	for _, t := range tmpl.Templates() {
		var filename string
		if t.Name() == "gh-command.tmpl" {
			filename = ghExtension
		} else {
			filename = strings.TrimRight(t.Name(), ".tmpl")
		}

		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("failed to create file: %v\n", err)
		}
		if strings.HasSuffix(filename, ".sh") {
			f.Chmod(0755)
		}
		tmpl.ExecuteTemplate(f, t.Name(), vals)
		f.Close()
	}
}
