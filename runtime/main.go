package main

import (
	"fmt"
	"github.com/JustinTimperio/osinfo"
	"github.com/jaypipes/ghw"
)

// functions
func getSystemInfo() (string, error) {
	info := ""
	cpu, err := ghw.CPU()
	if err != nil {
		return info, err
	}

	memory, err := ghw.Memory()
	if err != nil {
		return info, err
	}

	block, err := ghw.Block()
	if err != nil {
		return info, err
	}

	gpu, err := ghw.GPU()
	if err != nil {
		return info, err
	}

	for _, proc := range cpu.Processors {
		info += "CPU: " + proc.Model + "\n"
	}

	for _, vc := range gpu.GraphicsCards {
		info += "GPU: " + vc.DeviceInfo.Product.Name + "\n"
	}

	release := osinfo.GetVersion()
	info += release.String() + "\n"
	info += memory.String() + "\n"
	info += block.String() + "\n"
	total_threads := string(cpu.TotalThreads)
	info += total_threads + "\n"
	return info, nil

}

func main() {

	if sysinfo, err := getSystemInfo(); err == nil {
		fmt.Printf(sysinfo)
	}

}
