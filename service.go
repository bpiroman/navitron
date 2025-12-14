package main

import (
	"fmt"
	"log"
	"os/exec"
)

func stopService(serviceName string) {
	fmt.Println("Stopping service:", serviceName)

	cmd := exec.Command("sudo", "systemctl", "stop", serviceName)

	// Run the command and capture output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to stop service %s: %v\nOutput: %s", serviceName, err, output)
	}

	fmt.Println("Service stopped.")
}

func startService(serviceName string) {
	fmt.Println("Starting service:", serviceName)

	cmd := exec.Command("sudo", "systemctl", "start", serviceName)

	// Run the command and capture output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to stop service %s: %v\nOutput: %s", serviceName, err, output)
	}

	fmt.Println("Service started.")
}
