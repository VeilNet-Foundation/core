// core/optimization/adaptive.go
package optimization

import (
	"runtime"
	"time"
)

type Profile struct {
	Name      string
	Bandwidth int     // Кбит/с
	CPU       float64 // % загрузки
	Security  int     // 1–10
}

var Profiles = map[string]Profile{
	"light":  {Name: "light",  Bandwidth: 100,  CPU: 0.3, Security: 5},
	"normal": {Name: "normal", Bandwidth: 500,  CPU: 0.6, Security: 7},
	"high":   {Name: "high",   Bandwidth: 2000, CPU: 0.9, Security: 10},
}

func AutoTune() Profile {
	cpu := runtime.NumCPU()
	ram := getRAM()
	net := measureBandwidth()

	if cpu < 2 || ram < 2 {
		return Profiles["light"]
	}
	if net < 100 {
		return Profiles["light"]
	}
	return Profiles["high"]
}