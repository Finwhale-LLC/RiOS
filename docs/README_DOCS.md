# RiOS Documentation

This directory contains the complete documentation for RiOS (Intelligent Operating System) using GitBook format.

## 📚 Documentation Structure

```
docs/
├── README.md                 # Documentation home page
├── SUMMARY.md               # Table of contents
├── book.json                # GitBook configuration
├── .gitbook.yaml            # GitBook settings
├── getting-started/         # Getting started guides
│   ├── README.md
│   ├── prerequisites.md
│   ├── quick-start.md
│   └── installation.md
├── architecture/            # Technical architecture
│   ├── README.md
│   ├── system-design.md
│   ├── dmoe-engine.md
│   ├── security-model.md
│   └── network-protocol.md
├── user-guide/              # User guides
│   └── ...
├── worker-setup/            # Worker node setup
│   └── ...
├── api-reference/           # API documentation
│   └── ...
├── economic-model/          # Token economics
│   └── ...
├── faq.md                   # Frequently asked questions
├── glossary.md              # Terminology
├── contributing.md          # Contribution guide
└── changelog.md             # Version history
```

## 🚀 Quick Start

### View Documentation Locally

1. **Install GitBook CLI**
   ```bash
   npm install -g gitbook-cli
   ```

2. **Install Dependencies**
   ```bash
   cd docs
   gitbook install
   ```

3. **Serve Locally**
   ```bash
   gitbook serve
   ```
   
   Open [http://localhost:4000](http://localhost:4000) in your browser.

### Build Static HTML

```bash
gitbook build
```

Output will be in the `_book/` directory.

## 🌐 Deployment Options

### Option 1: GitBook.com (Recommended)

1. Create account at [GitBook.com](https://www.gitbook.com/)
2. Connect your GitHub repository
3. Set content source to `docs/` directory
4. Automatic deployment on push

**Custom Domain**: Configure `docs.rios.com.ai` in GitBook settings

### Option 2: GitHub Pages

```bash
# Build documentation
gitbook build

# Deploy to gh-pages branch
git subtree push --prefix docs/_book origin gh-pages
```

Access at: `https://yourusername.github.io/rios/`

### Option 3: Self-Hosted

```bash
# Build
gitbook build

# Copy to web server
rsync -avz _book/ user@server:/var/www/docs/
```

### Option 4: Netlify

```bash
npm install -g netlify-cli
netlify deploy --prod --dir=_book
```

### Option 5: Vercel

```bash
npm install -g vercel
vercel --prod
```

See [DOCS_DEPLOYMENT.md](DOCS_DEPLOYMENT.md) for detailed deployment instructions.

## 📝 Contributing to Documentation

### Adding a New Page

1. Create a new `.md` file in the appropriate directory
2. Add entry to `SUMMARY.md`
3. Test locally with `gitbook serve`
4. Submit pull request

### Documentation Standards

- Use clear, concise language
- Include code examples where applicable
- Add screenshots for UI-related topics
- Test all commands and code snippets
- Follow the existing structure and style

### Markdown Guidelines

```markdown
# H1 for page title (one per page)
## H2 for main sections
### H3 for subsections

**Bold** for emphasis
*Italic* for technical terms
`code` for inline code
```

**Code blocks with language:**
````markdown
```bash
rios deploy --name my-app
```

```javascript
const config = { ... };
```
````

### Linking

- **Internal links**: `[Getting Started](getting-started/README.md)`
- **External links**: `[RiOS Website](https://rios.com.ai)`
- **Anchor links**: `[DMoE Engine](#dmoe-engine)`

## 🔍 Documentation Coverage

Current documentation includes:

- ✅ Getting Started Guide
- ✅ Architecture Overview
- ✅ DMoE Engine Details
- ✅ API Reference
- ✅ Worker Setup Guide
- ✅ FAQ
- ✅ Glossary
- ✅ Contributing Guide
- ⏳ User Guide (in progress)
- ⏳ Economic Model (in progress)
- ⏳ Tutorials (planned)
- ⏳ Video Guides (planned)

## 🌍 Multi-language Support (Planned)

Future language support:
- 🇬🇧 English (current)
- 🇨🇳 中文 (planned)
- 🇯🇵 日本語 (planned)
- 🇰🇷 한국어 (planned)

## 📦 GitBook Plugins

Configured in `book.json`:

- `theme-default` - Default GitBook theme
- `search` - Full-text search
- `sharing` - Social sharing buttons
- `fontsettings` - Font customization
- `code` - Enhanced code blocks
- `prism` - Syntax highlighting

## 🛠️ Local Development

### Requirements

- Node.js 14+
- npm or yarn
- GitBook CLI

### Development Workflow

```bash
# Install dependencies
npm install

# Start development server
npm run docs:dev
# or
gitbook serve

# Build for production
npm run docs:build
# or
gitbook build

# Clean build artifacts
rm -rf _book .gitbook
```

### Troubleshooting

**Build fails:**
```bash
# Clear cache and rebuild
rm -rf _book .gitbook
gitbook install
gitbook build
```

**Plugin errors:**
```bash
# Reinstall plugins
gitbook install
```

**Port already in use:**
```bash
# Use different port
gitbook serve --port 4001
```

## 📊 Documentation Metrics

Track documentation quality:
- Coverage: What percentage of features are documented?
- Freshness: Are docs up-to-date with latest version?
- Clarity: User feedback and ratings
- Completeness: Are all topics covered?

## 🔗 Useful Links

- **Main Website**: [https://rios.com.ai](https://rios.com.ai)
- **Cloud Service**: [https://cloud.rios.com.ai](https://cloud.rios.com.ai)
- **GitHub**: [https://github.com/rios/rios](https://github.com/rios/rios)
- **Community Forum**: [https://community.rios.com.ai](https://community.rios.com.ai)
- **API Docs**: [https://api.rios.com.ai/docs](https://api.rios.com.ai/docs)

## 📧 Contact

For documentation questions:
- **Email**: docs@rios.com.ai
- **GitHub Issues**: [Report Documentation Issues](https://github.com/rios/rios/issues)
- **Discord**: Join our documentation channel

## 📄 License

Documentation is licensed under [Creative Commons Attribution 4.0 International (CC BY 4.0)](https://creativecommons.org/licenses/by/4.0/).

Code examples in documentation are licensed under [MIT License](../LICENSE).

---

**© 2015-2025 RiOS Foundation. All rights reserved.**

