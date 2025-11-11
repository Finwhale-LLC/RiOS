# GitBook.com 设置步骤指南

您已经注册并连接了 GitHub，接下来按照以下步骤完成设置。

## 第一步：创建新的 Space（空间）

1. **进入 GitBook 主页**
   - 登录后，点击右上角的 **"New Space"** 按钮
   - 或者在左侧边栏点击 **"+"** 创建新空间

2. **选择创建方式**
   - 选择 **"Import"** 选项卡
   - 点击 **"GitHub"** 图标

3. **配置导入设置**
   ```
   Repository: 选择您的 RiOS 仓库
   Branch: main (或您的主分支名称)
   Monorepo path: docs/
   ```
   
   ⚠️ **重要**：确保设置 `Monorepo path` 为 `docs/`，这样 GitBook 只会读取 docs 目录的内容

4. **命名您的 Space**
   - Name: `RiOS Documentation`
   - Slug: `rios-docs` (这将是您 URL 的一部分)

5. **点击 "Import"**
   - GitBook 会开始导入您的文档
   - 需要几分钟时间处理

## 第二步：验证导入

导入完成后，您应该看到：

✅ 左侧边栏显示文档目录结构
✅ 主内容区显示 README.md 的内容
✅ 右上角有 "Published" 状态

**检查项目：**
- [ ] 首页显示正常
- [ ] 左侧目录结构正确
- [ ] 可以点击导航到各个页面
- [ ] Getting Started 部分可访问
- [ ] Architecture 部分可访问

## 第三步：配置自动同步

1. **进入 Space 设置**
   - 点击右上角 **"..."** 菜单
   - 选择 **"Settings"**

2. **启用 GitHub 同步**
   - 在左侧菜单选择 **"Integrations"**
   - 找到 **"GitHub"** 部分
   - 确保 **"GitHub Sync"** 已启用
   - 设置同步选项：
     ```
     ✓ Sync on push to main branch
     ✓ Auto-publish changes
     Branch: main
     Path: docs/
     ```

3. **保存设置**

现在，每次您推送更新到 GitHub 的 `docs/` 目录，GitBook 会自动更新！

## 第四步：获取文档 URL

1. **查看已发布的文档**
   - 点击右上角的 **"Share"** 按钮
   - 或点击 **"Published"** 标签

2. **复制公开 URL**
   
   您的文档 URL 格式为：
   ```
   https://yourusername.gitbook.io/rios-docs
   ```
   或者
   ```
   https://app.gitbook.com/@yourusername/s/rios-docs/
   ```

3. **测试链接**
   - 在浏览器中打开该 URL
   - 确认文档显示正常
   - 测试导航和搜索功能

## 第五步：更新网站导航链接

现在需要更新您网站上的文档链接：

### 更新 index.html

**当前链接（本地）：**
```html
<a href="docs/index.html" class="nav-link" data-i18n="nav.docs">Documentation</a>
```

**更新为 GitBook URL：**
```html
<a href="https://yourusername.gitbook.io/rios-docs" class="nav-link" data-i18n="nav.docs" target="_blank">Documentation</a>
```

⚠️ 替换 `yourusername` 和 `rios-docs` 为您实际的 URL

**需要更新的位置：**
1. 第 469 行（桌面导航）
2. 第 501 行（移动端菜单）

### 添加 target="_blank"

建议添加 `target="_blank"` 以在新标签页打开文档：
```html
<a href="https://your-gitbook-url" class="nav-link" target="_blank">Documentation</a>
```

## 第六步：配置自定义域名（可选）

如果您想使用 `docs.rios.com.ai` 作为文档地址：

### 在 GitBook 中设置

1. **进入 Space 设置**
   - Settings → Customization → Custom Domain

2. **添加自定义域名**
   ```
   Custom domain: docs.rios.com.ai
   ```

3. **获取 DNS 配置信息**
   GitBook 会提供 CNAME 记录，通常是：
   ```
   CNAME docs.rios.com.ai → hosting.gitbook.io
   ```

### 配置 DNS

在您的域名服务商（如 Cloudflare, GoDaddy 等）添加 DNS 记录：

```
Type: CNAME
Name: docs
Value: hosting.gitbook.io
TTL: Auto (或 3600)
```

### 验证域名

1. **等待 DNS 传播**（可能需要 5-60 分钟）
   
   检查 DNS 是否生效：
   ```bash
   nslookup docs.rios.com.ai
   ```

2. **在 GitBook 中验证**
   - 返回 Custom Domain 设置
   - 点击 **"Verify"**
   - GitBook 会自动启用 HTTPS

3. **测试自定义域名**
   ```
   https://docs.rios.com.ai
   ```

### 更新网站链接为自定义域名

```html
<a href="https://docs.rios.com.ai" class="nav-link" target="_blank">Documentation</a>
```

## 第七步：优化和定制

### 1. 自定义外观

