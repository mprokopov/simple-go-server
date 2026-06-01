// Declarative Pipeline — top-level block required for `Jenkinsfile` syntax.
// The alternative is `node { ... }` (Scripted Pipeline), but Declarative is
// the recommended style for most CI/CD use cases.
pipeline {
    // `agent any` means: run on any available Jenkins agent (executor).
    // For tighter control you can target a label, e.g. `agent { label 'linux' }`,
    // or run inside a container with `agent { docker { image '...' } }`.
    agent any

    // `tools` auto-installs and PATH-injects a tool defined under
    // Manage Jenkins → Tools. The name "1.24.1" must match a Go installation
    // configured there; otherwise the build fails before any stage runs.
    tools {
       go "1.24.1"
    }

    stages {
        // Each `stage` shows up as a column in the Jenkins Blue Ocean / pipeline
        // view, making it easy to see where a build passed or failed.

        // Tests run *before* Build: fail-fast means we don't waste time
        // producing a binary when the logic is already known broken.
        stage('Test') {
            steps {
                // Single-file mode (no go.mod in this repo), so we name the
                // sources explicitly. With a module, `go test ./...` would
                // discover packages automatically.
                sh "go test main.go main_test.go"
            }
        }

        stage('Build') {
            steps {
                // `sh` runs a shell command on the agent. The Go toolchain
                // installed above is on PATH, so `go build` resolves correctly.
                // Output binary defaults to `main` (named after main.go).
                sh "go build main.go"
            }
        }

        stage('Deploy') {
            steps {
                // `withCredentials` binds a Jenkins credential to env vars only
                // for the duration of this block, then revokes them. The key
                // file is written to a temp path and removed on exit — it
                // never persists on the agent disk.
                //
                // credentialsId must match the ID you set when creating the
                // SSH credential in Jenkins (Manage Jenkins → Credentials).
                withCredentials([sshUserPrivateKey(
                    credentialsId: 'ubuntu',
                    keyFileVariable: 'SSH_KEY',
                    usernameVariable: 'SSH_USER'
                )]) {
                    // Single-quoted heredoc: Groovy does NOT interpolate
                    // here, so $SSH_KEY reaches the shell intact and
                    // Jenkins's credential-masker can scrub it from logs.
                    //
                    // ANSIBLE_HOST_KEY_CHECKING=False skips the interactive
                    // "are you sure you want to continue connecting" prompt
                    // on first contact — fine for an ephemeral lab host,
                    // would be a TOFU footgun in production.
                    //
                    // Ansible's `copy` module writes to a temp path and
                    // rename(2)s into place, so it dodges ETXTBSY on the
                    // running binary natively — no scp-then-mv dance needed.
                    sh '''
                        ANSIBLE_HOST_KEY_CHECKING=False \
                            ansible-playbook \
                                -i ansible/inventory.ini \
                                --private-key "$SSH_KEY" \
                                ansible/deploy.yml
                    '''
                }
            }
        }
    }
}
