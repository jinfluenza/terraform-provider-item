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
    default = "mochi"
}

variable "body" {
    type = string
    default = "hi"
}

resource "item_order" "example" {
    title = var.title
    body = var.body
}