package huecontroller

import (
	"encoding/json"
	"fmt"
)

type LightProps struct {
	Id         string
	Name       string
	On         bool
	Brightness float64
	ColorX     float64
	ColorY     float64
	ColorRgb   string
}

func (lp *LightProps) String() string {
	state := func() string {
		if lp.On {
			return "On"
		}

		return "Off"
	}()

	return fmt.Sprintf("\tId: %s,\n\tState: %s,\n\tBrightness: %0.1f,\n\tColor: %s\n", lp.Id, state, lp.Brightness, lp.ColorRgb)
}

func ParseLightResource(buf []byte) LightProps {
	type MetaData struct {
		Name string
	}

	type On struct {
		On bool
	}

	type Dimming struct {
		Brightness float64
	}

	type Xy struct {
		X float64
		Y float64
	}

	type Color struct {
		Xy Xy
	}

	type Data struct {
		Id       string
		MetaData MetaData
		Dimming  Dimming
		On       On
		Color    Color
	}

	type Result struct {
		Errors []string
		Data   []Data
	}

	var res = Result{}

	err := json.Unmarshal(buf, &res)
	if err != nil {
		fmt.Println(err)
		return LightProps{}
	}

	lightData := res.Data[0]

	rgbColor := hueColorToRgb(lightData.Color.Xy.X, lightData.Color.Xy.Y, lightData.Dimming.Brightness)

	return LightProps{
		Id:         lightData.Id,
		Name:       lightData.MetaData.Name,
		On:         lightData.On.On,
		Brightness: lightData.Dimming.Brightness,
		ColorX:     lightData.Color.Xy.X,
		ColorY:     lightData.Color.Xy.Y,
		ColorRgb:   rgbColor,
	}
}

func ParseDeviceResource(buf []byte) []string {
	type Service struct {
		Rid   string
		Rtype string
	}

	type Data struct {
		Services []Service
		Metadata struct {
			Name string
		}
	}

	type Result struct {
		Errors []string
		Data   []Data
	}

	var res = Result{}

	err := json.Unmarshal(buf, &res)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	lights := []string{}

	for _, light := range res.Data {
		for _, service := range light.Services {
			if service.Rtype == "light" {
				lights = append(lights, service.Rid)
			}
		}
	}

	return lights
}

func IsLinkButtonResponse(resp []byte) bool {
	type ErrorData struct {
		Type        int
		Address     string
		Description string
	}

	type Error struct {
		Error ErrorData
	}

	var error []Error
	err := json.Unmarshal(resp, &error)
	if err != nil {
		return false
	}

	if error[0].Error.Type == 101 || error[0].Error.Description == "link button not pressed" {
		return true
	}

	return false
}

func ParseNewUserResult(resp []byte) (string, error) {
	type SuccessData struct {
		Username  string
		Clientkey string
	}

	type Success struct {
		Success SuccessData
	}

	var success []Success
	err := json.Unmarshal(resp, &success)
	if err != nil {
		return "", err
	}

	return success[0].Success.Username, nil
}
