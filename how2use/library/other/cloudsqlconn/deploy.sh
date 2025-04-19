#!/bin/bash
set -e

PROJECT_ID="$1"
REGION="asia-northeast1"           # 東京リージョン
REPOSITORY="mytestrepository"      # Artifact Registry のリポジトリ名
IMAGE_NAME="my-go-webapi"
TAG="v1.0.0"

IMAGE_URI="${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPOSITORY}/${IMAGE_NAME}:${TAG}"

# プロジェクト設定
gcloud config set project ${PROJECT_ID}

# Artifact Registry 用の Docker 認証設定
gcloud auth configure-docker "${REGION}-docker.pkg.dev" --quiet

# イメージビルド
docker build -t ${IMAGE_URI} .

# イメージプッシュ
docker push ${IMAGE_URI}

echo "Artifact Registry へプッシュ完了: ${IMAGE_URI}"
