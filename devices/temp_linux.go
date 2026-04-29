//go:build linux
// +build linux

package devices

import (
	"log"

	"github.com/shirou/gopsutil/v4/sensors"
)

// All possible thermometers
func devs() []string {
	if sensorMap == nil {
		sensorMap = make(map[string]string)
	}
	sensors, err := sensors.SensorsTemperatures()
	if err != nil {
		log.Printf("gopsutil reports %s", err)
		if len(sensors) == 0 {
			log.Printf("no temperature sensors returned")
			return []string{}
		}
	}
	rv := make([]string, 0, len(sensors))
	for _, sensor := range sensors {
		rv = append(rv, sensor.SensorKey)
		sensorMap[sensor.SensorKey] = sensor.SensorKey
	}
	return rv
}

func defs() []string {
	// MUST be called AFTER init()
	rv := make([]string, 0, len(sensorMap))
	for _, v := range sensorMap {
		rv = append(rv, v)
	}
	return rv
}
