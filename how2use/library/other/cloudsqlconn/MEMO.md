
```
gcloud artifacts repositories create mytestrepository \
  --repository-format=docker \
  --location=asia-northeast1 \
  --description="Go WebAPI 用リポジトリ"
```

```
./deploy.sh xxxx
```

```
terraform init
```

```
export MY_GCP_PROJECT_ID=xxxx
export IMAGE_URL=asia-northeast1-docker.pkg.dev/${MY_GCP_PROJECT_ID}/mytestrepository/my-go-webapi:v1.0.0

terraform plan -var="project_id=${MY_GCP_PROJECT_ID}" -var="image=${IMAGE_URL}" -out=plan.tfplan

terraform apply plan.tfplan
```
