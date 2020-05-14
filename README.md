# Meeting Scheduler
Final project for INFO 441


Group Members: Nalin Gupta, Divit Jawa, Lizzy Chen, Huan Wang

# Project description

Meeting Schedule helps teams solve the difficulty of scheduling a meeting time. It’s a simple but powerful app for anyone who works or studies in a team. Each of the team members only needs to select their available time within the range the host set, then the app will automatically calculate the best available time for everyone. 
​There are many similar apps online, such as when2meet, Google Calendar. We want to build a new one that is more user-friendly and powerful than when2meet, but more simple and light-weighted than Google Calendar. We want to build a more instinctive and simple UI so that anyone with an invite link can use the app easily. At the same time, we want to offer more useful functionalities for the registered user, such as managing their events, creating new events, sharing, etc. 

### Audience 
Students/people who are working in a team, and want to schedule a time for meeting.

### Problem

* Most people work in teams or groups today and they always need to schedule a meeting. However, it’s very difficult to schedule an appropriate meeting time especially when the team is large. Asking each one’s opinion is low-efficient and puts a lot of burden on the hosts. 
* Usually, people belong to multiple groups, which makes scheduling more complex. We need to check our available time manually and it’s easy to get events overlapped.
  Solution
* Any user can create groups and meetings.
* Our application will allow the hosts of groups and meetings to set time ranges and possible places, and send links to invite people to join the meeting groups.
* Any user with a verified link can join in meeting groups. They can also select the best time slot and place according to their preference.
* Registered users can schedule meetings for the same group and with the same settings.
* Registered users can manage multiple meetings and groups. The app will automatically block the time slot that conflicts with other meetings.

### Main Features:

* Set available meeting time by simple selecting
* Create meeting with possible time range, location, participants and more detailed information (such as zoom link)
* Create groups and manage group meeting
* Automatically generate the earliest meeting time based on people’s availability 

### Stretch-goal features:

* Allow user to put preference in the time slot
* Automatically generate a suggested meeting time based on people’s preference

### Comparative analysis

* When2Meet: When2Meet has an un-intuitive UI and is difficult to operate on a mobile phone. Another disadvantage of When2Meet is the need to specify your schedule each time you use it for meeting with a group. 
* Google Calendar: Although it allows people to check free time spots of other people, it does not provide functionality to compare multiple people’s schedules. What’s more, it is not giving users the option to choose where they want to meet.
* Doodle: Doodle is an app that allows an event host to send a survey for the best meeting times. However, Doodle requires attendees to fill out their best times for each individual event.

# Architecture:

<img src="Architecture Diagram.png" alt="architecture_diagram" style="zoom:75%;" />

Details: https://app.lucidchart.com/invitations/accept/3428ba22-3cdc-4bc7-9db3-323686848dbc

# User Stories:

| Priority | User                       | Description                                                  | Implementation                                               |
| -------- | -------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| P0       | As a meeting group creator | I want to create a meeting group                             | Upon receiving a POST request to /v1/groups, the gateway creates a new group and stores it in the database. |
| P0       | As a meeting group creator | I want to create one or more meeting of the group            | Upon receiving POST requests to /v1/groups/{groupsid}/meetings/{meetingid}, the gateway creates one or more new meetings and stores them in the database. |
| P0       | As a meeting group creator | I want to have a meeting group url to share with people      | Upon receiving a POST request to /v1/groups, the groups microservice generates a unique groups URL, associates it with a new empty group in the DB, and opens a websocket with the client to enable the group. |
| P0       | As a meeting group creator | I want to invite people to a meeting group                   | Upon receiving PATCH request to /v1/groups/{groupsid}, the group microservices will check the DB and see if the new users are already added; if not, then send the emails of url to new users |
| P0       | As a user                  | I want to see the group details using the url                | The client will parse the URL to extract an ID, then send a GET request to /v1/groups/{id}, which will cause the groups microservice to retrieve the current state of the group and open a websocket with the client to enable the group. |
| P0       | As a user                  | I want to see a list of meetings created by the group        | Upon receiving a GET request to /v1/groups/{groupsid}/meetings, the groups microservice shows a list of meetings under the current user. |
| P0       | As a user                  | I want to see the details of a meeting by the meeting ID     | Upon receiving a GET request to /v1/groups/{groupsid}/meetings/{meetingid}, the groups microservice shows the meeting under the specific id. |
| P0       | As a user                  | I want to create an account                                  | Upon receiving a POST request to /v1/users, the gateway creates a new user account and stores it in the database. |
| P1       | As a registered user       | I want to log into my account                                | Upon receiving a POST request to /v1/sessions, the gateway verifies the user credentials and shows the user profile. |
| P1       | As a registered user       | I want to log out my account                                 | Upon receiving a DELETE request to /v1/sessions, the gateway verifies the user credentials and signs out the user. |
| P1       | As a user                  | I want to add my time availability to a meeting group        | Upon receiving PATCH request to /v1/groups/{groupsid}, the calendar microservices will check the DB and update the availability. |
| P1       | As a user                  | I want to know the earliest available time for group members before creating a meeting | Upon receiving a GET request to /v1/groups/{groupsid}, the calendar microservices will check users availability in the DB and return the earliest time for the meeting |
| P1       | As a user                  | I want to download an ics file of a meeting                  | Upon receiving GET request to /v1/groups/{groupsid}/meetings/{meetingsid}/ics-generate, the groups microservices will return a ics file generated from the meeting data |
| P1       | As a registered user       | I want to see a user&#39;s profile by given ID               | Upon receiving a GET request to /v1/users/{usersid}, the users microservice shows the user under the specific id. |
| P1       | As a registered user       | I want to update my user profile                             | Upon receiving a PATCH request to /v1/users/{usersid}, the users microservice will check in the DB and update the profile. |
| P1       | As a registered user       | I want to view all the meetings I have attended/will be attending/have created | Upon receiving a GET request to /v1/users/{usersid}/meetings?viewtype=all, the groups microservice shows a list of meetings under the current user. Viewtype could be &quot;all&quot;, &quot;past&quot; or &quot;future&quot;. Another parameter can specify to return only ones the user has created |
| P1       | As a registered user       | I want to view all my meeting groups                         | Upon receiving a GET request to /v1/users/{usersid}/groups, the groups microservice shows a list of groups under the current user. |
| P1       | As a registered user       | I want to delete a meeting under my account                  | Upon receiving a DELETE request to /v1/users/{usersid}/meetings/{meetingid}, the groups microservice deletes meetings under the current users. |
| P1       | As a meeting group creator | I want to delete a meeting I have created                    | Upon receiving a DELETE request to /v1/groups/{groupsid}/meetings/{meetingid}, the groups microservice deletes meetings from the DB and closes all websockets connected to the meetings. |
| P2       | As a registered user       | I want to set up my preference for meeting time              | Upon receiving a PATCH request to /v1/users/{usersid}, the users microservice will check in the DB and update the preference. The preference is part of the user profile, but belongs to more advanced features we will try to implement. |
| P2       | As a meeting group creator | I want to know the suggested available time based off group members&#39; preference before creating a meeting | Upon receiving a GET request to /v1/groups/{groupsid}, the calendar microservices will check users preference in the DB and return the suggested time for the meeting. |

