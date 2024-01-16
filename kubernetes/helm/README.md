Note, minikube was used.

In order to setup the Jenkins Controller using Helm we first need to add the Helm chart:
helm repo add jenkinsci https://charts.jenkins.io
helm repo update

After this we are going to create a separate namespace for our Jenkins stuff:
kubectl create ns jenkins

Instead of Helm automatically creating the volumes and service account for us, we are going to do it manually:
kubectl apply -f jenkins-volume.yaml
kubectl apply -f jenkins-sa.yaml

Install the Jenkins using Helm:
helm install jenkins -n jenkins -f jenkins-values.yaml jenkinsci/jenkins

Extra step:
minikube ssh
sudo chown -R 1000:1000 /data/jenkins-volume

Access the Jenkins controller locally:
kubectl port-forward service/jenkins -n jenkins 8080:8080

Retrieve the password:
jsonpath="{.data.jenkins-admin-password}"
secret=$(kubectl get secret -n jenkins jenkins -o jsonpath=$jsonpath)
echo $(echo $secret | base64 --decode)

Since our example pipeline uses the tools such as go and dockerTool, we are going to install those plugins on Jenkins controller, this can be either done via GUI:
Manage Jenkins -> Plugins -> Available Plugins, then install Go and Docker Commons Plugin and then
Manage Jenkins -> Tools and configure them to match the exact "versions" from the Jenkinsfile (note for docker tick the install automatically and specify the versions)
or editing the helm values file controller.installPlugins

Create the Credentials for Docker hub:
Manage Jenkins -> Manage Credentials -> Credential with username and password, has to match the Jenkinsfile

Keep in mind, the docker:dind side container inside our pod runs in non TLS mode, we have to modify the command since by default it runs in TLS mode. We do that by setting the env vars DOCKER_TLS_CERTDIR and DOCKER_TLS_VERIFY.





















