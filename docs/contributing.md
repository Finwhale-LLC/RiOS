# Contributing to RiOS

Thank you for your interest in contributing to RiOS! This document provides guidelines and instructions for contributing to the project.

## Ways to Contribute

### 1. Code Contributions
- Bug fixes
- New features
- Performance improvements
- Test coverage

### 2. Documentation
- Improve existing docs
- Add tutorials and guides
- Translate documentation
- Fix typos and errors

### 3. Community Support
- Answer questions on the forum
- Help other users troubleshoot issues
- Share your RiOS use cases
- Write blog posts and tutorials

### 4. Testing
- Report bugs
- Test new features
- Participate in beta programs
- Provide feedback

### 5. Infrastructure
- Run worker nodes
- Contribute to network stability
- Test in different environments

## Getting Started

### Prerequisites

**For Code Contributions**:
- Git
- Go 1.19+ (for backend work)
- Node.js 18+ (for frontend work)
- Docker
- Familiarity with the codebase

**For Documentation**:
- Markdown knowledge
- GitBook (optional)
- Good written English

### Setting Up Development Environment

```bash
# Clone the repository
git clone https://github.com/rios/rios.git
cd rios

# Install dependencies
npm install          # for frontend
go mod download      # for backend

# Run tests
npm test             # frontend tests
go test ./...        # backend tests

# Start development server
npm run dev          # frontend
go run main.go       # backend
```

## Contribution Workflow

### 1. Find or Create an Issue

