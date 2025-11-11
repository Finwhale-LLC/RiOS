# Installation

This guide covers detailed installation instructions for RiOS components.

## Installing the RiOS CLI

The RiOS CLI (Command Line Interface) is the primary tool for managing deployments from your terminal.

### macOS

#### Using Homebrew (Recommended)
```bash
brew tap rios/tap
brew install rios-cli
```

#### Manual Installation
```bash
# Download the latest release
curl -LO https://github.com/rios/cli/releases/latest/download/rios-darwin-amd64

# Make it executable
chmod +x rios-darwin-amd64

# Move to PATH
sudo mv rios-darwin-amd64 /usr/local/bin/rios

# Verify installation
rios version
```

### Linux

#### Debian/Ubuntu
```bash
# Add RiOS repository
curl -fsSL https://packages.rios.com.ai/gpg | sudo gpg --dearmor -o /usr/share/keyrings/rios-archive-keyring.gpg

echo "deb [signed-by=/usr/share/keyrings/rios-archive-keyring.gpg] https://packages.rios.com.ai/apt stable main" | sudo tee /etc/apt/sources.list.d/rios.list

# Install
sudo apt update
sudo apt install rios-cli
```

#### CentOS/RHEL/Fedora
```bash
# Add RiOS repository
sudo tee /etc/yum.repos.d/rios.repo << 'EOF'
[rios]
name=RiOS Repository
baseurl=https://packages.rios.com.ai/yum
enabled=1
gpgcheck=1
gpgkey=https://packages.rios.com.ai/gpg
EOF

# Install
sudo yum install rios-cli
```

#### Manual Installation
```bash
# Download the latest release
curl -LO https://github.com/rios/cli/releases/latest/download/rios-linux-amd64

# Make it executable
chmod +x rios-linux-amd64

# Move to PATH
sudo mv rios-linux-amd64 /usr/local/bin/rios

# Verify installation
rios version
```

### Windows

#### Using Chocolatey
```powershell
choco install rios-cli
```

#### Using Scoop
```powershell
scoop bucket add rios https://github.com/rios/scoop-bucket
scoop install rios-cli
```

#### Manual Installation
1. Download the latest release from [GitHub Releases](https://github.com/rios/cli/releases)
2. Extract `rios-windows-amd64.exe`
3. Rename to `rios.exe`
4. Add to your PATH
5. Verify: `rios version`

## Configuration

### Initial Setup

After installation, configure the CLI:

```bash
# Login to your account
rios login

# Or set API token manually
rios config set token YOUR_API_TOKEN

# Set default region (optional)
rios config set region us-west-1

# Verify configuration
rios config list
```

### Configuration File

The CLI stores configuration in `~/.rios/config.yaml`:

```yaml
current_context: default
contexts:
  default:
    api_endpoint: https://api.rios.com.ai
    token: your_api_token
    region: us-west-1
  production:
    api_endpoint: https://api.rios.com.ai
    token: prod_token
    region: eu-west-1
```

### Switch Between Contexts

```bash
# Create new context
rios config create-context production --region eu-west-1

# Switch context
rios config use-context production

# List contexts
rios config get-contexts
```

## Installing Docker (Required for Workers)

If you plan to run a worker node, Docker is required.

### macOS
```bash
# Download Docker Desktop from
# https://www.docker.com/products/docker-desktop

# Or using Homebrew
brew install --cask docker
```

### Linux (Ubuntu/Debian)
```bash
# Update package index
sudo apt update

# Install dependencies
sudo apt install -y apt-transport-https ca-certificates curl software-properties-common

# Add Docker GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Add Docker repository
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io

# Add user to docker group
sudo usermod -aG docker $USER

# Start Docker
sudo systemctl enable docker
sudo systemctl start docker
```

### Verify Docker Installation
```bash
docker --version
docker run hello-world
```

## Installing Worker Software

For detailed worker installation, see the [Worker Setup Guide](../worker-setup/installation.md).

Quick installation:

```bash
# Download worker binary
curl -LO https://github.com/rios/worker/releases/latest/download/rios-worker-linux-amd64

# Make executable
chmod +x rios-worker-linux-amd64
sudo mv rios-worker-linux-amd64 /usr/local/bin/rios-worker

# Verify
rios-worker version
```

## Building from Source

### Prerequisites
- Go 1.19 or later
- Git
- Make

### Clone and Build

```bash
# Clone repository
git clone https://github.com/rios/rios-cli.git
cd rios-cli

# Build
make build

# Install
sudo make install

# Verify
rios version
```

### Build Worker from Source

```bash
# Clone repository
git clone https://github.com/rios/rios-worker.git
cd rios-worker

# Build
make build

# Install
sudo make install

# Verify
rios-worker version
```

## Updating

### CLI Updates

```bash
# Using package manager
brew upgrade rios-cli           # macOS
sudo apt update && sudo apt upgrade rios-cli  # Ubuntu/Debian
sudo yum update rios-cli        # CentOS/RHEL

# Manual update
rios update
```

### Worker Updates

```bash
# Using package manager
sudo apt update && sudo apt upgrade rios-worker

# Or download latest binary
rios-worker update
```

## Uninstallation

### Remove CLI

```bash
# macOS (Homebrew)
brew uninstall rios-cli

# Ubuntu/Debian
sudo apt remove rios-cli

# Manual removal
sudo rm /usr/local/bin/rios
rm -rf ~/.rios
```

### Remove Worker

```bash
# Stop worker service
sudo systemctl stop rios-worker

# Remove package
sudo apt remove rios-worker

# Remove data
sudo rm -rf /var/lib/rios-worker
```

## Troubleshooting

### Command Not Found

If you get "command not found":

```bash
# Check if binary is in PATH
which rios

# Add to PATH (add to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:/usr/local/bin
```

### Permission Denied

```bash
# Make sure binary is executable
chmod +x /usr/local/bin/rios

# Or run with sudo (not recommended for regular use)
sudo rios
```

### Connection Issues

```bash
# Test API connectivity
curl https://api.rios.com.ai/health

# Check DNS resolution
nslookup api.rios.com.ai

# Verify firewall settings
sudo ufw status
```

## Next Steps

After installation:
- [Quick Start Guide](quick-start.md) - Deploy your first application
- [Worker Setup](../worker-setup/README.md) - Set up a worker node
- [API Reference](../api-reference/README.md) - Learn the API

## Getting Help

- **Documentation**: Full documentation available
- **GitHub Issues**: [Report bugs](https://github.com/rios/cli/issues)
- **Community**: [Join our forum](https://community.rios.com.ai)
- **Support**: support@rios.com.ai

