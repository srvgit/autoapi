# kubectl get svc argocd-server -n argocd -o yaml > argocd-server-service.yaml
# update port to 8090
# kubectl apply -f argocd-server-service.yaml
minikube service -n argocd argocd-server