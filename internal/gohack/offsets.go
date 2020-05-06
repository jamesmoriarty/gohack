package gohack

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

const (
	OffsetsURL = "https://raw.githubusercontent.com/frk1/hazedumper/master/csgo.yaml"
)

type Offsets struct {
	Timestamp  string `yaml:"timestamp"`
	Signatures struct {
		OffsetdwLocalPlayer uintptr `yaml:"dwLocalPlayer"`
		OffsetdwForceJump uintptr `yaml:"dwForceJump"`
	} `yaml:"signatures"`
	Netvars struct {
		Offsetm_fFlags uintptr `yaml:"m_fFlags"`
	} `yaml:"netvars"`
}

func GetOffsets() (*Offsets, error) {
	var offsets Offsets

	resp, err := http.Get(OffsetsURL)

	if err != nil {
		return nil, errors.New("Failed making offsets request")
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("Failed reading offsets request")
	}

	err = yaml.Unmarshal(bytes, &offsets)

	if err != nil {
		return nil, errors.New("Failed parsing offsets request")
	}

	return &offsets, nil
}
