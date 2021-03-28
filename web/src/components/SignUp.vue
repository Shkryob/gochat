<template>
  <v-container fluid class="fill-height">
    <v-row align="center" justify="center">
      <v-col cols="4" align-self="center">
        <h1>Sign Up</h1>
        <v-form @submit="signup($event)">
          <v-text-field label="Username" v-model="username"/>
          <v-text-field label="Email" v-model="email" type="email"/>
          <v-text-field label="Password" v-model="password" type="password"/>
          <v-text-field label="Confirm Password" v-model="password_confirm" type="password"/>

          <v-row>
            <v-col>
              <router-link :to="{name: 'login'}">Login</router-link>
            </v-col>
            <v-col class="text-right">
              <v-btn color="indigo" dark type="submit">
                Sign Up
              </v-btn>
            </v-col>
          </v-row>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import api from "../api";
import store from '../store';

export default {
  name: 'SignUp',

  data: function () {
    return {
      username: '',
      email: '',
      password: '',
      password_confirm: '',
    };
  },

  methods: {
    signup: function ($event) {
      $event.preventDefault();

      (new api()).signup(this.username, this.email, this.password).then((response) => {
        store.setUser(response.data.user);
        this.$router.push('/');
      });
    },
  },
};
</script>