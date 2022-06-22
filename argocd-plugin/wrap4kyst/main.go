package main

import (
	"flag"
	"log"

	"github.com/edge-experiments/wrap4kyst/flotta"
	"github.com/edge-experiments/wrap4kyst/ocm"
	"github.com/edge-experiments/wrap4kyst/util"
)

// hard coded the `configSpec` API to keep simple, instead of importing https://github.com/edge-experiments/kyst/tree/main/api/v1alpha1

func main() {
	target := flag.String("target", "k8s", "the target to which the workloads should be deployed") // k8s, ocm, flotta
	emptyConfigSpecFN := flag.String("empty-configspec", "./configspec-empty.yaml", "the configspec file whose content is to be populated")
	configSpecFN := flag.String("configspec", "./configspec.yaml", "the configspec file with content populated")
	manifestDir := flag.String("manifest-dir", "../manifests/", "the directory containing the to-be-wrapped manifests")
	extraManifestDir := flag.String("extra-manifest-dir", "./extra-manifests/", "the directory containing target-specific manifests (Custom Resources)")
	flag.Parse()

	if *target != "k8s" && *target != "ocm" && *target != "flotta" {
		log.Fatalf("invalid target %s", *target)
	}
	log.Println("target:", *target)

	if *target == "ocm" {
		ocm.WrapIntoConfigSpec(*emptyConfigSpecFN, *configSpecFN, *manifestDir, *extraManifestDir)
		return
	}

	if *target == "flotta" {
		flotta.WrapIntoConfigSpec(*emptyConfigSpecFN, *configSpecFN, *manifestDir)
		return
	}

	configSpec, err := util.ReadEmtpyConfigSpec(*emptyConfigSpecFN)
	if err != nil {
		log.Fatalf("error reading empty configspec: %v", err)
	}

	rawContent := util.ComposeRawContent(*manifestDir)
	content := []string{}
	for _, item := range rawContent {
		content = append(content, string(item))
	}
	configSpec["spec"].(map[string]interface{})["content"] = content

	err = util.WriteConfigSpec(*configSpecFN, configSpec)
	if err != nil {
		log.Fatalf("error writing configspec: %v", err)
	}
}
