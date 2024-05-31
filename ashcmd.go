package ashcmd

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// ScanResult represents the result of the scan performed by ashCmd.exe.
type ScanResult struct {
	IsDetected    bool   `json:"is_detected"`
	Malware       string `json:"malware"`
	EngineVersion string `json:"engine_version"`
}

type Avast struct {
	AshCmdPath string
}

func NewAvast(ashPath string) *Avast {
	// TODO: Check if ashPath is valid and contain a ashCmd.exe file
	// if no ashPath provided then just ignore

	return &Avast{
		AshCmdPath: "C:\\Program Files\\Avast Software\\Avast\\ashCmd.exe",
	}
}

func (avast *Avast) parseOutput(input string) *ScanResult {
	var result ScanResult

	var lines []string
	for _, line := range strings.Split(input, "\n") {
		if strings.Trim(line, "\n") == "" {
			continue
		}
		lines = append(lines, line)
	}

	// Compile the regex
	re := regexp.MustCompile(`^([^#][^\s]+)\s+(.*)`)

	// Find the matches
	matches := re.FindStringSubmatch(lines[0])

	// Check if the second group exists
	if len(matches) > 2 && matches[2] != "" {
		result.IsDetected = true
		result.Malware = matches[2]
	} else {
		result.IsDetected = false
	}

	re = regexp.MustCompile(`^#[^:]+\s*:\s*(.*)`)
	matches = re.FindStringSubmatch(lines[len(lines)-2])

	if len(matches) > 1 && matches[1] != "" {
		result.EngineVersion = matches[1]
	}

	return &result
}

// ExecuteScan runs the ashCmd.exe scan command and returns the parsed result.
func (avast *Avast) ScanFile() (*ScanResult, error) {
	// Execute the ashCmd.exe scan command
	cmd := exec.Command(fmt.Sprintf("& %s", avast.AshCmdPath), "/a /c /d /t=a /_")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, errors.New("failed to execute scan command: " + err.Error())
	}

	result := avast.parseOutput(out.String())

	return result, nil
}
