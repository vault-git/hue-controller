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
	lights := ParseDeviceResource(GetDeviceResource(config))

	for _, rid := range lights {
		light := ParseLightResource(GetLightResource(config, rid))

		if light.Name == lightProps.Name {
			currentLightProps := ParseLightResource(GetLightResource(config, rid))

			setNewLightProps(currentLightProps, &lightProps)

			PutLightResource(config, rid, lightProps)

			return
		}
	}
}
