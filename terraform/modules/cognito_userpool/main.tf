# Define the Cognito User Pool
resource "aws_cognito_user_pool" "myPool" {
  name = "mypool"

  # Configure email as the primary username
  username_attributes = ["email"]

  # Define custom attributes used in the Go function
  schema {
    name                = "first_name"
    attribute_data_type = "String"
    mutable             = true
    required            = false
  }

  schema {
    name                = "last_name"
    attribute_data_type = "String"
    mutable             = true
    required            = false
  }

  schema {
    name                = "google_login"
    attribute_data_type = "Boolean"
    mutable             = true
    required            = false
  }

  password_policy {
    minimum_length = 8
  }

  lifecycle {
    ignore_changes = [schema]
  }

  auto_verified_attributes = ["email"]

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email" # Use email for recovery
      priority = 1
    }
  }

  # Configure email settings (optional, for Cognito-managed emails)
  email_configuration {
    email_sending_account = "COGNITO_DEFAULT"
  }
}

# Define the App Client for the User Pool
resource "aws_cognito_user_pool_client" "myClient" {
  name                   = "my-app-client"
  user_pool_id           = aws_cognito_user_pool.myPool.id
  access_token_validity  = var.access_token_validity
  id_token_validity      = var.id_token_validity
  refresh_token_validity = var.refresh_token_validity


  explicit_auth_flows = [
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_ADMIN_USER_PASSWORD_AUTH"
  ]

  # Specify attributes the app client can write to
  write_attributes = [
    "email",
    "custom:first_name",
    "custom:last_name",
    "custom:google_login"
  ]


  read_attributes = [
    "email",
    "custom:first_name",
    "custom:last_name",
    "custom:google_login"
  ]
}

