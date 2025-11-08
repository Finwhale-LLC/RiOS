# RiOS Worker CLI

A decentralized GPU compute worker for the RiOS network. Contribute your GPU power and earn $ROS tokens.

## 📋 Prerequisites

### Required
- **NVIDIA GPU** (RTX 3060 or higher recommended)
- **NVIDIA Drivers** installed and working (`nvidia-smi` command available)
- **Docker** installed and running
- **nvidia-docker2** (NVIDIA Container Toolkit) for GPU support in Docker

### For Development Only
- **Go 1.21+** for building from source

> ⚠️ **Important**: RiOS Worker requires a real NVIDIA GPU to process tasks. Machines without NVIDIA GPUs cannot register as workers in production mode. Use `--skip-docker` flag for testing purposes only.

## 🚀 Quick Start

### One-Line Install (Recommended)

For Linux systems:

```bash
curl -fsSL https://install.rios.com.ai/worker.sh | bash
```

Or download and inspect the script first:

```bash
curl -fsSL https://install.rios.com.ai/worker.sh -o install.sh
chmod +x install.sh
./install.sh
```

**This automatically installs:**
- ✅ Docker
- ✅ NVIDIA Docker runtime (if GPU detected)
- ✅ RiOS Worker CLI
- ✅ All dependencies and configurations

### Manual Install

See [INSTALL.md](INSTALL.md) for detailed step-by-step instructions.

---

### After Installation

#### 1. Register Your Worker

```bash
rios-worker register --api https://api.rios.com.ai
```

You will be prompted to:
- Enter your $ROS wallet address (BSC chain)
- Optionally provide a name for your worker

The CLI will automatically detect your GPU configuration.

#### 2. Start Earning $ROS

```bash
rios-worker run
```

Your worker will:
- Connect to the RiOS orchestrator
- Fetch and execute jobs
- Earn $ROS tokens automatically

---

### Building from Source (Developers)

```bash
# Clone repository
git clone https://github.com/rios/worker.git
cd worker

# Build
make build

# Or use build script
./build.sh
```

## 🔧 Commands

### Register

Register your machine as a worker node:

```bash
rios-worker register [flags]
```

**Flags:**
- `--api <url>` - RiOS API endpoint (default: http://localhost:3000)

### Run

Start processing jobs and earning rewards:

```bash
rios-worker run [flags]
```

**Flags:**
- `--api <url>` - RiOS API endpoint (uses saved config by default)

Press `Ctrl+C` to gracefully stop the worker.

## 📁 Configuration

Configuration is stored in `~/.rios/config.json`:

```json
{
  "node_id": 123,
  "node_auth_token": "eyJhbGc...",
  "api_endpoint": "https://api.rios.com.ai",
  "wallet_address": "0x..."
}
```

## 🐳 Docker Requirements

The worker requires Docker with GPU support. Install nvidia-docker2:

### Ubuntu/Debian

```bash
distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | sudo apt-key add -
curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | \
  sudo tee /etc/apt/sources.list.d/nvidia-docker.list

sudo apt-get update && sudo apt-get install -y nvidia-docker2
sudo systemctl restart docker
```

### Test GPU Support

```bash
docker run --rm --gpus all nvidia/cuda:12.0.0-base-ubuntu22.04 nvidia-smi
```

## 💰 Earning Rewards

- Rewards are paid in $ROS tokens
- Earnings accumulate in your internal balance
- Withdraw to your BSC wallet address at any time
- 90% of job cost goes to workers, 10% platform fee

## 🔍 Monitoring

The worker displays:
- Current status (online/busy)
- Jobs completed
- Rewards earned per job
- Total session earnings

## 📊 Supported Task Types

### ComfyUI Image/Video Generation

Process AI image and video generation tasks using ComfyUI workflows.

### Model Training

Contribute GPU power for AI model training tasks.

## 🛠️ Development

### Project Structure

```
rios-worker/
├── cmd/           # CLI commands
│   ├── root.go    # Root command
│   ├── register.go # Registration command
│   └── run.go     # Worker run command
├── pkg/
│   ├── api/       # API client
│   ├── config/    # Configuration management
│   ├── docker/    # Docker utilities
│   ├── gpu/       # GPU detection
│   └── worker/    # Job executor
├── main.go
├── go.mod
└── README.md
```

### Build

```bash
# Development build
go build -o rios-worker main.go

# Production build (smaller binary)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o rios-worker main.go
```

### Cross-Platform Build

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o rios-worker-linux main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o rios-worker-macos main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o rios-worker.exe main.go
```

## ⚠️ Troubleshooting

### GPU Not Detected

**Error:** `❌ NVIDIA GPU not detected!`

**Cause:** No NVIDIA GPU found on your machine, or NVIDIA drivers not installed.

**Solutions:**

1. **Verify GPU Hardware:**
   ```bash
   # Check if NVIDIA GPU is installed
   lspci | grep -i nvidia
   ```

2. **Check NVIDIA Drivers:**
   ```bash
   # This should show GPU information
   nvidia-smi
   ```
   
   If command not found, install drivers:
   - Ubuntu/Debian: https://www.nvidia.com/drivers
   - Or use package manager: `sudo apt install nvidia-driver-535`

3. **Verify nvidia-docker:**
   ```bash
   docker run --rm --gpus all nvidia/cuda:12.0.0-base-ubuntu22.04 nvidia-smi
   ```

4. **For Testing Only:**
   ```bash
   # Skip GPU checks (for API testing only)
   ./rios-worker register --skip-docker
   ```

> ⚠️ **Note**: Machines without NVIDIA GPUs can only be used for testing with `--skip-docker` flag. They cannot process real tasks or earn rewards.

### Docker Not Running

```bash
# Start Docker daemon
sudo systemctl start docker

# Check status
docker ps
```

### Configuration Issues

```bash
# View configuration
cat ~/.rios/config.json

# Re-register if needed
rios-worker register --api <your-api-endpoint>
```

## 📝 License

MIT License - See LICENSE file for details

## 🤝 Support

- Documentation: https://docs.rios.com.ai
- Discord: https://discord.gg/rios
- Email: support@rios.com.ai

## 🌟 Contributing

Contributions welcome! Please submit issues and pull requests on GitHub.

