package util

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func ReadEmtpyConfigSpec(filename string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var dataStructure map[string]interface{}
	err = yaml.Unmarshal(bytes, &dataStructure)
	if err != nil {
		return nil, err
	}
	return dataStructure, nil
}

func WriteConfigSpec(filename string, dataStructure map[string]interface{}) error {
	bytes, err := yaml.Marshal(dataStructure)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ComposeRawContent(dir string) [][]byte {
	rawContent := [][]byte{}
	items, _ := ioutil.ReadDir(dir)
	for _, item := range items {
		if item.IsDir() {
			log.Printf("Found directory: %s, skipping", item.Name())
		} else {
			log.Printf("Found file: %s, adding to raw content", item.Name())
			bytes, err := ioutil.ReadFile(dir + item.Name())
			if err != nil {
				log.Fatal(err)
			}
			rawContent = append(rawContent, bytes)
		}
	}
	return rawContent
}
