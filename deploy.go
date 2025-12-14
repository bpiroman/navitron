package main

import (
	"fmt"
	"log"
	"os/exec"
)

// Deletes all contents of Paths.app
func cleanAppDir(appDir string) {
	fmt.Println("Cleaning app directory:", appDir)
	cmdStr := fmt.Sprintf("sudo rm -rf %s", appDir)
	cmd := exec.Command("sh", "-c", cmdStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to clean app directory %s: %v\nOutput: %s", appDir, err, output)
	}

	fmt.Println("App directory cleaned successfully.")
}

func ensureAppDir(appDir string) {
	cmdStr := fmt.Sprintf("sudo mkdir -p %s", appDir)
	cmd := exec.Command("sh", "-c", cmdStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to create app directory %s: %v\nOutput: %s", appDir, err, output)
	}

	fmt.Println("App directory ready:", appDir)
}

func copyFile(src, dst string) {
	cmdStr := fmt.Sprintf("sudo cp %s %s", src, dst)
	cmd := exec.Command("sh", "-c", cmdStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to copy file %s to %s: %v\nOutput: %s", src, dst, err, output)
	}

	fmt.Println("Copied file:", src, "to", dst)
}

func copyFolder(src, dst string) {
	cmdStr := fmt.Sprintf("sudo cp -r %s %s", src, dst)
	cmd := exec.Command("sh", "-c", cmdStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to copy folder %s to %s: %v\nOutput: %s", src, dst, err, output)
	}

	fmt.Println("Copied folder:", src, "to", dst)
}
