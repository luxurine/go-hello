pipeline {
    agent any
    stages {
        stage('Deploy') {
            steps {

                    retry(3) {
                        sh 'ls /root'
                    }
        }
    }
}
}