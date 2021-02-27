import axios from 'axios';

const api = function () {
    const basePath = '/api/';

    this.signup = (username, email, password) => {
        return axios
            .post(basePath + 'users', {
                user: {
                    username,
                    email,
                    password,
                },
            });
            // .then(response => (this.info = response))
    };

    this.login = (email, password) => {
        return axios
            .post(basePath + 'users/login', {
                user: {
                    email,
                    password,
                },
            });
    };
}

export default api;