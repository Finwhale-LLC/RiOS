package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/rios/worker/pkg/api"
	"github.com/rios/worker/pkg/config"
	"github.com/rios/worker/pkg/docker"
	"github.com/rios/worker/pkg/gpu"
	"github.com/rios/worker/pkg/worker"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the worker and begin processing jobs",
	Long: `Start the worker node and begin processing jobs from the RiOS network.
The worker will continuously fetch and execute jobs, earning $ROS tokens as rewards.`,
	RunE: runWorker,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runWorker(cmd *cobra.Command, args []string) error {
	PrintBanner()
	
	fmt.Println("🚀 Worker Starting")
	fmt.Println("========================")
	fmt.Println()

	// Load configuration
	fmt.Println("📋 Loading configuration...")
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w. Please run 'rios-worker register' first", err)
	}

	// Use configured API endpoint
	if cfg.APIEndpoint != "" {
		apiEndpoint = cfg.APIEndpoint
	}

	fmt.Printf("✅ Configuration loaded\n")
	fmt.Printf("   Node ID: %d\n", cfg.NodeID)
	fmt.Printf("   API Endpoint: %s\n", apiEndpoint)
	fmt.Printf("   Wallet: %s\n", cfg.WalletAddress)
	fmt.Println()

	// Check Docker
	fmt.Println("📦 Checking Docker...")
	if err := docker.CheckDockerInstalled(); err != nil {
		return err
	}
	if err := docker.CheckDockerRunning(); err != nil {
		return err
	}
	fmt.Println("✅ Docker is ready")
	fmt.Println()

	// Check GPU (required for production)
	fmt.Println("🎮 Verifying GPU configuration...")
	_, err = gpu.Detect()
	if err != nil {
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("  ⚠️  No NVIDIA GPU Detected")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		fmt.Println("This machine was registered but cannot run without a GPU.")
		fmt.Println()
		fmt.Println("📋 Requirements:")
		fmt.Println("   • NVIDIA GPU (RTX 3060 or higher recommended)")
		fmt.Println("   • NVIDIA drivers installed")
		fmt.Println("   • nvidia-smi command available")
		fmt.Println()
		fmt.Println("💡 Please switch to a machine with NVIDIA GPU hardware.")
		fmt.Println()
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("  Worker stopped.")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		os.Exit(0) // 优雅退出，不显示错误
	}
	fmt.Println("✅ GPU verified")
	fmt.Println()

	// Create API client
	client := api.NewClient(apiEndpoint)
	client.SetAuthToken(cfg.NodeAuthToken)

	// Create work directory
	homeDir, _ := os.UserHomeDir()
	workDir := filepath.Join(homeDir, ".rios", "work")
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return fmt.Errorf("failed to create work directory: %w", err)
	}

	// Create executor
	executor := worker.NewExecutor(workDir)

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start worker loop
	fmt.Println("💪 Worker is now online and ready to process jobs!")
	fmt.Println("   Press Ctrl+C to stop")
	fmt.Println()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	isProcessing := false
	totalJobsCompleted := 0
	totalRewardsEarned := 0.0

	for {
		select {
		case <-sigChan:
			fmt.Println()
			fmt.Println("⏹️  Shutting down gracefully...")
			
			// Send offline heartbeat
			if err := client.Heartbeat("offline"); err != nil {
				fmt.Printf("⚠️  Warning: Failed to send offline heartbeat: %v\n", err)
			}
			
			fmt.Println("👋 Worker stopped")
			fmt.Printf("📊 Session Summary:\n")
			fmt.Printf("   Jobs Completed: %d\n", totalJobsCompleted)
			fmt.Printf("   Total Rewards: %.8f $ROS\n", totalRewardsEarned)
			return nil

		case <-ticker.C:
			if !isProcessing {
				// Send heartbeat
				if err := client.Heartbeat("online"); err != nil {
					fmt.Printf("⚠️  Heartbeat failed: %v\n", err)
					continue
				}

				// Try to get a job
				job, err := client.GetJob()
				if err != nil {
					fmt.Printf("⚠️  Failed to get job: %v\n", err)
					continue
				}

				if job == nil {
					fmt.Printf("⏳ [%s] No jobs available, waiting...\n", time.Now().Format("15:04:05"))
					continue
				}

				// Process job
				isProcessing = true
				go func() {
					defer func() {
						isProcessing = false
					}()

					fmt.Println()
					fmt.Printf("🎯 New Job Received!\n")
					fmt.Printf("   Job ID: %s\n", job.JobID)
					fmt.Printf("   Type: %s\n", job.TaskType)
					fmt.Printf("   Docker Image: %s\n", job.Payload.DockerImage)
					fmt.Println()

					// Send busy heartbeat
					if err := client.Heartbeat("busy"); err != nil {
						fmt.Printf("⚠️  Warning: Failed to send busy heartbeat: %v\n", err)
					}

					// Execute job
					outputURL, err := executor.Execute(job)
					
					// Submit result
					req := &api.SubmitResultRequest{
						JobID: job.JobID,
					}

					if err != nil {
						fmt.Printf("❌ Job failed: %v\n", err)
						req.Status = "failed"
						req.ErrorMessage = err.Error()
					} else {
						fmt.Println("✅ Job completed successfully!")
						req.Status = "completed"
						req.OutputS3URL = outputURL
					}

					resp, submitErr := client.SubmitResult(req)
					if submitErr != nil {
						fmt.Printf("⚠️  Failed to submit result: %v\n", submitErr)
						return
					}

					if resp.Success {
						totalJobsCompleted++
						totalRewardsEarned += resp.RewardPaid
						
						fmt.Println()
						fmt.Printf("💰 Reward earned: %.8f $ROS\n", resp.RewardPaid)
						fmt.Printf("📊 Total earned this session: %.8f $ROS\n", totalRewardsEarned)
						fmt.Printf("✅ Total jobs completed: %d\n", totalJobsCompleted)
						fmt.Println()
					}
				}()
			}
		}
	}
}

