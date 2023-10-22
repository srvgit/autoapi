#!/bin/bash

# Check if minikube is installed
if ! command -v minikube &> /dev/null; then
    echo "Minikube is not installed. Please install it and try again."
    exit 1
fi

# Check if kubectl is installed
if ! command -v kubectl &> /dev/null; then
    echo "kubectl is not installed. Please install it and try again."
    exit 1
fi

# Start Minikube
echo "Starting Minikube..."
minikube start

# Set up ArgoCD
echo "Setting up ArgoCD on Minikube..."
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Wait for ArgoCD server pod to be ready
echo "Waiting for ArgoCD Server pod to be ready..."
kubectl wait --namespace argocd --for=condition=ready pod --selector=app.kubernetes.io/name=argocd-server --timeout=300s
echo "ArgoCD Server pod is ready!"

# Get and display default password
ARGOCD_SERVER_POD=$(kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2)
echo "ArgoCD default password: $ARGOCD_SERVER_POD"
echo "Setup complete!"

#VPA is not part og k8s yet, so we need to install it manually
git clone https://github.com/kubernetes/autoscaler.git
cd autoscaler/vertical-pod-autoscaler
./hack/vpa-up.sh



#verify :
kubectl get crd verticalpodautoscalers.autoscaling.k8s.io

kubectl create namespace api
