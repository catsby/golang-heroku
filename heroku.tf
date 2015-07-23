resource "heroku_formation" "default" {
  app  = "serene-eyrie-6062"
  type = "web"
  quantity = 2
  size = "2X"
}
