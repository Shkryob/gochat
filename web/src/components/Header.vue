<template>
  <v-app-bar
      app
      color="indigo"
      dark
  >
    <div class="d-flex align-center">
      <router-link :to="{name: 'chats'}" class="logo">
        <img
            alt="Vuetify Logo"
            class="shrink mr-2"
            contain
            src="/logo.svg"
            transition="scale-transition"
            width="40"
        />
        GoChat
      </router-link>
    </div>

    <v-spacer></v-spacer>

    <v-btn :to="{name: 'profile'}" v-if="sharedState.user" light>
      <span class="mr-2">{{sharedState.user.username}}</span>
      <v-icon>mdi-account-circle</v-icon>
    </v-btn>
    <v-btn @click="logout()" v-if="sharedState.user" light class="ml-2">
      <v-icon left>mdi-logout</v-icon>
      Logout
    </v-btn>
  </v-app-bar>
</template>

<style>
.v-application a.logo {
  font-size: 38px;
  color: white;
  text-decoration: none;
}
</style>

<script>
import store from "../store";

export default {
  name: 'Header',

  data: function () {
    return {
      sharedState: store.state
    };
  },

  methods: {
    logout: function () {
      store.setUser(null);
      this.$router.push({'name': 'login'})
    }
  }
};
</script>