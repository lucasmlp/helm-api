# Helm API

[![Tests](https://github.com/machado-br/helm-api/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/machado-br/helm-api/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/machado-br/helm-api)](https://goreportcard.com/report/github.com/machado-br/helm-api)
[![Documentation](https://godoc.org/github.com/machado-br/helm-api?status.svg)](http://godoc.org/github.com/machado-br/helm-api)
[![Gocover](http://gocover.io/_badge/github.com/machado-br/helm-api)](http://gocover.io/github.com/machado-br/helm-api)

This is a simple service that connects to an aws cluster and list all Helm releases on it.

## Requirements

- AWS CLI configured on your machine
- Golang 1.18 configured on your machine

## Usage

 - Modify the .env file with the desired values:
	 - CLUSTER_NAME: Name of the cluster you'll be working on;
	 - NAMESPACE: is the default namespace of your cluster;
	 - AWS_REGION: is the region that your cluster is created;
	 - GIN_MODE: is the gin logging's mode to be used.
