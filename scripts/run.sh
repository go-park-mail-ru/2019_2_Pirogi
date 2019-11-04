docker stop cinsear
docker rm cinsear
docker stop cinsear-db
docker rm cinsear-db

docker volume create media
docker volume create log
docker volume create db

docker-compose up --build --detach server mongo
