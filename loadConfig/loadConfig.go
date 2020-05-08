package loadConfig

import (
	"encoding/json"
	"io/ioutil"
)

// Config is a type that root of config
type Config struct {
	Courses  []Course `json:"courses"`
	Channels Channels `json:"channels"`
}

// Course is a type that root of couses
type Course struct {
	Alias    string  `json:"alias"`
	Subjects subject `json:"subjects"`
}

type subject struct {
	Major []string `json:"major"`
	Minor []string `json:"minor"`
}

// Channels is a type that channels setting
type Channels struct {
	Regist string `json:"regist"`
	Notify Notify `json:"notify"`
}

// Notify is a type that notify channels
type Notify struct {
	Major map[string]string `json:"major"`
	Minor map[string]string `json:"minor"`
}

// LoadConfig is a function that load config from config.json
func LoadConfig() Config {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var config Config

	json.Unmarshal(raw, &config)

	return config
}
