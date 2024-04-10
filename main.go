package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// VulnerabilityReport represents a simplified structure of Trivy's report.
// You might need to adjust it based on the actual output you need.
type VulnerabilityReport struct {
	Vulnerabilities []struct {
		VulnerabilityID  string `json:"VulnerabilityID"`
		PkgName          string `json:"PkgName"`
		InstalledVersion string `json:"InstalledVersion"`
		FixedVersion     string `json:"FixedVersion"`
		Severity         string `json:"Severity"`
		Title            string `json:"Title,omitempty"`
		Description      string `json:"Description,omitempty"`
	} `json:"Vulnerabilities"`
}

type Results struct {
	VulnerabilityReports []VulnerabilityReport `json:"Results"`
}

func main() {
	// Replace "alpine:3.10" with the image you want to scan
	// imageName := "quixcontainerregistry.azurecr.io/git-api:latest"
	imageName := "wordpress"

	// Construct the Trivy command
	cmd := exec.Command("trivy", "image", "--format", "json", imageName)

	// Execute the Trivy command
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing Trivy: %s\n", err)
		return
	}

	// Unmarshal the JSON output into the VulnerabilityReport struct
	var reports Results
	if err := json.Unmarshal(output, &reports); err != nil {
		fmt.Printf("Error parsing JSON output: %s\n", err)
		return
	}

	// fmt.Println(reports.VulnerabilityReports)
	// Process the _¿reports as needed??
	for _, report := range reports.VulnerabilityReports {
		for _, vuln := range report.Vulnerabilities {
			fmt.Printf("VulnerabilityID: %s, PkgName: %s, InstalledVersion: %s, FixedVersion: %s, Severity: %s\n",
				vuln.VulnerabilityID, vuln.PkgName, vuln.InstalledVersion, vuln.FixedVersion, vuln.Severity)
		}
	}
}
