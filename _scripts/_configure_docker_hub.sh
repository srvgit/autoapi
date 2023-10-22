#!/bin/bash

read -p "Enter Docker Hub Username: " DOCKERHUB_USERNAME
read -s -p "Enter Docker Hub Password: " DOCKERHUB_PASSWORD
echo ""

# Base64 encode the Docker Hub credentials
ENCODED_AUTH=$(echo -n "$DOCKERHUB_USERNAME:$DOCKERHUB_PASSWORD" | base64)

# Create the dockerconfigjson file
cat <<EOF > dockerconfigjson
{
  "auths": {
    "https://index.docker.io/v2/": {
      "auth": "$ENCODED_AUTH"
    }
  }
}
EOF

# Create the Kubernetes secret
kubectl create secret generic regcred \
  --from-file=.dockerconfigjson=dockerconfigjson \
  --type=kubernetes.io/dockerconfigjson \
  --namespace=api




# Clean up the dockerconfigjson file
rm -f dockerconfigjson

echo "Docker registry secret 'regcred' created in namespace 'namespace'"
