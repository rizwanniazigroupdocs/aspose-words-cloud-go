properties([
	gitLabConnection('gitlab'),
	[$class: 'ParametersDefinitionProperty', 
		parameterDefinitions: [
			[$class: 'StringParameterDefinition', name: 'branch', defaultValue: 'master', description: 'the branch to build'],
			[$class: 'StringParameterDefinition', name: 'apiUrl', defaultValue: 'https://api-qa.aspose.cloud', description: 'api url']
		]
	]
])

node('windows2019') {
	try {
		gitlabCommitStatus("checkout") {
			stage('checkout'){
				checkout([$class: 'GitSCM', branches: [[name: params.branch]], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'LocalBranch', localBranch: "**"]], submoduleCfg: [], userRemoteConfigs: [[credentialsId: '361885ba-9425-4230-950e-0af201d90547', url: 'https://git.auckland.dynabic.com/words-cloud/words-cloud-go.git']]])
			}
		}
		gitlabCommitStatus("tests") {
			stage('tests') {
				withCredentials([usernamePassword(credentialsId: '6839cbe8-39fa-40c0-86ce-90706f0bae5d', passwordVariable: 'AppKey', usernameVariable: 'AppSid')]) {
					def apiUrl = params.apiUrl
					bat "echo {\"AppSid\":\"%AppSid%\",\"AppKey\":\"%AppKey%\",\"BaseUrl\":\"%apiUrl%\" } > config.json"
				}
				try {
					bat 'docker run -v %cd%:c:/sdk -w="c:/sdk" --rm -t golang:1.14.0-windowsservercore-1809 go test ./tests/... -v 2>&1 | utils\\go-junit-report > testResults\\tests.xml'
				} 
				finally {
					junit '**\\testResults\\tests.xml'
				}
			}
		}
	} finally {
		bat 'docker system prune -f'
		deleteDir()
	}
}