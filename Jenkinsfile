pipeline {
    agent any
    stages {
        stage('Deploy') {
            steps {

                timeout(time: 3, unit: 'MINUTES') {
                    retry(3) {
                        sh 'ls /root'
                    }
                }
        }
    }
}
}