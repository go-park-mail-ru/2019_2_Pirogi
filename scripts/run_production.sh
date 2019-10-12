docker stop cinsear
docker rm cinsear
docker rmi cinsear
docker volume create log
docker volume create media
docker volume create ssl
docker build -t cinsear ../.
docker run -ti -p 8000:8000 \
        --env mode=production \
        --mount type=volume,source=media,target=/media \
        --mount type=volume,source=log,target=/log \
        --mount type=volume,source=ssl,target=/ssl \
        --name cinsear -d cinsear

