# GitBook Documentation Deployment Guide

This guide explains how to deploy the RiOS documentation using GitBook.

## Prerequisites

- Node.js 14+ installed
- npm or yarn package manager
- GitBook CLI or GitBook.com account

## Deployment Options

### Option 1: GitBook.com (Recommended)

GitBook.com provides the easiest way to host and deploy your documentation.

#### Setup Steps

1. **Create a GitBook Account**
   - Visit [https://www.gitbook.com/](https://www.gitbook.com/)
   - Sign up for a free account
   - Create a new space for RiOS documentation

2. **Connect to GitHub**
   - In GitBook, go to Integrations → GitHub
   - Connect your GitHub repository
   - Select the `docs/` directory as the content source

3. **Configure GitBook**
   - GitBook will automatically detect `.gitbook.yaml`
   - Set the primary branch (usually `main`)
   - Enable automatic sync

4. **Publish**
   - GitBook will automatically build and deploy
   - Get your documentation URL (e.g., `rios.gitbook.io/docs`)
   - Optionally configure custom domain

#### Custom Domain Setup

1. In GitBook settings, go to "Domain Settings"
2. Add your custom domain: `docs.rios.com.ai`
3. Configure DNS records:
   ```
   CNAME docs.rios.com.ai -> hosting.gitbook.io
   ```
4. Verify domain ownership
5. Enable HTTPS (automatic with GitBook)

### Option 2: Self-Hosted GitBook

Build and host the documentation yourself.

#### Installation

```bash
# Navigate to docs directory
cd docs/

# Install GitBook CLI globally
npm install -g gitbook-cli

# Install GitBook locally in the project
gitbook install

# Or if using package.json
npm install
```

#### Local Development

```bash
# Serve documentation locally
npm run docs:dev
# or
gitbook serve

# Open browser to http://localhost:4000
```

#### Build for Production

```bash
# Build static HTML files
npm run docs:build
# or
gitbook build

# Output will be in _book/ directory
```

#### Deploy to Web Server

**Using Nginx**:

```nginx
# /etc/nginx/sites-available/docs.rios.com.ai

server {
    listen 80;
    server_name docs.rios.com.ai;
    
    root /var/www/rios-docs/_book;
    index index.html;
    
    location / {
        try_files $uri $uri/ =404;
    }
    
    # Enable gzip compression
    gzip on;
    gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;
}
```

```bash
# Upload files
rsync -avz _book/ user@server:/var/www/rios-docs/

# Reload nginx
sudo systemctl reload nginx

# Setup SSL with Let's Encrypt
sudo certbot --nginx -d docs.rios.com.ai
```

**Using Apache**:

```apache
# /etc/apache2/sites-available/docs.rios.com.ai.conf

<VirtualHost *:80>
    ServerName docs.rios.com.ai
    DocumentRoot /var/www/rios-docs/_book
    
    <Directory /var/www/rios-docs/_book>
        Options -Indexes +FollowSymLinks
        AllowOverride All
        Require all granted
    </Directory>
    
    ErrorLog ${APACHE_LOG_DIR}/rios-docs-error.log
    CustomLog ${APACHE_LOG_DIR}/rios-docs-access.log combined
</VirtualHost>
```

### Option 3: Deploy to Netlify

Netlify offers easy deployment with continuous integration.

#### Setup

1. **Create netlify.toml**

```toml
[build]
  base = "docs/"
  command = "gitbook build"
  publish = "_book"

[[redirects]]
  from = "/*"
  to = "/index.html"
  status = 200
```

2. **Deploy**

```bash
# Install Netlify CLI
npm install -g netlify-cli

# Login to Netlify
netlify login

# Initialize site
netlify init

# Deploy
netlify deploy --prod
```

3. **Configure Custom Domain**
   - In Netlify dashboard: Settings → Domain Management
   - Add custom domain: `docs.rios.com.ai`
   - Configure DNS (Netlify provides instructions)
   - HTTPS is automatic

### Option 4: Deploy to GitHub Pages

Use GitHub Pages for free hosting.

#### Setup

1. **Create GitHub Action**

Create `.github/workflows/deploy-docs.yml`:

```yaml
name: Deploy Documentation

on:
  push:
    branches: [ main ]
    paths:
      - 'docs/**'

jobs:
  deploy:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v2
    
    - name: Setup Node.js
      uses: actions/setup-node@v2
      with:
        node-version: '18'
    
    - name: Install GitBook
      run: |
        cd docs
        npm install -g gitbook-cli
        gitbook install
    
    - name: Build GitBook
      run: |
        cd docs
        gitbook build
    
    - name: Deploy to GitHub Pages
      uses: peaceiris/actions-gh-pages@v3
      with:
        github_token: ${{ secrets.GITHUB_TOKEN }}
        publish_dir: ./docs/_book
```

2. **Enable GitHub Pages**
   - Repository Settings → Pages
   - Source: gh-pages branch
   - Custom domain: `docs.rios.com.ai`

3. **Configure DNS**
   ```
   CNAME docs.rios.com.ai -> yourusername.github.io
   ```

### Option 5: Deploy to Vercel

Vercel provides instant deployment with great performance.

#### Setup

```bash
# Install Vercel CLI
npm install -g vercel

# Navigate to docs
cd docs/

# Deploy
vercel --prod
```

**vercel.json**:

```json
{
  "version": 2,
  "builds": [
    {
      "src": "package.json",
      "use": "@vercel/static-build",
      "config": {
        "distDir": "_book"
      }
    }
  ],
  "routes": [
    {
      "src": "/(.*)",
      "dest": "/$1"
    }
  ]
}
```

## Update Documentation Link in Website

After deploying, update the main website to link to your documentation.

### Update index.html

Add to navigation (line ~467):

```html
<a href="https://docs.rios.com.ai" class="nav-link" data-i18n="nav.docs">Documentation</a>
```

Or link to your chosen deployment URL:
- GitBook.com: `https://rios.gitbook.io/docs`
- Self-hosted: `https://docs.rios.com.ai`
- GitHub Pages: `https://rios.github.io/docs`
- Netlify: `https://rios-docs.netlify.app`
- Vercel: `https://rios-docs.vercel.app`

## Continuous Deployment

### Automatic Updates

All hosting options support automatic deployment:

- **GitBook.com**: Auto-syncs with GitHub pushes
- **Netlify**: Triggers build on git push
- **Vercel**: Automatic deployment from git
- **GitHub Pages**: GitHub Actions workflow
- **Self-hosted**: Setup git hooks or CI/CD

### Example CI/CD with GitHub Actions

```yaml
name: Update Documentation

on:
  push:
    branches: [ main ]
    paths:
      - 'docs/**'

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Build and Deploy
        run: |
          cd docs
          npm install
          npm run docs:build
          
          # Deploy to your server
          rsync -avz _book/ ${{ secrets.DEPLOY_USER }}@${{ secrets.DEPLOY_HOST }}:/var/www/docs/
        env:
          SSH_PRIVATE_KEY: ${{ secrets.SSH_PRIVATE_KEY }}
```

## Maintenance

### Updating Documentation

1. Edit markdown files in `docs/`
2. Test locally: `npm run docs:dev`
3. Commit and push changes
4. Deployment happens automatically

### Adding New Pages

1. Create new .md file
2. Add entry to `SUMMARY.md`
3. Commit and push

### Translations (Future)

GitBook supports multi-language documentation:

```
docs/
  en/
    README.md
    SUMMARY.md
  zh/
    README.md
    SUMMARY.md
```

## Monitoring

- **GitBook.com**: Built-in analytics
- **Self-hosted**: Use Google Analytics or similar
- **Netlify/Vercel**: Built-in analytics available

## Troubleshooting

### Build Fails

```bash
# Clear GitBook cache
rm -rf _book .gitbook

# Reinstall plugins
gitbook install

# Rebuild
gitbook build
```

### Broken Links

```bash
# Check for broken links
npx broken-link-checker http://localhost:4000

# Or use markdown-link-check
npx markdown-link-check docs/**/*.md
```

### Styling Issues

Check `book.json` for plugin configurations and custom CSS in theme files.

## Recommended Setup

For RiOS, we recommend:

1. **Primary**: GitBook.com with custom domain `docs.rios.com.ai`
   - Easy to use
   - Professional appearance
   - Auto-deployment
   - Built-in search and analytics

2. **Backup**: GitHub Pages
   - Free hosting
   - Version controlled
   - Automatic deployment via Actions

## Next Steps

After deployment:
1. Update main website navigation
2. Add link to docs in README
3. Share documentation URL with community
4. Monitor for issues and feedback
5. Keep documentation up-to-date

## Support

For deployment issues:
- GitBook: [support.gitbook.com](https://support.gitbook.com)
- GitHub Pages: [docs.github.com/pages](https://docs.github.com/pages)
- Community: [community.rios.com.ai](https://community.rios.com.ai)

