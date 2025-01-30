package huecontroller

func GetAllLights(config BridgeConfig) []LightProps {
	lightRecources := ParseDeviceResource(GetDeviceResource(config))

	lights := []LightProps{}

	for _, rid := range lightRecources {
		lights = append(lights, ParseLightResource(GetLightResource(config, rid)))
	}

	return lights
}
