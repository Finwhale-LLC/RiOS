# Frequently Asked Questions (FAQ)

## General Questions

### What is RiOS?

RiOS (Intelligent Operating System) is a decentralized computing platform that enables:
- AI developers to deploy and run applications on distributed infrastructure
- Computing providers to contribute resources and earn rewards
- Organizations to access scalable, secure computing power

### How is RiOS different from traditional cloud providers?

| Feature | RiOS | Traditional Cloud |
|---------|------|-------------------|
| **Architecture** | Decentralized, P2P | Centralized datacenters |
| **Pricing** | Market-driven, token-based | Fixed pricing tiers |
| **Resources** | Global, community-provided | Limited to provider regions |
| **AI Optimization** | DMoE for distributed AI | General-purpose compute |
| **Entry Barrier** | Low (anyone can join) | High (enterprise focused) |

### What is DMoE?

Decentralized Mixture of Experts (DMoE) is RiOS's core technology that decomposes large AI models into specialized expert modules, enabling efficient distributed computation with significantly reduced bandwidth requirements.

[Learn more about DMoE](architecture/dmoe-engine.md)

### Is RiOS open source?

Yes, core components of RiOS are open source:
- Worker node software
- CLI tools
- API client libraries
- Documentation

Check our [GitHub repositories](https://github.com/rios) for more information.

## Getting Started

### How do I create an account?

1. Visit [https://cloud.rios.com.ai](https://cloud.rios.com.ai)
2. Click **Sign Up**
3. Enter your email and create a password
4. Verify your email
5. Complete your profile

[Detailed guide](getting-started/quick-start.md)

### Do I need ROS tokens to use RiOS?

Yes, ROS tokens are required to pay for computing resources. You can:
- **Purchase tokens** using crypto or fiat currency
- **Earn tokens** by running a worker node and contributing compute resources

### How much does it cost?

Pricing is dynamic and market-driven. Typical rates:
- **CPU**: 0.01-0.05 ROS per core per hour
- **RAM**: 0.005 ROS per GB per hour
- **GPU**: 0.5-2.0 ROS per GPU per hour (depending on model)

View current rates in your [dashboard](https://cloud.rios.com.ai/pricing).

### Can I try RiOS for free?

Yes! New users receive:
- **Free trial credits**: 10 ROS tokens
- **Free tier**: Limited CPU/RAM for testing
- **Sandbox environment**: For development and testing

## Using RiOS

### What applications can I deploy?

Any containerized application:
- AI/ML models (PyTorch, TensorFlow, etc.)
- Web applications
- APIs and microservices
- Data processing pipelines
- Batch jobs

### What programming languages are supported?

All languages that can run in Docker containers:
- Python (most common for AI)
- JavaScript/Node.js
- Go
- Java
- R
- C++
- And more...

### How do I deploy my application?

Three ways:

**1. Web Dashboard**
- Navigate to Deployments → New
- Configure your app
- Click Deploy

**2. CLI**
```bash
rios deploy --name my-app --image my-image:tag
```

**3. API**
```bash
curl -X POST https://api.rios.com.ai/v1/deployments \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"my-app","image":"my-image:tag"}'
```

[Quick Start Guide](getting-started/quick-start.md)

### How do I monitor my applications?

- **Dashboard**: Real-time metrics and logs
- **CLI**: `rios logs <app>` and `rios status <app>`
- **API**: Programmatic access to metrics
- **Webhooks**: Get notified of events

[Monitoring Guide](user-guide/monitoring-resources.md)

### Can I use custom domains?

Yes! You can:
- Configure custom domains
- Use SSL/TLS certificates
- Set up DNS records

[Custom Domains Guide](user-guide/custom-domains.md)

## Running a Worker Node

### Why should I run a worker node?

Benefits:
- **Earn ROS tokens** by contributing compute resources
- **Support the network** and help decentralize AI
- **Utilize idle resources** productively
- **Join the community** of computing providers

### What are the requirements?

**Minimum**:
- 4 CPU cores
- 8GB RAM
- 100GB storage
- Stable internet (10 Mbps+)

**Recommended**:
- 8+ CPU cores
- 16GB+ RAM
- 500GB+ SSD storage
- GPU (for AI workloads)
- 50 Mbps+ connection

[Full requirements](getting-started/prerequisites.md)

### How much can I earn?

Earnings depend on:
- **Resources provided**: More/better resources = higher earnings
- **Uptime**: Higher uptime = more reliability bonus
- **Demand**: Market demand affects rates
- **Location**: Certain regions may have higher demand

**Example**:
- 8-core CPU + 16GB RAM + RTX 3080 GPU
- 90% uptime
- Typical earning: 50-200 ROS/month (~$50-$200 equivalent)

### Is it safe to run a worker?

Yes! Security features:
- **Container isolation**: Workloads run in sandboxed containers
- **Zero trust model**: No implicit trust between nodes
- **Resource limits**: CPU, memory, and network caps
- **Monitoring**: Real-time security monitoring
- **Encrypted communication**: All data encrypted in transit

[Security Model](architecture/security-model.md)

### How do I set up a worker?

```bash
# Install worker software
curl -sSL https://get.rios.com.ai/worker | bash

# Configure
rios-worker init

# Start
rios-worker start
```

[Detailed Setup Guide](worker-setup/installation.md)

## Tokens and Billing

### What are ROS tokens?

ROS is the native utility token of the RiOS network:
- **Payment**: Pay for computing resources
- **Rewards**: Earn by providing resources
- **Governance**: Vote on network proposals (future)

### How do I get ROS tokens?

**Purchase**:
- Dashboard → Wallet → Buy ROS
- Supported: Crypto (ETH, BTC, USDT) or Credit Card

**Earn**:
- Run a worker node
- Participate in testnet
- Bug bounties and contributions

### How is billing calculated?

Billing is based on:
- **Resource usage**: CPU, RAM, GPU, storage
- **Time**: Per-second billing
- **Network**: Data transfer costs
- **Location**: Geographic pricing variations

**Formula**:
```
Cost = (CPU_hours × CPU_rate) + (RAM_GB_hours × RAM_rate) 
     + (GPU_hours × GPU_rate) + (Storage_GB_hours × Storage_rate)
     + (Network_GB × Network_rate)
```

[Billing Details](economic-model/pricing.md)

### Can I set a budget limit?

Yes! You can:
- Set daily/monthly spending limits
- Get alerts at thresholds (50%, 75%, 90%)
- Auto-stop deployments at budget limit
- Receive usage reports

Configure in Dashboard → Billing → Budget Settings.

### What happens if I run out of tokens?

1. You receive low balance warnings at 20%, 10%, 5%
2. At 0 balance:
   - Running deployments continue for 1 hour (grace period)
   - New deployments are blocked
   - You receive urgent notification
3. After grace period:
   - Deployments are paused (not deleted)
   - Data is preserved for 7 days
   - Resume by adding tokens

## Technical Questions

### What containerization technology does RiOS use?

- **Primary**: Docker
- **Security**: gVisor, Kata Containers for enhanced isolation
- **Orchestration**: Custom scheduler optimized for DMoE

### How does RiOS handle failures?

**Redundancy**:
- Multi-node deployment options
- Automatic failover
- Health checks

**Recovery**:
- Auto-restart on failure
- State preservation
- Backup worker selection

[High Availability Guide](user-guide/high-availability.md)

### What about data privacy?

- **Encryption**: All data encrypted in transit (TLS 1.3) and at rest (AES-256)
- **Isolation**: Container sandboxing prevents data leakage
- **Compliance**: GDPR, CCPA compliant
- **Private deployments**: Option for dedicated workers

[Security & Privacy](architecture/security-model.md)

### Can I use RiOS in production?

Yes! RiOS is production-ready with:
- 99.9% uptime SLA (for dedicated deployments)
- 24/7 monitoring
- Support options
- Regular security audits

## Troubleshooting

### My deployment failed. What should I check?

1. **Image accessibility**: Can RiOS pull your Docker image?
2. **Resource limits**: Are you requesting more than available?
3. **Balance**: Do you have sufficient ROS tokens?
4. **Configuration**: Check environment variables and ports
5. **Logs**: Review deployment logs for errors

[Troubleshooting Guide](user-guide/troubleshooting.md)

### My application is slow. How can I improve performance?

- **Scale up**: Increase CPU/RAM allocation
- **Scale out**: Add more replicas
- **Use GPU**: For AI workloads
- **Optimize code**: Profile and optimize your application
- **Choose closer regions**: Reduce latency
- **Enable caching**: Use Redis or similar

[Performance Tuning](user-guide/performance-tuning.md)

### I can't connect to my deployment. Help!

Check:
1. **Status**: Is deployment running? (`rios status <app>`)
2. **Ports**: Are ports correctly configured?
3. **Firewall**: Check your local firewall settings
4. **URL**: Using correct endpoint URL?
5. **Logs**: Check logs for connection errors

### How do I get support?

- **Documentation**: Comprehensive guides available
- **Community Forum**: [community.rios.com.ai](https://community.rios.com.ai)
- **Discord**: Join our Discord server
- **Email Support**: support@rios.com.ai
- **Paid Support**: Enterprise support plans available

## Roadmap and Future

### What's planned for RiOS?

**Short-term (Q1-Q2 2025)**:
- Enhanced auto-scaling
- More GPU types support
- Additional regions
- Mobile app

**Long-term (2025-2026)**:
- Federated learning support
- Cross-chain integration
- Governance DAO
- Marketplace for AI models

[Full Roadmap](roadmap.md)

### How can I contribute?

Ways to contribute:
- **Code**: Submit PRs on GitHub
- **Documentation**: Improve docs
- **Testing**: Join beta programs
- **Community**: Help other users
- **Resources**: Run worker nodes
- **Feedback**: Share your experience

[Contributing Guide](contributing.md)

## Still Have Questions?

- **Community Forum**: [community.rios.com.ai](https://community.rios.com.ai)
- **Discord**: [discord.gg/rios](https://discord.gg/rios)
- **Email**: support@rios.com.ai
- **Twitter**: [@RiOSOfficial](https://twitter.com/RiOSOfficial)

