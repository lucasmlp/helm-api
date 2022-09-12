package api

type InstallChartPayload struct {
	ReleaseName string	`json:"releaseName"  binding:"required"`
	ChartName	string	`json:"chartName"  binding:"required"`
	RepoURL		string	`json:"repoURL"  binding:"required"`
	Version 	string 	`json:"version"  binding:"required"`
}