#!/bin/bash

# Initialize ArgoCD
echo "Initializing ArgoCD..."
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
echo "ArgoCD Initialized!"

# Wait for ArgoCD server pod to be ready
echo "Waiting for ArgoCD Server pod to be ready..."
kubectl wait --namespace argocd --for=condition=ready pod --selector=app.kubernetes.io/name=argocd-server --timeout=300s
echo "ArgoCD Server pod is ready!"

# Get and display default password
ARGOCD_SERVER_POD=$(kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2)
echo "ArgoCD default password: $ARGOCD_SERVER_POD"

# Set repository credentials (Optional)
# Uncomment and modify the lines below according to your needs
# echo "Setting up Git repository credentials..."
# kubectl create secret generic git-creds --from-literal=username=YOUR_USERNAME --from-literal=password=YOUR_PASSWORD -n argocd


#run below scripts to DELETE ArgoCD
# kubectl delete -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
# kubectl delete namespace argocd
