package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/edge-experiments/wrap4kyst/ocm"
	"github.com/edge-experiments/wrap4kyst/util"
)

// hard coded the `configSpec` API to keep simple, instead of importing https://github.com/edge-experiments/kyst/tree/main/api/v1alpha1

func main() {
	target := flag.String("target", "k8s", "the target to which the workloads should be deployed") // k8s, ocm, flotta
	emptyConfigSpecFN := flag.String("empty-configspec", "./configspec-empty.yaml", "the configspec file whose content is to be populated")
	configSpecFN := flag.String("configspec", "./configspec.yaml", "the configspec file with content populated")
	manifestDir := flag.String("manifest-dir", "../manifests/", "the directory containing the to-be-wrapped manifests")
	flag.Parse()

	if *target != "k8s" && *target != "ocm" && *target != "flotta" {
		log.Fatalf("invalid target %s", *target)
	}
	log.Println("target:", *target)

	if *target == "ocm" {
		ocm.WrapConfigSpec(*emptyConfigSpecFN, *configSpecFN, *manifestDir)
		return
	}

	dataStructure, err := util.ReadEmtpyConfigSpec(*emptyConfigSpecFN)
	if err != nil {
		log.Fatalf("error reading empty configspec: %v", err)
	}

	content := []string{}
	items, _ := ioutil.ReadDir(*manifestDir)
	for _, item := range items {
		if item.IsDir() {
			log.Printf("Found directory: %s, skipping", item.Name())
		} else {
			log.Printf("Found file: %s, adding to content", item.Name())
			bytes, err := ioutil.ReadFile(*manifestDir + item.Name())
			if err != nil {
				log.Fatal(err)
			}
			content = append(content, string(bytes))
		}
	}
	dataStructure["spec"].(map[string]interface{})["content"] = content

	err = util.WriteConfigSpec(*configSpecFN, dataStructure)
	if err != nil {
		log.Fatalf("error writing configspec: %v", err)
	}
}
