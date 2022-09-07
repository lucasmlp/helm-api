package gcloud

import "github.com/machado-br/helm-api/adapters/models"

type CloudProviderAdapter interface {
	DescribeCluster() (models.Cluster, error)
}
