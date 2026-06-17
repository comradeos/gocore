# Go + Wails Minimal App

Minimal desktop app on:

- Go backend
- Wails bridge
- Vanilla `HTML/CSS/JS` frontend

The UI has one input and one button. Clicking the button calls a Go method through Wails and shows:

- a greeting from Go
- current platform and architecture
- current Go runtime version

## Why this stack works well

`Go` handles the backend logic and native packaging workflow.

`Wails` connects Go with a desktop window using the system webview, so the app stays lighter than Electron-style stacks.

`HTML/CSS/JS` keeps the frontend simple. You can start with plain files, then move to React/Vue/Svelte later without rewriting the Go side.

## Requirements

Official Wails docs currently list these base requirements:

- Go `1.21+`
- Node/NPM
- Wails CLI

Install the CLI:

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Platform notes

### macOS

- Install Xcode Command Line Tools: `xcode-select --install`

### Windows

- Install WebView2 Runtime
- Use `wails doctor` to verify the environment

### Linux

- Install `gcc`, `libgtk-3`, and WebKitGTK development packages
- On newer distros, Wails may require the `webkit2_41` build tag
- Use `wails doctor` to see distro-specific advice

## Run in development mode

```bash
wails dev
```

This gives you hot reload for the frontend and a live desktop window.

## Build the app

```bash
wails build
```

The packaged app will be created in the Wails build output directories for your current platform.

You can also use the helper scripts in this repository:

```bash
make build
```

On Windows:

```bat
build.bat
```

## Testing on macOS, Windows, and Linux

For the most reliable results, build and run the project on each target OS separately:

1. Clone or copy the project to that machine
2. Install the local Wails dependencies for that OS
3. Run `wails doctor`
4. Run `wails dev` for an interactive test
5. Run `wails build` for a packaged artifact

That means:

- on macOS you should test on macOS
- on Windows you should test on Windows
- on Linux you should test on Linux

You can also automate this later with CI runners for each platform.

## Project structure

- `main.go` boots the Wails app window
- `app.go` contains Go methods exposed to JavaScript
- `frontend/src/main.js` renders the UI and calls Go methods
- `frontend/src/*.css` styles the window
- `wails.json` defines Wails build commands

## Useful checks

```bash
wails doctor
wails build -clean
make doctor
```

## Official docs

- Installation: https://wails.io/docs/gettingstarted/installation/
- First project: https://wails.io/docs/gettingstarted/firstproject/
- Build config: https://wails.io/docs/reference/project-config/
