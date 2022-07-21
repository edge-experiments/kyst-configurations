package util

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

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

func WriteConfigSpecWithNameSuffix(original string, suffix string, dataStructure map[string]interface{}) error {
	filename := original
	if suffix != "" {
		filename = filename[0:len(filename)-5] + "-" + suffix + ".yaml" // this is a bit of a hack, but it works when we have a ".yaml" in the original filename
	}
	err := WriteConfigSpec(filename, dataStructure)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEmptyConfigSpec(filename string) error {
	err := os.Remove(filename)
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

func appendTimestamp(prefix string) string {
	suffix := strconv.FormatInt(time.Now().UnixNano(), 10)
	return prefix + "-" + suffix
}

func AppendTimestampToConfigSpecName(configSpec map[string]interface{}) map[string]interface{} {
	name := configSpec["metadata"].(map[string]interface{})["name"].(string)
	name = appendTimestamp(name)
	configSpec["metadata"].(map[string]interface{})["name"] = name
	return configSpec
}

func AppendSuffixToConfigSpecName(configSpec map[string]interface{}, suffix string) map[string]interface{} {
	name := configSpec["metadata"].(map[string]interface{})["name"].(string)
	name = name + "-" + suffix
	configSpec["metadata"].(map[string]interface{})["name"] = name
	return configSpec
}
