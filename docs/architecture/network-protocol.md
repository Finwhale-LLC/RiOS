# Network Protocol

RiOS uses multiple network protocols optimized for different use cases, with a focus on efficient P2P communication and low-latency distributed computing.

## Protocol Stack Overview

```
Application Layer:  HTTP/HTTPS, WebSocket, gRPC
Transport Layer:    TCP, QUIC
Network Layer:      IPv4, IPv6
P2P Layer:          libp2p
```

## HTTP/HTTPS API

### REST API

**Base URL:** `https://api.rios.com.ai/v1`

**Request Format:**
```http
POST /deployments HTTP/1.1
Host: api.rios.com.ai
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "my-app",
  "image": "nginx:latest",
  "resources": {
    "cpu": 2,
    "memory": 4096
  }
}
```

**Response Format:**
```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "id": "deploy_abc123",
  "status": "pending",
  "created_at": "2024-01-15T10:30:00Z"
}
```

### WebSocket Protocol

For real-time updates:

```javascript
const ws = new WebSocket('wss://api.rios.com.ai/v1/stream');

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log('Deployment status:', data);
};
```

**Message Format:**
```json
{
  "type": "deployment.update",
  "deployment_id": "deploy_abc123",
  "status": "running",
  "timestamp": "2024-01-15T10:31:00Z"
}
```

## gRPC Protocol

For high-performance inter-service communication:

**Service Definition (Protocol Buffers):**
```protobuf
syntax = "proto3";

service DeploymentService {
  rpc CreateDeployment(DeploymentRequest) returns (DeploymentResponse);
  rpc GetDeployment(GetRequest) returns (DeploymentResponse);
  rpc StreamLogs(StreamRequest) returns (stream LogMessage);
}

message DeploymentRequest {
  string name = 1;
  string image = 2;
  ResourceSpec resources = 3;
}

message DeploymentResponse {
  string id = 1;
  string status = 2;
  int64 created_at = 3;
}
```

## libp2p Protocol

For decentralized worker communication:

### Overview

libp2p provides:
- **Peer Discovery**: Find other workers
- **Connection Management**: Maintain connections
- **Stream Multiplexing**: Multiple streams over one connection
- **Security**: Encrypted channels
- **NAT Traversal**: Connect through firewalls

### Peer Identity

Each worker has a unique peer ID:
```
Peer ID: QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG
```

Generated from public key cryptography (Ed25519).

### Connection Establishment

```
Worker A                    Worker B
    │                           │
    ├─── Connect Request ──────>│
    │                           │
    │<──── Challenge ───────────┤
    │                           │
    ├─── Auth Response ────────>│
    │                           │
    │<──── Connection Ready ────┤
    │                           │
    ├═══ Encrypted Channel ════>│
```

### Protocol Handlers

**Discovery Protocol:**
```go
// Announce availability
dht.Provide(ctx, contentID)

// Find providers
peers := dht.FindProviders(ctx, contentID)
```

**Task Assignment Protocol:**
```go
protocol := "/rios/task/1.0.0"

// Register handler
host.SetStreamHandler(protocol, func(s network.Stream) {
    // Read task
    task := readTask(s)
    
    // Execute
    result := executeTask(task)
    
    // Send result
    writeResult(s, result)
})
```

## DMoE Communication Protocol

### Task Distribution

```
Orchestrator → Task Decomposition
      │
      ├─→ Expert 1: Task Part A
      ├─→ Expert 2: Task Part B
      └─→ Expert 3: Task Part C
            │
            └─→ Execute in parallel
```

**Message Format:**
```json
{
  "task_id": "task_xyz789",
  "expert_id": "expert_001",
  "input_tensor": {
    "shape": [1, 512],
    "dtype": "float32",
    "data": "<base64_encoded>"
  },
  "config": {
    "batch_size": 1,
    "temperature": 0.7
  }
}
```

### Result Aggregation

```json
{
  "task_id": "task_xyz789",
  "expert_id": "expert_001",
  "output_tensor": {
    "shape": [1, 512],
    "dtype": "float32",
    "data": "<base64_encoded>"
  },
  "metrics": {
    "execution_time_ms": 45,
    "memory_used_mb": 1024
  }
}
```

## Network Optimization

### Compression

**HTTP:**
- gzip compression for text responses
- Brotli compression (when supported)

**Binary Data:**
- Protocol Buffer serialization
- Tensor data compression (when applicable)

### Connection Pooling

```go
httpClient := &http.Client{
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}
```

### Request Batching

