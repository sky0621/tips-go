terraform {
  required_version = ">= 1.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

# 1. 必要な API を有効化
resource "google_project_service" "sqladmin" {
  service = "sqladmin.googleapis.com"
}
resource "google_project_service" "run" {
  service = "run.googleapis.com"
}
resource "google_project_service" "iam" {
  service = "iam.googleapis.com"
}
resource "google_project_service" "compute" {
  service = "compute.googleapis.com"
}

# 2. Cloud SQL：MySQL インスタンス
resource "google_sql_database_instance" "instance" {
  name             = var.db_instance_name
  database_version = "MYSQL_8_0"
  region           = var.region

  settings {
    tier = var.db_tier

    # 公開 IP は使わず、Cloud Run のデフォルト VPC 経由で接続
    ip_configuration {
      ipv4_enabled    = true
      require_ssl     = true
    }

    backup_configuration { enabled = true }
    activation_policy    = "ALWAYS"
  }
}

# 3. データベースとユーザー作成
resource "google_sql_database" "appdb" {
  name     = var.db_database_name
  instance = google_sql_database_instance.instance.name
}
resource "google_sql_user" "appuser" {
  name     = var.db_user
  instance = google_sql_database_instance.instance.name
  password = var.db_password
}

# 4. Cloud Run 用サービスアカウント
resource "google_service_account" "run_sa" {
  account_id   = "${var.service_name}-sa"
  display_name = "Service Account for Cloud Run"
}

# 5. サービスアカウントに Cloud SQL Client 権限を付与
resource "google_project_iam_member" "run_sa_sql_client" {
  project = var.project_id
  role    = "roles/cloudsql.client"
  member  = "serviceAccount:${google_service_account.run_sa.email}"
}

# 6. Cloud Run サービス
resource "google_cloud_run_service" "app" {
  name     = var.service_name
  location = var.region

  template {
    metadata {
      annotations = {
        # タイムアウトを 5 分（300 秒）に設定
        "run.googleapis.com/timeoutSeconds" = "300"
      }
    }
    spec {
      service_account_name = google_service_account.run_sa.email

      containers {
        image = var.image

        # 追加: Cloud Run 側がヘルスチェック・起動待機するポート
        ports {
          container_port = 8080
        }

        # Cloud SQL への接続情報を環境変数で渡す
        env {
          name  = "INSTANCE_CONNECTION_NAME"
          value = google_sql_database_instance.instance.connection_name
        }
        env {
          name  = "DB_USER"
          value = var.db_user
        }
        env {
          name  = "DB_PASS"
          value = var.db_password
        }
        env {
          name  = "DB_NAME"
          value = var.db_database_name
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

# 7. (オプション) 誰でもアクセス可能にする
resource "google_cloud_run_service_iam_member" "invoker" {
  location = var.region
  project  = var.project_id
  service  = google_cloud_run_service.app.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

# 8. 出力
output "cloud_run_url" {
  description = "Cloud Run の URL"
  value       = google_cloud_run_service.app.status[0].url
}

output "db_connection_name" {
  description = "Cloud SQL インスタンスの接続名"
  value       = google_sql_database_instance.instance.connection_name
}
