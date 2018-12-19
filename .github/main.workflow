workflow "new workflow" {
  on = "push"
  resolves = ["make bins"]
}

action "make bins" {
  uses = "docker://golang:1.11"
  runs = "make bins"
}