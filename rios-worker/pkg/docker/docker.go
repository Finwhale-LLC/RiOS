package docker

import (
	"fmt"
	"os/exec"
)

// CheckDockerInstalled checks if Docker is installed and available
func CheckDockerInstalled() error {
	cmd := exec.Command("docker", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Docker is not installed or not in PATH. Please install Docker: https://docs.docker.com/get-docker/")
	}
	return nil
}

// CheckDockerRunning checks if Docker daemon is running
func CheckDockerRunning() error {
	cmd := exec.Command("docker", "ps")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Docker daemon is not running. Please start Docker")
	}
	return nil
}

// CheckNVIDIADockerSupport checks if NVIDIA Docker runtime is available
func CheckNVIDIADockerSupport() error {
	cmd := exec.Command("docker", "run", "--rm", "--gpus", "all", "nvidia/cuda:12.0.0-base-ubuntu22.04", "nvidia-smi")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("NVIDIA Docker runtime not available. Please install nvidia-docker2: https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/install-guide.html")
	}
	return nil
}

