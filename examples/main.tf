provider "myprovider" {
  database_filename = "persons.db"
}

resource "myprovider_person" "wim" {
  person_id  = "1"
  last_name  = "Van den Wyngaert"
  first_name = "Wim"
}
