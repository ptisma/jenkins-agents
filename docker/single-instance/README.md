Running locally one instance of Jenkins inside a docker container, other container serves us a docker cli which jenkins can use.
docker run --name jenkins-docker --rm --detach \
  --privileged --network jenkins --network-alias docker \
  --env DOCKER_TLS_CERTDIR=/certs \
  --volume jenkins-docker-certs:/certs/client \
  --volume jenkins-data:/var/jenkins_home \
  --publish 2376:2376 \
  docker:dind --storage-driver overlay2



docker build -t myjenkins-blueocean:2.426.2-1 .

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