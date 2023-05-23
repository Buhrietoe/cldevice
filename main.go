package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"

	"github.com/Buhrietoe/go-opencl/cl"
	"github.com/xlab/tablewriter"
)

func main() {
	platforms, _ := cl.GetPlatforms()

	if len(platforms) == 0 {
		fmt.Printf("== Found %d OpenCL platforms ==\n", len(platforms))
		os.Exit(1)
	}

	for platformNum, platform := range platforms {
		devices, _ := platform.GetDevices(cl.DeviceTypeAll)

		for deviceNum, device := range devices {
			table := tablewriter.CreateTable()
			table.UTF8Box()
			table.AddTitle(strconv.Itoa(platformNum) + ": " + platform.Name() + " - " + strconv.Itoa(deviceNum) + ": " + device.Name())
			table.AddRow("Device Type", device.Type())
			table.AddRow("Supported OpenCL Version", device.OpenCLCVersion())
			table.AddRow("Profile", device.Profile())
			table.AddRow("Driver Version", device.DriverVersion())
			table.AddRow("Number of Compute Units", strconv.Itoa(device.MaxComputeUnits()))
			table.AddRow("Clock Frequency", strconv.Itoa(device.MaxClockFrequency())+" Mhz")
			table.AddRow("Global Mem Size", ByteCountIEC(device.GlobalMemSize()))
			table.AddRow("Allocatable Mem Size", ByteCountIEC(device.MaxMemAllocSize()))
			table.AddRow("Is Host Unified Memory", device.HostUnifiedMemory())
			fmt.Println("\n" + table.Render())
		}
	}
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
