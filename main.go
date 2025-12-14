package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Project string  `toml:"project"`
	Paths   Paths   `toml:"paths"`
	Service string  `toml:"service"`
	Build   Build   `toml:"build"`
	Include Include `toml:"include"`
}

type Paths struct {
	App string `toml:"app"`
}

type Build struct {
	Build []string `toml:"build"`
}

type Include struct {
	Bin     []string `toml:"bin"`
	Folders []string `toml:"folders"`
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current working director:", err)
		return
	}
	fmt.Println("current working directory: ", cwd)

	var config Config

	meta, err := toml.DecodeFile("navitron.toml", &config)
	if err != nil {
		log.Fatalf("failed to read config.toml: %v", err)
	}

	// Metadata checks
	fmt.Println("Undecoded keys:", meta.Undecoded())
	fmt.Println("Is project set?", meta.IsDefined("project"))
	fmt.Println("Is paths.app set?", meta.IsDefined("paths", "app"))
	fmt.Println("Is build.build set?", meta.IsDefined("build", "build"))

	// Print config values
	fmt.Println("Project name:", config.Project)
	fmt.Println("App path:", config.Paths.App)
	fmt.Println("Service file:", config.Service)

	fmt.Println("Build targets:")
	for _, b := range config.Build.Build {
		fmt.Println(" -", b)
	}

	fmt.Println("Include bins:")
	for _, b := range config.Include.Bin {
		fmt.Println(" -", b)
	}

	fmt.Println("Include folders:")
	for _, f := range config.Include.Folders {
		fmt.Println(" -", f)
	}

	// stop systemd service file
	stopService(config.Service)

	// build go binaries
	buildBins := config.Build.Build
	for _, buildBin := range buildBins {
		fmt.Println("Running build for:", buildBin)
		build(buildBin)
	}

	// clean app_dir
	cleanAppDir(config.Paths.App)

	// ensure app directory
	ensureAppDir(config.Paths.App)

	// copy files
	bins := config.Include.Bin
	for _, bin := range bins {
		dst := config.Paths.App
		copyFile(bin, dst)
	}

	// copy folders
	folders := config.Include.Folders
	for _, folder := range folders {
		dst := config.Paths.App
		copyFolder(folder, dst)
	}

	// stop systemd service file
	startService(config.Service)
}
