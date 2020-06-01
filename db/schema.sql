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

create table if not exists usermeeting (
    MeetingID int not null auto_increment primary key,
    Name varchar(255) not null,
    CreatorID int not null,
    GroupID int not null,
    Description varchar(320),
    StartTime datetime not null, 
    EndTime datetime not null,  
    Confirmed int not null,  
    CreateDate datetime not null 
)

create table if not exists meetingparticipant (
    MeetingID int not null primary key,
    UID int not null primary key
)

create table if not exists logentry (
    UserID int not null auto_increment primary key,
    SignInTime datetime not null,
    IPAddress varchar(32) not null
)

create table if not exists groups (
    GroupID int not null primary key,
    Description varchar(500),
    Name varchar(100),
    CreatorID int not null,
    CreateDate varchar(300)
)

create table if not exists schedule (
    ScheduleID int not null primary key,
    MeetingID int not null primary key,
    StartTime varchar(300) not null,
    EndTime varchar(300) not null,
    Votes int
)

create table if not exists membership (
    uid int not null primary key,
    GroupID int not null primary key
)