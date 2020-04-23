package k8s

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/ghodss/yaml"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func readSpecFile(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ext := filepath.Ext(path)
	if ext == ".json" {
		return bytes, nil
	}
	if ext == ".yml" || ext == ".yaml" {
		jsonBytes, err := yaml.YAMLToJSON(bytes)
		if err != nil {
			return nil, err
		}
		return jsonBytes, nil
	}
	return nil, errors.New(fmt.Sprintf("Unsupported extension: %s", ext))
}

func RunJobWithArgs(specFilePath string, args []string) {
	configPath := path.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		fmt.Printf("Tried to load K8s config from %s, but file not found.\n", configPath)
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Could not create client from K8s config: %s\n", err.Error())
		os.Exit(1)
	}

	bytes, err := readSpecFile(specFilePath)
	if err != nil {
		fmt.Printf("Failed to load K8s spec from file: %s\n", err.Error())
		os.Exit(1)
	}

	var job = batchv1.Job{}
	err = json.Unmarshal(bytes, &job)
	if err != nil {
		panic(err.Error())
	}

	containers := job.Spec.Template.Spec.Containers
	if len(containers) > 1 {
		fmt.Println("Expected a single container job.")
		os.Exit(1)
	}

	containers[0].Args = args
	jobsClient := clientset.BatchV1().Jobs(job.Namespace)

	res, err := jobsClient.Create(&job)

	if err != nil {
		fmt.Printf("Failed to create a new K8s job: %s\n", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Created job", res.Name)
	}
}
