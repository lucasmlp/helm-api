package services

import "errors"

var (
	ErrK8sSecrets = errors.New("failed while retrieving k8s secret")

	ErrInstallChart = errors.New("failed while installing chart")

	ErrWriteKubeConfig = errors.New("failed while writing kubeconfig file")

	ErrGetClusterInfo = errors.New("failed while retrieving cluster information")

	ErrListReleases = errors.New("failed while listing releases")
)
