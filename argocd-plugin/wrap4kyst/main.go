package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// hard coded to keep simple, instead of importing https://github.com/edge-experiments/kyst/tree/main/api/v1alpha1

var (
	emptyConfigSpecFN = "./configspec-empty.yaml"
	configSpecFN      = "./configspec.yaml"
	manifestsDir      = "../manifests/"
)

func main() {
	// read empty configSpec
	bytes, err := ioutil.ReadFile(emptyConfigSpecFN)
	if err != nil {
		log.Fatal(err)
	}
	dataStructure := make(map[string]interface{})
	error := yaml.Unmarshal(bytes, &dataStructure)
	if error != nil {
		log.Fatal(err)
	}

	// build content for configSpec
	content := []string{}
	items, _ := ioutil.ReadDir(manifestsDir)
	for _, item := range items {
		if item.IsDir() {
			log.Printf("Found directory: %s, skipping", item.Name())
		} else {
			log.Printf("Found file: %s, adding to content", item.Name())
			bytes, err := ioutil.ReadFile(manifestsDir + item.Name())
			if err != nil {
				log.Fatal(err)
			}
			content = append(content, string(bytes))
		}
	}
	dataStructure["spec"].(map[string]interface{})["content"] = content

	// write configSpec
	bytes, err = yaml.Marshal(dataStructure)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(configSpecFN, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
