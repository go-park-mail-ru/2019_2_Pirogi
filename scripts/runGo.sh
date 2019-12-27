echo "---Checking or updating vendors"
go mod vendor

echo "---Stopping containers"
docker stop cinsear
docker rm cinsear

echo "---Parsing models..."
easyjson --all -pkg app/domain/model

echo "---Starting..."
docker-compose up --build --detach server
