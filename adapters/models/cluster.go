package models

type Cluster struct {
	Arn         string
	Name        string
	Endpoint    string
	Certificate []byte
}
