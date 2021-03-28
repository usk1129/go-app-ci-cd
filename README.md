# golang-app-ci-cd

This is simple go application with CI-CD pipeline with Docker, Jenkins, Ansible and Kubernetes. The goal is to make automated testing in container app and automated deployment with the app running in container to AWS cloud provider and kubernetes cluster.

We start by dockerizing the application and build the image from it with the tag name where tag name is defined for envorinment variable with the build_number which can be used later on.

We will deploy that docker image to Dockerhub registry and it uses Dockerfile by default and adds the build number as the image tag.Later on, this will be of much importance when you need to determine which Jenkins build was the source of the currently running container.

After the image is built successfully, we push it to Docker Hub using the build number. Additionally, we add the “latest” tag to the image (a second tag) so that we allow users to pull the image without specifying the build number, should they need to.

I have used ansible as deployment tool to the AWS cloud provider. Ansilbe playbook will automatically take the latest image with the lates and deploy that container to the targeted servers and make sure that it is up and running.


There is another deployment in the stage in Jenkinsfile where we apply our deployment and service definition files to kubernetes the cluster(not recommended for production).
I invoked some script to connect to the cluster via ssh and create the deployment and service in the cluster. Kubernetes deployments, which ensures that we have zero downtime for the application when we are changing the container image. This is possible because Deployments use the rolling update method by default to terminate and recreate containers one at a time. Only when the new container is up and healthy does the Deployment terminate the old one.

