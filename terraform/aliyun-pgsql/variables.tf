variable "region" {
  description = "AliCloud resource Region"
  type        = string
  default     = "cn-shanghai"
}

variable "zone_id" {
  default     = "cn-shanghai-b"
  description = "The AZ for the db instance"
  type        = string
}

variable "tags" {
  default = {
    project = "example project"
    environment : "dev"
    costcenter : 7002
  }
}

variable "security_ips" {
  default = [
    "0.0.0.0/0",
    "100.104.128.192/26",
    "100.104.149.64/26",
    "100.104.175.0/24",
    "100.104.183.0/24",
    "100.104.188.0/24",
    "100.104.201.0/26",
    "100.104.205.0/24",
    "100.104.216.192/26",
    "100.104.226.128/26",
    "100.104.227.192/26",
    "100.104.236.128/26",
    "100.104.244.64/26",
    "100.104.35.192/26",
    "100.104.5.0/24",
    "100.104.52.0/24",
    "100.104.61.128/26",
    "100.104.72.0/24",
    "100.104.85.0/26"
  ]
  description = "List of IP addresses allowed to access all databases of an instance."
  type        = list(string)
}

variable "db_time_zone" {
  default     = "Asia/Shanghai"
  description = "The time zone of the instance. "
  type        = string
}

variable "environment" {
  default = "dev"
}

variable "engine" {
  default     = "PostgreSQL"
  description = "DataBase type"
  type        = string
}

variable "engine_version" {
  default     = "13.0"
  description = "Database version"
  type        = string
}

variable "instance_storage" {
  default     = 30
  description = "DataBase storage size"
  type        = number
}

variable "instance_type" {
  default     = "pg.n2.medium.1"
  description = "Database hardware configuration"
  type        = string
}

// cloud sql one name is an db instance and database name is the same
variable "instance_name" {
  default     = "example"
  description = "instance name"
  type        = string
}

variable "db_name" {
  default     = "db_example"
  description = "database name"
  type        = string
}

variable "db_port" {
  default = 1433
  type    = number
}

// database and account name whick same as name
variable "account_name" {
  default     = "postgres"
  description = "account name"
  type        = string
}

variable "instance_charge_type" {
  default     = "Postpaid" // Postpaid or Prepaid
  description = "Billing method."
  type        = string
}

variable "charge_period" {
  description = "when instance_charge_type is Prepaid can effective"
  default     = 1
  type        = number
}