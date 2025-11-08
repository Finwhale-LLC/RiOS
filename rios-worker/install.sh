#!/bin/bash

#################################################################
# RiOS Worker - One-Click Installation Script
# 
# This script automates the installation of:
# - Docker
# - NVIDIA Docker runtime (if GPU detected)
# - RiOS Worker CLI
#
# Usage:
#   curl -fsSL https://install.rios.com.ai | bash
#   OR
#   bash install.sh
#
#################################################################

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
WORKER_VERSION="latest"
WORKER_BINARY="rios-worker"
INSTALL_DIR="/usr/local/bin"
GITHUB_REPO="rios/worker"  # Update with actual repo

# Print colored output
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_header() {
    echo ""
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "  🚀 RiOS Worker - Automated Installation"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo ""
}

# Check if running as root
check_root() {
    if [ "$EUID" -eq 0 ]; then 
        print_warning "Running as root. This is not recommended for security."
        print_info "Press Ctrl+C to abort, or Enter to continue..."
        read
    fi
}

# Detect OS
detect_os() {
    print_info "Detecting operating system..."
    
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    case "$OS" in
        linux*)
            OS="linux"
            ;;
        darwin*)
            OS="darwin"
            print_warning "macOS detected. Docker Desktop is required."
            ;;
        *)
            print_error "Unsupported operating system: $OS"
            exit 1
            ;;
    esac
    
    case "$ARCH" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            print_error "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac
    
    print_success "Detected: $OS/$ARCH"
}

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Install Docker on Linux
install_docker_linux() {
    if command_exists docker; then
        print_success "Docker is already installed"
        docker --version
        return 0
    fi
    
    print_info "Installing Docker..."
    
    # Detect Linux distribution
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        DISTRO=$ID
    else
        print_error "Cannot detect Linux distribution"
        exit 1
    fi
    
    case "$DISTRO" in
        ubuntu|debian)
            # Ubuntu/Debian
            print_info "Installing Docker on Ubuntu/Debian..."
            curl -fsSL https://get.docker.com | sudo sh
            ;;
        centos|rhel|fedora)
            # CentOS/RHEL/Fedora
            print_info "Installing Docker on CentOS/RHEL/Fedora..."
            curl -fsSL https://get.docker.com | sudo sh
            ;;
        *)
            print_warning "Unknown distribution: $DISTRO"
            print_info "Attempting generic Docker installation..."
            curl -fsSL https://get.docker.com | sudo sh
            ;;
    esac
    
    # Add current user to docker group
    print_info "Adding user to docker group..."
    sudo usermod -aG docker $USER
    
    # Start Docker service
    print_info "Starting Docker service..."
    sudo systemctl start docker || true
    sudo systemctl enable docker || true
    
    print_success "Docker installed successfully"
    print_warning "You may need to log out and back in for docker group changes to take effect"
}

# Install Docker on macOS
install_docker_macos() {
    if command_exists docker; then
        print_success "Docker is already installed"
        docker --version
        return 0
    fi
    
    print_warning "Docker Desktop is required for macOS"
    print_info "Please install Docker Desktop from: https://docs.docker.com/desktop/install/mac-install/"
    print_info "After installation, run this script again."
    exit 1
}

# Check NVIDIA GPU
check_nvidia_gpu() {
    print_info "Checking for NVIDIA GPU..."
    
    if command_exists nvidia-smi; then
        print_success "NVIDIA GPU detected:"
        nvidia-smi --query-gpu=name,memory.total --format=csv,noheader | head -1
        return 0
    else
        print_warning "No NVIDIA GPU detected or drivers not installed"
        return 1
    fi
}

# Install NVIDIA Docker support
install_nvidia_docker() {
    if ! check_nvidia_gpu; then
        print_info "Skipping NVIDIA Docker installation (no GPU detected)"
        return 0
    fi
    
    # Check if nvidia-docker2 is already installed
    if docker info 2>/dev/null | grep -q "nvidia"; then
        print_success "NVIDIA Docker runtime already configured"
        return 0
    fi
    
    print_info "Installing NVIDIA Docker support..."
    
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        DISTRO=$ID
        VERSION_ID=$VERSION_ID
    else
        print_error "Cannot detect Linux distribution"
        return 1
    fi
    
    case "$DISTRO" in
        ubuntu|debian)
            distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
            
            # Add NVIDIA Docker GPG key
            curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | \
                sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg
            
            # Add NVIDIA Docker repository
            curl -s -L https://nvidia.github.io/libnvidia-container/$distribution/libnvidia-container.list | \
                sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
                sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list
            
            # Install
            sudo apt-get update
            sudo apt-get install -y nvidia-docker2
            
            # Restart Docker
            sudo systemctl restart docker
            
            print_success "NVIDIA Docker installed successfully"
            ;;
        centos|rhel|fedora)
            distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
            curl -s -L https://nvidia.github.io/libnvidia-container/$distribution/libnvidia-container.repo | \
                sudo tee /etc/yum.repos.d/nvidia-container-toolkit.repo
            
            sudo yum install -y nvidia-docker2
            sudo systemctl restart docker
            
            print_success "NVIDIA Docker installed successfully"
            ;;
        *)
            print_warning "Automatic NVIDIA Docker installation not supported for: $DISTRO"
            print_info "Please install manually: https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/install-guide.html"
            ;;
    esac
}