Combine multiple operations:
```json
{
  "batch": [
    {"op": "get", "path": "/deployments/1"},
    {"op": "get", "path": "/deployments/2"},
    {"op": "get", "path": "/deployments/3"}
  ]
}
```

## Error Handling

### HTTP Status Codes

- `200 OK`: Success
- `201 Created`: Resource created
- `400 Bad Request`: Invalid input
- `401 Unauthorized`: Missing/invalid auth
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `429 Too Many Requests`: Rate limit exceeded
- `500 Internal Server Error`: Server error
- `503 Service Unavailable`: Temporary unavailable

### Error Response Format

```json
{
  "error": {
    "code": "invalid_request",
    "message": "Invalid deployment configuration",
    "details": {
      "field": "resources.cpu",
      "reason": "must be between 1 and 64"
    }
  }
}
```

### Retry Strategy

Implement exponential backoff:
```python
def retry_with_backoff(func, max_retries=3):
    for attempt in range(max_retries):
        try:
            return func()
        except RetryableError:
            if attempt == max_retries - 1:
                raise
            time.sleep(2 ** attempt)  # 1s, 2s, 4s
```

## Network Security

### TLS Configuration

```go
tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS13,
    CipherSuites: []uint16{
        tls.TLS_AES_256_GCM_SHA384,
        tls.TLS_CHACHA20_POLY1305_SHA256,
    },
    PreferServerCipherSuites: true,
}
```

### Certificate Pinning

For critical connections:
```go
expectedCertHash := "sha256/..."
actualCertHash := computeHash(cert)

if actualCertHash != expectedCertHash {
    return ErrCertificateMismatch
}
```

## Rate Limiting

### Algorithm

Token bucket algorithm:
```
Bucket Capacity: 100 requests
Refill Rate: 10 requests per second
```

### Headers

```http
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 85
X-RateLimit-Reset: 1640995200
```

### Exceeded Response

```http
HTTP/1.1 429 Too Many Requests
Retry-After: 60

{
  "error": {
    "code": "rate_limit_exceeded",
    "message": "Rate limit exceeded. Retry after 60 seconds."
  }
}
```

## Load Balancing

### Strategies

**Round Robin:**
```
Request 1 → Server A
Request 2 → Server B
Request 3 → Server C
Request 4 → Server A (repeat)
```

**Least Connections:**
```
Server A: 10 connections
Server B: 5 connections  ← Route here
Server C: 8 connections
```

**Geographic:**
```
User in US → US servers
User in EU → EU servers
```

### Health Checks

```http
GET /health HTTP/1.1

HTTP/1.1 200 OK
{
  "status": "healthy",
  "version": "1.0.0",
  "uptime": 86400
}
```

## Monitoring

### Metrics

Tracked per endpoint:
- Request rate (requests/second)
- Response time (percentiles: p50, p90, p99)
- Error rate
- Bandwidth usage

### Tracing

Distributed tracing headers:
```http
X-Request-ID: req_abc123
X-Trace-ID: trace_xyz789
X-Span-ID: span_001
```

## Protocol Versioning

### API Versioning

**URL Versioning:**
```
https://api.rios.com.ai/v1/deployments
https://api.rios.com.ai/v2/deployments
```

**Header Versioning:**
```http
API-Version: 2024-01-15
```

### Deprecation

1. Announce deprecation 6 months in advance
2. Add deprecation headers:
   ```http
   Deprecation: true
   Sunset: Sat, 31 Dec 2024 23:59:59 GMT
   ```
3. Maintain old version for 12 months
4. Remove after sunset date

## Client Libraries

Official SDKs available:

**Python:**
```python
from rios import Client

client = Client(api_token="your_token")
deployment = client.deployments.create(
    name="my-app",
    image="nginx:latest"
)
```

**JavaScript:**
```javascript
const { Client } = require('@rios/sdk');

const client = new Client({ apiToken: 'your_token' });
const deployment = await client.deployments.create({
  name: 'my-app',
  image: 'nginx:latest'
});
```

**Go:**
```go
import "github.com/rios/go-sdk"

client := rios.NewClient("your_token")
deployment, err := client.Deployments.Create(&rios.DeploymentConfig{
    Name:  "my-app",
    Image: "nginx:latest",
})
```

## Related Documentation

- [API Reference](../api-reference/README.md) - Complete API docs
- [System Design](system-design.md) - Architecture overview
- [DMoE Engine](dmoe-engine.md) - Distributed computing protocol

---

For protocol specifications and RFCs, see the [specs repository](https://github.com/rios/specs).

