package describeCluster

import (
	"log"

	"github.com/machado-br/helm-api/adapters/aws"
	"github.com/machado-br/helm-api/adapters/models"
	"github.com/machado-br/helm-api/services"
)

type service struct {
	cloudProviderAdapter aws.Adapter
}

type Service interface {
	Run() (models.Cluster, error)
}

func NewService(
	cloudProviderAdapter aws.Adapter,
) (service, error) {
	return service{
		cloudProviderAdapter: cloudProviderAdapter,
	}, nil
}

func (s service) Run() (models.Cluster, error) {
	cluster, err := s.cloudProviderAdapter.DescribeCluster()
	if err != nil {
		log.Println("Failed while calling DescribeCluster: %v", err)
		return models.Cluster{}, services.ErrGetClusterInfo
	}

	return cluster, nil
}
