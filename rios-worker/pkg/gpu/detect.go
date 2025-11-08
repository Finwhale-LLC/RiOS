package gpu

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// GPUInfo represents GPU information
type GPUInfo struct {
	Type  string
	Count int
	VRam  int // in GB
}

// DetectNVIDIA detects NVIDIA GPU information using nvidia-smi
func DetectNVIDIA() (*GPUInfo, error) {
	// Check if nvidia-smi is available
	cmd := exec.Command("nvidia-smi", "--query-gpu=name,memory.total", "--format=csv,noheader")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("nvidia-smi not found or failed to execute. Please ensure NVIDIA drivers are installed: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("no NVIDIA GPUs detected")
	}

	// Parse first GPU info
	parts := strings.Split(lines[0], ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("unexpected nvidia-smi output format")
	}

	gpuName := strings.TrimSpace(parts[0])
	memoryStr := strings.TrimSpace(parts[1])

	// Extract memory in GB
	re := regexp.MustCompile(`(\d+)\s*MiB`)
	matches := re.FindStringSubmatch(memoryStr)
	
	var vramGB int
	if len(matches) >= 2 {
		vramMB, _ := strconv.Atoi(matches[1])
		vramGB = vramMB / 1024
	} else {
		vramGB = 0
	}

	return &GPUInfo{
		Type:  gpuName,
		Count: len(lines),
		VRam:  vramGB,
	}, nil
}

// Detect attempts to detect GPU information
func Detect() (*GPUInfo, error) {
	// Currently only supports NVIDIA
	info, err := DetectNVIDIA()
	if err != nil {
		return nil, fmt.Errorf("failed to detect GPU: %w", err)
	}

	return info, nil
}

