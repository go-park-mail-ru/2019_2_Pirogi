docker stop cinsear
docker rm cinsear
docker rmi cinsear
docker volume create log
docker volume create media
docker build -t cinsear ../.
docker run -ti -p 8000:8000 \
        --env mode=development \
        --mount type=volume,source=media,target=/media \
        --mount type=volume,source=log,target=/log \
        --name cinsear -d cinsear