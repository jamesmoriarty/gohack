package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://raw.githubusercontent.com/frk1/hazedumper/master/csgo.yaml"
)

type Offsets struct {
	Timestamp  string `yaml:"timestamp"`
	Signatures struct {
		OffsetLocalPlayer uintptr `yaml:"dwLocalPlayer"`
		OffsetForceJump   uintptr `yaml:"dwForceJump"`
	} `yaml:"signatures"`
	Netvars struct {
		OffsetLocalPlayerFlags uintptr `yaml:"m_fFlags"`
	} `yaml:"netvars"`
}

func GetOffsets() (*Offsets, error) {
	log.WithFields(log.Fields{"url": url}).Info("GetOffsets")

	var offsets Offsets

	resp, err := http.Get(url)

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
