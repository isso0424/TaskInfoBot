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
	Major map[string]string `json:"major"`
	Minor map[string]string `json:"minor"`
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
