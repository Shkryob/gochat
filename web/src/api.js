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

    this.getUsers = (search) => {
        return axios.get(basePath + 'users?search=' + search);
    }

    this.getUser = (id) => {
        return axios.get(basePath + 'users/' + id);
    }

    this.getAvatar = (id) => {
        return axios.get(basePath + 'users/' + id + '/avatar', { responseType: 'arraybuffer' });
    }

    this.getChats = () => {
        return axios
            .get(basePath + 'chats');
    };

    this.createChat = (participants) => {
        return axios
            .post(basePath + 'chats', {
                chat: {
                    participants
                }
            });
    };

    this.removeChat = (chatID) => {
        return axios
            .delete(basePath + 'chats/' + chatID);
    };

    this.getMessages = (chatID) => {
        return axios
            .get(basePath + 'chats/' + chatID + '/messages');
    };

    this.sendMessage = (chatID, body) => {
        return axios
            .post(basePath + 'chats/' + chatID + '/messages', {
                message: {
                    body
                }
            });
    };
}

export default api;