const store = {
    debug: true,
    localStorage: window.localStorage,
    state: {
        chats: [],
        user: null,
    },

    init () {
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
    },

    setUser (newValue) {
        if (this.debug) {
            console.log('Set user ', newValue);
        }
        this.state.user = newValue;
        if (typeof this.state.user === 'object') {
            this.localStorage.setItem('user', JSON.stringify(this.state.user));
        } else {
            this.localStorage.setItem('user', null);
        }

    },

    getUser () {
        if (this.debug) {
            console.log('Get user ', this.state.user);
        }
        return this.state.user;
    }
}

export default store;