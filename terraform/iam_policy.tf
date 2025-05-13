// IAM policy for lambdaBasicExecutionRole
data "aws_iam_policy" "lambda_basic_execution" {
  arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

// IAM policy cognitoAdminDelete
module "iam_policy_cognito_admin_delete" {
  source      = "./modules/iam/policy"
  policy_name = "cognito_admin_delete_user"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["cognito-idp:AdminDeleteUser"],
      "Resource" : [module.cognito_user_pool.user_pool_arn]
    }
  ]
}

// IAM policy for cognitoInitiateAuth
module "iam_policy_cognito_initiate_auth" {
  source      = "./modules/iam/policy"
  policy_name = "cognito_initiate_auth"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["cognito-idp:InitiateAuth"],
      "Resource" : [module.cognito_user_pool.user_pool_arn]
    }
  ]
}

// IAM policy for cognitoConfirmSignUP
module "iam_policy_cognito_confirm_sign_up" {
  source      = "./modules/iam/policy"
  policy_name = "cognito_confirm_sign_up"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["cognito-idp:ConfirmSignUp"],
      "Resource" : [module.cognito_user_pool.user_pool_arn]
    }
  ]
}


// IAM policy for cognitoConfirmForgetPassword
module "iam_policy_cognito_confirm_forget_password" {
  source      = "./modules/iam/policy"
  policy_name = "cognito_confirm_forget_password"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["cognito-idp:ConfirmForgotPassword"],
      "Resource" : [module.cognito_user_pool.user_pool_arn]
    }
  ]
}


// IAM policy for cognitoForgetPassword
module "iam_policy_cognito_forget_password" {
  source      = "./modules/iam/policy"
  policy_name = "cognito_forget_password"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["cognito-idp:ForgotPassword"],
      "Resource" : [module.cognito_user_pool.user_pool_arn]
    }
  ]
}


// IAM policy for cognitoRevokeToken
module "iam_policy_cognito_revoke_token" {
  source      = "./modules/iam/policy"
  policy_name = "cognito_revoke_token"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["cognito-idp:RevokeToken"],
      "Resource" : [module.cognito_user_pool.user_pool_arn]
    }
  ]
}


// IAM policy for cognitoAdminGetUser
module "iam_policy_cognito_admin_get_user" {
  source      = "./modules/iam/policy"
  policy_name = "cognito_admin_get_user"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["cognito-idp:AdminGetUser"],
      "Resource" : [module.cognito_user_pool.user_pool_arn]
    }
  ]
}




// IAM policy for cognitoAdminCreateUser
module "iam_policy_cognito_admin_create_user" {
  source      = "./modules/iam/policy"
  policy_name = "cognito_admin_create_user"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["cognito-idp:AdminCreateUser"],
      "Resource" : [module.cognito_user_pool.user_pool_arn]
    }
  ]
}



// IAM policy for cognitoSignUp
module "iam_policy_cognito_sign_up" {
  source      = "./modules/iam/policy"
  policy_name = "cognito_sign_up"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["cognito-idp:SignUp"],
      "Resource" : [module.cognito_user_pool.user_pool_arn]
    }
  ]
}




// IAM policy for dynamoGetItem
module "iam_policy_dynamo_get_item_user_table" {
  source      = "./modules/iam/policy"
  policy_name = "dynamo_get_item_user_table"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["dynamodb:GetItem"],
      "Resource" : [module.dynamo_user_meta_data.table_arn]
    }
  ]
}


// IAM policy for dynamoUpdateItem
module "iam_policy_dynamo_update_item_user_table" {
  source      = "./modules/iam/policy"
  policy_name = "dynamo_update_item_user_table"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["dynamodb:UpdateItem"],
      "Resource" : [module.dynamo_user_meta_data.table_arn]
    }
  ]
}



// IAM policy for dynamoPutItem
module "iam_policy_dynamo_put_item_user_table" {
  source      = "./modules/iam/policy"
  policy_name = "dynamo_put_item_user_table"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["dynamodb:PutItem"],
      "Resource" : [module.dynamo_user_meta_data.table_arn]
    }
  ]
}

// IAM policy for dynamoPutItem
module "iam_policy_dynamo_put_item_file_table" {
  source      = "./modules/iam/policy"
  policy_name = "dynamo_put_item_file_table"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["dynamodb:PutItem"],
      "Resource" : [module.dynamo_file_meta_data.table_arn]
    }
  ]
}

