npm run build
docker build -t nalin97/meetingscheduler .
docker push -t nalin97/meetingscheduler .
ssh root@138.68.43.86 < deploy.sh