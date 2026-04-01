// Package k8s handles dynamic operations with the Kubernetes API Server.
package k8s

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

// ApplyYAML takes raw YAML configuration data, decodes it into an unstructured Kubernetes object,
// and applies it to the cluster (create or update). It builds a dynamic client using in-cluster configuration.
func ApplyYAML(yamlData []byte) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			return err
		}
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return err
	}

	dec := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

	obj := &unstructured.Unstructured{}
	_, _, err = dec.Decode(yamlData, nil, obj)
	if err != nil {
		return err
	}

	gvr, _ := GetGVR(obj)

	resource := client.Resource(gvr).Namespace(obj.GetNamespace())

	_, err = resource.Create(context.TODO(), obj, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("Create failed, trying update:", err)

		_, err = resource.Update(context.TODO(), obj, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	}

	fmt.Println("Applied:", obj.GetKind(), obj.GetName())
	return nil
}

// GetGVR dynamically resolves the GroupVersionResource (GVR) associated with the specified unstructured Kubernetes object.
// GVRs are required by the Kubernetes dynamic client to locate the correct REST API endpoint for the resource.
// Note: This relies on a simplified heuristic appending "s" to the Resource kind and won't work perfectly for irregular plurals.
func GetGVR(obj *unstructured.Unstructured) (schema.GroupVersionResource, error) {
	gvk := obj.GroupVersionKind()

	return schema.GroupVersionResource{
		Group:    gvk.Group,
		Version:  gvk.Version,
		Resource: strings.ToLower(gvk.Kind) + "s",
	}, nil
}