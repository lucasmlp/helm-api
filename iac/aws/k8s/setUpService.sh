kubectl apply -f ./iac/k8s/serviceAccount.yml
kubectl apply -f ./iac/k8s/clusterRole.yml
kubectl apply -f ./iac/k8s/clusterRoleBinding.yml
kubectl apply -f ./iac/k8s/deployment.yml