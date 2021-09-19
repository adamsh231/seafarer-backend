pipeline{
    agent any

    triggers {
        githubPush()
    }

    environment{
        // ----------- project initiate ------------ //
        PROJECT_NAME = 'seafarer-backend'
        PROJECT_TYPE = 'services'
        // ----------------------------------------- //

        // credential
        TELEGRAM_BOT_TOKEN = credentials('telegram-bot-token')
        TELEGRAM_GROUP_CHAT_ID = credentials('telegram-group-chat-id')
        SSH_PASS = credentials('qword-development-root-password')

        // compute variable
        SSH_COMMAND = sh (
            script: "echo -n 'sshpass -p ${SSH_PASS} ssh root@103.102.153.44'",
            returnStdout: true
        )
        TELEGRAM_COMMAND = sh (
            script: "echo -n 'curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage?parse_mode=html -d chat_id=${TELEGRAM_GROUP_CHAT_ID} -d'",
            returnStdout: true
        )
        AUTHOR_NAME = sh (
            script: "git show -s --pretty=%an",
            returnStdout: true
        ).trim()
    }

    stages{

        // -------------------------------- Master ----------------------------------- //
        stage("build-master"){
            when {
                branch 'master'
            }
            environment{
                // ----------- branch initiate ------------ //
                BRANCH_DIR = 'production'
                // ----------------------------------------- //

                // compute variable
                GO_TO_DIR = sh (
                    script: "echo -n 'cd /home/${BRANCH_DIR}/${PROJECT_TYPE}/${PROJECT_NAME}; ls -l;'",
                    returnStdout: true
                )
            }
            steps{
                sh '${TELEGRAM_COMMAND} text="Build started by <b>${AUTHOR_NAME}</b>%0A%0A<b>${JOB_NAME}</b> - ${BUILD_DISPLAY_NAME}%0A%0A(<a href=\'${BUILD_URL}console\'>Open</a>)"'
                sh '${SSH_COMMAND} "${GO_TO_DIR} git pull origin ${GIT_BRANCH};"'
                sh '${SSH_COMMAND} "${GO_TO_DIR} docker-compose up -d --build;"'
                sh '${SSH_COMMAND} "docker image prune -f;"'
            }
            post{
                success{
                    sh '${TELEGRAM_COMMAND} text="Successfully build%0A%0A<b>${JOB_NAME}</b> - ${BUILD_DISPLAY_NAME}%0A%0A(<a href=\'${BUILD_URL}console\'>Open</a>)"'
                }
                failure{
                    sh '${TELEGRAM_COMMAND} text="Failed to build%0A%0A<b>${JOB_NAME}</b> - ${BUILD_DISPLAY_NAME}%0A%0A(<a href=\'${BUILD_URL}console\'>Open</a>)"'
                }
            }
        }
        // --------------------------------------------------------------------------- //

    }

}