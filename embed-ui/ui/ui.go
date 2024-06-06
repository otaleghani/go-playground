package ui

import (
  "embed"
  "io/fs"
)

//go:embed dist/*
var content embed.FS

func Ui() (fs.FS, error) {
  distFS, err := fs.Sub(content, "dist")
  if err != nil {
    return nil, nil
  }
  return distFS, nil
}