// IAM policy for dynamoQuery
module "iam_policy_dynamo_query_file_table" {
  source      = "./modules/iam/policy"
  policy_name = "dynamo_query_file_table"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["dynamodb:Query"],
      "Resource" : [module.dynamo_file_meta_data.table_arn]
    }
  ]
}




// IAM policy for s3_file_get_item
module "iam_policy_s3_user_file_get_object" {
  source      = "./modules/iam/policy"
  policy_name = "s3_user_file_get_object"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["s3:GetObject"],
      "Resource" : [module.s3_file_upload.s3_bucket_arn]
    }
  ]
}

// IAM policy for s3_user_file_put_object
module "iam_policy_s3_user_file_put_object" {
  source      = "./modules/iam/policy"
  policy_name = "s3_user_file_put_object"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["s3:PutObject"],
      "Resource" : [module.s3_file_upload.s3_bucket_arn]
    }
  ]
}


// IAM policy for secretManagerGetItem
module "iam_policy_secret_manager_get_item_cognito_region" {
  source      = "./modules/iam/policy"
  policy_name = "secret_manager_get_item_cognito_region"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["secretsmanager:GetSecretValue"],
      "Resource" : [module.secret_manager_cognito_region.secret_arn]
    }
  ]
}


// IAM policy for secretManagerGetItem
module "iam_policy_secret_manager_get_item_app_client_id" {
  source      = "./modules/iam/policy"
  policy_name = "secret_manager_get_item_app_client_id"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["secretsmanager:GetSecretValue"],
      "Resource" : [module.secret_manager_app_client_id.secret_arn]
    }
  ]
}


// IAM policy for secretManagerGetItem
module "iam_policy_secret_manager_get_item_user_pool_id" {
  source      = "./modules/iam/policy"
  policy_name = "secret_manager_get_item_user_pool__id"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["secretsmanager:GetSecretValue"],
      "Resource" : [module.secret_manager_user_pool_id.secret_arn]
    }
  ]
}

// IAM policy for secretManagerGetItem
module "iam_policy_secret_manager_get_item_dynamo_region" {
  source      = "./modules/iam/policy"
  policy_name = "secret_manager_get_item_dynamo_region"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["secretsmanager:GetSecretValue"],
      "Resource" : [module.secret_manager_dynamo_region.secret_arn]
    }
  ]
}

// IAM policy for secretManagerGetItem
module "iam_policy_secret_manager_get_item_dynamo_user_table_name" {
  source      = "./modules/iam/policy"
  policy_name = "secret_manager_get_item_dynamo_user_table_name"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["secretsmanager:GetSecretValue"],
      "Resource" : [module.secret_manager_dynamo_user_table_name.secret_arn]
    }
  ]
}



// IAM policy for secretManagerGetItem
module "iam_policy_secret_manager_get_item_dynamo_file_name" {
  source      = "./modules/iam/policy"
  policy_name = "secret_manager_get_item_dynamo_file_name"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["secretsmanager:GetSecretValue"],
      "Resource" : [module.secret_manager_dynamo_file_name.secret_arn]
    }
  ]
}


// IAM policy for secretManagerGetItem
module "iam_policy_secret_manager_get_item_google_client_id" {
  source      = "./modules/iam/policy"
  policy_name = "secret_manager_get_item_google_client_id"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["secretsmanager:GetSecretValue"],
      "Resource" : [module.secret_manager_google_client_id.secret_arn]
    }
  ]
}






module "iam_policy_secret_manager_get_item_google_client_secret" {
  source      = "./modules/iam/policy"
  policy_name = "secret_manager_get_item_google_client_secret"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["secretsmanager:GetSecretValue"],
      "Resource" : [module.secret_manager_google_client_secret.secret_arn]
    }
  ]
}


module "iam_policy_secret_manager_get_item_google_redirect_url" {
  source      = "./modules/iam/policy"
  policy_name = "secret_manager_get_item_google_redirect_url"
  env_name    = var.env_name
  policy_statements = [
    {
      "Effect" : "Allow",
      "Action" : ["secretsmanager:GetSecretValue"],
      "Resource" : [module.secret_manager_google_redirect_url.secret_arn]
    }
  ]
}
