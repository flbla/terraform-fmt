variable "zvar" {
  type        = string
  description = "z var"
  default     = "z"
}

variable "avar" {
  type        = string
  description = "a var"
  default     = "a"
  validation {
    condition     = length(var.avar) == 32
    error_message = "Must be a 32 character long string."
  }
}

variable "optionalvar" {
  type        = string
  description = "x var"
  default     = null
}

variable "optionalvar2" {
  type        = string
  description = "x var"
  default     = ""
}

variable "optionalvar3" {
  type        = string
  description = "x var"
  default     = {}
}

variable "optionalvar4" {
  type        = string
  description = "x var"
  default     = []
}

variable "xvar" {
  type        = string
  description = "x var"
}

variable "bvar" {
  type        = string
  default     = "B"
  sensitive   = true
  description = "B var"
}

variable "evar" {
  type = object({
    name   = string
    adress = string
    nr     = number
  })
  description = "e var"
  default = {
    name   = "FOO"
    adress = "street"
    nr     = 101
  }
}
