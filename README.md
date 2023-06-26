# Stairwell Kubernetes Takehome

If you're reading this, you probably like Kubernetes.
Or not.
I don't know; I'm just a README file.
This repository contains starter code for two servers.
In a nutshell, your goal is to deploy them and make them talk to each other.

The first server is in the `entrypoint` folder.
You can think of it as a user-facing API server.
It should be on Kubernetes and available to the public Internet.
Is has the second server as a dependency.

The second server is in the `generator` folder.
You can think of it as a backend server with high compute requirements.
Since cloud costs (in this hypothetical scenario) can get expensive, we DO NOT
want this on Kubernetes.

## Prerequisites

There are a couple of dependencies that you'll need:
- `python`
- `go`
- `docker`
- `kubectl`

This takehome assumes working knowledge over all of these.

## Requirements

There are many ways to solve this takehome.
The actual code and implementation matters less than the end behavior.
Here are the specific conditions we require:

1. The `entrypoint` server must be deployed to Kubernetes.
2. The `entrypoint` server must be publicly reachable.
3. The `generator` server must NOT be deployed to Kubernetes.
4. The `generator` server must NOT be publicly reachable.
5. The `entrypoint` server must be able to reach the `generator` server.

You will be expected to demo this and show that it works.

## Other Pro Tips
- When developing with Kubernetes, you can use `minikube` or some local
  single-node cluster.
  At least that's what I did when creating this takehome challenge.
- Feel free to modify any of the code as necessary.
  The solution is not meant to be prescriptive in any way.
- Reach out if you have questions.
  This is meant to be a collaborative challenge, not some form of medieval
  torture.
