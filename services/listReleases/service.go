package listReleases

import (
	"github.com/machado-br/helm-api/adapters/helm"
	"github.com/machado-br/helm-api/adapters/models"
)

type service struct {
	helmAdapter helm.Adapter
}

type Service interface {
	Run() ([]models.Release, error)
}

func NewService(
	helmAdapter helm.Adapter,
) (service, error) {
	return service{
		helmAdapter: helmAdapter,
	}, nil
}

func (s service) Run() ([]models.Release, error) {
	releases := s.helmAdapter.ListReleases()
	// if err != nil {
	// 	return []models.Release{}, errors.New("failed while listing releases")
	// }

	return releases, nil
}
