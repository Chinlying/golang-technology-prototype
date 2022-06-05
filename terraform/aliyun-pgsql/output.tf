output "vswitches" {
  value = data.alicloud_vswitches.vsws.vswitches
}

output "resource_group" {
  value = data.alicloud_resource_manager_resource_groups.groups.groups[0].id
}

output "vswitch_id" {
  value = data.alicloud_vswitches.vsws.vswitches[0].id
}

output "db_instance_name" {
  value = "psql-dev-example"
}

output "db_connection_url" {
  value = alicloud_db_instance.postgres.connection_string
}

output "db_name" {
  value = alicloud_db_database.database.name
}

output "db_port" {
  value = alicloud_db_instance.postgres.port
}

output "db_username" {
  value = alicloud_db_account.account.account_name
}

output "db_password" {
  value = nonsensitive(alicloud_db_account.account.account_password)
}
