Prerequisite:
The Jenkins controller has already been setup (see /docker/single-instance)
Add a new Node:
This option allows you to add a new individual Jenkins agent (or node) to your Jenkins environment. A node in Jenkins represents a machine (physical or virtual) that can execute tasks as part of your Jenkins jobs or pipelines. Adding a new node involves defining the details of a single machine that Jenkins can use to offload work from the master Jenkins instance. You can specify the node's name, labels, connection details (such as its IP address, credentials, or connection method), and other configurations specific to that node.


Generate an SSH key pair:
ssh-keygen -f ~/.ssh/jenkins_agent_key

Add Jenkins SSH credential inside Manage Credentials

Create a Docker agent (node for our controller in Jenkins context):
docker run -d --rm --name=agent1 --network jenkins -p 22:22 \
-e "JENKINS_AGENT_SSH_PUBKEY=[your-public-key]" \
jenkins/ssh-agent:alpine-jdk17


Public key can be found by cating the  ~/.ssh/jenkins_agent_key.pub


ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDg32MF5bsXDeK5B6QcS4rM5fFxD5JxtTgU2/nk4X5ZXYcuZU4NqVRJiMIHqh+fnCfMWYLBkuC1C9ksIQjCgTj7C3RaqYuRiFIOC/4lOnRtE4R1LemnwwuxHcdUOVoODV0iVeKI9pUsfEkJHS5vmO/BHOa9RnANnz8NGeVupphJPfyhsb2+aHsZut1LDqEdP/e6lSTG5DJ2zCDUG4gqd2yCxPeTEHWPFpYHAyY234pc+YCBLuofuXuQLwZYK9zQyq+wldi1dpev8Pdah8mhnnlDqWfA5ZO6bpZI8vQE+ZyxHhFoLyimKbdqCBa6lAov5Z7AwlxRQcPYm7azBi06ptcBYc8tpVywdWYKUJ7cNzvxQD0RWfFjx07lECtSefsYVveSImUE5ckfJOA1VY/IsJGjrHGGL1RvNTv9I9Bio5z2kmzP74o3JiVDzkfSCQGzLD9k6nvha9Qb23uu3kXeEnbdnfFS4y4OpMKayhfQ2i9ZfxxHMGOtsuII+fvEV1T5fg0= ptisma@LENPF3XNVRK

Setup Jenkins agent:
Fill the Node/agent name and select the following options:
Name: agent1, Type: Permanent Agent

Now fill the fields:
Remote root directory: /home/jenkins
label: agent1
usage: only build jobs with label expression
Launch method: Launch agents by SSH
Host: host.docker.internal
Credentials: jenkins
Host Key verification Strategy: Manually trusted key verification



