docker rm -f group
# docker pull ice2meu/group
docker run -d --name group --network test -p 8100:8100 -e ADDR=:8100 -e DSN="root:password@tcp(mysql:3306)/INFO441" ice2meu/group