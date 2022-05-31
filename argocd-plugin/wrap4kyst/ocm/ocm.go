package ocm

import (
	"log"

	"github.com/edge-experiments/wrap4kyst/util"
)

func WrapConfigSpec(inputFile, outputFile, manifestDir string) {
	dataStructure, err := util.ReadEmtpyConfigSpec(inputFile)
	if err != nil {
		log.Fatalf("error reading empty configspec: %v", err)
	}

	err = util.WriteConfigSpec(outputFile, dataStructure)
	if err != nil {
		log.Fatalf("error writing configspec: %v", err)
	}
}
