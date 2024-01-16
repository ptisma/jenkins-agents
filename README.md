I wanted to explore more in depth about Jenkins controller-agent architecture but without hosting the VM-s on cloud, rather get dirty and use the containers on my own local host machine (WSL2). Spoiler alert it does get complicated :D.

## source-repo

Contains the simple Go based application which is used by Jenkins for testing, building and pushing to the container registry.

## docker

Contains the Jenkins controller-agent setup using docker containers (single instance running Jenkins on one node bad practice spoiler alert, Jenkins controller using static VM-s as agents or using containers as dynamically spun agents)

## kubernetes

We use Helm to deploy a Jenkins controller. Its basically a stateful set. It's preconfigured to use the cloud plugin for Kubernetes, esentially everytime we trigger the build, it's run inside the new temporary pod. It follows the dynamic agent architecture.

Inside the pod we also have to follow up a sidecar pattern (we are using the docker build in pipeline duh) since the default agent image doesn't have the docker cli (we don't care about that tbh since it can be installed via tools) and daemon installed.

## dind-test

Its a simple pod yaml with dind sidecar for local testing. docker:dind runs on TLS mode by default, use docker:18.05-dind for default non TLS.

