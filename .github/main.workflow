workflow "new workflow" {
  on = "push"
  resolves = ["make"]
}

action "make" {
  uses = "golang:1.11"
  runs = "make bins"
}