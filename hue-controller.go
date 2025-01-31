package huecontroller

import (
	"fmt"
	"math"
	"strconv"
)

func gammaCorrection(normalizedColor float64) float64 {
	if normalizedColor > 0.04045 {
		return math.Pow((normalizedColor+0.055)/1.055, 2.4)
	}

	return normalizedColor / 12.92
}

func reverseGammaCorrection(rgbValue float64) float64 {
	if rgbValue <= 0.0031308 {
		return rgbValue * 12.92
	}

	return 1.055*math.Pow(rgbValue, 1.0/2.4) - 0.055
}

// converts #affa00 -> 0.4, 0.5 hue coordinates
func rgbToHueColor(color string) (float64, float64) {
	if len(color) != 7 {
		return -1.0, -1.0
	}

	r, _ := strconv.ParseInt(color[1:3], 16, 32)
	g, _ := strconv.ParseInt(color[3:5], 16, 32)
	b, _ := strconv.ParseInt(color[5:7], 16, 32)

	rFinal := gammaCorrection(float64(r) / 255.0)
	gFinal := gammaCorrection(float64(g) / 255.0)
	bFinal := gammaCorrection(float64(b) / 255.0)

	x := rFinal*0.649926 + gFinal*0.103455 + bFinal*0.197109
	y := rFinal*0.234327 + gFinal*0.743075 + bFinal*0.022598
	z := rFinal*0.000000 + gFinal*0.053077 + bFinal*1.035763

	sum := x + y + z
	if sum == 0 {
		return -1.0, -1.0
	}

	return x / sum, y / sum
}

// converts hue coordinates 0.4, 0.5 -> #aaff00
func hueColorToRgb(x, y, br float64) string {
	if x == 0 && y == 0 {
		return "#FFFFFF"
	}

	z := 1.0 - x - y

	yCoord := br
	xCoord := (yCoord / y) * x
	zCoord := (yCoord / y) * z

	r := xCoord*1.656492 - yCoord*0.354851 - zCoord*0.255038
	g := -xCoord*0.707196 + yCoord*1.655397 + zCoord*0.036152
	b := xCoord*0.051713 - yCoord*0.121364 + zCoord*1.01153

	if r > b && r > g && r > 1.0 {
		// red is too big
		g = g / r
		b = b / r
		r = 1.0
	} else if g > b && g > r && g > 1.0 {
		// green is too big
		r = r / g
		b = b / g
		g = 1.0
	} else if b > r && b > g && b > 1.0 {
		// blue is too big
		r = r / b
		g = g / b
		b = 1.0
	}

	r = reverseGammaCorrection(r)
	g = reverseGammaCorrection(g)
	b = reverseGammaCorrection(b)

	if r > b && r > g {
		// red is biggest
		if r > 1.0 {
			g = g / r
			b = b / r
			r = 1.0
		}
	} else if g > b && g > r {
		// green is biggest
		if g > 1.0 {
			r = r / g
			b = b / g
			g = 1.0
		}
	} else if b > r && b > g {
		// blue is biggest
		if b > 1.0 {
			r = r / b
			g = g / b
			b = 1.0
		}
	}

	return fmt.Sprintf("#%X%X%X",
		int32(math.Round(r*255)),
		int32(math.Round(g*255)),
		int32(math.Round(b*255)),
	)
}

func GetAllLights(config BridgeConfig) []LightProps {
	lightRecources := ParseDeviceResource(GetDeviceResource(config))

	lights := []LightProps{}

	for _, rid := range lightRecources {
		lights = append(lights, ParseLightResource(GetLightResource(config, rid)))
	}

	return lights
}

func setNewLightProps(currentProps LightProps, newProps *LightProps) {
	if newProps.Brightness == -1.0 {
		newProps.Brightness = currentProps.Brightness
	}

	if newProps.Brightness >= 5.0 {
		newProps.On = true
	}

	// either set the color from the rgb value, or the x y coordinates if both are set
	if len(newProps.ColorRgb) != 0 {
		newProps.ColorX, newProps.ColorY = rgbToHueColor(newProps.ColorRgb)
	} else if newProps.ColorX == -1.0 || newProps.ColorY == -1.0 {
		newProps.ColorX = currentProps.ColorX
		newProps.ColorY = currentProps.ColorY
	}
}

func SetLight(config BridgeConfig, lightProps LightProps) {
	currentLightProps := ParseLightResource(GetLightResource(config, lightProps.Id))
	setNewLightProps(currentLightProps, &lightProps)
	PutLightResource(config, lightProps)
}
