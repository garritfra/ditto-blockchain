workflow "New workflow" {
  on = "push"
  resolves = ["GitHub Action for Zeit"]
}

action "GitHub Action for Zeit" {
  uses = "actions/zeit-now@9fe84d5"
  secrets = ["GITHUB_TOKEN"]
}
