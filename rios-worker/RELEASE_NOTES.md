# 🚀 RiOS Worker CLI - Release v0.1.0

**Release Date:** 2025-11-07  
**Status:** ✅ Production Ready

---

## 📦 Available Downloads

### Individual Platform Packages (Recommended)

Download only the version you need:

| Platform | Download | Size | SHA256 |
|----------|----------|------|--------|
| **Linux x86_64** | [rios-worker-v0.1.0-linux-amd64.tar.gz](dist/rios-worker-v0.1.0-linux-amd64.tar.gz) | 3.2 MB | See SHA256SUMS.txt |
| **Linux ARM64** | [rios-worker-v0.1.0-linux-arm64.tar.gz](dist/rios-worker-v0.1.0-linux-arm64.tar.gz) | 2.9 MB | See SHA256SUMS.txt |
| **macOS Intel** | [rios-worker-v0.1.0-darwin-amd64.tar.gz](dist/rios-worker-v0.1.0-darwin-amd64.tar.gz) | 3.3 MB | See SHA256SUMS.txt |
| **macOS Apple Silicon** | [rios-worker-v0.1.0-darwin-arm64.tar.gz](dist/rios-worker-v0.1.0-darwin-arm64.tar.gz) | 3.0 MB | See SHA256SUMS.txt |
| **Windows x86_64** | [rios-worker-v0.1.0-windows-amd64.zip](dist/rios-worker-v0.1.0-windows-amd64.zip) | 3.3 MB | See SHA256SUMS.txt |

### All Platforms Bundle

| Download | Size | Description |
|----------|------|-------------|
| [rios-worker-v0.1.0-all-platforms.tar.gz](dist/rios-worker-v0.1.0-all-platforms.tar.gz) | 16 MB | All binaries (tar.gz) |
| [rios-worker-v0.1.0-all-platforms.zip](dist/rios-worker-v0.1.0-all-platforms.zip) | 32 MB | All binaries (zip) |

---

## 🆕 What's New in v0.1.0

### Features

- ✅ **GPU Auto-Detection** - Automatically detects NVIDIA GPU configuration
- ✅ **Docker Integration** - Seamless Docker container execution
- ✅ **JWT Authentication** - Secure API communication
- ✅ **Automatic Task Fetching** - Polls for jobs every 10 seconds
- ✅ **Real-time Rewards** - Shows $ROS earnings after each task
- ✅ **Graceful Shutdown** - Ctrl+C for safe exit
- ✅ **Config Management** - Stores configuration in ~/.rios/config.json
- ✅ **Heartbeat Mechanism** - Keeps node status updated

### CLI Commands

- `rios-worker register` - Register your GPU node
- `rios-worker run` - Start processing tasks and earning $ROS
- `rios-worker --help` - Show help information
- `rios-worker --version` - Show version (future)

### Flags

