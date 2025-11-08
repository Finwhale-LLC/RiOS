# RiOS Worker - Deployment Guide

Complete deployment guide for different scenarios.

## 📦 Available Files

| File | Purpose | Size |
|------|---------|------|
| `install.sh` | One-line automated installer | ~15 KB |
| `setup-service.sh` | Systemd service setup | ~3 KB |
| `rios-worker.service` | Systemd service file | ~1 KB |
| `INSTALL.md` | Detailed installation guide | ~20 KB |
| `TESTING_INSTALL.md` | Installation testing guide | ~8 KB |
| `README.md` | Main documentation | ~10 KB |

---

## 🚀 Deployment Scenarios

### Scenario 1: Single GPU Server (Recommended)

**Use Case:** Individual or small-scale mining with 1 GPU

**Steps:**

```bash
# 1. One-line install
curl -fsSL https://install.rios.com.ai/worker.sh | bash

# 2. Register
rios-worker register --api https://api.rios.com.ai

# 3. Setup as service (optional but recommended)
curl -fsSL https://install.rios.com.ai/setup-service.sh | sudo bash

# 4. Verify
sudo systemctl status rios-worker@$USER
```

**Advantages:**
- ✅ Automatic restart on failure
- ✅ Starts on boot
- ✅ Easy monitoring via journalctl
- ✅ Resource limits configured

---

### Scenario 2: Multi-GPU Server

**Use Case:** Server with multiple GPUs

**Option A: Multiple Workers (Recommended)**

Run one worker instance per GPU:

```bash
# Register worker for GPU 0
CUDA_VISIBLE_DEVICES=0 rios-worker register --api https://api.rios.com.ai

# Copy config for each GPU
cp ~/.rios/config.json ~/.rios/config-gpu0.json
cp ~/.rios/config.json ~/.rios/config-gpu1.json
# ... etc

# Create service for each GPU
# Edit service file to set CUDA_VISIBLE_DEVICES
sudo systemctl enable rios-worker-gpu0@$USER
sudo systemctl enable rios-worker-gpu1@$USER

# Start all
sudo systemctl start rios-worker-gpu0@$USER
sudo systemctl start rios-worker-gpu1@$USER
```

**Option B: Single Worker (Simpler)**

Let Docker manage all GPUs:

```bash
# Register once
rios-worker register --api https://api.rios.com.ai

# Setup service
sudo ./setup-service.sh

# Worker will use all GPUs via --gpus all
```

---

### Scenario 3: Data Center / Mining Farm

**Use Case:** Large-scale deployment with many servers

**Tools:**
- Ansible/Terraform for automation
- Kubernetes for orchestration
- Monitoring with Prometheus/Grafana

**Ansible Playbook Example:**

```yaml
---
- name: Deploy RiOS Workers
  hosts: gpu_servers
  become: yes
  tasks:
    - name: Install Worker
      shell: curl -fsSL https://install.rios.com.ai/worker.sh | bash

    - name: Register Worker
      shell: |
        su - {{ worker_user }} -c "rios-worker register --api https://api.rios.com.ai"
      environment:
        WALLET_ADDRESS: "{{ wallet_address }}"

    - name: Setup Service
      shell: ./setup-service.sh
      args:
        chdir: /home/{{ worker_user }}

    - name: Verify Service
      systemd:
        name: "rios-worker@{{ worker_user }}"
        state: started
        enabled: yes
```

**Kubernetes Deployment Example:**

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: rios-worker
spec:
  selector:
    matchLabels:
      app: rios-worker
  template:
    metadata:
      labels:
        app: rios-worker
    spec:
      hostNetwork: true
      containers:
      - name: worker
        image: rios/worker:latest
        env:
        - name: API_ENDPOINT
          value: "https://api.rios.com.ai"
        - name: NODE_AUTH_TOKEN
          valueFrom:
            secretKeyRef:
              name: worker-secrets
              key: auth-token
        resources:
          limits:
            nvidia.com/gpu: 1
        volumeMounts:
        - name: docker-sock
          mountPath: /var/run/docker.sock
      volumes:
      - name: docker-sock
        hostPath:
          path: /var/run/docker.sock
```

---

### Scenario 4: Cloud Deployment

**Use Case:** Deploy on cloud providers (AWS, GCP, Azure)

#### AWS EC2 with GPU

```bash
# 1. Launch EC2 instance (p3.2xlarge or g4dn.xlarge)
# 2. SSH into instance
# 3. Run install script
curl -fsSL https://install.rios.com.ai/worker.sh | bash

# 4. Register and start
rios-worker register --api https://api.rios.com.ai
sudo ./setup-service.sh
```

**Auto-Scaling Group UserData:**

```bash
#!/bin/bash
# UserData script for AWS EC2 Auto Scaling

# Install worker
curl -fsSL https://install.rios.com.ai/worker.sh | bash

# Register with stored credentials
WALLET_ADDRESS=$(aws ssm get-parameter --name /rios/wallet-address --query Parameter.Value --output text)
echo "$WALLET_ADDRESS" | rios-worker register --api https://api.rios.com.ai

# Setup service
./setup-service.sh

