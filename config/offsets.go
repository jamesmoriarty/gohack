package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type Offsets struct {
	Timestamp  string `yaml:"timestamp"`
	Signatures struct {
		OffsetLocalPlayer int `yaml:"dwLocalPlayer"`
		OffsetForceJump   int `yaml:"dwForceJump"`
	} `yaml:"signatures"`
	Netvars struct {
		OffsetLocalPlayerFlags int `yaml:"m_fFlags"`
	} `yaml:"netvars"`
}

func GetLatestOffsets(url string) (*Offsets, error) {
	var offsets Offsets

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("Failed getting offsets")
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	err = yaml.Unmarshal(bytes, &offsets)

	return &offsets, nil
}
