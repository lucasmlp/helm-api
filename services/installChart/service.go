package installChart

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
	Run(releaseName string, dryRun bool, chart models.Chart) error
}

func NewService(
	helmAdapter helm.Adapter,
) (service, error) {
	return service{
		helmAdapter: helmAdapter,
	}, nil
}

func (s service) Run(releaseName string, dryRun bool, chart models.Chart) error {
	opName := "installChart.Run"
	log.Printf("entering %v", opName)

	
	err := s.helmAdapter.InstallChart(releaseName, dryRun, chart)
	if err != nil {
		log.Println(err)
		return services.ErrInstallChart
	}

	return nil
}