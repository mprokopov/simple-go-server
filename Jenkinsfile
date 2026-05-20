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
       go "1.26"
    }

    stages {
        // Each `stage` shows up as a column in the Jenkins Blue Ocean / pipeline
        // view, making it easy to see where a build passed or failed.
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
