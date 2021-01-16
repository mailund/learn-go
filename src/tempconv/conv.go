package tempconv

// CToF Converts Celsius to Farenheit
func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }

// FToC Convers Farenheit to Celsius
func FToC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
