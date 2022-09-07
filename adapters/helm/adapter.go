package helm

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/machado-br/helm-api/adapters"
	"github.com/machado-br/helm-api/adapters/models"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"
)

type adapter struct {
	action *action.Configuration
	settings *cli.EnvSettings
	namespace string
	chartDirectory string
}

type Adapter interface {
	ListReleases() ([]models.Release, error)
	InstallChart(releaseName string, dryRun bool, chart models.Chart) error
	UninstallRelease(releaseName string, dryRun bool) error
}

func NewAdapter(
	namespace string,
	configPath string,
	driver string,
	chartDirectory string,
) (adapter, error) {

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(configPath, "", namespace), namespace, driver, log.Printf); err != nil {
		log.Fatalln(err)
	}

	settings := cli.New()

	return adapter{
		settings: settings,
		action: actionConfig,
		namespace: namespace,
		chartDirectory: chartDirectory,
	}, nil
}

func (a adapter) ListReleases() ([]models.Release, error) {
	opName := "ListReleases"
	log.Printf("entering %v", opName)

	listAction := action.NewList(a.action)
	releases, err := listAction.Run()
	if err != nil {
		log.Println(err)
		return []models.Release{}, adapters.ErrListReleases
	}

	return mapToReleaseModel(releases), nil
}

func mapToReleaseModel(releases []*release.Release) []models.Release {
	opName := "mapToReleaseModel"
	log.Printf("entering %v", opName)

	releaseList := []models.Release{}
	for _, release := range releases {
		releaseList = append(releaseList, models.Release{
			Name: release.Name,
		})
	}

	return releaseList
}

func (a adapter) InstallChart(releaseName string, dryRun bool, chart models.Chart) error {
	opName := "InstallChart"
	log.Printf("entering %v", opName)

	chartDirectory, err := ioutil.TempDir("", "charts")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer os.RemoveAll(chartDirectory)

	pulledChart := a.pullChart(chart, chartDirectory)

	client := action.NewInstall(a.action)

	client.ReleaseName = releaseName
	client.DryRun = dryRun
	client.Namespace = a.namespace

	_, err = client.Run(pulledChart, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	
	return nil
}

func (a adapter) pullChart(chart models.Chart, chartDirectory string) *chart.Chart {
	opName := "pullChart"
	log.Printf("entering %v", opName)

	client := action.NewPullWithOpts(action.WithConfig(a.action))

	client.RepoURL = chart.RepoURL
	client.Settings = a.settings
	client.DestDir = chartDirectory
	client.Version = chart.Version

	_, err := client.Run(chart.Name)
	if err != nil {
		log.Println(err)
		return nil
	}
	
	validatedChart, err := loadAndValidate(chartDirectory + "/" + chart.Name + "-" + chart.Version + ".tgz")
	if err != nil {
		log.Println(err)
	}

	log.Printf("chart %s is valid \n", validatedChart.Metadata.Name)
	return validatedChart
}

func loadAndValidate(chartPath string) (*chart.Chart, error) {
	opName := "loadAndValidate"
	log.Printf("entering %v", opName)

	chart, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}

	err = chart.Validate()
	if err != nil {
		return nil, err
	}

	return chart, nil
}

func (a adapter) UninstallRelease(releaseName string, dryRun bool) error {
	opName := "UninstallRelease"
	log.Printf("entering %v", opName)

	client := action.NewUninstall(a.action)
	client.DryRun = dryRun
	
	_, err := client.Run(releaseName)
	if err != nil {
		log.Println(err)
		return err
	}
	
	return nil
}