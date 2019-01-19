workflow "new workflow" {
  on = "push"
  resolves = ["golint", "yamllint", "pbpush"]
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