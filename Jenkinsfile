pipeline{
    agent{
        kubernetes()
        {
            cloud 'kubernetes'
            inheritFrom 'golang-agent-template'

        }
    }
    stages{
        stage('checkout'){
            steps{
                checkout scm
            }
        }
        stage('Download Dependencies'){
            steps{
                container('golang-agent'){
                    sh 'go mod tidy'
                }
            }
        }
    }
}