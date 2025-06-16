group "default" {
  targets = ["frontend", "backend"]
}

target "frontend" {
  context = "./frontend"
  dockerfile = "../infra/docker/frontend/Dockerfile"
  tags = ["llm-frontend:latest"]
}

target "backend" {
  context = "./backend"
  dockerfile = "../infra/docker/backend/Dockerfile"
  tags = ["llm-backend:latest"]
}