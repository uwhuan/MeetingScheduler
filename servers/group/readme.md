# Group and Meeting Server

This server handles request of /groups

The server manages group meeting management, group member management

## APIs

**All APIs below require a `X-user` header with the authenticated userID from gateway**

- `/v1/groups`

  - POST: Insert a group into db, returns the JSON format group struct

    - Accept input: (Other information will be filled by the server)

      ```json
      {"name":"Group name","description": "details"}
      OR {"name":"Group name"}
      ```
      
    - Success: Return `successfully create group, id: `

    - Status code:

      - 201, 415, 400, 500

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
              ...
      
          },
          "Members": {
              {
              	"username":"user",
              	"firstname":"ad",
              	"lastname":"ad"
          	},
      		...
          }
      }
      ```
      
  - Status code
    
  - PUT: Update the group's name and description

    - Accept Input:

      ```json
      {"name":"New name","description": "New details"}
      ```

    - Success Return:

      ```JSON
      {
          "groupID": 3,
          "description": "test message",
          "name": "test33",
          "creatorID": 1,
          "createDate": "Thu Jun  4 21:30:19 UTC 2020"
      }
      ```

  - PATCH: Invite another user and generate a invitation link 

    - Accept Input:

      ```JSON
      {"name":"guest", "email":"test@mail.com"}
      ```

    - Success Return Example:

      path: host/version/guest/

      g: group_id

      id: random generated guestID

      ```
      localhost:8100/v1/guest/g=1&name=guest1&email=ade@mail.com&id=6645600
      ```

  - DELETE: Delete the whole group and any meetings under the group. Only group creator can use this method

    - Success return: `Delete success`

- `/v1/groups/{group_id}/meetings`

  - GET: Get all meetings under the group

  - POST: Create new meetings under the group

    - Accept Input:

      ```JSON
      {"name":"Meeting name","description": "details"}
      OR {"name":"Meeting name"}
      ```

    - Success Return Example:

      ```
      {"meetingID":5,"groupID":1,"name":"Meeting Name","description":"","creatorID":12,"startTime":"","endTime":"","createDate":"Thu Jun  4 21:44:38 UTC 2020","confirmed":0}
      ```

- `/v1/groups/{group_id}/meetings/{meeting_id}`

  - GET: Get all information of a meeting

    - Success Return Example:

      ```JSON
      {
          "MeetingInfo": {
              "meetingID": 5,
              "groupID": 1,
              "name": "Meeting Name",
              "description": "",
              "creatorID": 12,
              "startTime": "",
              "endTime": "",
              "createDate": "Thu Jun  4 21:44:38 UTC 2020",
              "confirmed": 0
          },
          "Schedules": null,
          "Participants": null
      }
      ```

  - PUT: Update information of a current meeting

    - Accept Input:

      ```
      {"name":"Meeting name","description": "details"}
      OR {"name":"Meeting name"}
      ```

    - Success Return Example:

      ```JSON
      {
          "meetingID": 5,
          "groupID": 1,
          "name": "New Meeting Name",
          "description": "",
          "creatorID": 12,
          "startTime": "",
          "endTime": "",
          "createDate": "Thu Jun  4 21:44:38 UTC 2020",
          "confirmed": 0
      }
      ```

  - PATCH: Invite another user into this meeting

    -  Accept Input:

      ```JSON
      {"name":"guest", "email":"test@mail.com"}
      ```

    - Success Return Example:

      ```
      localhost:8100/v1/guest/m=1&name=guest1&email=ade@mail.com&id=4377141
      ```

      m: meetingid

      id: random generated guestID

  - DELETE: will return `Successfully deleted` if success

- `/v1/groups/{group_id}/meetings/{meeting_id}/schedule`

  - POST
  
    - Accept Input:
  
      ```JSON
      {
          "startTime": "Mon Jun 5 19:00:00 -0700 MST 2020",
          "endTime": "Mon Jun 5 19:30:00 -0700 MST 2020"
      }
      ```
  
    - Success Return Example:
  
      ```JSON
      {
          "scheduleID": 3,
          "meetingID": 1,
          "startTime": "Mon Jun 5 19:00:00 -0700 MST 2020",
          "endTime": "Mon Jun 5 19:30:00 -0700 MST 2020",
          "votes": 0
      }
      ```
  
  - GET
  
- `/v1/groups/{group_id}/meetings/{meeting_id}/schedule/{schedule_id}`

  - PATCH: This is used for voting the current schedule
  - DELETE: Delete the current schedule

## Notice

1. The server will read the "X-user" header as the current user. When the client create a meeting or a group, the creator will set to the current user.
2. This server doesn't modify any information about users or guests

## File Structure

### handler

- `context.go`: Store context information, including signing keys and database connections
- `groupsHandler.go`: The main handler
- `handlerUtil.go`: Common modules that are used in handlers, such as reading request body

### dao

- `groupServer.go`: conduct operations on group table
- `meetingsServer.go`: conduct operations on meeting table
- `invitationServer.go`: TBD
- `scheduleServer.go`: conduct operations on schedule table
- `store.go`: (deprecated) store interface

### model 

​	refer to 

[https://github.com/uwhuan/MeetingScheduler/blob/master/db/readme.md]: db/readme.md

- `groups.go`: 
  - Group struct: contains general information from userGroups table in database
  - GroupReturnBody struct: is used to return the response
- `invitation.go`: contains two functions to generate random url
- `meetings.go`: contains general information of a meeting according to the  meetings table
- `schedule.go`: contains an option of schedule. It is used for users to choose from.
- `user.go`: It's only used for display user information.

## TODO

1. Time zone issues