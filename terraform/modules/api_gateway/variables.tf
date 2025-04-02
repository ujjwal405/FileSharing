variable "endpoint_type" {
  type    = string
  default = "REGIONAL"
  validation {
    condition     = contains(["EDGE", "REGIONAL", "PRIVATE"], var.endpoint_type)
    error_message = "Invalid endpoint type. Must be one of: EDGE, REGIONAL, PRIVATE."
  }
}
