# Group and Meeting Server

This server handles request of /groups

The server manages group meeting management, group member management

## APIs

**All APIs below require a `X-user` header with the authenticated userID from gateway**

- `/v1/groups`

  - POST: Insert a group into db, returns the JSON format group struct

    - Accept input: (Other information will fill by the server)

      ```json
      {
          "name":"meeting name",
          "description": "details"
      }
      ```

    - Status code:

  - GET: return a JSON array of all groups in server (only used by the site admin for testing)

- `/v1/groups/{group_id}`

  - GET

    - Response body

      ```JSON
      {
          "GroupInfo": {
              "groupID": 1,
              "description": "some words",
              "name": "meeting name",
              "creatorID": 1,
              "createDate": "Wed Jun  3 19:39:05 UTC 2020"
          },
          "Meetings": {
              {
                  "meetingID": 1,
                  "groupID": 0,
                  "name": "test",
                  "description": "",
                  "creatorID": 12,
                  "startTime": "",
                  "endTime": "",
                  "createDate": "Wed Jun  3 19:45:19 UTC 2020",
                  "confirmed": 0
          	},
              ...,
      
          },
          "Members": {
              {
              	"username":"user",
              	"firstname":"ad",
              	"lastname":"ad"
              	...
          	},
      		...
          }
      }
      ```

    - Status code

  - PATCH

  - DELETE

- `/v1/groups/{group_id}/meetings`

  - GET
  - POST

- `/v1/groups/{group_id}/meetings/{meeting_id}`

  - GET
  - PATCH
  - DELETE

- `/v1/groups/{group_id}/meetings/{meeting_id}/schedule`

  - POST
  - DELETE