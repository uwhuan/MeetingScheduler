export default {
    base: "https://api.info441deploy.me",
    testbase: "https://localhost:4000",
    handlers: {
        users: "/v1/users",
        myuser: "/v1/users/me",
        sessions: "/v1/sessions",
        sessionsMine: "/v1/sessions/mine",
        resetPasscode: "/v1/resetcodes",
        passwords: "/v1/passwords/"
    }
}