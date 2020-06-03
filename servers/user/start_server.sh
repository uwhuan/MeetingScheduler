docker system prune --all --force
docker system prune --volumes --force
docker network disconnect -f mainnetwork ljchen17gateway
docker network disconnect -f mainnetwork ljchen17redis
docker network disconnect -f mainnetwork ljchen17mysqldemo
docker network disconnect -f mainnetwork ljchen17summary
docker network disconnect -f mainnetwork ljchen17message
docker network disconnect -f mainnetwork customMongoContainer
docker network rm mainnetwork
docker network create mainnetwork

docker rm -f ljchen17redis
docker run -p 6379:6379 --name ljchen17redis -d redis
docker rm -f ljchen17mysqldemo
docker pull ljchen17/mysqldemo
docker run -d -e MYSQL_ROOT_PASSWORD=Hdkme7294 --name ljchen17mysqldemo -p 3306:3306 ljchen17/mysqldemo

docker rm -f ljchen17summary
docker pull ljchen17/summaryservice
docker run -d --name ljchen17summary -p 5100:5100 -e ADDR=:5100 ljchen17/summaryservice

docker rm -f customMongoContainer
docker run -d -p 27017:27017 --name customMongoContainer --network=mainnetwork mongo

docker rm -f ljchen17message
docker pull ljchen17/messageservice
docker run -d --name ljchen17message -p 5200:5200 --network=mainnetwork ljchen17/messageservice

docker rm -f ljchen17gateway
docker pull ljchen17/gateway
docker run -d --name ljchen17gateway -p 443:443 -v /etc/letsencrypt:/etc/letsencrypt:ro -e TLSCERT=/etc/letsencrypt/live/api.ljchen17.me/fullchain.pem -e TLSKEY=/etc/letsencrypt/live/api.ljchen17.me/privkey.pem -e SESSIONKEY=hfewi1 -e REDISADDR=ljchen17redis:6379 -e DSN="root:Hdkme7294@tcp(ljchen17mysqldemo:3306)/INFO441" -e MESSAGESADDR=ljchen17message:5200 -e SUMMARYADDR=ljchen17summary:5100 ljchen17/gateway

docker network connect mainnetwork ljchen17gateway
docker network connect mainnetwork ljchen17redis
docker network connect mainnetwork ljchen17mysqldemo
docker network connect mainnetwork ljchen17summary

