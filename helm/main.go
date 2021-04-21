package main

import (
	"fmt"
	"log"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
)

func main() {
	// Load chart
	chrt, err := loader.Load("resources/dummy")
	if err != nil {
		log.Panicf("error while loading chart: %s", err)
	}

	err = chrt.Validate()
	if err != nil {
		log.Panicf("error while validatiing the chart: %s", err)
	}

	// Prepare values
	overrides := values.Options{
		ValueFiles: []string{"resources/dummy/landscape-dev.yaml"},
	}
	settings := cli.New()
	vals, err := overrides.MergeValues(getter.All(settings))
	if err != nil {
		log.Panicf("error while preparing overrides: %s", err)
	}

	// Prepare helm install command in dry-run mode
	// in order to render yaml from chart
	actionConfig := new(action.Configuration)
	installCmd := action.NewInstall(actionConfig)
	installCmd.DryRun = true
	installCmd.ClientOnly = true
	installCmd.ReleaseName = "RELEASE" // TODO: read from arg

	release, err := installCmd.Run(chrt, vals)
	if err != nil {
		log.Panicf("error while rendering chart using helm install: %s", err)
	}

	// At this point we can apply the YAML
	yaml := release.Manifest
	fmt.Println(yaml)
}
