resource "aws_lambda_function" "auth_fiap_food" {
  function_name = "auth-fiap-food"

  s3_bucket = aws_s3_bucket.lambda_bucket.id
  s3_key    = "testV0.0.1/${aws_s3_object.lambda_main.key}"
  s3_object_version = aws_s3_object.lambda_main.version_id

  runtime = "go1.x"
  handler = "main"

  source_code_hash = aws_s3_object.lambda_main.content_base64

  role = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      CLIENT_ID = aws_cognito_user_pool_client.cognito_user_pool_client.id
    }
  }

}

resource "aws_lambda_permission" "apigw_lambda_token" {

  statement_id  = "inter-lambda"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.auth_fiap_food.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "arn:aws:execute-api:${var.aws_region}:${var.account_id}:${aws_apigatewayv2_api.auth_fiapfood_api.id}/*/*/{proxy+}"
}


resource "aws_cloudwatch_log_group" "fiap_food_log_group" {
  name = "/aws/lambda/${aws_lambda_function.auth_fiap_food.function_name}"

  retention_in_days = 30
}

resource "aws_iam_role" "lambda_exec" {
  name = "serverless_lambda_roles"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }
    ]
  })
}

resource "aws_iam_policy" "function_logging_policy" {
  name   = "lambda-logging-policy"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        Action : [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Effect : "Allow",
        Resource : "arn:aws:logs:*:*:*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.function_logging_policy.arn

}

