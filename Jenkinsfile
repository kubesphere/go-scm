pipeline{

  agent {
    node {
      label 'go'
    }
  }

    parameters {
        string(name:'TAG_NAME',defaultValue: '',description:'')
    }

    environment {
        GITHUB_CREDENTIAL_ID = 'github-id'
    }

  stages {
          stage ('checkout scm') {
              steps {
                  checkout(scm)
              }
          }

        stage ('unit test') {
            steps {
                container ('go') {
                    sh 'export KUBEBUILDER_CONTROLPLANE_START_TIMEOUT=1m; go test ./scm/... -covermode=atomic -coverprofile=coverage.txt'
                }
            }
        }

        stage('push with tag'){
          when{
            expression{
              return params.TAG_NAME =~ /v.*/
            }
          }
          steps {
              container ('go') {
                input(id: 'release-image-with-tag', message: 'release image with tag?')
                  withCredentials([usernamePassword(credentialsId: "$GITHUB_CREDENTIAL_ID", passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]) {
                    sh 'git config --global user.email "kubesphere@yunify.com" '
                    sh 'git config --global user.name "kubesphere" '
                    sh 'git tag -a $TAG_NAME -m "$TAG_NAME" '
                    sh 'git push http://$GIT_USERNAME:$GIT_PASSWORD@github.com/kubesphere/go-scm.git --tags --ipv4'
                  }
          }
          }
        }
  }
}