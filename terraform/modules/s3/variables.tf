variable "bucket_name" {
  description = "Name of the S3 bucket"
  type        = string
}


variable "force_destroy" {
  description = "Whether to allow bucket deletion even if it has objects"
  type        = bool
  default     = false
}


variable "enable_sse" {
  description = "Enable server-side encryption (SSE)"
  type        = bool
  default     = false
}


variable "block_public_acls" {
  type    = bool
  default = true
}


variable "block_public_policy" {
  type    = bool
  default = true
}


variable "ignore_public_acls" {
  type    = bool
  default = true
}


variable "restrict_public_buckets" {
  type    = bool
  default = true
}


variable "enable_static_website" {
  description = "Enable S3 static website hosting"
  type        = bool
  default     = false
}

variable "index_document" {
  description = "The index document for the static website (can be empty)"
  type        = string
  default     = ""
}

variable "error_document" {
  description = "The error document for the static website (can be empty)"
  type        = string
  default     = ""
}

variable "enable_public_access" {
  type    = bool
  default = false
}


variable "enable_lifecycle" {
  description = "Enable S3 lifecycle rules"
  type        = bool
  default     = false
}
