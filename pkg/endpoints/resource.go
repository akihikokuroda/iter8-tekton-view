package endpoints
import (
    "fmt"
    //metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    //"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
    "k8s.io/apimachinery/pkg/runtime/schema"
    "k8s.io/client-go/dynamic"
    "k8s.io/client-go/rest"
)

var (
	iter8GVR = schema.GroupVersionResource{
		Group:    "iter8.tools",
		Version:  "v1alpha1",
		Resource: "experiments",
//		Group:    "apps",
//		Version:  "v1",
//		Resource: "deployments",
	}
)


type Resource struct {
    experimentClient dynamic.NamespaceableResourceInterface
}

func NewResource() (Resource, error) {
    config, err := rest.InClusterConfig()
    if err != nil {
        fmt.Printf("Error %v", err)
    }
    fmt.Println(config)
    dynClient, errClient := dynamic.NewForConfig(config)
    if errClient != nil {
        fmt.Printf("Error received creating client %v", errClient)
    }

    crdClient := dynClient.Resource(iter8GVR)
    r := Resource{
        experimentClient: crdClient,
    }
    fmt.Println(crdClient)
    return r, nil
}