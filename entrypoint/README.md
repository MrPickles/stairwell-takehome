# Entrypoint

This folder contains starter code for a simple Golang server, as well as
configs to containerize it and deploy it to Kubernetes.

The server will reach out to the "generator" server (configured via a flag) and
forward its response to the client.
The code for the other server is in the sibling folder.

Your goal is to deploy this to Kubernetes while making it publically available.
Feel free to modify any of the code as necessary.

## Development

You can run the server locally to try it out.
You will need to pass a `uri` flag.

```sh
go run entrypoint.go --uri=http://localhost:3000
```

## Deployment

Deploying this server is a matter of building a Docker image and then deploying
it to a Kubernetes cluster.

```sh
docker build -t stairwell/entrypoint .
kubectl apply -f entrypoint.yaml
```

The Kubernetes configs expose a `NodePort` service.
In its current state, you will not be able to reach it from the public.

```
❯ kubectl --namespace=stairwell get service entrypoint-service
NAME                 TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)           AGE
entrypoint-service   NodePort   10.107.105.71   <none>        42069:30196/TCP   73m
```
