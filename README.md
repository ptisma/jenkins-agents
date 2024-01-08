I wanted to explore more in depth about Jenkins controller-agent architecture but without hosting the VM-s on cloud, rather use the containers on my own local host machine (WSL2)

## source-repo

Contains the simple Go based application which is used by Jenkins for testing, building and pushing to the container registry.

## docker

Contains the Jenkins controller-agent setup using docker containers (single instance running Jenkins on one node, Jenkins controller using static VM-s as agents or using containers as dynamically spun agents)
