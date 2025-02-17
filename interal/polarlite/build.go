package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var targets = []struct {
	OS   string
	Arch string
	Ext  string
}{
	{"windows", "amd64", ".exe"},
	{"linux", "amd64", ""},
	{"linux", "arm64", ""},
	{"darwin", "arm64", ""},
}

func main() {
	// Get version from Git
	version, err := getGitVersion()
	if err != nil {
		fmt.Println("⚠️  Failed to retrieve version:", err)
		version = "dev" // Fallback version
	}

	fmt.Println("🔧 Building version:", version)

	// Create output directory
	outDir := "./dist"
	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Println("⚠️  Failed to create output directory:", err)
		return
	}

	// Build for each target
	for _, target := range targets {
		buildTarget(version, target.OS, target.Arch, target.Ext, outDir)
	}

	fmt.Println("✅ Build completed!")
}

// Get Git version (tag or commit hash)
func getGitVersion() (string, error) {
	cmd := exec.Command("git", "describe", "--tags", "--always", "--dirty")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// Build for a specific OS/Architecture
func buildTarget(version, goos, goarch, ext, outDir string) {
	fmt.Printf("🚀 Building for %s/%s...\n", goos, goarch)

	// Set environment variables for cross-compilation
	env := os.Environ()
	env = append(env, "GOOS="+goos, "GOARCH="+goarch)

	// Generate output filename
	outputName := filepath.Join(outDir, goos, goarch, fmt.Sprintf("polarlite%s", ext))

	// Run build command
	cmd := exec.Command("go", "build", "-ldflags", "-X 'main.version="+version+"'", "-o", outputName, "./cmd/polarlite/main.go")
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute build
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Failed to build for %s/%s: %v\n", goos, goarch, err)
	} else {
		fmt.Printf("✅ Successfully built: %s\n", outputName)
	}
}
