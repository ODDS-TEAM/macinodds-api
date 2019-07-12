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
                    withDockerRegistry(credentialsId: credential_id, url: registry) {
                        to_push.push()
                    }
                }
            }
        }
        stage ('Deploy') {
            steps {
                sshPublisher(publishers: [sshPublisherDesc(configName: 'macinodds.tk', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: '''docker rm -f macinodds-api
                docker login -u ${docker_user} -p ${docker_password} registry.odds.team 2>/dev/null
                docker pull registry.odds.team/internship/macinodds-api:${BUILD_NUMBER}
                docker run --name macinodds-api -p 1323:1323 -v /home/nc-user/images:/app/mac -d registry.odds.team/internship/macinodds-api:${BUILD_NUMBER}''', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
            }
        }
    }
}