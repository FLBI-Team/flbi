package core

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// MeterReading represents single reading
type MeterReading struct {
	MeterID     string
	Timestamp   string
	Current     string
	Voltage     string
	Power       string
	PowerFactor string
	Frequency   string
	Energy      string
}

// CreateMeters creates a list of random meter IDs
func CreateMeters(n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = CreateMeterID()
	}
	return result
}

// CreateMeterID creates a random meter ID
func CreateMeterID() string {
	a := strconv.Itoa(rand.Intn(10000))
	b := strconv.Itoa(rand.Intn(100))
	c := strconv.Itoa(rand.Intn(1000000))
	return "M-" + a + "-" + b + "-" + c
}

// ReadMeter reads meter values
func ReadMeter(meterID string) MeterReading {
	I := rand.Float64()
	V := 215 + 20*rand.Float64()
	PF := 0.9 + math.Exp(-rand.Float64())/10
	P := PF * I * V
	F := 45 + 10*rand.Float64()
	E := PF * I * V * 5

	return MeterReading{
		MeterID:     meterID,
		Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
		Current:     strconv.FormatFloat(I, 'f', 4, 64),
		Voltage:     strconv.FormatFloat(V, 'f', 4, 64),
		Power:       strconv.FormatFloat(P, 'f', 4, 64),
		PowerFactor: strconv.FormatFloat(PF, 'f', 4, 64),
		Frequency:   strconv.FormatFloat(F, 'f', 4, 64),
		Energy:      strconv.FormatFloat(E, 'f', 4, 64),
	}
}
