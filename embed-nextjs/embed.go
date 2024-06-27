package main

import (
    "embed"
    "io/fs"
    "io/ioutil"
    "os"
    "path/filepath"
    "os/exec"
)

//go:embed ui/*
var uiFiles embed.FS

func embedFiles() {
    // Create a temporary directory
    tmpDir, err := ioutil.TempDir("", "nextjs")
    if err != nil {
        panic(err)
    }
    defer os.RemoveAll(tmpDir) // Clean up

    // Write the embedded files to the temporary directory
    fs.WalkDir(uiFiles, "ui", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        if d.IsDir() {
            return nil
        }

        data, err := uiFiles.ReadFile(path)
        if err != nil {
            return err
        }

        destPath := filepath.Join(tmpDir, path)
        os.MkdirAll(filepath.Dir(destPath), 0755)

        return ioutil.WriteFile(destPath, data, 0644)
    })

    // Run `npx next start` in the temporary directory
    cmd := exec.Command("npm", "run", "start")
    cmd.Dir = filepath.Join(tmpDir, "ui")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        panic(err)
    }
}
