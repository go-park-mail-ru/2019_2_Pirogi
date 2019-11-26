echo "---Stopping containers"
docker stop prometheus
docker rm prometheus
docker stop nodeexporter
docker rm nodeexporter
docker stop grafana
docker rm grafana

echo "---Starting..."
docker-compose up --build --detach prometheus nodeexporter grafana
