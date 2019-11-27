echo "---Checking or updating vendors"
go mod vendor

echo "---Stopping containers"
docker stop cinsear
docker rm cinsear
docker stop cinsear-db
docker rm cinsear-db
docker stop cinsear-sessions
docker rm cinsear-sessions
docker stop cinsear-users
docker rm cinsear-users

echo "---Checking or creating volumes..."
docker volume create media
docker volume create log
docker volume create db

echo "---Parsing models..."
easyjson -all -pkg app/domain/model

echo "---Starting..."
docker-compose up --build --detach server sessions users mongo

if [ "$1" = "-first-time" ]; then
  echo "---Filling db..."

  go run cmd/database/initDB.go
fi