# Report success to CloudWatch
aws cloudwatch put-metric-data --metric-name WorkerDeployment --value 1 --namespace RiOS
```

#### GCP Compute Engine

```bash
# Create instance with GPU
gcloud compute instances create rios-worker-1 \
  --zone=us-central1-a \
  --machine-type=n1-standard-4 \
  --accelerator=type=nvidia-tesla-t4,count=1 \
  --maintenance-policy=TERMINATE \
  --metadata=startup-script='#!/bin/bash
    curl -fsSL https://install.rios.com.ai/worker.sh | bash
    rios-worker register --api https://api.rios.com.ai
    ./setup-service.sh
  '
```

---

### Scenario 5: Docker-Only Deployment

**Use Case:** Prefer containerized deployment

**Single Container:**

```bash
docker run -d \
  --name rios-worker \
  --restart unless-stopped \
  --gpus all \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v ~/.rios:/root/.rios \
  -e API_ENDPOINT=https://api.rios.com.ai \
  rios/worker:latest
```

**Docker Compose:**

```yaml
version: '3.8'
services:
  rios-worker:
    image: rios/worker:latest
    container_name: rios-worker
    restart: unless-stopped
    runtime: nvidia
    environment:
      - API_ENDPOINT=https://api.rios.com.ai
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - ~/.rios:/root/.rios
      - ./work:/work
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: all
              capabilities: [gpu]
```

---

## 📊 Monitoring & Management

### Logs

```bash
# Systemd service logs
sudo journalctl -u rios-worker@$USER -f

# Docker logs
docker logs -f rios-worker

# Worker-specific logs
tail -f ~/.rios/worker.log
```

### Metrics

```bash
# GPU usage
watch -n 1 nvidia-smi

# System resources
htop

# Network traffic
iftop
```

### Alerts

Setup alerts for:
- Worker offline > 5 minutes
- GPU temperature > 80°C
- GPU utilization < 10% (idle)
- Disk space < 10 GB
- Memory usage > 90%

---

## 🔄 Update Strategy

### Manual Update

```bash
# Stop worker
sudo systemctl stop rios-worker@$USER

# Update binary
curl -fsSL https://install.rios.com.ai/worker.sh | bash

# Start worker
sudo systemctl start rios-worker@$USER
```

### Automated Update

Create update script:

```bash
#!/bin/bash
# auto-update.sh

# Stop worker
systemctl stop rios-worker@$USER

# Backup current version
cp /usr/local/bin/rios-worker /usr/local/bin/rios-worker.backup

# Download new version
wget https://github.com/rios/worker/releases/latest/download/rios-worker-linux-amd64 -O /tmp/rios-worker
chmod +x /tmp/rios-worker

# Test new version
if /tmp/rios-worker --version; then
    mv /tmp/rios-worker /usr/local/bin/rios-worker
    systemctl start rios-worker@$USER
    echo "Update successful"
else
    echo "Update failed, reverting"
    mv /usr/local/bin/rios-worker.backup /usr/local/bin/rios-worker
    systemctl start rios-worker@$USER
fi
```

Add to cron:

```bash
# Check for updates daily at 3 AM
0 3 * * * /home/user/auto-update.sh
```

---

## 🛡️ Security Best Practices

### 1. Network Security

```bash
# Allow only necessary ports
sudo ufw allow 22/tcp  # SSH
sudo ufw allow 443/tcp # HTTPS
sudo ufw enable
```

### 2. User Isolation

```bash
# Create dedicated user
sudo useradd -m -s /bin/bash riosworker
sudo usermod -aG docker riosworker

# Run worker as this user
sudo ./setup-service.sh
# Select: riosworker
```

### 3. Resource Limits

Edit `/etc/systemd/system/rios-worker@.service`:

```ini
[Service]
# CPU limit (50%)
CPUQuota=50%

# Memory limit (8GB)
MemoryLimit=8G

# Max file descriptors
LimitNOFILE=65536
```

### 4. Monitoring

```bash
# Install fail2ban
sudo apt-get install fail2ban

# Monitor for suspicious activity
sudo tail -f /var/log/auth.log
```

---

## 📈 Performance Optimization

### 1. GPU Settings

```bash
# Set persistent mode (NVIDIA)
sudo nvidia-smi -pm 1

# Set power limit (adjust as needed)
sudo nvidia-smi -pl 250

# Set compute mode
sudo nvidia-smi -c EXCLUSIVE_PROCESS
```

### 2. Docker Optimization

Edit `/etc/docker/daemon.json`:

```json
{
  "storage-driver": "overlay2",
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "10m",
    "max-file": "3"
  },
  "default-runtime": "nvidia"
}
```

### 3. System Tuning

```bash
# Increase file watchers
echo "fs.inotify.max_user_watches=524288" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p

# TCP tuning
sudo sysctl -w net.core.rmem_max=26214400
sudo sysctl -w net.core.wmem_max=26214400
```

---

## 🆘 Troubleshooting

See [INSTALL.md#troubleshooting](INSTALL.md#troubleshooting) for common issues and solutions.

---

## 📞 Support

- **Documentation:** https://docs.rios.com.ai
- **Community:** https://discord.gg/rios
- **Issues:** https://github.com/rios/worker/issues

---

**Happy Deployment! 🚀**

