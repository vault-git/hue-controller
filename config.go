package huecontroller

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

const CONFIG_FILE = "config.json"

type BridgeConfig struct {
	Ip     string `json:"ip"`
	ApiKey string `json:"apikey"`
}

func (config *BridgeConfig) Load() {
	f, err := os.Open(CONFIG_FILE)
	if err != nil {
		log.Fatal("error opening config file: ", err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	content := ""

	for scanner.Scan() {
		content += scanner.Text()
	}

	err = json.Unmarshal([]byte(content), config)
	if err != nil {
		log.Fatal("error parsing config file: ", err.Error())
	}
}

func (config *BridgeConfig) Save() {
	f, err := os.Create(CONFIG_FILE)
	if err != nil {
		log.Fatal("error creating config file: ", err.Error())
	}
	defer f.Close()

	configBytes, err := json.Marshal(*config)
	if err != nil {
		log.Fatal("error saving config file: ", err.Error())
	}

	f.Write(configBytes)
}
