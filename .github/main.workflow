workflow "New workflow" {
  on = "push"
  resolves = ["Docker Tag"]
}

action "Setup Go for use with actions" {
  uses = "actions/setup-go@632d18fc920ce2926be9c976a5465e1854adc7bd"
}

action "Docker Tag" {
  uses = "actions/docker/tag@fe7ed3ce992160973df86480b83a2f8ed581cd50"
  needs = ["Setup Go for use with actions"]
}

workflow "New workflow 1" {
  on = "push"
}
