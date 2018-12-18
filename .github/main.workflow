workflow "new workflow" {
  on = "push"
  resolves = ["push-gen"]
}

action "gen" {
  uses = "./.github/action/gen"
}

action "push-gen" {
  needs = ["gen"]
  uses = "./.github/action/gen"
  runs = ".github/push-gen.sh"
  secrets = ["GITHUB_TOKEN"]
}