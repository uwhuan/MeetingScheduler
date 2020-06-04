CREATE DATABASE INFO441;
USE INFO441;

create table if not exists user (
    UID int not null auto_increment primary key,
    Email varchar(320) not null,
    PassHash varchar(255) not null,
    UserName varchar(255) not null,
    FirstName varchar(64) not null,
    LastName varchar(128) not null,
    UNIQUE KEY unique_email (Email),
    UNIQUE KEY unique_username (UserName)
);

create table if not exists meetings (
    MeetingID int not null auto_increment primary key,
    Name varchar(255) not null,
    CreatorID int not null,
    GroupID int not null,
    Description varchar(320),
    StartTime varchar(100), 
    EndTime varchar(100),  
    Confirmed int,  
    CreateDate varchar(100) not null,
    CONSTRAINT mg UNIQUE (MeetingID,GroupID)
);

create table if not exists meetingparticipant (
    MeetingID int not null,
    UID int not null,
    CONSTRAINT mu UNIQUE (MeetingID,UID)
);

create table if not exists logentry (
    UserID int not null auto_increment primary key,
    SignInTime datetime not null,
    IPAddress varchar(32) not null
);

create table if not exists schedule (
    ScheduleID int not null auto_increment primary key,
    MeetingID int not null,
    StartTime varchar(300) not null,
    EndTime varchar(300) not null,
    Votes int
);

create table if not exists userGroups (
    GroupID int not null auto_increment primary key,
    Description varchar(500),
    Name varchar(100),
    CreatorID int not null,
    CreateDate varchar(300)
);

create table if not exists membership (
    uid int not null,
    GroupID int not null,
    CONSTRAINT gu UNIQUE (GroupID, uid)
);

create table if not exists guests (
    GuestID int not null primary key,
    Email varchar(100) not null,
    DisplayName varchar(100), 
    GroupID int not null,
    MeetingID int,
    InvitedBy int,
    Confirmed int,
    CONSTRAINT single UNIQUE (Email, GroupID)
);