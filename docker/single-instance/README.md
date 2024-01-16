This is a controller only version, which means our controller does builds as well. We will call it controller-builder.

First we spin up a docker in docker container which we are going to use for docker image builds:
docker run --name jenkins-docker --rm --detach \
  --privileged --network jenkins --network-alias docker \
  --env DOCKER_TLS_CERTDIR=/certs \
  --volume jenkins-docker-certs:/certs/client \
  --volume jenkins-data:/var/jenkins_home \
  --publish 2376:2376 \
  docker:dind --storage-driver overlay2

Then we build the image for our Jenkins controller-builder:
docker build -t myjenkins-blueocean:2.426.2-1 .

Keep in mind we have installed only the Docker CLI on it, not the Docker daemon as well, which means we can't build images without the Docker in Docker container.


Run it and preconfigure the DOCKER_HOST variable so we use the Docker in Docker container as a Docker host. We also hook up the volumes for certs, since it the dind runs on TLS by default.
docker run --name jenkins-blueocean --restart=on-failure --detach \
  --network jenkins --env DOCKER_HOST=tcp://docker:2376 \
  --env DOCKER_CERT_PATH=/certs/client --env DOCKER_TLS_VERIFY=1 \
  --publish 8080:8080 --publish 50000:50000 \
  --volume jenkins-data:/var/jenkins_home \
  --volume jenkins-docker-certs:/certs/client:ro \
  myjenkins-blueocean:2.426.2-1


docker exec -it b36edc6cbabd /bin/bash
cat /var/jenkins_home/secrets/initialAdminPassword

Setup Jenkins after initial login using password
44d97ae789b64f95a7822795734b4269

Setup Credentials: DOCKER_USERNAME, DOCKER_PASSWORD
