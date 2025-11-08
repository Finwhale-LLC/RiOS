# Testing the Installation Script

This guide is for developers who want to test the `install.sh` script locally.

## 🧪 Testing on Local Machine

### Test 1: Dry Run (Review Script)

```bash
# Review the script without executing
cat install.sh

# Or use a linter
shellcheck install.sh
```

### Test 2: Local Installation

The script detects if a local binary exists and uses it:

```bash
# Build the worker first
make build

# Run installation script
./install.sh
```

The script will:
- Use the local `rios-worker` binary instead of downloading
- Install Docker (if needed)
- Install NVIDIA Docker (if GPU detected)
- Copy binary to /usr/local/bin

### Test 3: Test with Docker Container

Test the script in a clean Ubuntu container:

```bash
# Start Ubuntu container
docker run -it --rm ubuntu:22.04 bash

# Inside container, install prerequisites
apt-get update
apt-get install -y curl wget

# Test the script
curl -fsSL https://raw.githubusercontent.com/rios/worker/main/install.sh | bash
```

### Test 4: Test Individual Functions

You can source the script and test individual functions:

```bash
# Source without running
source install.sh

# Test specific functions
detect_os
check_nvidia_gpu
```

## 🐧 Testing on Different Linux Distributions

### Ubuntu 22.04

```bash
docker run -it --rm ubuntu:22.04 bash
apt-get update && apt-get install -y curl
curl -fsSL https://YOUR_SERVER/install.sh | bash
```

### Ubuntu 20.04

```bash
docker run -it --rm ubuntu:20.04 bash
apt-get update && apt-get install -y curl
curl -fsSL https://YOUR_SERVER/install.sh | bash
```

### Debian 11

```bash
docker run -it --rm debian:11 bash
apt-get update && apt-get install -y curl
curl -fsSL https://YOUR_SERVER/install.sh | bash
```

### CentOS 8

```bash
docker run -it --rm centos:8 bash
yum install -y curl
curl -fsSL https://YOUR_SERVER/install.sh | bash
```

## 🔍 What to Test

### ✅ Checklist

- [ ] Script detects OS correctly
- [ ] Script detects architecture (amd64/arm64)
- [ ] Docker installation works
- [ ] Docker service starts
- [ ] User added to docker group
- [ ] NVIDIA GPU detection works
- [ ] NVIDIA Docker installation (if GPU present)
- [ ] Worker CLI downloads correctly
- [ ] Worker CLI is executable
- [ ] Worker CLI `--help` works
- [ ] Colored output works correctly
- [ ] Error handling works (bad URL, no permissions, etc.)

### Error Scenarios to Test

1. **No internet connection:**
   ```bash
   # Disconnect network and test
   ```

2. **Insufficient permissions:**
   ```bash
   # Run without sudo
   ./install.sh
   ```

3. **Already installed Docker:**
   ```bash
   # Install Docker first, then run script
   ```

4. **No NVIDIA GPU:**
   ```bash
   # Test on system without GPU
   ```

## 📊 Testing Matrix

| OS | Architecture | Docker | NVIDIA | Expected Result |
|----|--------------|--------|--------|-----------------|
| Ubuntu 22.04 | amd64 | ❌ | ❌ | Install Docker + CLI |
| Ubuntu 22.04 | amd64 | ✅ | ❌ | Install CLI only |
| Ubuntu 22.04 | amd64 | ❌ | ✅ | Install all |
| Ubuntu 22.04 | amd64 | ✅ | ✅ | Install NVIDIA Docker + CLI |
| Ubuntu 20.04 | amd64 | ❌ | ❌ | Install Docker + CLI |
| Debian 11 | amd64 | ❌ | ❌ | Install Docker + CLI |
| CentOS 8 | amd64 | ❌ | ❌ | Install Docker + CLI |
| Ubuntu 22.04 | arm64 | ❌ | ❌ | Install Docker + CLI (ARM) |
| macOS | arm64 | ❌ | ❌ | Prompt for Docker Desktop |

## 🐛 Common Issues

### Issue 1: Script fails to download

**Solution:** Check internet connection and URL

```bash
curl -I https://install.rios.com.ai/worker.sh
```

### Issue 2: Docker installation fails

**Solution:** Check system logs

```bash
journalctl -xe
```

### Issue 3: Permission denied

**Solution:** Ensure script has execute permissions

```bash
chmod +x install.sh
```

## 📝 Manual Testing Steps

1. **Fresh Ubuntu VM:**
   ```bash
   # Create fresh Ubuntu 22.04 VM
   # SSH into it
   
   # Run one-line install
   curl -fsSL https://install.rios.com.ai/worker.sh | bash
   
   # Verify installation
   docker --version
   rios-worker --help
   
   # Try to register (will fail without GPU, expected)
   rios-worker register --api http://test-api
   ```

2. **With NVIDIA GPU:**
   ```bash
   # On machine with NVIDIA GPU
   
   # Run install
   ./install.sh
   
   # Verify GPU access
   docker run --rm --gpus all nvidia/cuda:12.0.0-base-ubuntu22.04 nvidia-smi
   
   # Register and test
   rios-worker register --api https://api.rios.com.ai
   ```

## 🔄 CI/CD Testing

Add to `.github/workflows/test-install.yml`:

```yaml
name: Test Install Script

on: [push, pull_request]

jobs:
  test-ubuntu:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Test install script
        run: |
          cd rios-worker
          ./install.sh
          
      - name: Verify installation
        run: |
          docker --version
          rios-worker --help
```

## 📦 Preparing for Production

Before releasing:

1. **Update URLs in script:**
   - Change GitHub repo URL
   - Update API endpoint
   - Update documentation links

2. **Host script on CDN:**
   ```bash
   # Upload to your CDN
   aws s3 cp install.sh s3://cdn.rios.com.ai/worker.sh
   ```

3. **Create DNS record:**
   ```
   install.rios.com.ai → CDN URL
   ```

4. **Test production URL:**
   ```bash
   curl -fsSL https://install.rios.com.ai/worker.sh | bash
   ```

## 🎯 Success Criteria

Installation is successful if:

- ✅ Script completes without errors
- ✅ Docker is installed and running
- ✅ Worker CLI is in PATH
- ✅ `rios-worker --help` works
- ✅ NVIDIA Docker configured (if GPU present)
- ✅ User can run `rios-worker register`

