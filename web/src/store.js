import axios from "axios";

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

    init (Vue) {
        this.vue = Vue;
        this.eventBus = new Vue;

        const user = this.localStorage.getItem('user');
        if (typeof user === 'string') {
            try {
                this.state.user = JSON.parse(user);
            } catch (e) {
                console.log('Can\'t parse user state', e);
            }

        }
        if (this.debug) {
            console.log('Init ', this.state.user);
        }

        axios.interceptors.request.use((config) => {
            if (this.state.user) {
                config.headers.Authorization = 'Token ' + this.state.user.token;
            }

            return config;
        });
    },

    setUser (newValue) {
        if (this.debug) {
            console.log('Set user ', newValue);
        }
        this.state.user = newValue;
        if (typeof this.state.user === 'object') {
            this.localStorage.setItem('user', JSON.stringify(this.state.user));
            this.authorizeSocket();
        } else {
            this.localStorage.setItem('user', null);
        }

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

    addMessage(message) {
        const messageData = JSON.parse(message.data);
        this.eventBus.$emit('message-received', messageData.message);
    },

    onSocketEvent(target, event) {
        if (target === 'SOCKET_ONOPEN') {
            this.authorizeSocket();
        } else if (target === 'SOCKET_ONMESSAGE') {
            this.addMessage(event)
        }
    },

    hideLeftMenu() {
        this.state.showLeftMenu = false;
        this.eventBus.$emit('toggle-left-menu', this.state.showLeftMenu);
    },

    toggleLeftMenu() {
        this.state.showLeftMenu = !this.state.showLeftMenu;
        console.log('this.state.showLeftMenu', this.state.showLeftMenu);
        this.eventBus.$emit('toggle-left-menu', this.state.showLeftMenu);
    },
}

export default store;