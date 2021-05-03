pipeline {
    agent any
    tools {
        go 'go1.16.3'
    }
 
    environment {
        GO111MODULE = 'on'
        //CGO_ENABLED = 0 
        //GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    } 
    
    stages {   
        stage('Compile') {
            steps {
                shell 'go version'
                shell 'go build'
            }
        } 
        stage('Test') {
            environment {
                CODECOV_TOKEN = credentials('codecov_token')
            }
            steps {
                sh 'go test ./... -coverprofile=coverage.txt'
                sh "curl -s https://codecov.io/bash | bash -s -"
            }
        }
        stage('Code Analysis') {
            steps {
                sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.12.5'
                sh 'golangci-lint run'
            }
        }
        stage('Release') {
            when {
                buildingTag()
            }
            environment {
                GITHUB_TOKEN = credentials('github_token')
            }
            steps {
                sh 'curl -sL https://git.io/goreleaser | bash'
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
                    sh ' go test -v' //cd test &&
                }
            }
        }
        */
    }
    post {
        always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "${params.RECIPIENTS}",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"
            
        }
    }  
}