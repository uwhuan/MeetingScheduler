import React, { Component } from 'react';
import api from '../../../../Constants/APIEndpoints/APIEndpoints';
import PageTypes from '../../../../Constants/PageTypes/PageTypes';
import '../../Styles/MainPageContent.css'
import Errors from '../../../Errors/Errors';

class Meetings extends Component {
    constructor(props) {
        super(props);
        this.state = {
            meetings: [],
            meetingdetails: [],
            error: ''
        }
    }

    sendRequest = async (e) => {
        const response = await fetch(api.base + api.handlers.meetings, {
            method: "GET",
            headers: new Headers({
                "Authorization": localStorage.getItem("Authorization")
            })
        });
        if (response.status >= 300) {
            const error = await response.text();
            this.setError(error);
            return;
        }
        const meetingResponse = await response.json();
        this.setState({
            meetings: meetingResponse
        })
    }

    sendRequestTwo = async (e) => {
        const response = await fetch(api.base + api.handlers.meetings + this.props.user.uid, {
            method: "GET",
            headers: new Headers({
                "Authorization": localStorage.getItem("Authorization")
            })
        });
        if (response.status >= 300) {
            const error = await response.text();
            this.setError(error);
            return;
        }
        const specificMeetingResponse = await response.json();
        this.setState({
            meetingdetails: specificMeetingResponse
        })
    }

    setError = (error) => {
        this.setState({error})
    }

    render() {
        const {error} = this.state;
        const meetingInfo = this.state.meetingdetails;
        console.log(meetingInfo)
        const meetinglist = this.state.meetings;
        console.log(meetinglist)
        return <div className="meetings">
            <Errors error={error} setError={this.setError}/>
            <h1>Meeting Details:</h1>
            <h2>current user meetings:</h2>
            {meetingInfo}
            <h2>All meeting details:</h2>
            {meetinglist}
        </div>
    }

}
export default Meetings;