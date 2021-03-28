import axios from 'axios';

const api = function () {
    const basePath = '/api/';

    this.signup = (username, email, password) => {
        return axios
            .post(basePath + 'users/signup', {
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

    this.refreshToken = () => {
        return axios.get(basePath + 'users/me');
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

    this.addBlacklist = (id) => {
        return axios.post(basePath + 'users/' + id + '/blacklist');
    }

    this.removeBlacklist = (id) => {
        return axios.delete(basePath + 'users/' + id + '/blacklist');
    }

    this.addFriend = (id) => {
        return axios.post(basePath + 'users/' + id + '/friend');
    }

    this.removeFriend = (id) => {
        return axios.delete(basePath + 'users/' + id + '/friend');
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

    this.updateMessage = (chatID, messageID, body) => {
        return axios
            .put(basePath + 'chats/' + chatID + '/messages/' + messageID, {
                message: {
                    body
                }
            });
    };

    this.removeMessage = (chatID, messageID) => {
        return axios
            .delete(basePath + 'chats/' + chatID + '/messages/' + messageID);
    };
}

export default api;