# RiOS Worker - Installation Guide

Complete installation guide for RiOS Worker CLI.

## 🚀 Quick Install (Recommended)

### One-Line Install

For Linux systems with internet connection:

```bash
curl -fsSL https://install.rios.com.ai/worker.sh | bash
```

Or if you prefer to download and inspect first:

```bash
curl -fsSL https://install.rios.com.ai/worker.sh -o install.sh
chmod +x install.sh
./install.sh
```

**What this does:**
- ✅ Detects your operating system
- ✅ Installs Docker
- ✅ Installs NVIDIA Docker runtime (if GPU detected)
- ✅ Downloads and installs Worker CLI
- ✅ Configures permissions

---

## 📋 System Requirements

### Hardware Requirements

**Minimum:**
- CPU: 4 cores
- RAM: 8 GB
- Storage: 50 GB free space
- GPU: NVIDIA RTX 3060 or higher

**Recommended:**
- CPU: 8+ cores
- RAM: 16+ GB
- Storage: 100+ GB SSD
- GPU: NVIDIA RTX 4090 or A100

### Software Requirements

- **Operating System:**
  - Ubuntu 20.04+ (recommended)
  - Debian 11+
  - CentOS 8+
  - RHEL 8+
  - Fedora 35+

- **NVIDIA Requirements:**
  - NVIDIA GPU with CUDA support
  - NVIDIA drivers (version 525+ recommended)
  - nvidia-smi working

- **Network:**
  - Stable internet connection
  - Ports 80/443 outbound (for API communication)

---

## 🔧 Manual Installation

If the automated script doesn't work for your system, follow these manual steps:

### Step 1: Install Docker

#### Ubuntu/Debian

```bash
# Update package index
sudo apt-get update

# Install prerequisites
sudo apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# Add Docker's official GPG key
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | \
    sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

# Set up repository
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] \
  https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io

# Add user to docker group
sudo usermod -aG docker $USER
```

#### CentOS/RHEL

```bash
# Install required packages
sudo yum install -y yum-utils

# Add Docker repository
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

# Install Docker
sudo yum install -y docker-ce docker-ce-cli containerd.io

# Start Docker
sudo systemctl start docker
sudo systemctl enable docker

# Add user to docker group
sudo usermod -aG docker $USER
```

**After installation, log out and back in for group changes to take effect.**

---

### Step 2: Install NVIDIA Drivers

Check if drivers are installed:

```bash
nvidia-smi
```

If not installed:

#### Ubuntu/Debian

```bash
# Add NVIDIA driver PPA
sudo add-apt-repository ppa:graphics-drivers/ppa
sudo apt-get update

# Install recommended driver
sudo ubuntu-drivers autoinstall

# Or install specific version
sudo apt-get install -y nvidia-driver-535

# Reboot
sudo reboot
```

#### CentOS/RHEL

```bash
# Install EPEL repository
sudo yum install -y epel-release

# Install NVIDIA driver
sudo yum install -y nvidia-driver

# Reboot
sudo reboot
```

---

### Step 3: Install NVIDIA Container Toolkit

#### Ubuntu/Debian

```bash
# Add NVIDIA container toolkit repository
distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | \
    sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg
curl -s -L https://nvidia.github.io/libnvidia-container/$distribution/libnvidia-container.list | \
    sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
    sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list

# Install
sudo apt-get update
sudo apt-get install -y nvidia-docker2

# Restart Docker
sudo systemctl restart docker
```

#### CentOS/RHEL

```bash
# Add repository
distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
curl -s -L https://nvidia.github.io/libnvidia-container/$distribution/libnvidia-container.repo | \
    sudo tee /etc/yum.repos.d/nvidia-container-toolkit.repo

# Install
sudo yum install -y nvidia-docker2

# Restart Docker
sudo systemctl restart docker
```

**Verify installation:**

```bash
docker run --rm --gpus all nvidia/cuda:12.0.0-base-ubuntu22.04 nvidia-smi
```

You should see your GPU information.

---

### Step 4: Install RiOS Worker CLI

#### Download Binary

```bash
# Set version and architecture
VERSION="latest"
ARCH="amd64"  # or "arm64"

# Download
wget https://github.com/rios/worker/releases/download/$VERSION/rios-worker-linux-$ARCH

# Make executable
chmod +x rios-worker-linux-$ARCH

# Move to PATH
sudo mv rios-worker-linux-$ARCH /usr/local/bin/rios-worker
```

#### Build from Source (Optional)

```bash
# Install Go 1.21+
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Clone repository
git clone https://github.com/rios/worker.git
cd worker

# Build
make build

# Install
sudo cp rios-worker /usr/local/bin/
```

---

## ✅ Verify Installation

Run these commands to verify everything is working:

