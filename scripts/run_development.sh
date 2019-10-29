docker stop cinsear
docker rm cinsear
docker stop cinsear-db
docker rm cinsear-db
docker-compose up --build --detach server-dev mongo
