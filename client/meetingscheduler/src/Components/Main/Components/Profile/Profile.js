import React, { Component } from 'react';
import api from '../../../../Constants/APIEndpoints/APIEndpoints';
import PageTypes from '../../../../Constants/PageTypes/PageTypes';
import Errors from '../../../Errors/Errors';
import '../../Styles/MainPageContent.css'
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';

class Profile extends Component {
    constructor(props) {
        super(props);
        this.state = {
            error: ''
        }
    }

    render() {
        return <div className="profile-page">
            <div className="profile">
                <h1>Here's your Profile:</h1>
                <Card style={{width: '40rem'}}>
                    <Card.Body>
                        <h2>{this.props.user.firstName} {this.props.user.lastName}</h2>
                        <h2>Username: {this.props.user.username} </h2>
                        <h4>Username: <span className="red">{this.props.user.userName}</span></h4>
                        <Button variant="primary" onClick={(e) => {
                            this.props.setPage(e, PageTypes.signedInUpdateName)
                        }}>EDIT PROFILE</Button>
                    </Card.Body>
                </Card>
                <div className="space"/>
            </div>
        </div>
    }
}
export default Profile;
