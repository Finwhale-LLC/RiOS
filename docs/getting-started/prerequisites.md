# Prerequisites

Before you start using RiOS, ensure your system meets the following requirements.

## For Users (AI Developers)

### System Requirements
- **Operating System**: Any modern OS with web browser support
- **Browser**: Chrome 90+, Firefox 88+, Safari 14+, or Edge 90+
- **Internet Connection**: Stable broadband connection (minimum 10 Mbps)

### Account Requirements
- Valid email address for registration
- ROS tokens for paying for computing resources (can be purchased or earned)

### Development Tools (Optional)
- **Docker**: For local testing of containerized applications
- **Git**: For version control
- **Code Editor**: VSCode, Sublime Text, or your preferred editor

## For Workers (Computing Providers)

### Minimum Hardware Requirements
- **CPU**: 4+ cores (8+ recommended)
- **RAM**: 8GB (16GB+ recommended)
- **Storage**: 100GB free space (SSD recommended)
- **GPU** (optional): NVIDIA GPU with CUDA support for AI workloads
- **Network**: Stable connection with public IP or port forwarding capability

### Operating System
- **Linux**: Ubuntu 20.04 LTS or later (recommended)
- **macOS**: 10.15 (Catalina) or later
- **Windows**: Windows 10/11 with WSL2

### Software Dependencies
- **Docker**: Version 20.10 or later
- **Docker Compose**: Version 1.29 or later
- **Go**: Version 1.19 or later (for building from source)

### Network Requirements
- Open ports for P2P communication (configurable, default: 9000-9010)
- Minimum upload speed: 10 Mbps
- Minimum download speed: 10 Mbps
- Low latency connection preferred

## Security Requirements

### SSL/TLS
- Valid SSL certificate for production deployments
- Self-signed certificates acceptable for development

### Firewall
- Ability to configure firewall rules
- Understanding of network security basics

### Access Control
- SSH key-based authentication for remote servers
- Strong password policy
- Two-factor authentication recommended

## Knowledge Prerequisites

### For Developers
- Basic understanding of:
  - RESTful APIs
  - Container technology (Docker)
  - Command-line interface
  - Git version control

### For Workers
- System administration basics
- Docker fundamentals
- Network configuration
- Basic security practices

## Optional Tools

### Monitoring
- **Prometheus**: For metrics collection
- **Grafana**: For visualization
- **ELK Stack**: For log analysis

### Development
- **Postman**: For API testing
- **kubectl**: For Kubernetes management (if using K8s deployment)

## Getting Help

If you're unsure about any requirements:
- Check our [FAQ](../faq.md)
- Visit the [Community Forum](https://community.rios.com.ai)
- Contact support at support@rios.com.ai

## Next Steps

Once you've verified all prerequisites, proceed to:
- **For Users**: [Quick Start Guide](quick-start.md)
- **For Workers**: [Worker Installation](../worker-setup/installation.md)

