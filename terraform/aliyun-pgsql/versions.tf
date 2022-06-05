terraform {
  required_version = ">= 0.12.2"
  required_providers {
    alicloud = {
      source  = "aliyun/alicloud"
      version = ">=1.56.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.2.0"
    }
  }
}
