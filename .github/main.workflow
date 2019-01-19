workflow "new workflow" {
  on = "push"
  resolves = ["make bins", "yamllint"]
}

action "make bins" {
  uses = "docker://golang:1.11"
  runs = "make bins"
}

action "yamllint" {
  uses = "./.github/action/yamllint"
  runs = "yamllint config/*/*.yaml"
}