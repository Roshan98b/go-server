# Go-Server
A basic web server using Mux to return a JSON response

## Local server instance

Edit go mod file
```
go mod edit -go=1.19
go mod edit -module=github.com/roshan98b/go-server
```

Clean the project, install dependencies and update `go.mod` file
```
go mod tidy
```

Install dependencies without updating `go.mod` file
```
go mod download
```

Install individual dependencies and update `go.mod` file
```
go get -u github.com/gorilla/mux
```

Build and run the application
```
go build -o ./go-server
./go-server
```

Run the aplication using `go`
```
go run .
```

## Containerized server instance

Build docker image from `Dockerfile`
```
docker build -t go-server:1.0.0 -t go-server:latest .
```

If build fails, optional environment variables to be added
```
export DOCKER_BUILDKIT=0
export COMPOSE_DOCKER_CLI_BUILD=0
```

Run app on a docker container
```
docker run -d --rm --name go-server -p 8080:8080 go-server:latest
```

## Kubernetes configurations

Create the Nginx Ingress Controller using Helm in a new namespace.
```
kubectl create namespace ingress-nginx

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

helm install ingress-nginx ingress-nginx/ingress-nginx -n ingress-nginx --set controller.replicaCount=2
```

After build, Tag and push the application docker images to ACR
```
# Backend build and push image
docker build -t go-server:1.0.0 -t go-server:latest .
docker tag go-server:latest robadrinacr1.azurecr.io/go-server:latest
docker push robadrinacr1.azurecr.io/go-server:latest
```

Deploy the ingress and then the application, creates a deployment and a service. Uses a new namespace `esg-hack` for the deployment.
```
kubectl apply -f ./manifests/Go-Server.ingress.yaml -n esg-hack
kubectl apply -f ./manifests/Go-Server.yaml -n esg-hack
```

Other troubleshooting commands
```
# Map running pod shell to terminal
kubectl get pods -n esg-hack
kubectl exec -i -t -n esg-hack go-server-podname -- /bin/sh

# Get logs of a pod
kubectl get pods -n esg-hack
kubectl logs  -n esg-hack go-server-podname

# Restart or redeploy a deployment
kubectl rollout restart deployment go-server -n esg-hack 
```
