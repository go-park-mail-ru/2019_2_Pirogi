echo "---Stopping containers"
docker stop cinsear-chat
docker rm cinsear-chat

echo "---Checking or creating volumes..."
docker volume create log
docker volume create db

echo "---Starting..."
docker-compose up --build --detach chat

if [ "$1" = "-first-time" ]; then
  echo "---Filling db..."
  go run cmd/database/initDB.go
fi
