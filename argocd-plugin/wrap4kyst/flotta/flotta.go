package flotta

import (
	"encoding/json"
	"log"

	"github.com/edge-experiments/wrap4kyst/util"
	v1alpha1 "github.com/project-flotta/flotta-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/yaml"
)

var edgeWorkloadName string = "wrap4kyst-generated"
var edgeWorkloadNamespace string = "default"

// WrapIntoConfigSpec tries to find the one piece of user manifest which contains a pod,
// extract the pod,
// wrap the pod into an EdgeWorkload,
// then wrap again the EdgeWorkload into a ConfigSpec.
func WrapIntoConfigSpec(inputFile, outputFile, manifestDir string) {
	configSpec, err := util.ReadEmtpyConfigSpec(inputFile)
	if err != nil {
		log.Fatalf("error reading empty configspec: %v", err)
	}

	found, podSpec := false, corev1.Pod{}.Spec
	rawContent := util.ComposeRawContent(manifestDir)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	for _, item := range rawContent {
		obj, _, err := decode(item, nil, nil)
		if err != nil {
			log.Fatalf("error decoding line in raw content: %v", err)
		}
		if obj.GetObjectKind().GroupVersionKind().Kind == "Deployment" {
			found, podSpec = true, obj.(*appsv1.Deployment).Spec.Template.Spec
			log.Printf("found podSpec: %v", podSpec)
			break
		}
	}
	if !found {
		log.Fatalf("no pod found in manifestDir: %v", manifestDir)
	}

	edgeWorkload := v1alpha1.EdgeWorkload{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "management.project-flotta.io/v1alpha1",
			Kind:       "EdgeWorkload",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      edgeWorkloadName,
			Namespace: edgeWorkloadNamespace,
		},
		Spec: v1alpha1.EdgeWorkloadSpec{
			Pod: v1alpha1.Pod{
				Spec: podSpec,
			},
			Type: v1alpha1.PodWorkloadType,
		},
	}

	bytes := []byte{}
	bytes, err = json.Marshal(edgeWorkload)
	bytes, err = yaml.JSONToYAML(bytes)
	log.Printf("wrapped EdgeWorkload:\n%v\n", string(bytes))

	configSpec["spec"].(map[string]interface{})["content"] = []string{string(bytes)}
	err = util.WriteConfigSpec(outputFile, configSpec)
	if err != nil {
		log.Fatalf("error writing ConfigSpec: %v", err)
	}
}
