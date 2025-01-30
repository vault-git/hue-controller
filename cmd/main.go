package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	hc "github.com/vault-git/hue-controller"
)

func registerApiKey(ip string) (string, error) {
	resp, err := hc.PostNewClient(ip)
	if err != nil {
		return "", err
	}

	if resp != nil && hc.IsLinkButtonResponse(resp) {
		log.Println("please press the link button on the hue bridge, then press any button...")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		resp, err := hc.PostNewClient(ip)
		if err != nil {
			return "", err
		}

		return hc.ParseNewUserResult(resp)
	}

	return "", errors.New("error registering new api key")
}

func createNewConfig(config *hc.BridgeConfig) {
	fmt.Print("input hue bridge ip: ")

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	ip := net.ParseIP(stdin.Text())
	if ip == nil {
		log.Fatal("ip out of range")
	}

	config.Ip = ip.To4().String()

	apiKey, err := registerApiKey(config.Ip)
	if err != nil {
		log.Fatal(err)
	}

	config.ApiKey = apiKey
}

func setNewLightProps(currentProps hc.LightProps, newProps *hc.LightProps) {
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

func printAllLights(config hc.BridgeConfig) {
	lights := hc.ParseDeviceResource(hc.GetDeviceResource(config))

	for _, rid := range lights {
		light := hc.ParseLightResource(hc.GetLightResource(config, rid))

		fmt.Printf("%s: %s", light.Name, light.String())
	}
}

func controlLights(config hc.BridgeConfig, lightName string, newLightProps hc.LightProps) {
	lights := hc.ParseDeviceResource(hc.GetDeviceResource(config))

	for _, rid := range lights {
		light := hc.ParseLightResource(hc.GetLightResource(config, rid))

		if light.Name == lightName {
			currentLightProps := hc.ParseLightResource(hc.GetLightResource(config, rid))

			setNewLightProps(currentLightProps, &newLightProps)

			hc.PutLightResource(config, rid, newLightProps)

			return
		}
	}

	log.Fatalf("light name \"%v\" not registered", lightName)
}

func main() {
	listLights := flag.Bool("list", false, "Lists all registered lights.")
	createConfig := flag.Bool("register", false, "Creates a config and registers a new HUE api key.")
	lightName := flag.String("light", "", "Name of the light to control.")
	brightness := flag.Float64("br", -1.0, "Controls the brightness of the given light. [0 - 100]")
	colorX := flag.Float64("colorx", -1.0, "Controls the X Coordinate in the color diagram. [0.0 - 1.0]")
	colorY := flag.Float64("colory", -1.0, "Controls the Y Coordinate in the color diagram. [0.0 - 1.0]")

	flag.Parse()

	config := hc.BridgeConfig{}

	if *createConfig {
		createNewConfig(&config)
		config.Save()
		log.Println("created config file config.json")
		return
	}

	config.Load()

	if *listLights {
		printAllLights(config)
		return
	}

	if len(*lightName) == 0 {
		flag.Usage()
		return
	}

	controlLights(config, *lightName, hc.LightProps{
		On:         false,
		Brightness: *brightness,
		ColorX:     *colorX,
		ColorY:     *colorY,
	})
}