# Download and install Worker CLI
install_worker_cli() {
    print_info "Installing RiOS Worker CLI..."
    
    # Determine download URL based on OS and architecture
    BINARY_NAME="rios-worker-${OS}-${ARCH}"
    
    # For local testing, use the built binary
    if [ -f "./rios-worker" ]; then
        print_info "Using local binary for installation..."
        sudo cp ./rios-worker "$INSTALL_DIR/$WORKER_BINARY"
    else
        # In production, download from GitHub releases
        DOWNLOAD_URL="https://github.com/$GITHUB_REPO/releases/download/$WORKER_VERSION/$BINARY_NAME"
        
        print_info "Downloading from: $DOWNLOAD_URL"
        
        # Download
        if command_exists wget; then
            wget -q --show-progress -O "/tmp/$WORKER_BINARY" "$DOWNLOAD_URL" || {
                print_error "Download failed. Please check your internet connection."
                exit 1
            }
        elif command_exists curl; then
            curl -L -o "/tmp/$WORKER_BINARY" "$DOWNLOAD_URL" || {
                print_error "Download failed. Please check your internet connection."
                exit 1
            }
        else
            print_error "Neither wget nor curl is available. Please install one."
            exit 1
        fi
        
        # Install
        sudo mv "/tmp/$WORKER_BINARY" "$INSTALL_DIR/$WORKER_BINARY"
    fi
    
    # Make executable
    sudo chmod +x "$INSTALL_DIR/$WORKER_BINARY"
    
    print_success "Worker CLI installed to: $INSTALL_DIR/$WORKER_BINARY"
}

# Verify installation
verify_installation() {
    print_info "Verifying installation..."
    
    # Check Docker
    if ! command_exists docker; then
        print_error "Docker installation failed"
        return 1
    fi
    
    # Check Worker CLI
    if ! command_exists $WORKER_BINARY; then
        print_error "Worker CLI installation failed"
        return 1
    fi
    
    # Test Worker CLI
    if ! $WORKER_BINARY --help >/dev/null 2>&1; then
        print_error "Worker CLI is not working properly"
        return 1
    fi
    
    print_success "All components verified successfully"
}

# Print next steps
print_next_steps() {
    echo ""
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "  🎉 Installation Complete!"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo ""
    print_success "RiOS Worker has been installed successfully!"
    echo ""
    echo "📋 Next Steps:"
    echo ""
    echo "  1. Register your worker node:"
    echo "     ${GREEN}rios-worker register --api https://api.rios.com.ai${NC}"
    echo ""
    echo "  2. Start earning $ROS tokens:"
    echo "     ${GREEN}rios-worker run${NC}"
    echo ""
    echo "📚 Documentation:"
    echo "  - Worker Guide: https://docs.rios.com.ai/worker"
    echo "  - Troubleshooting: https://docs.rios.com.ai/troubleshooting"
    echo ""
    echo "💬 Support:"
    echo "  - Discord: https://discord.gg/rios"
    echo "  - Email: support@rios.com.ai"
    echo ""
    
    if [ "$NEW_DOCKER_USER" = true ]; then
        print_warning "IMPORTANT: You were added to the 'docker' group."
        print_warning "Please log out and log back in for changes to take effect."
        echo ""
        echo "Or run: ${YELLOW}newgrp docker${NC}"
        echo ""
    fi
    
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
}

# Main installation flow
main() {
    NEW_DOCKER_USER=false
    
    print_header
    
    # Check root
    check_root
    
    # Detect OS
    detect_os
    
    # Install Docker
    if [ "$OS" = "linux" ]; then
        install_docker_linux
        NEW_DOCKER_USER=true
    elif [ "$OS" = "darwin" ]; then
        install_docker_macos
    fi
    
    # Install NVIDIA Docker (Linux only)
    if [ "$OS" = "linux" ]; then
        install_nvidia_docker
    fi
    
    # Install Worker CLI
    install_worker_cli
    
    # Verify
    verify_installation
    
    # Print next steps
    print_next_steps
}

# Run main function
main "$@"

