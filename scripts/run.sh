echo "---Checking or updating vendors"
go mod vendor

echo "---Stopping containers"
docker stop cinsear
docker rm cinsear
docker stop cinsear-db
docker rm cinsear-db

echo "---Checking or creating volumes..."
docker volume create media
docker volume create log
docker volume create db

echo "---Starting..."
docker-compose up --build --detach server mongo


if [ "$1" = "-first-time" ]; then
  echo "---Filling db..."
  cd cmd/database/ && go run initDB.go
fi