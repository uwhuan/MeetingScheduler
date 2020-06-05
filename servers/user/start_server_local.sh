docker system prune --all --force
docker system prune --volumes --force
docker network disconnect -f mainnetwork ljchen17gateway
docker network disconnect -f mainnetwork ljchen17redis
docker network disconnect -f mainnetwork ljchen17mysqldemo
docker network rm mainnetwork
docker network create mainnetwork

docker rm -f ljchen17redis
docker run -p 6379:6379 --name ljchen17redis -d redis
docker rm -f ljchen17mysqldemo
docker pull ljchen17/mysqldemo
docker run -d -e MYSQL_ROOT_PASSWORD=Hdkme7294 --name ljchen17mysqldemo -p 3306:3306 ljchen17/mysqldemo

docker rm -f ljchen17gateway
docker pull ljchen17/usergateway
docker run -d --name ljchen17gateway -p 5200:5200 -e SESSIONKEY=hfewi1 -e REDISADDR=ljchen17redis:6379 -e DSN="root:Hdkme7294@tcp(ljchen17mysqldemo:3306)/INFO441" ljchen17/usergateway

docker network connect mainnetwork ljchen17gateway
docker network connect mainnetwork ljchen17redis
docker network connect mainnetwork ljchen17mysqldemo