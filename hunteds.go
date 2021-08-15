package main

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Hunteds struct {
	hunteds []string `yaml:"hunteds"`
}

func (h *Hunteds) NewHunteds() {
	filename, _ := filepath.Abs("./hunteds.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &h.hunteds)
	if err != nil {
		panic(err)
	}
}
