pipeline {
    agent any
    environment
    {
        registry = "usko/golang"
        GOCACHE = "/tmp"
    }
    stages{
        stage("Build"){
            agent{
                docker {
                    image 'golang:1.15-buster'
                }
            }
            steps{
                // Create our project directory.               
                sh 'cd ${GOPATH}/src'               
                sh 'mkdir -p ${GOPATH}/src/hello-world'               
                // Copy all files in our Jenkins workspace to our project directory.                              
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/hello-world'               
                // Build the app.
                sh 'go build'

            }
        }
         stage("Test"){
            agent{
                docker {
                    image 'golang:1.15-buster'
                }
            }
            steps{
                                   
                sh 'cd ${GOPATH}/src'               
                sh 'mkdir -p ${GOPATH}/src/hello-world'               
                // Copy all files in our Jenkins workspace to our project directory.                              
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/hello-world'               
                // Remove cached test results.               
                sh 'go clean -cache'               
                // Run Unit Tests.               
                sh 'go test ./... -v -short' 

            }
        }
         stage("Publish"){
            environment{
                registryCredential = "dockerhub"
            }
            steps{
               script 
                {                   
                    def appimage = docker.build registry + ":$BUILD_NUMBER"                   
                    docker.withRegistry( '', registryCredential ) 
                    {                       
                        appimage.push()                       
                        appimage.push('latest')                   
                        
                    }               
                    
                }   
            }
        }
        stage('Deploy to aws server'){
            steps{
                ansiblePlaybook extras: "-e tag=${env.BUILD_NUMBER}", 
                                credentialsId: 'private-key',
                                disableHostKeyChecking: true, 
                                playbook: 'ansible/docker-vagrant.yaml',
                                inventory: 'ansible/inventory'
            }
        }
        /** local minikube cluster
        stage('Deploy t') {
            steps {
                sh 'chmod +x changeTag.sh'
                sh './changeTag.sh ${BUILD_NUMBER}'
                withCredentials([
                    string(credentialsId: 'my-kubernetes', variable: 'api_token')
                    ]) {
                     sh 'kubectl --token $api_token --server https://192.168.64.6:6443 --insecure-skip-tls-verify=true apply -f go-deployment.yaml'
                       }
                    }
                   }
        **/
        stage('Deploy to k8s'){
            steps{
                 sh 'chmod +x changeTag.sh'
                 sh './changeTag.sh ${BUILD_NUMBER}'
                 script {
                    kubernetesDeploy(configs: "go-deployment.yaml", kubeconfigId: "my-kubernetes") 
                    }
                }
              }
    }
}

