# add the necessary instructions
# to create a Docker container image
# for your Go API gateway server

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
#RUN update-ca-certificates

COPY user user
EXPOSE 5200
ENTRYPOINT [ "/user" ]
