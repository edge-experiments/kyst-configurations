package ocm

import (
	"encoding/json"
	"log"

	"github.com/edge-experiments/wrap4kyst/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	workv1 "open-cluster-management.io/api/work/v1"
	ioyaml "sigs.k8s.io/yaml"
)

var manifestWorkName string = "wrap4kyst-generated"
var manifestWorkNamespace string = "cluster1"

func WrapIntoConfigSpec(inputFile, outputFile, manifestDir string) {
	configSpec, err := util.ReadEmtpyConfigSpec(inputFile)
	if err != nil {
		log.Fatalf("error reading empty configspec: %v", err)
	}

	manifests := []workv1.Manifest{}
	rawContent := util.ComposeRawContent(manifestDir)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	for _, item := range rawContent {
		m := workv1.Manifest{}
		obj, _, err := decode(item, nil, nil)
		if err != nil {
			log.Fatalf("error decoding line in raw content: %v", err)
		}
		m.Object = obj
		manifests = append(manifests, m)

		// newFile, err := os.Create("lastobj.yaml")
		// y := printers.YAMLPrinter{}
		// defer newFile.Close()
		// y.PrintObj(obj, newFile)
	}

	manifestwork := workv1.ManifestWork{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "work.open-cluster-management.io/v1",
			Kind:       "ManifestWork",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      manifestWorkName,
			Namespace: manifestWorkNamespace,
		},
		Spec: workv1.ManifestWorkSpec{
			Workload: workv1.ManifestsTemplate{
				Manifests: manifests,
			},
		},
	}

	// s := serializerjson.NewYAMLSerializer(serializerjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
	// err = s.Encode(&manifestwork, os.Stdout)

	bytes := []byte{}
	bytes, err = json.Marshal(manifestwork)
	bytes, err = ioyaml.JSONToYAML(bytes)
	log.Printf("wrapped ManifestWork: %v\n", string(bytes))
	configSpec["spec"].(map[string]interface{})["content"] = []string{string(bytes)}

	err = util.WriteConfigSpec(outputFile, configSpec)
	if err != nil {
		log.Fatalf("error writing ConfigSpec: %v", err)
	}
}
