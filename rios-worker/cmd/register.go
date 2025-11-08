package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rios/worker/pkg/api"
	"github.com/rios/worker/pkg/config"
	"github.com/rios/worker/pkg/docker"
	"github.com/rios/worker/pkg/gpu"
	"github.com/spf13/cobra"
)

var skipDocker bool

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register this machine as a worker node",
	Long: `Register this machine as a worker node in the RiOS network.
This will detect your GPU configuration and register with the orchestrator.`,
	RunE: runRegister,
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().BoolVar(&skipDocker, "skip-docker", false, "Skip Docker checks (for testing)")
}

func runRegister(cmd *cobra.Command, args []string) error {
	PrintBanner()

	fmt.Println("🚀 Worker Registration")
	fmt.Println("============================")
	fmt.Println()

	// Step 1: Check Docker
	if !skipDocker {
		fmt.Println("📦 Checking Docker installation...")
		if err := docker.CheckDockerInstalled(); err != nil {
			return err
		}
		fmt.Println("✅ Docker is installed")

		if err := docker.CheckDockerRunning(); err != nil {
			return err
		}
		fmt.Println("✅ Docker daemon is running")
	} else {
		fmt.Println("⚠️  Skipping Docker checks (--skip-docker enabled)")
	}

	// Step 2: Detect GPU
	fmt.Println()
	var gpuInfo *gpu.GPUInfo

	if skipDocker {
		// 测试模式：允许使用 mock GPU
		fmt.Println("🎮 Using mock GPU configuration (--skip-docker mode)...")
		gpuInfo = &gpu.GPUInfo{
			Type:  "Mock GPU (Testing)",
			Count: 1,
			VRam:  8,
		}
		fmt.Printf("   GPU: %s (%d GPU, %d GB VRAM)\n", gpuInfo.Type, gpuInfo.Count, gpuInfo.VRam)
	} else {
		// 生产模式：必须有真实的 NVIDIA GPU
		fmt.Println("🎮 Detecting GPU configuration...")
		var err error
		gpuInfo, err = gpu.Detect()
		if err != nil {
			// GPU detection failed - 优雅退出
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("  ⚠️  No NVIDIA GPU Detected")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			fmt.Println("RiOS Worker requires an NVIDIA GPU to process tasks.")
			fmt.Println()
			fmt.Println("📋 Requirements:")
			fmt.Println("   • NVIDIA GPU (RTX 3060 or higher recommended)")
			fmt.Println("   • NVIDIA drivers installed")
			fmt.Println("   • nvidia-smi command available")
			fmt.Println()
			fmt.Println("💡 Solutions:")
			fmt.Println("   1. Install NVIDIA drivers: https://www.nvidia.com/drivers")
			fmt.Println("   2. Switch to a machine with NVIDIA GPU")
			fmt.Println("   3. For testing API only: use --skip-docker flag")
			fmt.Println()
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println("  Registration cancelled.")
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Println()
			os.Exit(0) // 优雅退出，不显示错误
		}

		fmt.Printf("✅ Detected GPU: %s\n", gpuInfo.Type)
		fmt.Printf("   Count: %d\n", gpuInfo.Count)
		fmt.Printf("   VRAM: %d GB\n", gpuInfo.VRam)
	}

	// Step 3: Get wallet address
	fmt.Println()
	fmt.Print("💰 Enter your $ROS wallet address (BSC): ")
	reader := bufio.NewReader(os.Stdin)
	walletAddress, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read wallet address: %w", err)
	}
	walletAddress = strings.TrimSpace(walletAddress)

	if len(walletAddress) != 42 || !strings.HasPrefix(walletAddress, "0x") {
		return fmt.Errorf("invalid wallet address format. Must be a valid BSC address (0x...)")
	}

	// Step 4: Optional contributor name
	fmt.Print("📝 Enter a name for your worker (optional, press Enter to skip): ")
	contributorName, _ := reader.ReadString('\n')
	contributorName = strings.TrimSpace(contributorName)

	// Step 5: Register with API
	fmt.Println()
	fmt.Println("📡 Registering with RiOS Orchestrator...")
	fmt.Printf("   API Endpoint: %s\n", apiEndpoint)

	client := api.NewClient(apiEndpoint)

	req := &api.RegisterRequest{
		GPUType:          gpuInfo.Type,
		GPUVram:          gpuInfo.VRam,
		GPUCount:         gpuInfo.Count,
		RosWalletAddress: walletAddress,
		ContributorName:  contributorName,
	}

	resp, err := client.Register(req)
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}

	if !resp.Success {
		return fmt.Errorf("registration failed: %s", resp.Message)
	}

	// Step 6: Save configuration
	fmt.Println()
	fmt.Println("💾 Saving configuration...")

	cfg := &config.Config{
		NodeID:        resp.NodeID,
		NodeAuthToken: resp.NodeAuthToken,
		APIEndpoint:   apiEndpoint,
		WalletAddress: walletAddress,
	}

	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	configPath, _ := config.GetConfigPath()
	fmt.Printf("✅ Configuration saved to: %s\n", configPath)

	// Success!
	fmt.Println()
	fmt.Println("🎉 Registration Successful!")
	fmt.Println("============================")
	fmt.Printf("Node ID: %d\n", resp.NodeID)
	fmt.Printf("Wallet: %s\n", walletAddress)
	fmt.Println()
	fmt.Println("🚀 Next Steps:")
	fmt.Println("   Run 'rios-worker run' to start contributing and earning $ROS!")
	fmt.Println()

	return nil
}
