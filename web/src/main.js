import Vue from 'vue';
import VueRouter from 'vue-router';
import VueNativeSock from 'vue-native-websocket';

import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './router';
import store from "./store";
import './main.css';

store.init(Vue);

Vue.config.productionTip = false;

Vue.use(VueRouter);

Vue.use(VueNativeSock, 'ws://localhost/api/sockets', {
  store,
  reconnection: true,
  reconnectionAttempts: 5,
  reconnectionDelay: 3000,
  passToStoreHandler: function (eventName, event) {
    if (!eventName.startsWith('SOCKET_')) {
      return;
    }
    let target = eventName.toUpperCase();

    this.store['onSocketEvent'](target, event);
  },
});

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
