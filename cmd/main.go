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

func main() {
	listLights := flag.Bool("list", false, "Lists all registered lights.")
	createConfig := flag.Bool("register", false, "Creates a config and registers a new HUE api key.")
	lightId := flag.String("light", "", "ID of the light to control.")
	brightness := flag.Float64("br", -1.0, "Controls the brightness of the given light. [0 - 100]")
	colorX := flag.Float64("colorx", -1.0, "Controls the X Coordinate in the color diagram as a floating point number.")
	colorY := flag.Float64("colory", -1.0, "Controls the Y Coordinate in the color diagram as a floating point number.")
	colorRGB := flag.String("rgb", "", "Controls the color of the lamp in RGB format (Format: #00ff99)")

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
		for _, light := range hc.GetAllLights(config) {
			fmt.Printf("%s:\n%s", light.Name, light.String())
		}

		return
	}

	if len(*lightId) == 0 {
		flag.Usage()
		return
	}

	hc.SetLight(config, hc.LightProps{
		Id:         *lightId,
		Brightness: *brightness,
		ColorX:     *colorX,
		ColorY:     *colorY,
		ColorRgb:   *colorRGB,
	})
}
