package main

import (
	"fmt"
	"log"

	"sigs.k8s.io/kustomize/api/filesys"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/api/types"
)

func main() {
	// Initialize Kustomize API with Helm plugin enabled
	fs := filesys.MakeFsOnDisk()
	opts := &krusty.Options{
		PluginConfig: &types.PluginConfig{
			HelmConfig: types.HelmConfig{
				Enabled: true,
				Command: "helm",
			},
		},
	}
	kustomizer := krusty.MakeKustomizer(opts)

	// Run Kustomize on the resources
	res, err := kustomizer.Run(fs, "resources/nginx-from-helm")
	if err != nil {
		log.Panicf("error while calling Run on kustomizer: %s", err)
	}

	// Prepare YAML representation
	yamlData, err := res.AsYaml()
	if err != nil {
		log.Panicf("error while rendering yaml from kustomizer ResMap: %s", err)
	}

	// At this point we can apply the yaml
	yaml := string(yamlData)
	fmt.Println(yaml)
}
