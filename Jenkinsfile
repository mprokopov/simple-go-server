// Declarative Pipeline — top-level block required for `Jenkinsfile` syntax.
// The alternative is `node { ... }` (Scripted Pipeline), but Declarative is
// the recommended style for most CI/CD use cases.
pipeline {
    // `agent any` means: run on any available Jenkins agent (executor).
    // For tighter control you can target a label, e.g. `agent { label 'linux' }`,
    // or run inside a container with `agent { docker { image '...' } }`.
    agent any

    // `tools` auto-installs and PATH-injects a tool defined under
    // Manage Jenkins → Tools. The name "1.26" must match a Go installation
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
    }
}
