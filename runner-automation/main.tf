provider "aws" {
  region = "eu-west-1"
}

resource "aws_s3_bucket" "example" {
  bucket = "tester-bucket"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
