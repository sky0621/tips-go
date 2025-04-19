variable "project_id" {
  description = "GCP プロジェクト ID"
  type        = string
}
variable "region" {
  description = "リージョン（例：asia-northeast1）"
  type        = string
  default     = "asia-northeast1"
}
variable "db_instance_name" {
  description = "Cloud SQL インスタンス名"
  type        = string
  default     = "tf-cloudsql-instance"
}
variable "db_tier" {
  description = "DB インスタンスタイプ（例：db-f1-micro）"
  type        = string
  default     = "db-f1-micro"
}
variable "db_database_name" {
  description = "作成するデータベース名"
  type        = string
  default     = "cloudsqlconnmysql01"
}
variable "db_user" {
  description = "DB に作成するユーザー名"
  type        = string
  default     = "yuckyjuice"
}
variable "db_password" {
  description = "DB ユーザーのパスワード"
  type        = string
  sensitive   = true
  default     = "yuckyjuice"
}
variable "service_name" {
  description = "Cloud Run サービス名"
  type        = string
  default     = "my-app-service"
}
variable "image" {
  description = "Cloud Run にデプロイするコンテナイメージ"
  type        = string
}
