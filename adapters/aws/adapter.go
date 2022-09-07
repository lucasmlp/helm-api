package aws

import (
	"encoding/base64"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/machado-br/helm-api/adapters/models"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

type adapter struct {
	region      string
	clusterName string
	session     *session.Session
	eks         *eks.EKS
}

type Adapter interface {
	GetToken() (string, error)
	DescribeCluster() (models.Cluster, error)
}

func NewAdapter(
	region string,
	clusterName string,
) (adapter, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	eksSvc := eks.New(sess)

	return adapter{
		region:      region,
		clusterName: clusterName,
		session:     sess,
		eks:         eksSvc,
	}, nil
}

func (a adapter) DescribeCluster() (models.Cluster, error) {
	opName := "DescribeCluster"
	log.Printf("entering %v", opName)

	input := &eks.DescribeClusterInput{
		Name: aws.String(a.clusterName),
	}

	result, err := a.eks.DescribeCluster(input)
	if err != nil {
		return models.Cluster{}, err
	}

	ca, err := base64.StdEncoding.DecodeString(aws.StringValue(result.Cluster.CertificateAuthority.Data))
	if err != nil {
		return models.Cluster{}, err
	}

	return models.Cluster{
		Arn:         aws.StringValue(result.Cluster.Arn),
		Name:        aws.StringValue(result.Cluster.Name),
		Endpoint:    aws.StringValue(result.Cluster.Endpoint),
		Certificate: ca,
	}, nil
}

func (a adapter) GetToken() (string, error) {
	opName := "GetToken"
	log.Printf("entering %v", opName)

	gen, err := token.NewGenerator(true, false)
	if err != nil {
		return "", err
	}

	opts := &token.GetTokenOptions{
		ClusterID: a.clusterName,
	}

	tok, err := gen.GetWithOptions(opts)
	if err != nil {
		return "", err
	}

	return tok.Token, nil
}
