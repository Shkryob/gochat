import VueRouter from "vue-router";

import Login from './components/Login';
import Chat from './components/Chat';
import SignUp from './components/SignUp';
import store from "./store";
import Profile from "./components/Profile";
import Users from "./components/Users";
import User from "./components/User";

const routes = [
    { path: '/login', component: Login, name: 'login' },
    { path: '/sign-up', component: SignUp, name: 'sign-up' },
    { path: '/profile', component: Profile, name: 'profile' },
    { path: '/users/:id', component: User, name: 'user' },
    { path: '/users', component: Users, name: 'users' },
    { path: '/:id', component: Chat, name: 'chat' },
    { path: '/', component: Chat, name: 'chat' },
];

const router = new VueRouter({
    routes,
    mode: 'history',
});

router.beforeEach((to, from, next) => {
    if (to.name !== 'login' && to.name !== 'sign-up' && !store.state.user) {
        next({ name: 'login' });
        return;
    }

    if ((to.name === 'login' || to.name === 'sign-up') && store.state.user) {
        next({ name: 'chat' });
        return;
    }

    next();
});

export default router;