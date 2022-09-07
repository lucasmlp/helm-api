package listReleases

import (
	"log"

	"github.com/machado-br/helm-api/adapters/helm"
	"github.com/machado-br/helm-api/adapters/models"
	"github.com/machado-br/helm-api/services"
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
	opName := "listReleases.Run"
	log.Printf("entering %v", opName)

	releases, err := s.helmAdapter.ListReleases()
	if err != nil {
		log.Println(err)
		return []models.Release{}, services.ErrListReleases
	}

	return releases, nil
}
