pipeline {
    agent {
        docker {
            image 'golang:1.24.3'
        }
    }

    stages {
        stage('Deps') {
            steps {
                sh 'go mod tidy && go mod download'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o app'
            }
        }
    }
}
