variable "zvar" {
  type        = string
  description = "z var"
  default     = "z"
}

variable "avar" {
  type        = string
  description = "a var"
  default     = "a"
}

variable "bvar" {
  type        = string
  description = "B var"
  default     = "B"
  sensitive = true
}

variable "evar" {
  type        = object({
    name = string
    adress = string
    nr = number
  })
  description = "e var"
  default     = {
    name = "FOO"
    adress = "street"
    nr = 101
  }
}
