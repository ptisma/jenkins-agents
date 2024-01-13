How to setup a Jenkins Controller on single-node minikube cluster:
helm repo add jenkinsci https://charts.jenkins.io
helm repo update

kubectl create ns jenkins

kubectl apply -f jenkins-volume.yaml

minikube ssh
sudo chown -R 1000:1000 /data/jenkins-volume


kubectl apply -f jenkins-sa.yaml

helm install jenkins -n jenkins -f jenkins-values.yaml jenkinsci/jenkins

Get your 'admin' user password by running:
jsonpath="{.data.jenkins-admin-password}"
secret=$(kubectl get secret -n jenkins jenkins -o jsonpath=$jsonpath)
echo $(echo $secret | base64 --decode)

Get the Jenkins URL to visit by running these commands in the same shell:
jsonpath="{.spec.ports[0].nodePort}"
NODE_PORT=$(kubectl get -n jenkins -o jsonpath=$jsonpath services jenkins)
jsonpath="{.items[0].status.addresses[0].address}"
NODE_IP=$(kubectl get nodes -n jenkins -o jsonpath=$jsonpath)
echo http://$NODE_IP:$NODE_PORT/login




















