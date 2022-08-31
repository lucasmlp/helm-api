package createKubeConfig

import (
	"log"

	"github.com/machado-br/helm-api/adapters/k8s"
	"github.com/machado-br/helm-api/services"
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
		log.Println(err)
		return services.ErrK8sSecrets
	}

	err = s.k8sAdapter.WriteToFile(secret)
	if err != nil {
		log.Println(err)
		return services.ErrWriteKubeConfig
	}

	return nil
}
