workflow "new workflow" {
  on = "push"
  resolves = ["make bins", "yamllint"]
}

action "make bins" {
  uses = "./.github/action/make"
  args = "bins"
}

action "yamllint" {
  uses = "./.github/action/yamllint"
  args = "config/*/*.yaml"
}