package main

// add icon
// go-winres simply --icon icon.ico
// go-winres simply --icon icon.ico --manifest gui
// go build -ldflags="-H windowsgui"
import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/webview/webview"
)

func show_output(r io.Reader) {
	reader := bufio.NewReader(r)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Println(line)
		line, err = reader.ReadString('\n')
		if strings.Contains(line, "localhost") {
			start_webview()
			break
		}

	}
	// scanner := bufio.NewScanner(r)
	// for scanner.Scan() {
	// 	if strings.Contains(scanner.Text(), "localhost") {
	// 		go start_webview()
	// 		break
	// 	}
	// fmt.Println(scanner.Text())
	// }
}

func start_webview() {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Property Management")
	w.SetSize(900, 620, webview.HintNone)
	w.Navigate("http://localhost:5000/")
	w.Run()
}
func main() {
	cmd := exec.Command("pythonw.exe", "main.py")

	stdout, err := cmd.StdoutPipe()
	// cmd.Start()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go show_output(stdout)
	cmd.Wait()
}