Before starting work:
- Check [existing issues](https://github.com/rios/rios/issues)
- Create a new issue if none exists
- Discuss your approach in the issue comments
- Wait for maintainer approval for large changes

### 2. Fork and Branch

```bash
# Fork the repository on GitHub
# Clone your fork
git clone https://github.com/YOUR_USERNAME/rios.git
cd rios

# Add upstream remote
git remote add upstream https://github.com/rios/rios.git

# Create a feature branch
git checkout -b feature/your-feature-name
```

### 3. Make Changes

**Code Style**:
- Follow existing code style
- Use meaningful variable names
- Add comments for complex logic
- Write tests for new features

**Commit Messages**:
```
<type>(<scope>): <subject>

<body>

<footer>
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `style`: Formatting changes
- `refactor`: Code restructuring
- `test`: Adding tests
- `chore`: Maintenance tasks

Example:
```
feat(api): add user profile update endpoint

- Implement PATCH /api/v1/users/:id
- Add validation for profile fields
- Update API documentation

Closes #123
```

### 4. Test Your Changes

```bash
# Run all tests
npm test
go test ./...

# Run specific tests
go test ./pkg/worker
npm test -- --testPathPattern=deployment

# Check test coverage
go test -cover ./...
npm test -- --coverage

# Run linters
golangci-lint run
npm run lint
```

### 5. Submit Pull Request

```bash
# Update your branch with latest upstream
git fetch upstream
git rebase upstream/main

# Push to your fork
git push origin feature/your-feature-name
```

On GitHub:
1. Navigate to your fork
2. Click "New Pull Request"
3. Select your feature branch
4. Fill in PR template
5. Submit PR

**PR Checklist**:
- [ ] Tests pass locally
- [ ] Code follows style guidelines
- [ ] Documentation updated
- [ ] CHANGELOG.md updated (for significant changes)
- [ ] Commits are well-formed
- [ ] PR description is clear

### 6. Code Review

After submission:
- Maintainers will review your PR
- Address feedback promptly
- Make requested changes
- Keep discussion professional and constructive

**Making Changes After Review**:
```bash
# Make changes
git add .
git commit -m "address review comments"

# Update PR
git push origin feature/your-feature-name
```

### 7. Merge

Once approved:
- Maintainer will merge your PR
- Your contribution will be in the next release
- You'll be credited in release notes

## Code Guidelines

### Go Code Style

```go
// Good: Clear function name and documentation
// ProcessPayment processes a payment transaction and returns the result
func ProcessPayment(userID string, amount float64) (*Payment, error) {
    if amount <= 0 {
        return nil, errors.New("invalid amount")
    }
    
    // Business logic here
    payment := &Payment{
        UserID: userID,
        Amount: amount,
        Status: StatusPending,
    }
    
    return payment, nil
}

// Bad: Unclear naming, no documentation
func pp(u string, a float64) (*Payment, error) {
    p := &Payment{}
    // ...
    return p, nil
}
```

### JavaScript/TypeScript Style

```javascript
// Good: Modern ES6+, clear logic
async function deployApplication(config) {
  try {
    const deployment = await createDeployment(config);
    await waitForReady(deployment.id);
    return {
      success: true,
      deployment,
    };
  } catch (error) {
    logger.error('Deployment failed:', error);
    throw new DeploymentError(error.message);
  }
}

// Bad: Callback hell, unclear error handling
function deployApplication(config, callback) {
  createDeployment(config, function(err, deployment) {
    if (err) {
      callback(err);
      return;
    }
    // ...
  });
}
```

### Testing Standards

```go
// Good: Clear test name, proper setup/teardown
func TestPaymentProcessing(t *testing.T) {
    tests := []struct {
        name        string
        userID      string
        amount      float64
        wantErr     bool
        expectedMsg string
    }{
        {
            name:    "valid payment",
            userID:  "user123",
            amount:  10.00,
            wantErr: false,
        },
        {
            name:        "invalid amount",
            userID:      "user123",
            amount:      -5.00,
            wantErr:     true,
            expectedMsg: "invalid amount",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ProcessPayment(tt.userID, tt.amount)
            if (err != nil) != tt.wantErr {
                t.Errorf("ProcessPayment() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            // Additional assertions...
        })
    }
}
```

## Documentation Guidelines

### Markdown Style

```markdown
# Main Title (H1) - One per page

Brief introduction paragraph.

## Section (H2)

Content for this section.

### Subsection (H3)

More detailed content.

#### Sub-subsection (H4)

Use sparingly.

## Code Examples

Always include:
- Language identifier
- Comments explaining non-obvious parts
- Complete, runnable examples when possible

\`\`\`bash
# Install RiOS CLI
curl -sSL https://get.rios.com.ai | bash

# Verify installation
rios version
\`\`\`

## Lists

- Use hyphens for unordered lists
- Keep items parallel in structure
- Use consistent punctuation

1. Use numbers for ordered lists
2. Ensure logical sequence
3. Keep steps clear and concise

## Links

- Use [descriptive text](https://example.com) for links
- Prefer relative links for internal docs: [Getting Started](getting-started/README.md)
- Check that all links work
```

### Documentation Checklist

- [ ] Clear title and introduction
- [ ] Logical structure with proper headings
- [ ] Code examples are tested and work
- [ ] Links are valid
- [ ] Screenshots are up-to-date (if applicable)
- [ ] Grammar and spelling checked
- [ ] Consistent terminology
- [ ] Follows existing doc style

## Community Guidelines

### Code of Conduct

We follow a [Code of Conduct](CODE_OF_CONDUCT.md) to ensure a welcoming community:

- **Be respectful**: Treat everyone with respect
- **Be constructive**: Provide helpful feedback
- **Be inclusive**: Welcome newcomers
- **Be professional**: No harassment or discrimination

### Communication Channels

- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: General questions and ideas
- **Discord**: Real-time community chat
- **Forum**: Long-form discussions
- **Twitter**: Updates and announcements

## Recognition

Contributors are recognized through:
- Credits in release notes
- Contributor list on website
- Special roles on Discord
- Swag for significant contributions (when available)
- Annual contributor awards

## Questions?

- Read the [FAQ](faq.md)
- Ask on [GitHub Discussions](https://github.com/rios/rios/discussions)
- Join our [Discord](https://discord.gg/rios)
- Email: contribute@rios.com.ai

## License

By contributing to RiOS, you agree that your contributions will be licensed under the [MIT License](../LICENSE).

---

**Thank you for contributing to RiOS! Together, we're building the future of decentralized computing.** 🚀

