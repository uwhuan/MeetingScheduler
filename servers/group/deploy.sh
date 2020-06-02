docker rm -f group
docker pull ice2meu/group
docker run -d --name summary --network test -p 8100:8100 -e ADDR=:8100 ice2meu/group