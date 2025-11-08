#!/bin/bash

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  🚀 Pushing RiOS to GitHub"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Check if gh CLI is installed
if command -v gh &> /dev/null; then
    echo "✅ GitHub CLI detected"
    echo ""
    echo "Logging in to GitHub..."
    gh auth login
    echo ""
fi

echo "📤 Pushing code to GitHub..."
git push -u origin main

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Code pushed successfully!"
    echo ""
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "  📦 Creating GitHub Release"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo ""
    
    if command -v gh &> /dev/null; then
        echo "Creating release v0.1.0 with binaries..."
        gh release create v0.1.0 \
          backend/downloads/rios-worker-v0.1.0-linux-amd64.tar.gz \
          backend/downloads/rios-worker-v0.1.0-linux-arm64.tar.gz \
          backend/downloads/rios-worker-v0.1.0-windows-amd64.zip \
          backend/downloads/rios-worker-v0.1.0-all-platforms.tar.gz \
          backend/downloads/SHA256SUMS.txt \
          backend/downloads/README.md \
          backend/downloads/INSTALLATION_GUIDE.txt \
          --title "RiOS Worker CLI v0.1.0 - Initial Release" \
          --notes "# RiOS Worker CLI v0.1.0 - Initial Release

## 📦 Download

Choose your platform:
- [Linux (x86_64)](https://github.com/Finwhale-LLC/RiOS/releases/download/v0.1.0/rios-worker-v0.1.0-linux-amd64.tar.gz) - 3.2 MB ⭐
- [Linux (ARM64)](https://github.com/Finwhale-LLC/RiOS/releases/download/v0.1.0/rios-worker-v0.1.0-linux-arm64.tar.gz) - 2.9 MB
- [Windows](https://github.com/Finwhale-LLC/RiOS/releases/download/v0.1.0/rios-worker-v0.1.0-windows-amd64.zip) - 3.3 MB

## 🚀 Quick Start

\`\`\`bash
# Extract
tar -xzf rios-worker-v0.1.0-linux-amd64.tar.gz
chmod +x rios-worker-linux-amd64
sudo mv rios-worker-linux-amd64 /usr/local/bin/rios-worker

# Register
rios-worker register --api https://api.rios.com.ai

# Start earning
rios-worker run
\`\`\`

## 🔒 Verification

SHA256 checksums: [SHA256SUMS.txt](https://github.com/Finwhale-LLC/RiOS/releases/download/v0.1.0/SHA256SUMS.txt)

## 📚 Documentation

- [Worker Guide](https://github.com/Finwhale-LLC/RiOS/blob/main/rios-worker/README.md)
- [Installation](https://github.com/Finwhale-LLC/RiOS/blob/main/rios-worker/INSTALL.md)
- [Deployment](https://github.com/Finwhale-LLC/RiOS/blob/main/rios-worker/DEPLOYMENT.md)

## 🎯 Features

✅ GPU auto-detection
✅ Docker integration
✅ Automatic task fetching
✅ Real-time rewards
✅ Graceful shutdown
✅ Systemd service support

## 📋 Requirements

- NVIDIA GPU (RTX 3060+)
- NVIDIA Drivers (525+)
- Docker + nvidia-docker2
- RAM: 8 GB+
- Storage: 50 GB+"
        
        echo ""
        echo "✅ Release created successfully!"
        echo ""
        echo "🌐 View at: https://github.com/Finwhale-LLC/RiOS/releases/tag/v0.1.0"
    else
        echo "⚠️  GitHub CLI not found. Please create release manually:"
        echo "   https://github.com/Finwhale-LLC/RiOS/releases/new"
        echo ""
        echo "Upload these files:"
        echo "   • backend/downloads/rios-worker-v0.1.0-linux-amd64.tar.gz"
        echo "   • backend/downloads/rios-worker-v0.1.0-linux-arm64.tar.gz"
        echo "   • backend/downloads/rios-worker-v0.1.0-windows-amd64.zip"
        echo "   • backend/downloads/SHA256SUMS.txt"
    fi
    
    echo ""
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "  ✅ Upload Complete!"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo ""
    echo "🌐 Repository: https://github.com/Finwhale-LLC/RiOS"
    echo "📦 Releases:   https://github.com/Finwhale-LLC/RiOS/releases"
    echo ""
    echo "Users can now download Worker CLI from GitHub Releases!"
    echo ""
else
    echo ""
    echo "❌ Push failed. Please check your GitHub credentials."
    echo ""
    echo "💡 Try one of these methods:"
    echo "   1. Use GitHub CLI: gh auth login"
    echo "   2. Use Personal Access Token"
    echo "   3. Use SSH: git remote set-url origin git@github.com:Finwhale-LLC/RiOS.git"
    echo ""
    echo "See GITHUB_UPLOAD_GUIDE.md for detailed instructions."
fi

