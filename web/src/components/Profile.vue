<template>
  <v-container fluid class="fill-height">
    <v-row align="center" justify="center">
      <v-col cols="4" align-self="center">
        <h1>Profile</h1>
        <v-form @submit="updateProfile($event)">
          <v-text-field label="Username" v-model="username"/>
          <v-text-field label="Email" v-model="email" type="email" disabled/>
          <v-row>
            <v-col class="text-right">
              <v-btn color="indigo" dark type="submit">
                Update
              </v-btn>
            </v-col>
          </v-row>
        </v-form>

        <v-form @submit="updatePassword($event)">
          <v-text-field label="New Password" v-model="password" type="password"/>
          <v-text-field label="Confirm New Password" v-model="password_confirm" type="password"/>

          <v-row>
            <v-col class="text-right">
              <v-btn color="indigo" dark type="submit">
                Update
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
  name: 'Profile',

  data: function () {
    return {
      username: '',
      email: '',

      password: '',
      password_confirm: '',
    };
  },

  created: function () {
    const user = store.getUser();

    this.username = user.username + '';
    this.email = user.email + '';
  },

  methods: {
    updateProfile: function ($event) {
      $event.preventDefault();

      (new api()).signup(this.username, this.email, this.password).then((response) => {
        store.setUser(response.data.user);
        this.$router.push('/');
      });
    },

    updatePassword: function ($event) {
      $event.preventDefault();
    },
  },
};
</script>