echo "---Stopping containers"
docker stop cinsear-db
docker rm cinsear-db

echo "---Checking or creating volumes..."
docker volume create media
docker volume create log
docker volume create db

echo "---Starting..."
docker-compose up --build --detach mongo


if [ "$1" = "-first-time" ]; then
  echo "---Filling db..."

  go run cmd/database/initDB.go
fi
