import React, { Component } from 'react';
import PageTypes from '../../Constants/PageTypes/PageTypes';
import FirstPageContent from './Content/MainPageContent/FirstPageContent';
import SignOutButton from './Components/SignOutButton/SignOutButton';
import UpdateName from './Components/UpdateName/UpdateName';
import UpdateAvatar from './Components/UpdateAvatar/UpdateAvatar';
import Profile from './Components/Profile/Profile';
import { BrowserRouter, Route, Switch, Link, NavLink, Redirect } from 'react-router-dom';
const Main = ({ page, setPage, setAuthToken, setUser, user, thread, setThread, post, setPost}) => {
    let content = <></>
    let contentPage = true;
    window.scrollTo(0, 0);
    switch (page) {
        case PageTypes.signedInMain:
            content = <FirstPageContent user={user} setPage={setPage} />;
            break;
        case PageTypes.signedInUpdateName:
            content = <UpdateName user={user} setUser={setUser} setPage={setPage} />;
            break;
        case PageTypes.signedInUpdateAvatar:
            content = <UpdateAvatar user={user} setUser={setUser} />;
            break;
        case PageTypes.profile:
            content = <Profile user={user} setUser={setUser} setPage={setPage}/>;
            break;
        default:
            content = <>Error, invalid path reached</>;
            contentPage = false;
            break;
    }
    return <>
        <div>
            <nav>
                <div className="one">
                    <ul>
                        <li id="home">
                            <div><button onClick={(e) => { setPage(e, PageTypes.signedInMain) }}>HOME</button></div>
                        </li>
                        <li id="profile">
                            <div><button onClick={(e) => { setPage(e, PageTypes.profile) }}>PROFILE</button></div>
                        </li>
                    </ul>
                </div>
                        <div className="display-user">
                            <h1>Logged in as: <span className="red">{user.userName}</span></h1>
                        </div>
                        <SignOutButton setUser={setUser} setAuthToken={setAuthToken} />
            </nav>
        </div>
        {content}
    </>
}

export default Main;