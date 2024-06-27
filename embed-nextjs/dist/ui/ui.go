package ui

import (
  "embed"
  "io/fs"
)

//go:embed .next public package.json
var content embed.FS

struct FileContent {
  build: fs.FS,
  public: fs.FS,
  package: fs.FS,
}

func Ui() (fs.FS[], error) {
  nextFS, err := fs.Sub(content, ".next")
  if err != nil {
    return nil, err
  }

  publicFS, err := fs.Sub(content, "public")
  if err != nil {
    return nil, err
  }

  packageFS, err := fs.Sub(content, "package.json")
  if err != nil {
    return nil, err
  }

  return FileContent{
    build: nextFS,
    public: publicFS,
    package: packageFS,
  }, nil
}
