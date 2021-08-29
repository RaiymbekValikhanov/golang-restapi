package main

import (
	"io/ioutil"
	"log"

	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/apiserver"
	"gopkg.in/yaml.v2"
)

var (
	configPath = "./configs/apiserver.yaml"
)

func main() {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	config := apiserver.NewConfig()
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}