<div align="center">

<img src="./build/appicon.png" width="120" alt="buddybuddy-toolkit" />

# buddybuddy-toolkit

A desktop tool for batch-processing image filenames

[![Release](https://img.shields.io/github/v/release/Wiiiiill/buddybuddy-toolkit?sort=semver)](https://github.com/Wiiiiill/buddybuddy-toolkit/releases)
[![License](https://img.shields.io/github/license/Wiiiiill/buddybuddy-toolkit)](./LICENSE)
[![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-brightgreen)]()

[简体中文](README.md) · [English](README.en.md)

</div>

## Introduction

buddybuddy-toolkit is a small desktop utility built with [Wails](https://wails.io) for batch-processing image filenames. Pick a folder and it can decode garbled filenames, archive files by prefix, and replace strings in filenames.

## Features

- **Decode filenames** — Restore URL-encoded filenames (e.g. `%E4%B8%AD.jpg`) to readable text
- **Archive by prefix** — Create subfolders from the part before the underscore and move files in
- **Replace in filenames** — Find and replace a substring, with quick presets (e.g. `.jpg` → `_首图.jpg`)
- **Configurable** — Supported image formats and preset rules are editable in Settings and persisted
- **Themes** — All 32 built-in daisyUI themes with color previews
- **Bilingual** — Switch between English and Chinese at any time
- **Feedback** — Shows how many files were actually processed

## Download

Grab the package for your platform from [Releases](https://github.com/Wiiiiill/buddybuddy-toolkit/releases):

| Platform | File |
| --- | --- |
| Windows | `buddybuddy-toolkit-windows-amd64.zip` |
| macOS (Intel / Apple Silicon) | `buddybuddy-toolkit-macos-universal.zip` |
| Linux | `buddybuddy-toolkit-linux-amd64.tar.gz` |

## Build from source

### Requirements

- [Go](https://go.dev/dl/) 1.26+
- [Node.js](https://nodejs.org/) 18+
- [Wails CLI](https://wails.io/docs/gettingstarted/installation) v2.17.1

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@v2.17.1
```

### Development

```bash
git clone https://github.com/Wiiiiill/buddybuddy-toolkit.git
cd buddybuddy-toolkit
wails dev
```

### Package

```bash
wails build
```

Output goes to `build/bin/`. The repository is set up with GitHub Actions: pushing a `v*` tag builds all three platforms and creates a release automatically.

## Tech stack

| Layer | Technology |
| --- | --- |
| Backend | Go, Wails v2 |
| Frontend | Vue 3, TypeScript, Vite, Pinia, Vue Router, Vue I18n |
| Styling | TailwindCSS, daisyUI |

## Configuration file

Settings are stored in `buddybuddy-toolkit/config.json` inside the user config directory:

- Windows: `%AppData%\buddybuddy-toolkit\config.json`
- macOS: `~/Library/Application Support/buddybuddy-toolkit/config.json`
- Linux: `~/.config/buddybuddy-toolkit/config.json`

## License

[MIT](./LICENSE)
