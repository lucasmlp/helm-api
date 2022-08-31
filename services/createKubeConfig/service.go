package createKubeConfig

import (
	"log"

	"github.com/machado-br/helm-api/adapters/k8s"
)

type service struct {
	k8sAdapter k8s.Adapter
}

type Service interface {
	Run() error
}

func NewService(
	k8sAdapter k8s.Adapter,
) (service, error) {
	return service{
		k8sAdapter: k8sAdapter,
	}, nil
}

func (s service) Run() error {

	secret, err := s.k8sAdapter.RetrieveSecret()
	if err != nil {
		log.Fatalf("Failed while retrieving k8s secret: %v", err)
	}

	err = s.k8sAdapter.WriteToFile(secret)
	if err != nil {
		log.Fatalf("Failed while writing kubeconfig file: %v", err)
	}

	return nil
}
