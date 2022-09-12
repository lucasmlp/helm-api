package uninstallRelease

import (
	"log"

	"github.com/machado-br/helm-api/adapters/helm"
	"github.com/machado-br/helm-api/services"
)

type service struct {
	helmAdapter helm.Adapter
}

type Service interface {
	Run(releaseName string, dryRun bool) error
}

func NewService(
	helmAdapter helm.Adapter,
) (service, error) {
	return service{
		helmAdapter: helmAdapter,
	}, nil
}

func (s service) Run(releaseName string, dryRun bool) error {
	opName := "uninstallChart.Run"
	log.Printf("entering %v", opName)

	
	err := s.helmAdapter.UninstallRelease(releaseName, dryRun)
	if err != nil {
		log.Println(err)
		return services.ErrUninstallRelease
	}

	return nil
}