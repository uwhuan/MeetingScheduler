import React, { Component } from 'react';
import PageTypes from '../../../../Constants/PageTypes/PageTypes';

export class FirstPageContent extends Component {
    render() {
        let content = (
            <div className= "intro">
                <h1>Welcome to Meeting Scheduler</h1>
                <p><span>Meeting Schedule helps teams solve the difficulty of scheduling a meeting time. It&rsquo;s
                    a simple but powerful app for anyone who works or studies in a team. Each of the team members
                    only needs to select their available time within the range the host set, then the app will
                    automatically calculate the best available time for everyone.There are many similar apps online,
                    such as when2meet, Google Calendar. We want to build a new one that is more user-friendly and
                    powerful than when2meet, but more simple and light-weighted than Google Calendar. We want to build
                    a more instinctive and simple UI so that anyone with an invite link can use the app easily. At the
                    same time, we want to offer more useful functionalities for the registered user, such as managing
                    their events, creating new events, sharing, etc.</span>&nbsp;</p>
                <h3>Main Features:</h3>
                <ul>
                    <li><u>Create meeting with possible time range, location, participants and more detailed information (such as zoom link)</u></li>
                    <li><u>Create groups and manage group meeting</u></li>
                </ul>
            </div>
        );
        return content;
    }
}

export default FirstPageContent;