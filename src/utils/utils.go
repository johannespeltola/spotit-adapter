package utils

import (
	"entsoe/src/config"
	"fmt"
	"time"
)

func GetHour() int {
	return time.Now().Hour()
}

// Converts €/MWh to c/KWh and adds VAT (24 %)
func ConvertPrice(MWh float32) float32 {
	return (MWh / 10) * config.GetVat()
}

func GetEntsoeURL() string {
	time := time.Now().Format(config.GetEntsoeDateFormat()) + "0100"
	return config.GetEntsoeBase() + fmt.Sprintf("&periodStart=%v&periodEnd=%v", time, time)
}