**Settings → Customization**
- **Logo**: 上传 RiOS logo
- **Favicon**: 设置网站图标
- **Primary Color**: `#f66666` (与主站一致)
- **Theme**: Light/Dark 模式

### 2. 配置 SEO

**Settings → SEO**
```
Title: RiOS Documentation - Intelligent Operating System
Description: Comprehensive documentation for RiOS, the next-generation decentralized computing platform
```

### 3. 启用搜索分析

**Settings → Integrations**
- Google Analytics（可选）
- 查看搜索统计
- 了解用户行为

### 4. 配置导航

**Settings → Customization → Header**
```
Add external links:
- RiOS Website: https://rios.com.ai
- Cloud Service: https://cloud.rios.com.ai
- GitHub: https://github.com/rios/rios
```

## 测试检查清单

完成设置后，进行以下测试：

- [ ] ✅ 文档 URL 可以访问
- [ ] ✅ 主页内容显示正确
- [ ] ✅ 左侧导航可以展开/折叠
- [ ] ✅ 所有页面链接正常工作
- [ ] ✅ 搜索功能正常
- [ ] ✅ 代码块语法高亮正确
- [ ] ✅ 移动端显示正常
- [ ] ✅ 主网站导航链接正确
- [ ] ✅ 新标签页打开文档
- [ ] ✅ HTTPS 证书有效（如使用自定义域名）

## 常见问题解决

### Q1: 导入失败
**原因**：GitBook 找不到 `.gitbook.yaml` 或 `SUMMARY.md`

**解决方案**：
1. 确认 GitHub 仓库中有 `docs/.gitbook.yaml`
2. 确认 `docs/SUMMARY.md` 存在
3. 重新导入，确保路径设置为 `docs/`

### Q2: 目录结构不正确
**原因**：`SUMMARY.md` 格式问题

**解决方案**：
1. 检查 `docs/SUMMARY.md` 的格式
2. 确保缩进正确（2 个空格）
3. 确保所有文件路径正确

### Q3: 图片不显示
**原因**：图片路径不正确

**解决方案**：
1. 将图片放在 `docs/` 目录下
2. 使用相对路径：`![Logo](./images/logo.png)`
3. 重新推送到 GitHub

### Q4: 自定义域名验证失败
**原因**：DNS 配置不正确或未生效

**解决方案**：
1. 检查 DNS 记录是否正确
2. 等待 DNS 传播（最多 24 小时）
3. 使用 `nslookup` 检查 DNS
4. 清除浏览器缓存

### Q5: 更新不同步
**原因**：GitHub 同步未启用或分支不对

**解决方案**：
1. 检查 Settings → Integrations → GitHub Sync
2. 确认分支设置正确
3. 手动触发同步：Settings → GitHub → Sync Now

## 性能优化建议

### 1. 图片优化
```bash
# 压缩图片
npm install -g imagemin-cli
imagemin docs/images/*.png --out-dir=docs/images/optimized
```

### 2. 减少页面大小
- 避免单页内容过长
- 拆分大文档为多个小页面
- 使用目录和子章节

### 3. 加快加载速度
- GitBook 自动优化
- 使用 CDN（GitBook 内置）
- 启用浏览器缓存

## 下一步行动

### 立即执行（今天）

1. ✅ **完成 GitBook 设置**
   - 按照上述步骤操作
   - 验证所有功能正常

2. ✅ **更新网站链接**
   - 修改 index.html 中的文档链接
   - 测试从主站点击文档链接

3. ✅ **分享给团队**
   - 发送文档 URL 给团队成员
   - 收集初步反馈

### 本周内完成

4. **添加更多内容**
   - 完善 User Guide 部分
   - 添加 API Reference
   - 补充 Worker Setup 详细步骤

5. **优化用户体验**
   - 添加截图和图表
   - 录制视频教程
   - 添加更多代码示例

6. **推广文档**
   - 在社交媒体分享
   - 发布公告
   - 邀请用户反馈

### 本月内完成

7. **多语言支持**（可选）
   - 创建中文版本
   - 日文、韩文版本

8. **高级功能**
   - 集成 API playground
   - 添加交互式示例
   - 视频嵌入

## 支持资源

- **GitBook 官方文档**: [docs.gitbook.com](https://docs.gitbook.com)
- **GitBook 社区**: [community.gitbook.com](https://community.gitbook.com)
- **支持邮箱**: support@gitbook.com

## 联系方式

如有问题：
- **技术问题**: 在 GitHub 创建 issue
- **文档内容**: docs@rios.com.ai
- **紧急支持**: support@rios.com.ai

---

## 🎉 恭喜！

完成上述步骤后，您的 RiOS 文档将：
- ✅ 在 GitBook.com 上专业托管
- ✅ 从主网站可访问
- ✅ 自动与 GitHub 同步
- ✅ 支持全文搜索
- ✅ 移动端友好
- ✅ SEO 优化
- ✅ 免费 HTTPS

**您的文档系统已经生产就绪！** 🚀

---

有任何问题请随时询问！

