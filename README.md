# golang-app-ci-cd

This is simple go application with CI-CD pipeline with Docker, Jenkins, Ansible and Kubernetes. The goal is to make automated testing in container app and automated deployment with the app running in container to AWS cloud provider and kubernetes cluster.

This project provides necessary config files to create docker images for jenkins with pre-installed role-based startegy plugin, validates few test cases and deploy the container image to dockerhub and deploy that image to the running kubernetes cluster in AWS and also ansible to deploy that image to the remote node in AWS and create container with the app in it.

![Untitled Document (1)](https://user-images.githubusercontent.com/32932279/112898649-63b85c00-90e1-11eb-91bf-d4a621b78086.jpg)

1. Jenkins pipeline will automatically trigger whenever there is a commit to the Github repository. GitHub Repository : SCM for fetching the necessary config file

2. It will deploy that docker image to Dockerhub registry and it uses Dockerfile by default and adds the build number as the image tag.Later on, this will be of much importance when you need to determine which Jenkins build was the source of the currently running container.
After the image is built successfully, we push it to Docker Hub using the build number. Additionally, we add the “latest” tag to the image (a second tag) so that we allow users to pull the image without specifying the build number, should they need to.

3. We will deploy that container image to the kubernetes cluster from Jenkins. 
I invoked some script to connect to the cluster via ssh and to create the deployment and node-port service in the cluster. It will replace the tag version from the latest version and create deployment in the running kubernetes cluster. Don't use this for production and ingress-controll service is much secure. 

3. I have also used ansible as deployment tool to the AWS cloud provider. Ansilbe playbook will automatically take the latest image with the latest version and deploy that container to the targeted servers and make sure that it is up and running.