# Endpoints

## Users

### /v1/users

* Post: Creates a new user account 
  * 201: Successfully created user.
  * 400: The request body is not a valid user.
  * 415: Content-Type not JSON
  * 500: Internal server error.

### /v1/users/{user_id}

* GET:  Get the user with the given ID or current user with me
  * 200: Successfully Return of user.
  * 400: The id parameter is not a valid user ID.
  * 401: The user is not logged in.
  * 403: Not Authorized to get the user
  * 500: Internal server error.
* PATCH:  Update the user with the current ID 
  * 200: Successfully Updated User.
  * 400: The id parameter is not a valid user ID.
  * 401: The user is not logged in.
  * 403: Not Authorized to get the user
  * 415: Content Type Provided is not JSON
  * 500: Internal server error.

## Sessions        

### /v1/sessions

* POST: Creates a new user session
  * 201: Successfully created session.
  * 400: The request body is not valid.
  * 401: The email/password combo given was incorrect.
  * 415: Content-Type not JSON
  * 500: Internal server error.

### /v1/session/{sessionid}

* DELETE : Ends the current user session. Should be the current session id or mine. 
  * 200: Successfully ended session.
  * 403: The user is attempting to end another user's session.
  * 500: Internal server error

## Meetings user specific 

### /v1/user/{userid}/meetings 

* GET : Gets a list of all the meetings created by the user
  * 200: Successfully retrieved meeting data
  * 401: No valid user with the given ID
  * 500: Internal server error
* POST : Creates a new meeting 
  * 201: Successfully created meeting
  * 500: Internal server error

### /v1/user/{userid}/meetings/{meetingid} 

* GET:  Get the current status of a specific meeting
  * 200 Returns the current state of the meeting
  * 401 Could not verify player, or they are not in the game
  * 404 The meeting wasn’t found
  * 415: Unsupported media type
  * 500: Internal server error
 * PUT/PATCH: Edit the current meeting
   * 200 ok: Returns the updated state of the meeting
   * 400 no updated: No new information was provided
   * 401 unauthorized: Could not verify user
 * DELETE : Delete the meeting or ends the meeting
   * 200: Successfully ends the meeting session
   * 401: Could not verify player the user
   * 404: The meeting  wasn’t found
   * 500: Internal Server Error 
     Group

### /v1/groups

* Post: Creates a new group 
  * 201: Successfully created user.
  * 400: The request body is not a valid user.
  * 415: Content-Type not JSON
  * 500: Internal server error.

### /v1/groups/{group_id}

* GET:  Get the user with the given ID or current group with me
  * 200: Successfully Return of group.
  * 400: The id parameter is not a valid group ID.
  * 401: The user is not logged in.
  * 403: Not Authorized to get the group
  * 500: Internal server error.
* PATCH:  Update the group with the current ID 
  * 200: Successfully Updated group.
  * 400: The id parameter is not a valid group ID.
  * 401: The user is not logged in.
  * 403: Not Authorized to change the group
  * 415: Content Type Provided is not JSON

### /v1/groups/{group_id}/meetings 

* GET : Gets a list of all the meetings created by the group
  * 200: Successfully retrieved meeting data
  * 401: No valid group with the given ID
  * 500: Internal server error
* POST : Creates a new meeting 
  * 201: Successfully created meeting
  * 500: Internal server error

### /v1/groups/{groupsid}/meetings/{meetingid} 

* GET:  Get the current status of a specific meeting
  * 200 Returns the current state of the meeting
  * 401 Could not verify player, or they are not in the game
  * 404 The meeting wasn’t found
  * 415: Unsupported media type
  * 500: Internal server error
 * PUT/PATCH: Edit the current meeting
   * 200  Returns the updated state of the meeting
   * 400 no updated: No new information was provided
   * 401 unauthorized: Could not verify user
 * DELETE : Delete the meeting or ends the meeting
   * 200: Successfully ends the meeting session
   * 401: Could not verify player the user
   * 404: The meeting  wasn’t found
   * 500: Internal Server Error
