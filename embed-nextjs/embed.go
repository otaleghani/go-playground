package main

import (
    "embed"
    "fmt"
    "io/fs"
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
)

var content embed.FS

func embedFiles() {
    // Create a temporary directory
    tempDir, err := ioutil.TempDir("", "nextjs-app")
    if err != nil {
        fmt.Println("Error creating temp directory:", err)
        return
    }
    defer os.RemoveAll(tempDir) // Clean up

    fmt.Println("Temporary directory created at:", tempDir)

    // Extract embedded files to the temporary directory
    if err := extractFiles(tempDir); err != nil {
        fmt.Println("Error extracting files:", err)
        return
    }

    // Change to the temporary directory
    if err := os.Chdir(tempDir); err != nil {
        fmt.Println("Error changing directory:", err)
        return
    }

    // Run 'npm install'
    if err := runCommand("npm", "install"); err != nil {
        fmt.Println("Error running npm install:", err)
        return
    }

    // Run 'npx next start'
    if err := runCommand("npx", "next", "start"); err != nil {
        fmt.Println("Error running npx next start:", err)
        return
    }
}

func extractFiles(targetDir string) error {
    return fs.WalkDir(content, ".", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }

        data, err := content.ReadFile(path)
        if err != nil {
            return err
        }

        // Remove 'ui/' prefix from the path
        relativePath := path[len("ui/"):]

        targetPath := filepath.Join(targetDir, relativePath)
        os.MkdirAll(filepath.Dir(targetPath), 0755)
        return ioutil.WriteFile(targetPath, data, 0644)
    })
}

func runCommand(name string, arg ...string) error {
    cmd := exec.Command(name, arg...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
