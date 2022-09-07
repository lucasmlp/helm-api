package k8s

import (
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
) (adapter, error) {
	clientSet, err := newClientset(cluster, token)
	if err != nil {
		return adapter{}, err
	}

	return adapter{
		cluster:   cluster,
		clientSet: clientSet,
		namespace: namespace,
		region:    region,
	}, nil
}

func newClientset(cluster models.Cluster, token string) (*kubernetes.Clientset, error) {
	log.Printf("Cluster name: %+v", cluster.Name)

	clientset, err := kubernetes.NewForConfig(
		&rest.Config{
			Host:        cluster.Endpoint,
			BearerToken: token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: cluster.Certificate,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func (a adapter) RetrieveSecret() ([]byte, error) {
	secretList, err := a.clientSet.CoreV1().Secrets(a.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	if len(secretList.Items) == 0 {
		log.Panicln("Secret list is empty")
	}

	secret := secretList.Items[0].Data["ca.crt"]

	return secret, nil
}

func (a adapter) WriteToFile(certificate []byte) error {
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

	exec := api.ExecConfig{
		Command:    "aws",
		Args:       []string{"eks", "get-token", "--region", a.region, "--cluster-name", a.cluster.Name},
		APIVersion: "client.authentication.k8s.io/v1beta1",
	}

	authInfoList := map[string]*api.AuthInfo{
		a.cluster.Arn: {
			Exec: &exec,
		},
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
