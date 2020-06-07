docker rm -f meetingscheduler

docker pull nalin97/meetingscheduler

docker run -d \
-p 443:443 \
-p 80:80 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
--name meetingscheduler \
nalin97/meetingscheduler

exit