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
    required            = true
  }

  schema {
    name                = "last_name"
    attribute_data_type = "String"
    mutable             = true
    required            = true
  }

  schema {
    name                = "google_login"
    attribute_data_type = "Boolean"
    mutable             = true
    required            = true
  }

  password_policy {
    minimum_length = 8
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
  name                  = "my-app-client"
  user_pool_id          = aws_cognito_user_pool.myPool.id
  access_token_validity = 3600 # Access token validity in seconds (e.g., 1 hour)
  id_token_validity     = 3600

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

