#!/bin/bash

#################################################################
# RiOS Worker - Systemd Service Setup
# 
# This script sets up RiOS Worker as a systemd service
# that starts automatically on boot.
#
# Usage:
#   sudo ./setup-service.sh
#
#################################################################

set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  🔧 RiOS Worker - Systemd Service Setup"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Check if running as root
if [ "$EUID" -ne 0 ]; then 
    print_error "This script must be run as root (use sudo)"
    exit 1
fi

# Check if rios-worker is installed
if ! command -v rios-worker &> /dev/null; then
    print_error "rios-worker is not installed"
    echo "Please run the installation script first:"
    echo "  curl -fsSL https://install.rios.com.ai/worker.sh | bash"
    exit 1
fi

# Get the user who will run the service
if [ -n "$SUDO_USER" ]; then
    WORKER_USER="$SUDO_USER"
else
    echo "Enter the username that will run the worker:"
    read -r WORKER_USER
fi

# Validate user exists
if ! id "$WORKER_USER" &>/dev/null; then
    print_error "User $WORKER_USER does not exist"
    exit 1
fi

# Check if user has registered
CONFIG_FILE="/home/$WORKER_USER/.rios/config.json"
if [ ! -f "$CONFIG_FILE" ]; then
    print_warning "Worker not registered yet"
    echo ""
    echo "Please register first as user $WORKER_USER:"
    echo "  su - $WORKER_USER"
    echo "  rios-worker register --api https://api.rios.com.ai"
    exit 1
fi

print_success "Worker is registered (user: $WORKER_USER)"

# Copy service file
print_warning "Installing systemd service..."

SERVICE_FILE="/etc/systemd/system/rios-worker@.service"
cat > "$SERVICE_FILE" << 'EOF'
[Unit]
Description=RiOS Worker - Decentralized GPU Compute Node
Documentation=https://docs.rios.com.ai
After=network-online.target docker.service
Wants=network-online.target
Requires=docker.service

[Service]
Type=simple
User=%i
Group=docker

# Environment
Environment="PATH=/usr/local/bin:/usr/bin:/bin"

# Worker configuration is loaded from ~/.rios/config.json
ExecStart=/usr/local/bin/rios-worker run

# Restart policy
Restart=always
RestartSec=10

# Resource limits
LimitNOFILE=65536
LimitNPROC=4096

# Logging
StandardOutput=journal
StandardError=journal
SyslogIdentifier=rios-worker

# Security
NoNewPrivileges=true
PrivateTmp=true

[Install]
WantedBy=multi-user.target
EOF

print_success "Service file installed: $SERVICE_FILE"

# Reload systemd
systemctl daemon-reload
print_success "Systemd daemon reloaded"

# Enable service
systemctl enable "rios-worker@$WORKER_USER"
print_success "Service enabled (will start on boot)"

# Start service
systemctl start "rios-worker@$WORKER_USER"
print_success "Service started"

# Show status
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  🎉 Service Setup Complete!"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
print_success "RiOS Worker is now running as a system service"
echo ""
echo "📋 Useful Commands:"
echo ""
echo "  Check status:"
echo "    sudo systemctl status rios-worker@$WORKER_USER"
echo ""
echo "  View logs:"
echo "    sudo journalctl -u rios-worker@$WORKER_USER -f"
echo ""
echo "  Stop service:"
echo "    sudo systemctl stop rios-worker@$WORKER_USER"
echo ""
echo "  Restart service:"
echo "    sudo systemctl restart rios-worker@$WORKER_USER"
echo ""
echo "  Disable service:"
echo "    sudo systemctl disable rios-worker@$WORKER_USER"
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Show current status
systemctl status "rios-worker@$WORKER_USER" --no-pager

