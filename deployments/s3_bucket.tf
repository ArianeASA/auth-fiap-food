resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "auth-lambda-bucket-go"
}

resource "aws_s3_bucket_ownership_controls" "lambda_bucket_controls" {
  bucket = aws_s3_bucket.lambda_bucket.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_acl" "lambda_bucket_acl" {
  depends_on = [aws_s3_bucket_ownership_controls.lambda_bucket_controls]

  bucket = aws_s3_bucket.lambda_bucket.id
  acl    = "private"
}

locals {
  source_file = "../${path.module}/main"
  output_path_zip = "../${path.module}/main.zip"

}

data "archive_file" "lambda_zip" {
  type = "zip"

  source_file  = local.source_file
  output_path = local.output_path_zip
}


resource "aws_s3_object" "lambda_main" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "main.zip"
  source = local.output_path_zip
  acl   = "private"
  etag = filebase64sha256(local.output_path_zip)
}
