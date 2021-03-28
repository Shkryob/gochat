<template>
  <v-app>
    <Header/>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
import Header from "./components/Header";
import store from "./store";

export default {
  name: 'App',

  components: {
    Header,
  },

  created() {

    store.eventBus.$on('message_created', (data) => {
      if (data.message.user.id !== store.state.user.id) {
        this.$notification.show('New message from ' + data.message.user.username, {
          body: data.message.body,
        }, {
          onclick: () => {
            this.$router.push({
              name: 'chat', params: { id: data.message.chat.id }
            });
          }
        });
      }
    });
  },

  data: () => ({
    //
  }),
};
</script>
