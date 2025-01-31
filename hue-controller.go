package huecontroller

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

	if newProps.ColorX == -1.0 {
		newProps.ColorX = currentProps.ColorX
	}

	if newProps.ColorY == -1.0 {
		newProps.ColorY = currentProps.ColorY
	}
}

func SetLight(config BridgeConfig, lightProps LightProps) {
	currentLightProps := ParseLightResource(GetLightResource(config, lightProps.Id))
	setNewLightProps(currentLightProps, &lightProps)
	PutLightResource(config, lightProps)
}
