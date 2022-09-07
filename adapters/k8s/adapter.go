package k8s

import (
	"io/ioutil"
	"log"

	"github.com/machado-br/helm-api/adapters/models"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type adapter struct {
	cluster   models.Cluster
	clientSet *kubernetes.Clientset
	namespace string
	region    string
	token     string
	deployed  bool
}

type Adapter interface {
	RetrieveSecret() ([]byte, error)
	WriteToFile(certificate []byte) error
}

func NewAdapter(
	cluster models.Cluster,
	namespace string,
	region string,
	token string,
	deployed bool,
) (adapter, error) {
	clientSet, err := newClientset(cluster, token, deployed)
	if err != nil {
		return adapter{}, err
	}

	return adapter{
		deployed:  deployed,
		token:     token,
		cluster:   cluster,
		clientSet: clientSet,
		namespace: namespace,
		region:    region,
	}, nil
}

func newClientset(cluster models.Cluster, token string, deployed bool) (*kubernetes.Clientset, error) {
	opName := "newClientset"
	log.Printf("entering %v", opName)

	clientset := &kubernetes.Clientset{}
	config := &rest.Config{}

	if deployed {
		var err error
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	} else {
		config = &rest.Config{
			Host:        cluster.Endpoint,
			BearerToken: token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: cluster.Certificate,
			},
		}
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func (a adapter) RetrieveSecret() ([]byte, error) {
	opName := "RetrieveSecret"
	log.Printf("entering %v", opName)

	rest.InClusterConfig()
	secretList, err := a.clientSet.CoreV1().Secrets(a.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	if len(secretList.Items) == 0 {
		return []byte{}, err
	}

	secret := secretList.Items[0].Data["ca.crt"]

	return secret, nil
}

func (a adapter) WriteToFile(certificate []byte) error {
	opName := "WriteToFile"
	log.Printf("entering %v", opName)

	clustersList := map[string]*api.Cluster{
		a.cluster.Arn: {
			Server:                   a.cluster.Endpoint,
			CertificateAuthorityData: certificate,
		},
	}

	contextList := map[string]*api.Context{
		a.cluster.Arn: {
			Cluster:  a.cluster.Arn,
			AuthInfo: a.cluster.Arn,
		},
	}

	var content []byte
	authInfoList := map[string]*api.AuthInfo{}

	if a.deployed {
		var err error
		content, err = ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
		if err != nil {
			return err
		}

		authInfoList = map[string]*api.AuthInfo{
			a.cluster.Arn: {
				Token: string(content),
			},
		}
	} else {
		authInfoList = map[string]*api.AuthInfo{
			a.cluster.Arn: {
				Token: a.token,
			},
		}
	}

	clientConfig := api.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       clustersList,
		Contexts:       contextList,
		AuthInfos:      authInfoList,
		CurrentContext: a.cluster.Arn,
	}

	err := clientcmd.WriteToFile(clientConfig, "./config/kube")
	if err != nil {
		return err
	}

	return nil
}
