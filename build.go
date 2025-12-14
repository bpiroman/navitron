package main

import (
	"fmt"
	"log"
	"os/exec"
)

// go build -o ./tmp/webapp ./cmd/web
// go build -o ./tmp/latest ./cmd/latest

func build(buildBin string) {
	fmt.Println("building binary:", buildBin)

	cmd := exec.Command("sh", "-c", buildBin)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to build the binary %s: %v\nOutput: %s", buildBin, err, output)
	}

	fmt.Println("binary built.")
}
