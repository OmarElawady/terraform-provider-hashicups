terraform {
  required_providers {
    hashicups = {
      version = "0.2"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

variable "coffee_name" {
  type    = string
  default = "Vagrante espresso"
}

resource "hashicups_coffees" "all" {
  name = "khaled6"

}

resource "hashicups_menu" "all1" {
  name = resource.hashicups_coffees.all.fancy_name
  depends_on = [
    hashicups_coffees.all
  ]
  triggers = {
    coffe_name = hashicups_coffees.all
  }
}

# Returns all coffees
output "all_coffees" {
  value = resource.hashicups_coffees.all.coffees
}

# Only returns packer spiced latte
output "coffee" {
  value = {
    for coffee in resource.hashicups_coffees.all.coffees :
    coffee.id => coffee
    if coffee.name == var.coffee_name
  }
}
