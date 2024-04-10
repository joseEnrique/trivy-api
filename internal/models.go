package internal

// Or simply:
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

// Or simply:
type Results struct {
	VulnerabilityReports []VulnerabilityReport `json:"Results"`
}
