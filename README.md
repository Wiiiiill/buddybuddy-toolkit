<div align="center">

<img src="./build/appicon.png" width="120" alt="buddybuddy-toolkit" />

# 好哥哥工具

批量处理图片文件名的桌面小工具

[![Release](https://img.shields.io/github/v/release/Wiiiiill/buddybuddy-toolkit?sort=semver)](https://github.com/Wiiiiill/buddybuddy-toolkit/releases)
[![License](https://img.shields.io/github/license/Wiiiiill/buddybuddy-toolkit)](./LICENSE)
[![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-brightgreen)]()

[简体中文](README.md) · [English](README.en.md)

</div>

## 简介

「好哥哥工具」是一个基于 [Wails](https://wails.io) 构建的桌面小工具，用来批量整理图片文件名。选一个目录，即可完成文件名解码、按前缀归档、字符串替换等常见操作。

## 功能

- **处理文件名乱码** — 将 URL 编码的文件名（如 `%E4%B8%AD.jpg`）还原为可读文字
- **按前缀归档** — 根据文件名下划线前的部分自动创建子文件夹并归入
- **替换文件名** — 查找并替换文件名中的一段字符串，支持快捷预设（如 `.jpg` → `_首图.jpg`）
- **可配置** — 支持的图片格式、快捷预设规则均可在设置页中增删改，配置持久化保存
- **多主题** — 内置 daisyUI 全部 32 个主题，带配色预览，一键切换
- **中英双语** — 界面语言可随时切换
- **结果反馈** — 操作完成后显示实际处理的文件数量

## 下载

前往 [Releases](https://github.com/Wiiiiill/buddybuddy-toolkit/releases) 下载对应平台的安装包：

| 平台 | 文件 |
| --- | --- |
| Windows | `buddybuddy-toolkit-windows-amd64.zip` |
| macOS（Intel / Apple Silicon） | `buddybuddy-toolkit-macos-universal.zip` |
| Linux | `buddybuddy-toolkit-linux-amd64.tar.gz` |

## 从源码构建

### 环境要求

- [Go](https://go.dev/dl/) 1.26+
- [Node.js](https://nodejs.org/) 18+
- [Wails CLI](https://wails.io/docs/gettingstarted/installation) v2.17.1

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@v2.17.1
```

### 开发

```bash
git clone https://github.com/Wiiiiill/buddybuddy-toolkit.git
cd buddybuddy-toolkit
wails dev
```

### 打包

```bash
wails build
```

产物位于 `build/bin/`。仓库已配置 GitHub Actions，推送 `v*` 标签会自动构建三平台并创建 Release。

## 技术栈

| 层 | 技术 |
| --- | --- |
| 后端 | Go、Wails v2 |
| 前端 | Vue 3、TypeScript、Vite、Pinia、Vue Router、Vue I18n |
| 样式 | TailwindCSS、daisyUI |

## 配置文件

设置项保存在用户配置目录下的 `buddybuddy-toolkit/config.json`：

- Windows：`%AppData%\buddybuddy-toolkit\config.json`
- macOS：`~/Library/Application Support/buddybuddy-toolkit/config.json`
- Linux：`~/.config/buddybuddy-toolkit/config.json`

## 许可证

[MIT](./LICENSE)
