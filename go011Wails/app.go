package main

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name.
func (a *App) Greet(name string) string {
	cleanName := strings.TrimSpace(name)
	if cleanName == "" {
		cleanName = "friend"
	}

	return fmt.Sprintf(
		"Hello, %s! Go says hi from %s at %s.",
		cleanName,
		runtime.GOOS,
		time.Now().Format("15:04:05"),
	)
}

// SystemInfo returns a short runtime summary that we can show in the UI.
func (a *App) SystemInfo() string {
	return fmt.Sprintf(
		"Platform: %s/%s | Go: %s",
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
	)
}
