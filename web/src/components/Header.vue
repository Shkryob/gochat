<template>
  <v-app-bar
      app
      color="indigo"
      dark
  >
    <v-btn @click="toggleLeftMenu()"
           light
           class="ml-2 hidden-md-and-up"
           v-if="['chat', 'chats'].indexOf($route.name) > -1">
      <v-icon>mdi-menu</v-icon>
    </v-btn>
    <v-spacer class="hidden-md-and-up"></v-spacer>
    <div class="d-flex align-center">
      <router-link :to="{name: 'chats'}" class="logo">
        <img
            alt="Vuetify Logo"
            class="shrink mr-2 hidden-sm-and-down"
            src="/logo.svg"
            width="40"
        />
        GoChat
      </router-link>
    </div>

    <v-spacer></v-spacer>

    <v-btn :to="{name: 'profile'}" v-if="sharedState.user" light>
      <span class="mr-2 hidden-sm-and-down">{{sharedState.user.username}}</span>
      <v-icon>mdi-account-circle</v-icon>
    </v-btn>
    <v-btn @click="logout()" v-if="sharedState.user" light class="ml-2">
      <span class="mr-2 hidden-sm-and-down">Logout</span>
      <v-icon>mdi-logout</v-icon>
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

  data() {
    return {
      sharedState: store.state
    };
  },

  methods: {
    logout() {
      store.setUser(null);
      this.$router.push({'name': 'login'})
    },

    toggleLeftMenu() {
      store.toggleLeftMenu();
    },
  }
};
</script>