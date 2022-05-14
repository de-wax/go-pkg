/*
Mit diesem Paket lässt sich der Taupunkt aus Temperatur und Luftfeuchte berechnen.
Es wird die die Magnus-Tetens-Formel (Sonntag90) benutzt die relativ genaue Ergebnisse (mit einer Unsicherheit von 0,35 °C) für Temperaturen von -45 °C bis 60 °C berechnet.
--
With this package the dew point can be calculated from temperature and humidity.
It uses the Magnus-Tetens formula (Sonntag90) which gives relatively accurate results (with an uncertainty of 0.35 °C) for temperatures from -45 °C to 60 °C.
*/
package dewpoint

import (
	// Ein paar Standard-Pakete werden benötigt
	// A few standard packages are required
	//"fmt"
	"errors"
	"math"	
)
func Calculate(T float64, H float64) (float64, error) {
	// Überprüfen ob der übergebene Wert für die Temperatur im gültigen Bereich liegt
	// Check if the transferred value for the temperature is within the valid range
	if T < -45 || T > 60 {
		return 0, errors.New("Temperatur nicht im gültigen Bereich (-45 - +60°C)")
	} else {
		// Überprüfen ob der übergebene Wert für die Feuchtigkeit im gültigen Bereich liegt
		// Check if the transferred value for humidity is within the valid range
		if H < 0 || H > 100 {
			return 0, errors.New("Feuchtigkeit nicht im gültigen Bereich (0 - 100%)")
		} else {
			// Konstanten für die Magnus Formel
			// Constants for the Magnus formula
			var a,b float64 = 0,0
			if T > 0 {
				a = 7.5
				b = 237.3
			} else {
				a = 7.6
				b = 240.7
			}
			// Magnus Formel
			// Magnus formula
			dd := math.Log10(math.Pow(10, ( a * T) / ( b + T)) * (H/100))
			return (b*dd)/(a-dd), nil
		}
	}
}
