pipeline {
    agent any

    tools {
       go "1.26"
    }

    stages {
        stage('Build') {
            steps {
                sh "go build main.go"
            }
        }
    }
}
