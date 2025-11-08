package worker

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rios/worker/pkg/api"
)

// Executor handles job execution
type Executor struct {
	WorkDir string
}

// NewExecutor creates a new executor
func NewExecutor(workDir string) *Executor {
	return &Executor{
		WorkDir: workDir,
	}
}

// Execute executes a job
func (e *Executor) Execute(job *api.Job) (outputURL string, err error) {
	// Create temporary work directory for this job
	jobWorkDir := filepath.Join(e.WorkDir, job.JobID)
	if err := os.MkdirAll(jobWorkDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create job work directory: %w", err)
	}
	defer func() {
		// Clean up on error
		if err != nil {
			os.RemoveAll(jobWorkDir)
		}
	}()

	inputDir := filepath.Join(jobWorkDir, "input")
	outputDir := filepath.Join(jobWorkDir, "output")
	
	if err := os.MkdirAll(inputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create input directory: %w", err)
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	// Download input files
	fmt.Println("   📥 Downloading input files...")
	if job.Payload.InputS3URL != "" {
		inputFile := filepath.Join(inputDir, "workflow.json")
		if err := e.downloadFile(job.Payload.InputS3URL, inputFile); err != nil {
			return "", fmt.Errorf("failed to download input: %w", err)
		}
	}

	// Execute Docker command
	fmt.Println("   🐳 Running Docker container...")
	if err := e.runDocker(job, inputDir, outputDir); err != nil {
		return "", fmt.Errorf("docker execution failed: %w", err)
	}

	// Upload output files
	fmt.Println("   📤 Uploading output files...")
	outputURL, err = e.uploadOutput(outputDir, job.Payload.OutputS3Path)
	if err != nil {
		return "", fmt.Errorf("failed to upload output: %w", err)
	}

	// Clean up work directory
	os.RemoveAll(jobWorkDir)

	return outputURL, nil
}

// downloadFile downloads a file from URL to local path
func (e *Executor) downloadFile(url, filepath string) error {
	// For demo purposes, we'll just create a dummy file
	// In production, this should download from S3
	
	// If URL starts with http/https, actually download
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		out, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		return err
	}

	// For S3 URLs, create a placeholder
	// In production, use AWS SDK to download from S3
	return os.WriteFile(filepath, []byte("{}"), 0644)
}

// runDocker runs the Docker container for the job
func (e *Executor) runDocker(job *api.Job, inputDir, outputDir string) error {
	dockerImage := job.Payload.DockerImage
	
	// Build docker command
	args := []string{
		"run",
		"--rm",
		"--gpus", "all",
		"-v", fmt.Sprintf("%s:/workspace/input", inputDir),
		"-v", fmt.Sprintf("%s:/workspace/output", outputDir),
		dockerImage,
	}

	// Add task-specific arguments
	if job.TaskType == "comfyui" {
		args = append(args,
			"--input", "/workspace/input/workflow.json",
			"--output", "/workspace/output/",
		)
		
		if job.Payload.Prompt != "" {
			args = append(args, "--prompt", job.Payload.Prompt)
		}
	} else if job.TaskType == "training" {
		args = append(args,
			"--dataset", "/workspace/input/",
			"--output", "/workspace/output/",
		)
	}

	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// uploadOutput uploads output files to S3
func (e *Executor) uploadOutput(outputDir, s3Path string) (string, error) {
	// For demo purposes, just check if output files exist
	// In production, use AWS SDK to upload to S3
	
	files, err := os.ReadDir(outputDir)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", fmt.Errorf("no output files generated")
	}

	// Return a mock S3 URL
	// In production, this should be the actual S3 URL after upload
	return s3Path + "result.mp4", nil
}

