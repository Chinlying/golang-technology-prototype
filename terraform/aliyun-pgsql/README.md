+ Knowledge Points 
  + Create PostgreSQL instance on Alicloud
  + Use `random_password` to generate random passwords for database
+ Notes
  + If you don't have a switch already, you need to set one first or set one by using terraform script, please refer [alicloud_vpc, alicloud_vswitch](https://registry.terraform.io/providers/aliyun/alicloud/latest/docs/resources/db_account) to imitate
  + Set up accurate switch `alicloud_vswitches` according to your Alicloud recourses
  + Set up accurate charge type `instance_charge_type`, if you set up it as 'Prepaid', you must pay attention to the variable `charge_period`
+ How to use
```shell
terraform init
terraform plan
terraform apply
```
+ Refernces
  + [Create RDS instance on Alicloud](https://registry.terraform.io/providers/aliyun/alicloud/latest/docs/resources/db_account)
  + [RDS PostgreSQL主实例规格列表](https://www.alibabacloud.com/help/zh/apsaradb-for-rds/latest/primary-apsaradb-rds-for-postgresql-instance-types)