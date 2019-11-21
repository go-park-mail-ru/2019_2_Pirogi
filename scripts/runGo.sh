echo "---Checking or updating vendors"
go mod vendor

echo "---Stopping containers"
docker stop cinsear
docker rm cinsear

echo "---Starting..."
docker-compose up --build --detach server
