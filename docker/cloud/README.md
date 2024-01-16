##"Configure Clouds" refers to setting up cloud-based agents or dynamic provisioning of resources. 
Jenkins supports various cloud providers (like AWS, Azure, Google Cloud, etc.) or virtualization solutions (such as Docker, Kubernetes, OpenStack, etc.) to dynamically create and manage Jenkins agents based on demand. 
This feature enables Jenkins to automatically provision agents in the cloud or virtual environments when there's a workload, and then remove them when the workload decreases. 
You can define cloud configurations, specifying the cloud provider details, templates for agents, and other settings for dynamic agent provisioning.

Spin up the local Docker build server which is going to dynamically spawn docker container which are going to serve as agents for our Jenkins Controller:
docker run --name jenkins-docker --rm --detach \
  --privileged --network jenkins --network-alias docker \
  --env DOCKER_TLS_CERTDIR=/certs \
  --volume jenkins-docker-certs:/certs/client \
  --volume jenkins-data:/var/jenkins_home \
  --publish 2376:2376 \
  docker:dind --storage-driver overlay2

It's the same one we use in single-instance-folder.

Since inside the setup for docker cloud we use a TLS connection we need to setup the credentials inside Jenkins.

Go to Manage Credentials -> System -> Global credentials -> Create a new credential -> X.509 Client Certificate -> "docker-tls".

In Client Key paste the key.pem

In Client Certificate paste the cert.pem

In Server CA certificate paste the ca.pem

You can find the *.pem files inside the docker volumes. Since we are using in single-instance folder a docker-in-docker build server "under the hood", we hook up the CLI with the daemon, here the setup uses the TLS connection for check-up so we need to have the credentials.

Setup docker cloud with following settings:

Host: tcp://docker:2376 (docker is an network alias for previously created docker-in-docker build server).

Server credentials: docker-tls (one we created in previous step)

Keep in mind, restarting the docker-in-docker will change the certiifates.

What about the volumes for docket sockets?


  
