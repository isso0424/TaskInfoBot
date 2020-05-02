package loadConfig

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Courses  []Course `json:"courses"`
	Channels Channels `json:"channels"`
}

type Course struct {
	Name     string  `json:"name"`
	Alias    string  `json:"alias"`
	Subjects subject `json:"subjects"`
}

type subject struct {
	Major []string `json:"major"`
	Minor []string `json:"minor"`
}

type Channels struct {
	Regist string `json:"regist"`
	Notify Notify `json:"notify"`
}

type Notify struct {
	Major major `json:"major"`
	Minor minor `json:"minor"`
}

type major struct {
	General string `json:"General"`
	M       string `json:"M"`
	E       string `json:"E"`
	I       string `json:"I"`
	C       string `json:"C"`
}

type minor struct {
	M string `json:"M"`
	E string `json:"E"`
	I string `json:"I"`
	C string `json:"C"`
	G string `json:"G"`
}

func LoadConfig() Config {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var config Config

	json.Unmarshal(raw, &config)

	return config
}
