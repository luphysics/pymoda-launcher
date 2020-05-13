package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	wd := correctWorkingDirectory()
	_ = os.MkdirAll(wd, os.ModeDir)

	os.Chdir(wd)
	println(wd)

	files := listFiles()
	version := getVersionToLaunch(files)

	launchPyMODA(version)
}

func listFiles() []string {
	files, err := ioutil.ReadDir("./")

	if err != nil {
		log.Fatal(err)
	}

	out := []string{}
	for i := 0; i < len(files); i++ {
		out = append(out, files[i].Name())
	}

	return out
}

func getVersionToLaunch(files []string) string {
	for _, f := range files {
		if strings.Contains(f, "latest-") {
			return f[len("latest-"):]
		}
	}
	return "PyMODA"
}

func launchPyMODA(directory string) {
	args := strings.Join(os.Args[1:], " ")
	args += "--launcher"

	if isLinux() {
		cmd := exec.Command(directory+"/PyMODA", args)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			println("Failed to launch PyMODA.")
		}
	} else {
		cmd := exec.Command("open "+directory+"/PyMODA.app", "--args", args)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			println("Failed to launch PyMODA.")
		}
	}
}

func home() string {
	home, _ := os.UserHomeDir()

	return home
}

func correctWorkingDirectory() string {
	return filepath.Join(home(), ".pymoda")
}

func isLinux() bool {
	return runtime.GOOS == "linux"
}
