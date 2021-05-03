pipeline {
    agent any
    tools {
        go 'go1.16.3'
    }
 
    environment {
        GO111MODULE = 'on'
        //CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    } 
    
    stages {   
        stage('Build') {
            steps {
                shell 'go mod init'
                shell 'go mod tidy'
                shell 'go build'
            }
        } 
        stage('Test') {
            steps {
                shell 'go test ./... -coverprofile=coverage.txt'
                shell "curl -s https://codecov.io/bash | bash -s -"
            }
        }
        stage('Code Analysis') {
            steps {
                shell 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.12.5'
                shell 'golangci-lint run'
            }
        }
        stage('Deploy') {
            when {
              expression {
                currentBuild.result == null || currentBuild.result == 'SUCCESS' 
              }
            }
            steps {
                sh 'make publish'
            }
        }
  
/*
        stage('Build') {
            steps {
                echo 'Compiling and building'
                //sh "${root}/bin/go build"
                sh 'go build'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet .'
                    echo 'Running test'
                    sh 'go test -v' //cd test &&
                }
            }
        }
        */
    }
  /*  post {
        always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "${params.RECIPIENTS}",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"
            
        }
    } */ 
}