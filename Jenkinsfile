#!/usr/bin/env groovy
node {
    stage('Build') {
        dir('git.home.foxienet.com/hostnotes') {
            checkout scm
        }
        sh('pwd && ls -Rl')
        // withDockerRegistry([credentialsId: 'jenkins-docker-reg']) {

    }
}