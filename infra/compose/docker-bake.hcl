group "default" {
  targets = ["frontend", "backend"]
}

target "frontend" {
  context = "./frontend"
  dockerfile = "../infra/docker/frontend/Dockerfile"
  tags = ["llm-frontend:latest"]
  output = ["type=docker"]
}

target "backend" {
  context = "./backend"
  dockerfile = "../infra/docker/backend/Dockerfile.dev"
  tags = ["llm-backend:latest"]
  output = ["type=docker"]
}