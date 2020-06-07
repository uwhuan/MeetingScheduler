import React, { Component } from 'react';
import api from '../../../../Constants/APIEndpoints/APIEndpoints';
import Errors from '../../../Errors/Errors';
import PageTypes from '../../../../Constants/PageTypes/PageTypes';

class Groups extends Component {
    constructor(props) {
        super(props);
        this.state = {
            url: "",
            group: "",
            error: ''
        }
    }

    sendRequest = async (e) => {
        const response = await fetch(api.base + api.handlers., {
            method: "GET",
            headers: new Headers({
                "Authorization": localStorage.getItem("Authorization")
            })
        });
}