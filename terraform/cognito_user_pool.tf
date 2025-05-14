// creating cognito userPool and userPool client
module "cognito_user_pool" {
  source                 = "./modules/cognito_userpool"
  access_token_validity  = var.access_token_validity
  id_token_validity      = var.id_token_validity
  refresh_token_validity = var.refresh_token_validity
}
