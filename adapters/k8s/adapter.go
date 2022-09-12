package k8s

import (
	"log"

	"github.com/machado-br/helm-api/adapters/models"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type adapter struct {
	clusterName           string
	deployedConfiguration models.DeployedConfiguration
}

type Adapter interface {
	WriteToFile() error
}

func NewAdapter(
	clusterName string,
	deployedConfiguration models.DeployedConfiguration,
) (adapter, error) {

	return adapter{
		deployedConfiguration: deployedConfiguration,
		clusterName:           clusterName,
	}, nil
}

func (a adapter) WriteToFile() error {
	opName := "WriteToFile"
	log.Printf("entering %v", opName)

	clustersList := map[string]*api.Cluster{
		a.clusterName: {
			Server:               a.deployedConfiguration.K8sAPIEndpoint,
			CertificateAuthority: a.deployedConfiguration.CertificatePath,
		},
	}

	contextList := map[string]*api.Context{
		a.clusterName: {
			Cluster:  a.clusterName,
			AuthInfo: a.clusterName,
		},
	}

	authInfoList := map[string]*api.AuthInfo{
		a.clusterName: {
			TokenFile: a.deployedConfiguration.TokenPath,
		},
	}

	clientConfig := api.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       clustersList,
		Contexts:       contextList,
		AuthInfos:      authInfoList,
		CurrentContext: a.clusterName,
	}

	err := clientcmd.WriteToFile(clientConfig, "./config/kube")
	if err != nil {
		return err
	}

	return nil
}
