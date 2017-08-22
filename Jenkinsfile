#!/usr/bin/env groovy
node {
    stage('Build') {
        dir('git.home.foxienet.com/hostnotes/projectservice') {
            checkout scm
        }
        // sh('pwd && ls -Rl')
        withDockerRegistry([credentialsId: 'jenkins-docker-reg', url: 'https://registry.svc.brettmack.me']) {

            sh('docker run --rm -v $(pwd):/go/src registry.svc.brettmack.me/gobuild:0.1 /bin/bash -c \"cd /go/src/git.home.foxienet.com/hostnotes/projectservice && dep ensure -v\"')
        }
    }
}