workflow "go" {
  on = "push"
  resolves = ["golint"]
}

workflow "pb" {
  on = "push"
  resolves = ["pbpush"]
}

workflow "yaml" {
  on = "push"
  resolves = ["yamllint"]
}


action "golint" {
  uses = "./.github/action/go"
  runs = ".github/golint.sh"
}

action "pbpush" {
  uses = "./.github/action/prototool"
  runs = ".github/pbpush.sh"
  secrets = ["PUSH_TOKEN"]
}

action "yamllint" {
  uses = "./.github/action/yamllint"
  runs = ".github/yamllint.sh"
}