provider "myprovider" {
  database_filename = "db.json"
}

resource "myprovider_person" "wim" {
  name_id    = "1"
  last_name  = "Van den Wyngaert"
  first_name = "Wim"
}
