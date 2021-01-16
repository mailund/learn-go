// Package tempconv converts between Celsius and Farenheit
package tempconv

import "fmt"

// Celsius Temperature scale
type Celsius float64

// Farenheit Temperature scale
type Farenheit float64

const (
	// AbsoluteZeroC Absolute zero temperature in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC Freezing point of water
	FreezingC Celsius = 0
	// BoilingC Boiling point of water
	BoilingC Celsius = 100
)

func (c Celsius) String() string   { return fmt.Sprintf("%g℃", c) }
func (f Farenheit) String() string { return fmt.Sprintf("%g℉", f) }
