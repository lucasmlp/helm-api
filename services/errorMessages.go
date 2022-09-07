package services

import "errors"

var (
	ErrK8sSecrets = errors.New("failed while retrieving k8s secret")

	ErrWriteKubeConfig = errors.New("failed while writing kubeconfig file")

	ErrGetClusterInfo = errors.New("failed while retrieving cluster information")

	ErrListReleases = errors.New("failed while listing releases")

	ErrInstallChart = errors.New("failed while installing chart")

	ErrUninstallRelease = errors.New("failed while uninstalling release")
)
