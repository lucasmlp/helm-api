aws ecr create-repository \
    --repository-name helm-api \
    --image-scanning-configuration scanOnPush=true \
    --region us-west-2