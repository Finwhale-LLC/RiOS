# Quick Start

Get up and running with RiOS in just a few minutes!

## Step 1: Create an Account

1. Visit [https://cloud.rios.com.ai](https://cloud.rios.com.ai)
2. Click **Sign Up** in the top right corner
3. Fill in your email and create a strong password
4. Verify your email address
5. Complete your profile

## Step 2: Get ROS Tokens

You need ROS tokens to pay for computing resources.

### Option A: Purchase Tokens
1. Navigate to **Wallet** in your dashboard
2. Click **Buy ROS Tokens**
3. Choose your payment method (crypto or fiat)
4. Complete the transaction

### Option B: Earn Tokens
1. Set up a worker node (see [Worker Setup](../worker-setup/README.md))
2. Contribute computing resources
3. Earn ROS tokens as rewards

## Step 3: Deploy Your First Application

### Using the Web Interface

1. **Go to Dashboard**
   - Log into [https://cloud.rios.com.ai](https://cloud.rios.com.ai)
   - Navigate to **Deployments** → **New Deployment**

2. **Configure Your Application**
   ```
   Name: my-first-app
   Image: your-docker-image:tag
   Resources:
     - CPU: 2 cores
     - RAM: 4 GB
     - GPU: Optional
   ```

3. **Set Environment Variables** (if needed)
   ```
   API_KEY=your_api_key
   DATABASE_URL=your_db_url
   ```

4. **Review and Deploy**
   - Review pricing (estimated ROS/hour)
   - Click **Deploy**
   - Wait for deployment to complete (~1-2 minutes)

5. **Access Your Application**
   - Copy the provided URL or IP address
   - Your application is now running!

### Using the API

```bash
# Set your API token
export RIOS_API_TOKEN="your_api_token_here"

# Create a deployment
curl -X POST https://api.rios.com.ai/v1/deployments \
  -H "Authorization: Bearer $RIOS_API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-first-app",
    "image": "nginx:latest",
    "resources": {
      "cpu": 2,
      "memory": 4096
    }
  }'
```

## Step 4: Monitor Your Application

### View Logs
```bash
# Using the web interface
Dashboard → Deployments → Select your app → Logs

# Using the CLI
rios logs my-first-app
```

### Check Resource Usage
```bash
# View metrics in dashboard
Dashboard → Deployments → Select your app → Metrics

# Using the API
curl https://api.rios.com.ai/v1/deployments/my-first-app/metrics \
  -H "Authorization: Bearer $RIOS_API_TOKEN"
```

### Monitor Costs
- Go to **Billing** in your dashboard
- View real-time spending
- Set budget alerts

## Step 5: Scale Your Application

### Scale Up
```bash
# Increase resources
rios scale my-first-app --cpu 4 --memory 8192
```

### Scale Out
```bash
# Increase replicas
rios scale my-first-app --replicas 3
```

## Example: Deploy a Simple Web Server

Here's a complete example deploying a Node.js application:

```bash
# 1. Create your application (app.js)
cat > app.js << 'EOF'
const http = require('http');
const server = http.createServer((req, res) => {
  res.writeHead(200, {'Content-Type': 'text/plain'});
  res.end('Hello from RiOS!\n');
});
server.listen(3000, () => {
  console.log('Server running on port 3000');
});
EOF

# 2. Create Dockerfile
cat > Dockerfile << 'EOF'
FROM node:18-alpine
WORKDIR /app
COPY app.js .
EXPOSE 3000
CMD ["node", "app.js"]
EOF

# 3. Build and push to registry
docker build -t your-registry/hello-rios:v1 .
docker push your-registry/hello-rios:v1

# 4. Deploy to RiOS
rios deploy \
  --name hello-rios \
  --image your-registry/hello-rios:v1 \
  --port 3000 \
  --cpu 1 \
  --memory 512
```

## Common Commands

```bash
# List all deployments
rios list

# Get deployment details
rios describe my-first-app

# Update deployment
rios update my-first-app --image new-image:v2

# Stop deployment
rios stop my-first-app

# Delete deployment
rios delete my-first-app

# Check balance
rios balance

# View billing history
rios billing history
```

## Troubleshooting

### Deployment Failed
- Check your Docker image is accessible
- Verify sufficient ROS token balance
- Review deployment logs for errors

### Application Not Responding
- Check if deployment is running: `rios status my-first-app`
- View logs: `rios logs my-first-app`
- Verify network settings and ports

### Out of Tokens
- Purchase more tokens from the dashboard
- Set up a worker node to earn tokens

## Next Steps

Now that you have your first application running:

1. **Explore Advanced Features**
   - [GPU Computing](../user-guide/gpu-computing.md)
   - [Auto-scaling](../user-guide/auto-scaling.md)
   - [Custom Domains](../user-guide/custom-domains.md)

2. **Optimize Your Deployment**
   - [Performance Tuning](../user-guide/performance-tuning.md)
   - [Cost Optimization](../user-guide/cost-optimization.md)

3. **Learn the Architecture**
   - [How DMoE Works](../architecture/dmoe-engine.md)
   - [Security Model](../architecture/security-model.md)

## Getting Help

- **Documentation**: Browse the full docs
- **API Reference**: [API Documentation](../api-reference/README.md)
- **Community**: [Forum](https://community.rios.com.ai)
- **Support**: support@rios.com.ai

