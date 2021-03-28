import Vue from 'vue';
import VueRouter from 'vue-router';
import VueNativeSock from 'vue-native-websocket';
import VueNativeNotification from 'vue-native-notification';

import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './router';
import store from "./store";
import "./filters";
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

Vue.use(VueNativeNotification, {
  // Automatic permission request before
  // showing notification (default: true)
  requestOnNotify: true
});

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