- `--api <url>` - Custom API endpoint (default: http://localhost:3000)
- `--skip-docker` - Skip Docker/GPU checks (testing only)

---

## 🎯 System Requirements

### Hardware

- **CPU:** 4+ cores
- **RAM:** 8+ GB
- **Storage:** 50+ GB free space
- **GPU:** NVIDIA RTX 3060 or higher (required)

### Software

- **Operating System:**
  - Ubuntu 20.04+ ✅
  - Debian 11+ ✅
  - CentOS 8+ ✅
  - RHEL 8+ ✅
  - Fedora 35+ ✅
  - Windows 10+ ✅
  - macOS 11+ ⚠️ (testing only, no NVIDIA GPU)

- **Dependencies:**
  - Docker (latest stable)
  - NVIDIA Drivers (525+)
  - nvidia-docker2 (NVIDIA Container Toolkit)

---

## 🚀 Quick Install

### One-Line Install (Linux)

```bash
curl -fsSL https://install.rios.com.ai/worker.sh | bash
```

### Manual Install

#### Linux (Ubuntu/Debian)

```bash
# 1. Download
wget https://github.com/rios/worker/releases/download/v0.1.0/rios-worker-v0.1.0-linux-amd64.tar.gz

# 2. Extract
tar -xzf rios-worker-v0.1.0-linux-amd64.tar.gz

# 3. Install
chmod +x rios-worker-linux-amd64
sudo mv rios-worker-linux-amd64 /usr/local/bin/rios-worker

# 4. Register
rios-worker register --api https://api.rios.com.ai

# 5. Start
rios-worker run
```

#### macOS (Apple Silicon)

```bash
# Download
curl -L -O https://github.com/rios/worker/releases/download/v0.1.0/rios-worker-v0.1.0-darwin-arm64.tar.gz

# Extract
tar -xzf rios-worker-v0.1.0-darwin-arm64.tar.gz

# Remove quarantine
xattr -d com.apple.quarantine rios-worker-darwin-arm64

# Install
chmod +x rios-worker-darwin-arm64
sudo mv rios-worker-darwin-arm64 /usr/local/bin/rios-worker

# Register (testing mode)
rios-worker register --skip-docker --api https://api.rios.com.ai
```

#### Windows

```powershell
# Download
Invoke-WebRequest -Uri "https://github.com/rios/worker/releases/download/v0.1.0/rios-worker-v0.1.0-windows-amd64.zip" -OutFile "rios-worker.zip"

# Extract
Expand-Archive -Path rios-worker.zip -DestinationPath .

# Run
.\rios-worker-windows-amd64.exe register --api https://api.rios.com.ai
```

---

## 🔒 Security & Verification

### Verify Download Integrity

**Linux/macOS:**
```bash
# Check SHA256
sha256sum rios-worker-linux-amd64

# Compare with SHA256SUMS.txt
grep linux-amd64 SHA256SUMS.txt
```

**Windows (PowerShell):**
```powershell
Get-FileHash rios-worker-windows-amd64.exe -Algorithm SHA256
```

### SHA256 Checksums

```
6bb446d42bf39cc790b0d09d8db5fd4d6c19d75363a30ca9a50120aa5771d5f3  rios-worker-darwin-amd64
1cca28f7e42d281769b6892f9eef682ada9efeb576b5b923fab21e238234019c  rios-worker-darwin-arm64
c56b0268e296f65b09499a1de7fd29448fbd5828b9010de51f70ae3e9621d1c5  rios-worker-linux-amd64
ef57e45256710ab076c363c9136080353437b918933e70fd8dc39e5200e7ce4c  rios-worker-linux-arm64
cf0f072a023e1b8355769ddbdf984067a55a2e5f3e6564fa5391b16e3cd5714a  rios-worker-windows-amd64.exe
```

---

## 📊 Binary Information

All binaries are:
- ✅ **Statically Linked** - No external Go dependencies
- ✅ **Stripped** - Optimized for size (-ldflags="-s -w")
- ✅ **Cross-Compiled** - Built for all platforms
- ✅ **Portable** - Single file, no installation needed

### Technical Specifications

```
Language:         Go 1.21+
Build Type:       Static (CGO_ENABLED=0)
Binary Format:
  - Linux:        ELF 64-bit LSB executable
  - macOS:        Mach-O 64-bit executable
  - Windows:      PE32+ executable (console)
Compression:      Not applied (can reduce by ~40% with UPX)
Code Signing:     Not yet (planned for future releases)
```

---

## 🔄 Upgrade from Previous Versions

This is the initial release (v0.1.0). No previous versions exist.

For future updates:

```bash
# Stop worker
sudo systemctl stop rios-worker@$USER

# Download new version
wget https://github.com/rios/worker/releases/latest/download/rios-worker-linux-amd64

# Replace binary
sudo mv rios-worker-linux-amd64 /usr/local/bin/rios-worker

# Restart
sudo systemctl start rios-worker@$USER
```

---

## 🐛 Known Issues

### None

This is a stable release with no known critical issues.

### Limitations

- ⚠️ Requires NVIDIA GPU (no AMD GPU support yet)
- ⚠️ macOS cannot process tasks (no NVIDIA GPUs on modern Macs)
- ⚠️ S3 integration is placeholder (full implementation pending)

---

## 🔮 Roadmap

### v0.2.0 (Planned)

- [ ] AMD GPU support
- [ ] Real S3 file upload/download
- [ ] Task progress reporting
- [ ] Auto-update mechanism
- [ ] Web UI for monitoring

### v0.3.0 (Planned)

- [ ] Multi-GPU support
- [ ] Task prioritization
- [ ] Offline task queuing
- [ ] Bandwidth optimization
- [ ] Code signing for binaries

---

## 📝 Changelog

### v0.1.0 (2025-11-07) - Initial Release

**Added:**
- GPU auto-detection using nvidia-smi
- Docker environment verification
- Worker registration with API
- Automatic job fetching and execution
- Real-time $ROS reward display
- Heartbeat mechanism
- Graceful shutdown
- Configuration management
- ASCII art banner
- Cross-platform support (Linux/macOS/Windows)

**Technical:**
- Go 1.21+ codebase
- Cobra CLI framework
- JWT authentication
- HTTP API client
- Docker container execution
- Systemd service integration

---

## 📞 Support

- **Documentation:** https://docs.rios.com.ai
- **Discord:** https://discord.gg/rios
- **GitHub Issues:** https://github.com/rios/worker/issues
- **Email:** support@rios.com.ai

---

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra)
- Inspired by [Prime Intellect](https://primeintellect.ai)
- Community contributors

---

## ⚖️ License

MIT License - See [LICENSE](../LICENSE) file for details.

---

**Thank you for contributing to the RiOS Network!** 🚀

Start earning $ROS tokens with your GPU today!

