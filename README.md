# autoapi
A simple experimental code to enable auto provision with predefined entities as an  isolated API service

Setup:
get gqlgen tool kit
go get github.com/99designs/gqlgen
go get github.com/99designs/gqlgen@v0.17.39
initialie the project
go run github.com/99designs/gqlgen init

enable apollo federation:
federation:
  filename: graph/federation.go
  package: graph

to create federated server - 
go run github.com/99designs/gqlgen

# use below mutation to persist server provision request @http://localhost:8080/
mutation {
    createService(config: {
        graphPackagePath: "autoapi/graph",
        playgroundPath: "/",
        queryPath: "/query",
        ginMode:"INFO",
        port: 8080
    }) {
        id
        graphPackagePath
        playgroundPath
        queryPath
        ginMode
        port
    }
}

# To view persisted requests use below query :
{
  allServerConfigs {
    id
    graphPackagePath
    playgroundPath
    queryPath
    ginMode
    port
  }
}

# To delete by ID:

# To delete all requests
mutation {
  deleteAllServerConfigs
}


# To test the project locally please follow below instructions::
1. Install MiniKube
https://minikube.sigs.k8s.io/docs/start/
2. brew install minikube (MAC)
3. start minikube:
    minikube start
    To use docker instead of VM use below command to start :
    minikube start --driver=docker
4. check status :
    minikube status
5. to view dashboard :
  minikube dashboard

ArgoCD setup:

1. Create seperate namespance for ArgoCD:
kubectl create namespace argocd
install with :
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
to access ui :
minikube service -n argocd argocd-server

get the admin password with below command :

kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2

to update pass word :
ARGO_PWD=$(bcrypt-hash YOUR_NEW_PASSWORD)
kubectl patch secret argocd-secret -n argocd -p '{"stringData": {"admin.password": "test123232324313412341", "admin.passwordMtime": "'$(date +%FT%T%Z)'"}}'

kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

use rancher for k8 management :
sudo docker run --privileged -d --restart=unless-stopped -p 80:80 -p 443:443 rancher/rancher

to get initial password: 26mnwv4fx9rdz9v9r9hld47kbgpbmc75b59wwc9ktzbz4f59wwj8tv


docker logs  754d3e12cb70  2>&1 | grep "Bootstrap Password:"



To delete ns:
kubectl delete namespace argocd
To get Initial password from argocd terminal:
argocd admin initial-password



//TODO:: Add ingress
//remove service from Kustomization base so that we can define seperate service names 
//Update Graphlets to serve data seperately based on features
//Add test cases


#To create a service use below mutation(please note service name should be lower case with - or _):
mutation {
  createService(
    config: {apiserverName: "vehicle-service", contextPath: "vehicle", features: [VEHICLE, DEALER], performanceRequirements: {apiUsageFrequency: LOW, requestVolume: SMALL, highAvailability: false, batchLoad: false}}
  ) {
    id
    apiserverName
  }
}