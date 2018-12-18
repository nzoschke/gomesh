workflow "new workflow" {
  on = "push"
  resolves = ["push-gen"]
}

action "gen" {
  uses = "./.github/action/gen"
}

action "push-gen" {
  needs = ["generate"]
  uses = "docker://debian:9-slim"
  runs = ".github/push-gen.sh"
  secrets = ["GITHUB_TOKEN"]
}