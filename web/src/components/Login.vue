<template>
  <v-container fluid class="fill-height">
    <v-row align="center" justify="center">
      <v-col cols="4" align-self="center">
        <h1>Login</h1>
        <v-form @submit="login($event)">
          <v-text-field label="Email" v-model="email" type="email"/>
          <v-text-field label="Password" v-model="password" type="password"/>
          <v-row>
            <v-col>
              <router-link :to="{name: 'sign-up'}">Sign Up</router-link>
            </v-col>
            <v-col class="text-right">
              <v-btn color="indigo" dark type="submit">
                <v-icon left>mdi-login</v-icon>
                Login
              </v-btn>
            </v-col>
          </v-row>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import store from '../store';
import api from "../api";

export default {
  name: 'Login',

  data: function () {
    return {
      email: '',
      password: '',
    };
  },

  methods: {
    login: function ($event) {
      $event.preventDefault();

      (new api()).login(this.email, this.password).then((response) => {
        store.setUser(response.data.user);
        this.$router.push('/');
      });
    },
  },
};
</script>