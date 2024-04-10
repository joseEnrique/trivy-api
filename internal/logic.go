package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

func trivy() (Results, error) {
	// Replace "alpine:3.10" with the image you want to scan
	// imageName := "quixcontainerregistry.azurecr.io/git-api:latest"
	imageName := "alpine"
	var reports Results

	// Construct the Trivy command
	cmd := exec.Command("trivy", "image", "--format", "json", imageName)

	// Execute the Trivy command
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing Trivy: %s\n", err)
		return reports, errors.New("error executing Trivy")
	}

	// Unmarshal the JSON output into the VulnerabilityReport struct
	if err := json.Unmarshal(output, &reports); err != nil {
		fmt.Printf("Error parsing JSON output: %s\n", err)
		return reports, errors.New("error parsing json output")
	}

	return reports, nil

}
