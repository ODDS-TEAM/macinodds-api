def to_push
pipeline {
    agent any
    options {
        ansiColor('xterm')
    }
    environment {
        host="registry.odds.team"
        registry="https://$host"
        group="internship"
        image="macinodds-api:${BUILD_NUMBER}"
        img="${host}/${group}/${image}"
        credential_id="f3045e57-33d2-4b1f-8cf1-fca92b6df613"
    }
    stages {
        stage ('Build') {
            steps {
                script {
                    to_push = docker.build(img)
                }   
            }
        }
        stage ('Push') {
            steps {
                script {
                    withDockerRegistry(
                        credentialsId: credential_id, 
                        url: registry) {
                            to_push.push()
                    }
                }
            }
        }
        stage ('Deploy') {
            steps {
                sshPublisher(
                    publishers:
                     [
                         sshPublisherDesc(
                             configName: 'macinodds.tk', 
                             transfers: [
                                 sshTransfer(
                                    cleanRemote: false, 
                                    excludes: '', 
                                    execCommand: 'BUILD_NUMBER=${BUILD_NUMBER} ./reload.sh', 
                                    execTimeout: 120000, 
                                    flatten: false,
                                    makeEmptyDirs: false, 
                                    noDefaultExcludes: false, 
                                    patternSeparator: '[, ]+', 
                                    remoteDirectory: 'api', 
                                    remoteDirectorySDF: false, 
                                    removePrefix: '', 
                                    sourceFiles: 'docker-compose.yaml')
                            ], 
                            usePromotionTimestamp: false, 
                            useWorkspaceInPromotion: false, 
                            verbose: false
                        )
                    ]
                )
            }
        }
    }
    post {
        success {
            slackSend teamDomain: 'for-odds-team', 
                tokenCredentialId: 'slack-for-odds-team', 
                username: 'admin', 
                color: "good", 
                message: "🎉SUCCESS: ${env.JOB_NAME} #${env.BUILD_NUMBER} 😀 (<${env.BUILD_URL}|Open>)"
        }
        failure {
            slackSend teamDomain: 'for-odds-team', 
                tokenCredentialId: 'slack-for-odds-team', 
                username: 'admin', 
                color: "danger", 
                message: "❗️FAILURE: ${env.JOB_NAME} #${env.BUILD_NUMBER} 🤢 (<${env.BUILD_URL}|Open>)"
        }
        unstable {
            slackSend teamDomain: 'for-odds-team', 
                tokenCredentialId: 'slack-for-odds-team', 
                username: 'admin', 
                color: "warning", 
                message: "⚠️UNSTABLE: ${env.JOB_NAME} #${env.BUILD_NUMBER} 😕 (<${env.BUILD_URL}|Open>)"
        }
    }
}