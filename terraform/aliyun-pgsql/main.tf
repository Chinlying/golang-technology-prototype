data "alicloud_vswitches" "vsws" {
  vswitch_name = "dev-dscore-vsw-b"
}

data "alicloud_resource_manager_resource_groups" "groups" {
  name_regex = "dscore"
}

resource "alicloud_db_instance" "postgres" {
  engine           = var.engine
  engine_version   = var.engine_version
  instance_name    = "psql-${var.environment}-${var.instance_name}"
  instance_storage = var.instance_storage
  instance_type    = var.instance_type
  vswitch_id       = data.alicloud_vswitches.vsws.vswitches[0].id
  /*
  It is valid when instance_charge_type is PrePaid.
  Valid values: [1~9], 12, 24, 36. -> NOTE:
  The attribute period is only used to create Subscription instance or modify the PayAsYouGo instance to Subscription.
  Once effect, it will not be modified that means running terraform apply will not effect the resource.
  */
  instance_charge_type   = var.instance_charge_type
  period                 = var.charge_period
  port                   = var.db_port
  security_ips           = var.security_ips
  db_time_zone           = var.db_time_zone
  resource_group_id      = data.alicloud_resource_manager_resource_groups.groups.groups[0].id
  whitelist_network_type = "VPC"
  zone_id                = var.zone_id
  tags = var.tags
}

resource "alicloud_db_database" "database" {
  instance_id   = alicloud_db_instance.postgres.id
  character_set = "UTF8"
  name          = var.db_name
  lifecycle {
    ignore_changes = [character_set]
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "_"
}

resource "alicloud_db_account" "account" {
  db_instance_id   = alicloud_db_instance.postgres.id
  account_name     = var.account_name
  account_password = random_password.password.result
  depends_on = [
    alicloud_db_instance.postgres
  ]
}

resource "alicloud_db_connection" "connection" {
  instance_id       = alicloud_db_instance.postgres.id
  connection_prefix = alicloud_db_instance.postgres.instance_name
  port              = alicloud_db_instance.postgres.port
}
