pipeline {
    agent any

    environment {
        APP_NAME = "app"
        DOCKER_IMAGE = "golang:1.24.3"
        GOCACHE_DIR = "${WORKSPACE}/.cache"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build inside Docker') {
            steps {
                sh """
                mkdir -p ${GOCACHE_DIR}
                docker run --rm \
                  -v ${WORKSPACE}:/app -w /app \
                  -e GOCACHE=/app/.cache \
                  ${DOCKER_IMAGE} \
                  sh -c "go mod tidy && go mod download && go test ./... && go build -buildvcs=false -o ${APP_NAME}"
                """
            }
        }

        stage('Archive') {
            steps {
                archiveArtifacts artifacts: "${APP_NAME}", fingerprint: true
            }
        }
    }

    post {
        success {
            echo "✅ Build, test, and archive successful!"
        }
        failure {
            echo "❌ Build failed."
        }
    }
}
