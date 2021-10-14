package main

// go-winres simply --icon icon.ico --manifest gui
// go build -ldflags="-H windowsgui"

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("pythonw.exe", "main.py")
	cmd.Run()

}