```bash
# Check Docker
docker --version
docker ps

# Check NVIDIA
nvidia-smi

# Check NVIDIA Docker
docker run --rm --gpus all nvidia/cuda:12.0.0-base-ubuntu22.04 nvidia-smi

# Check Worker CLI
rios-worker --help
```

---

## 🎯 Getting Started

After installation:

### 1. Register Your Worker

```bash
rios-worker register --api https://api.rios.com.ai
```

You'll be prompted for:
- Your $ROS wallet address (BSC)
- Optional worker name

### 2. Start Earning

```bash
rios-worker run
```

The worker will:
- Connect to RiOS network
- Fetch and execute jobs
- Automatically earn $ROS tokens

---

## 🔄 Running as System Service (Recommended for Servers)

To run Worker as a systemd service that starts automatically on boot:

```bash
# Download setup script
curl -fsSL https://install.rios.com.ai/setup-service.sh -o setup-service.sh
chmod +x setup-service.sh

# Run as root
sudo ./setup-service.sh
```

Or manually:

```bash
# Copy service file
sudo cp rios-worker.service /etc/systemd/system/rios-worker@.service

# Enable and start for your user
sudo systemctl enable rios-worker@$USER
sudo systemctl start rios-worker@$USER

# Check status
sudo systemctl status rios-worker@$USER
```

**Service Management:**

```bash
# View logs
sudo journalctl -u rios-worker@$USER -f

# Restart
sudo systemctl restart rios-worker@$USER

# Stop
sudo systemctl stop rios-worker@$USER
```

---

## 🐳 Docker Compose (Alternative)

If you prefer Docker Compose for deployment:

Create `docker-compose.yml`:

```yaml
version: '3.8'

services:
  rios-worker:
    image: rios/worker:latest
    runtime: nvidia
    restart: unless-stopped
    environment:
      - API_ENDPOINT=https://api.rios.com.ai
      - NODE_AUTH_TOKEN=${NODE_AUTH_TOKEN}
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - ./work:/work
      - ./config:/root/.rios
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: all
              capabilities: [gpu]
```

Deploy:

```bash
docker-compose up -d
```

---

## 🔒 Security Considerations

### Firewall Rules

Allow outbound connections:
- Port 443 (HTTPS) to api.rios.com.ai
- Docker registry ports (if pulling custom images)

### User Permissions

The worker needs:
- Docker group membership (for container operations)
- Read access to GPU devices
- Network access for API communication

### Wallet Security

- Never share your private key
- Store wallet addresses securely
- Use hardware wallets when possible

---

## 📊 System Monitoring

### Check Worker Status

```bash
# View logs
journalctl -u rios-worker -f

# Check resource usage
docker stats

# GPU usage
nvidia-smi -l 1
```

### Performance Tuning

```bash
# Limit Docker resources
docker run --memory="8g" --cpus="4" ...

# GPU memory allocation
export CUDA_VISIBLE_DEVICES=0  # Use specific GPU
```

---

## 🛠️ Troubleshooting

### Docker Permission Denied

```bash
sudo usermod -aG docker $USER
newgrp docker
```

### NVIDIA Driver Issues

```bash
# Reinstall drivers
sudo apt-get purge nvidia-*
sudo apt-get install nvidia-driver-535
sudo reboot
```

### Worker Won't Start

```bash
# Check logs
rios-worker run --verbose

# Verify configuration
cat ~/.rios/config.json

# Test Docker
docker run hello-world
```

### GPU Not Detected

```bash
# Check if GPU is visible
lspci | grep -i nvidia

# Check drivers
nvidia-smi

# Check Docker GPU support
docker run --rm --gpus all nvidia/cuda:12.0.0-base-ubuntu22.04 nvidia-smi
```

---

## 📞 Support

- **Documentation:** https://docs.rios.com.ai
- **Discord:** https://discord.gg/rios
- **GitHub Issues:** https://github.com/rios/worker/issues
- **Email:** support@rios.com.ai

---

## 🔄 Updating

To update to the latest version:

```bash
# Using install script
curl -fsSL https://install.rios.com.ai/worker.sh | bash

# Or manually
wget https://github.com/rios/worker/releases/latest/download/rios-worker-linux-amd64
sudo mv rios-worker-linux-amd64 /usr/local/bin/rios-worker
sudo chmod +x /usr/local/bin/rios-worker
```

---

## 🗑️ Uninstallation

To completely remove RiOS Worker:

```bash
# Stop worker
sudo systemctl stop rios-worker

# Remove binary
sudo rm /usr/local/bin/rios-worker

# Remove configuration
rm -rf ~/.rios

# Optional: Remove Docker
sudo apt-get purge docker-ce docker-ce-cli containerd.io
sudo rm -rf /var/lib/docker
```

---

**Happy Mining! 🚀**

