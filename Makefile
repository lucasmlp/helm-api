include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

server:
	@ echo
	@ echo "Spinning up server..."
	@ echo
	@ rm -rf ./config/kube
	@ go run ./cmd/main.go

docker-image:
	@ echo
	@ echo "Building docker image..."
	@ echo
	@ docker build -t machado-br/helm-api:latest .

docker-tag-aws:
	@ echo
	@ echo "Tagging docker image for AWS..."
	@ echo
	@ docker tag machado-br/helm-api:latest 774429751797.dkr.ecr.us-west-2.amazonaws.com/helm-api:latest

login-aws-ecr:
	@ echo
	@ echo "Logging in AWS ECR..."
	@ echo
	@ aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin 774429751797.dkr.ecr.us-west-2.amazonaws.com

docker-push-aws:
	@ echo
	@ echo "Pushing docker image to AWS ECR..."
	@ echo
	@ docker push 774429751797.dkr.ecr.us-west-2.amazonaws.com/helm-api:latest