# Glossary

## A

### API (Application Programming Interface)
A set of protocols and tools for building software applications. RiOS provides RESTful APIs for programmatic access.

### Auto-scaling
Automatic adjustment of computing resources based on demand. RiOS supports both horizontal (more instances) and vertical (more resources per instance) scaling.

## B

### Bandwidth
The rate of data transfer, typically measured in Mbps or Gbps. RiOS's DMoE architecture significantly reduces bandwidth requirements.

### Billing
The system for calculating and charging for resource usage. RiOS uses per-second billing with ROS tokens.

### Blockchain
A distributed ledger technology used by RiOS for transparent token transactions and reward distribution.

## C

### CLI (Command Line Interface)
A text-based interface for interacting with RiOS services. The `rios` CLI provides commands for deployment and management.

### Container
A lightweight, standalone executable package that includes everything needed to run software. RiOS uses Docker containers.

### Computing Node
A server or machine that provides computing resources to the RiOS network. Also called a "worker node."

## D

### Dashboard
The web-based graphical interface for managing RiOS services at [cloud.rios.com.ai](https://cloud.rios.com.ai).

### Decentralization
Distribution of control and resources across multiple nodes rather than a central authority. Core principle of RiOS architecture.

### Deployment
An instance of an application running on RiOS infrastructure.

### DMoE (Decentralized Mixture of Experts)
RiOS's core technology for distributing AI model computation across multiple nodes efficiently.

### Docker
A platform for developing, shipping, and running applications in containers.

## E

### Endpoint
The URL or address where an application or API can be accessed.

### Expert Module
A specialized component of an AI model in the DMoE architecture, designed to handle specific types of tasks.

## F

### Failover
Automatic switching to a backup system when the primary system fails. RiOS provides automatic failover for high availability.

### Fault Tolerance
The ability of a system to continue operating properly in the event of failure of some components.

## G

### GPU (Graphics Processing Unit)
Specialized hardware for parallel processing, essential for AI and machine learning workloads.

### gVisor
A container runtime sandbox that provides enhanced security isolation for RiOS workloads.

## H

### Health Check
Automated monitoring to verify that a service is running correctly. RiOS performs regular health checks on all deployments.

### High Availability (HA)
System design approach that ensures a certain level of operational performance, usually uptime, for a higher than normal period.

### Horizontal Scaling
Adding more instances of an application to handle increased load.

## I

### Infrastructure
The underlying foundation of computing resources (hardware, network, storage) that supports applications.

### IPFS (InterPlanetary File System)
A distributed file storage protocol that influences RiOS's storage architecture.

## J

### JWT (JSON Web Token)
A compact token format used for authentication in RiOS APIs.

## K

### Kubernetes (K8s)
A container orchestration platform. RiOS can integrate with Kubernetes for certain deployment scenarios.

## L

### Latency
The time delay between a request and response. RiOS's distributed architecture aims to minimize latency.

### Load Balancer
A system that distributes network traffic across multiple servers to optimize resource use and maximize throughput.

## M

### Metrics
Quantitative measurements of system performance, such as CPU usage, memory consumption, and request rates.

### Microservices
An architectural style that structures an application as a collection of loosely coupled services.

### Mining
In RiOS context, the process of earning ROS tokens by providing computing resources to the network.

## N

### Node
A single computing unit in the RiOS network. Can be a user's deployment or a worker providing resources.

### NFT (Non-Fungible Token)
While not currently used, RiOS may integrate NFTs for unique digital asset management in future versions.

## O

### Orchestration
Automated configuration, coordination, and management of computer systems and software. RiOS's DMoE orchestrator manages distributed computations.

### OAuth
An open standard for access delegation, used as an authentication method in RiOS.

## P

### P2P (Peer-to-Peer)
A decentralized network architecture where each node can act as both client and server. Used in RiOS worker communication.

### Pod
In Kubernetes terminology, a group of one or more containers. RiOS uses similar concepts for deployment units.

### PoS (Proof of Stake)
The consensus mechanism used in RiOS's blockchain layer for validating transactions.

## R

### Rate Limiting
Controlling the rate of requests to an API to prevent abuse. RiOS implements rate limiting at the API gateway.

### Replica
A copy of an application instance. Running multiple replicas provides redundancy and scalability.

### REST (Representational State Transfer)
An architectural style for APIs. RiOS provides RESTful APIs for all services.

### ROS Token
The native cryptocurrency of the RiOS network, used for payments and rewards.

### Router Network
In DMoE, the intelligent system that determines which expert modules should handle a given request.

## S

### Sandbox
An isolated environment for running untrusted code. RiOS uses sandboxing for security.

### Scalability
The capability to handle increasing amounts of work or to be enlarged to accommodate growth.

### SLA (Service Level Agreement)
A commitment between a service provider and client regarding aspects like uptime and performance. RiOS offers SLAs for dedicated deployments.

### Smart Contract
Self-executing contracts with terms directly written into code, used in RiOS's blockchain layer.

### SSL/TLS
Security protocols for encrypted communication over networks. RiOS uses TLS 1.3 for all data transmission.

## T

### Tensor
A multi-dimensional array used in machine learning computations. In DMoE, tensors are passed between expert modules.

### Throughput
The amount of work completed in a given time period, often measured in requests per second.

### Token
In blockchain context, a digital asset. ROS tokens are used within the RiOS ecosystem.

## V

### Vertical Scaling
Increasing the resources (CPU, RAM) of an existing instance rather than adding more instances.

### Virtual Machine (VM)
An emulation of a computer system. RiOS primarily uses containers, which are more lightweight than VMs.

## W

### Wallet
A digital tool for storing and managing cryptocurrency. Users need a ROS wallet to pay for and receive tokens.

### Worker Node
A computing resource provider in the RiOS network that executes user workloads and earns ROS tokens.

### Workload
A specific task or application running on computing infrastructure.

## Z

### Zero Trust
A security model based on the principle of "never trust, always verify." RiOS implements zero trust architecture across all components.

---

**Don't see a term you're looking for?** Please let us know at docs@rios.com.ai or contribute to this glossary on [GitHub](https://github.com/rios/docs).

