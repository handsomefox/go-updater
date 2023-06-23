package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed archive.tar.gz
var zipped []byte

func main() {
	if len(zipped) == 0 {
		log.Fatal("No files were embedded, what?")
	}

	f, err := os.CreateTemp(os.TempDir(), "goupdate")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(zipped)
	if err != nil {
		log.Fatal(err)
	}

	p, err := filepath.Abs(f.Name())
	if err != nil {
		log.Fatal(err)
	}

	temp, err := os.MkdirTemp(os.TempDir(), "goupdate")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("tar", "xf", p, "--directory="+temp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	args := []string{
		"-c",
		temp + "/vendor/update-golang/update-golang.sh",
	}
	args = append(args, os.Args[1:]...)

	cmd = exec.Command("/bin/sh", args...)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
