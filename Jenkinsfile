pipeline {
    agent any

    stages {
        stage('Build inside Docker') {
            steps {
                sh '''
                docker run --rm -v $WORKSPACE:/app -w /app golang:1.24.3 \
                  sh -c "export GOCACHE=/app/.cache && mkdir -p /app/.cache && go mod tidy && go mod download && go test ./... && go build -o app"
                '''
            }
        }
    }
}
