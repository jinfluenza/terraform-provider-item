terraform {
  required_providers {
    item = {
        version = "0.1.0"
        source = "jinfluenza.com/edu/item"
    }
  }
}

variable "title" {
    type = string
    default = "Marquel"
}

variable "body" {
    type = string
    default = "Bye"
}

resource "item_order" "example" {
    title = var.title
    body = var.body
}