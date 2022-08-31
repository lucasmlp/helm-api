package helm

import (
	"log"

	"github.com/machado-br/helm-api/adapters"
	"github.com/machado-br/helm-api/adapters/models"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"
)

type adapter struct {
	action *action.Configuration
}

type Adapter interface {
	ListReleases() ([]models.Release, error)
}

func NewAdapter(
	namespace string,
	configPath string,
	driver string,
) (adapter, error) {

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(configPath, "", namespace), namespace, driver, log.Printf); err != nil {
		log.Fatalln(err)
	}

	return adapter{
		action: actionConfig,
	}, nil
}
func (a adapter) ListReleases() ([]models.Release, error) {

	listAction := action.NewList(a.action)
	releases, err := listAction.Run()
	if err != nil {
		log.Println(err)
		return []models.Release{}, adapters.ErrListReleases
	}

	return mapToReleaseModel(releases), nil
}

func mapToReleaseModel(releases []*release.Release) []models.Release {
	releaseList := []models.Release{}
	for _, release := range releases {
		releaseList = append(releaseList, models.Release{
			Name: release.Name,
		})
	}

	return releaseList
}
