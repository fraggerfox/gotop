//go:build !linux
// +build !linux

package devices

import "github.com/shirou/gopsutil/v4/cpu"

func CpuCount() (int, error) {
	return cpu.Counts(false)
}
