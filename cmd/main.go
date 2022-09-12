package main

import (
	"log"
	"os"
	"strconv"

	"github.com/machado-br/helm-api/adapters/helm"
	"github.com/machado-br/helm-api/adapters/k8s"
	"github.com/machado-br/helm-api/adapters/models"
	"github.com/machado-br/helm-api/api"
	"github.com/machado-br/helm-api/services/createKubeConfig"
	"github.com/machado-br/helm-api/services/installChart"
	"github.com/machado-br/helm-api/services/listReleases"
	"github.com/machado-br/helm-api/services/uninstallRelease"
)

func main() {

	name := os.Getenv("CLUSTER_NAME")
	namespace := os.Getenv("NAMESPACE")
	chartDirectory := os.Getenv("CHART_DIRECTORY")
	kubeconfigPath := os.Getenv("KUBECONFIG_PATH")
	helmDriver := os.Getenv("HELM_DRIVER")
	deployedK8sApiEndpoint := os.Getenv("DEPLOYED_K8S_ENDPOINT")
	deployedCertificatePath := os.Getenv("DEPLOYED_CERTIFICATE_PATH")
	deployedTokenPath := os.Getenv("DEPLOYED_TOKEN_PATH")

	deployed, err := strconv.ParseBool(os.Getenv("DEPLOYED"))
	if err != nil {
		log.Fatalf("failed to parse deployed env var: %v", err)
	}

	deployedConfiguration := models.DeployedConfiguration{
		K8sAPIEndpoint:  deployedK8sApiEndpoint,
		CertificatePath: deployedCertificatePath,
		TokenPath:       deployedTokenPath,
	}

	k8sAdapter, err := k8s.NewAdapter(name, deployedConfiguration)
	if err != nil {
		log.Fatalf("failed while creating k8s adapter: %v", err)
	}

	createKubeConfigService, err := createKubeConfig.NewService(k8sAdapter)
	if err != nil {
		log.Fatalf("failed while creating createKubeConfig service: %v", err)
	}

	if deployed {
		err = createKubeConfigService.Run()
		if err != nil {
			log.Fatalf("failed while creating createKubeConfig file: %v", err)
		}
		log.Println("Kubernetes configuration file created successfully")
	}

	helmAdapter, err := helm.NewAdapter(namespace, kubeconfigPath, helmDriver, chartDirectory, deployed)
	if err != nil {
		log.Fatalf("failed while creating helm adapter: %v", err)
	}

	listReleasesService, err := listReleases.NewService(helmAdapter)
	if err != nil {
		log.Fatalf("failed while creating list releases service: %v", err)
	}

	installChartService, err := installChart.NewService(helmAdapter)
	if err != nil {
		log.Fatalf("failed while creating install chart service: %v", err)
	}

	uninstallReleaseService, err := uninstallRelease.NewService(helmAdapter)
	if err != nil {
		log.Fatalf("failed while creating uninstall release service: %v", err)
	}

	api, err := api.NewApi(listReleasesService, installChartService, uninstallReleaseService)
	if err != nil {
		log.Fatalf("failed while creating api: %v", err)
	}

	api.Run()
}
