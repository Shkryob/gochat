import axios from "axios";
import jwt_decode from "jwt-decode";
import api from "./api";

const store = {
    debug: true,
    localStorage: window.localStorage,
    state: {
        chats: [],
        user: null,
    },
    vue: null,
    eventBus: null,
    showLeftMenu: false,
    refreshTimeout: null,

    init (Vue) {
        this.vue = Vue;
        this.eventBus = new Vue;

        this.restoreUserFromStorage();

        axios.interceptors.request.use((config) => {
            if (this.state.user) {
                config.headers.Authorization = 'Token ' + this.state.user.token;
            }

            return config;
        });

    },

    restoreUserFromStorage() {
        const user = this.localStorage.getItem('user');
        if (typeof user === 'string') {
            try {
                const parsedUser = JSON.parse(user);
                if (this.isTokenValid(parsedUser)) {
                    this.setUser(parsedUser);
                } else {
                    this.setUser(null);
                }
            } catch (e) {
                console.log('Can\'t parse user state', e);
            }
        }
    },

    isTokenValid(user) {
        const decodedToken = jwt_decode(user.token);
        const tokenExpiresIn = decodedToken['exp'] * 1000 - Date.now();

        return tokenExpiresIn > 0;
    },

    setUser (newValue) {
        if (this.debug) {
            console.log('Set user ', newValue);
        }
        if (this.refreshTimeout) {
            clearTimeout(this.refreshTimeout);
            this.refreshTimeout = null;
        }
        this.state.user = newValue;
        if (typeof this.state.user === 'object' && this.state.user) {
            this.localStorage.setItem('user', JSON.stringify(this.state.user));
            this.authorizeSocket();
            const decodedToken = jwt_decode(this.state.user.token);
            const tokenExpiresIn = decodedToken['exp'] * 1000 - Date.now();

            this.refreshTimeout = setTimeout(() => {
                this.refreshTimeout = null;
                this.refreshToken();
            }, (tokenExpiresIn - 2000) > 0 ? (tokenExpiresIn - 2000) : 0);
        } else {
            this.localStorage.setItem('user', null);
        }
    },

    refreshToken() {
        (new api()).refreshToken().then((response) => {
            this.setUser(response.data.user);
        });
    },

    getUser () {
        if (this.debug) {
            console.log('Get user ', this.state.user);
        }
        return this.state.user;
    },

    authorizeSocket() {
        if (!this.state.user || !this.vue.prototype.$socket) {
            return;
        }
        this.vue.prototype.$socket.send(JSON.stringify({
            'action': 'authorize',
            'jwt': this.state.user.token,
        }));
    },

    emitSocketEvent(message) {
        const messageData = JSON.parse(message.data);
        this.eventBus.$emit(messageData.event, messageData);
    },

    onSocketEvent(target, event) {
        if (target === 'SOCKET_ONOPEN') {
            this.authorizeSocket();
        } else if (target === 'SOCKET_ONMESSAGE') {
            this.emitSocketEvent(event)
        }
    },

    hideLeftMenu() {
        this.state.showLeftMenu = false;
        this.eventBus.$emit('toggle-left-menu', this.state.showLeftMenu);
    },

    toggleLeftMenu() {
        this.state.showLeftMenu = !this.state.showLeftMenu;
        this.eventBus.$emit('toggle-left-menu', this.state.showLeftMenu);
    },
}

export default store;