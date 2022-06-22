package ocm

import (
	"encoding/json"
	"log"

	"github.com/edge-experiments/wrap4kyst/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	workv1 "open-cluster-management.io/api/work/v1"
	"sigs.k8s.io/yaml"
)

var manifestWorkName string = "wrap4kyst-generated"

// WrapIntoConfigSpec wraps all user's k8s manifests into a single ManifestsWork,
// then wraps again the ManifestsWork into a ConfigSpec.
func WrapIntoConfigSpec(inputFile, outputFile, manifestDir, extraManifestDir string) {
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
	}

	manifestwork := workv1.ManifestWork{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "work.open-cluster-management.io/v1",
			Kind:       "ManifestWork",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: manifestWorkName,
		},
		Spec: workv1.ManifestWorkSpec{
			Workload: workv1.ManifestsTemplate{
				Manifests: manifests,
			},
		},
	}

	// // this method only outputs to io.Writer
	// // "k8s.io/apimachinery/pkg/runtime/serializer/json"
	// s := serializerjson.NewYAMLSerializer(serializerjson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
	// err = s.Encode(&manifestwork, os.Stdout)

	// // this method only outputs to io.Writer
	// // "k8s.io/cli-runtime/pkg/printers"
	// p := printers.YAMLPrinter{}
	// p.PrintObj(&manifestwork, os.Stdout)

	// this method can output to []byte
	// "sigs.k8s.io/yaml"
	bytes := []byte{}
	bytes, err = json.Marshal(manifestwork)
	bytes, err = yaml.JSONToYAML(bytes)
	log.Printf("wrapped ManifestWork:\n%v\n", string(bytes))

	content := []string{string(bytes)}

	extraRawContent := util.ComposeRawContent(extraManifestDir)
	for _, item := range extraRawContent {
		content = append(content, string(item))
	}

	configSpec["spec"].(map[string]interface{})["content"] = content
	err = util.WriteConfigSpec(outputFile, configSpec)
	if err != nil {
		log.Fatalf("error writing ConfigSpec: %v", err)
	}
}
