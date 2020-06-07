import React, { Component } from 'react';
import api from '../../../../Constants/APIEndpoints/APIEndpoints';
import Errors from '../../../Errors/Errors';
import PageTypes from '../../../../Constants/PageTypes/PageTypes';

class UpdateName extends Component {
    constructor(props) {
        super(props);
        this.state = {
            firstName: '',
            lastName: '',
            error: ''
        }
    }

    sendRequest = async (e) => {
        e.preventDefault();
        const { firstName, lastName } = this.state;
        const sendData = { firstName, lastName };
        const response = await fetch(api.base + api.handlers.myuser, {
            method: "PATCH",
            body: JSON.stringify(sendData),
            headers: new Headers({
                "Authorization": localStorage.getItem("Authorization"),
                "Content-Type": "application/json"
            })
        });
        if (response.status >= 300) {
            const error = await response.text();
            console.log(error);
            this.setError(error);
            return;
        }
        alert("Profile updated successfully!") // TODO make this better by refactoring errors
        const user = await response.json();
        this.props.setUser(user);
        this.props.setPage(e, PageTypes.profile);
    }

    setValue = (e) => {
        this.setState({ [e.target.name]: e.target.value });
    }

    setError = (error) => {
        this.setState({ error })
    }

    render() {
        const { firstName, lastName, error } = this.state;
        return <div className="updateProfile">
            <Errors error={error} setError={this.setError} />
            <h2>Edit your profile:</h2>
            <form onSubmit={this.sendRequest}>
                <div>
                    <span>First name: </span>
                    <input name={"firstName"} value={firstName} onChange={this.setValue} />
                </div>
                <div>
                    <span>Last name: </span>
                    <input name={"lastName"} value={lastName} onChange={this.setValue} />
                </div>
                <input type="submit" value="Update Profile" />
                <input type="submit" value="Go Back" onClick={(e) => this.props.setPage(e, PageTypes.profile)}/>
            </form>
        </div>
    }

}

export default UpdateName;