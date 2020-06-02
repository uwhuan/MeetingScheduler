GOOS=linux go build
docker build -t ice2meu/group .
docker push ice2meu/group