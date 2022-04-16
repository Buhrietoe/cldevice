package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/Buhrietoe/go-opencl/cl"
	"github.com/xlab/tablewriter"
)

func main() {
	platforms, err := cl.GetPlatforms()
	if err != nil {
		fmt.Println(err)
	}

	if len(platforms) == 0 {
		panic("GetPlatforms returned 0 devices")
	}

	devices, err := platforms[0].GetDevices(cl.DeviceTypeAll)
	if err != nil {
		fmt.Println(err)
	}

	if len(devices) == 0 {
		panic("GetDevices returned 0 devices")
	}
	device := devices[0]

	table := tablewriter.CreateTable()
	table.UTF8Box()
	table.AddTitle("OPENCL DEVICE")
	table.AddRow("Physical Device", device.Name())
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
