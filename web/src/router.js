import VueRouter from "vue-router";

import Login from './components/Login';
import Chat from './components/Chat';
import SignUp from './components/SignUp';
import store from "./store";
import Profile from "./components/Profile";

const routes = [
    { path: '/login', component: Login, name: 'login' },
    { path: '/sign-up', component: SignUp, name: 'sign-up' },
    { path: '/', component: Chat, name: 'chat' },
    { path: '/profile', component: Profile, name: 'profile' },
]

const router = new VueRouter({
    routes,
    mode: 'history',
})

router.beforeEach((to, from, next) => {
    if (to.name !== 'login' && to.name !== 'sign-up' && !store.state.user) {
        next({ name: 'login' });
        return;
    }

    if ((to.name === 'login' || to.name === 'sign-up') && store.state.user) {
        next({ name: 'chat' });
        return;
    }

    next()
});

export default router;