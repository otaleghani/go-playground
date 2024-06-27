package main

import (
	"fmt"
	"os"
	"os/exec"
)

func buildAndStart() {
	cmd := exec.Command("npm", "run", "build")
	cmd.Dir = "./ui"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	cmd2 := exec.Command("npm", "run", "start")
	cmd2.Dir = "./ui"
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	err = cmd2.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func build() {
	cmd := exec.Command("npm", "run", "build")
	cmd.Dir = "./ui"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func start() {
	cmd2 := exec.Command("npm", "run", "start")
	cmd2.Dir = "./ui"
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
  err := cmd2.Run()
	if err != nil {
		fmt.Println(err)
	}
}
