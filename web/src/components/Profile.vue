<template>
  <v-container fluid class="fill-height">
    <v-row align="center" justify="center">
      <v-col cols="4" align-self="center">
        <h1>Profile</h1>
        <Avatar :id="id" />
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

        <v-form @submit="submitFile($event)">
          <v-file-input v-model="file"
                        prepend-icon="mdi-camera"
                        accept="image/png, image/jpeg, image/bmp"/>
          <v-row>
            <v-col class="text-right">
              <v-btn color="indigo" dark type="submit">
                Upload
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
import axios from "axios";
import Avatar from "./Avatar";

export default {
  name: 'Profile',

  components: {
    Avatar,
  },

  data: function () {
    return {
      id: '',
      username: '',
      email: '',

      password: '',
      password_confirm: '',

      file: null,
    };
  },

  created: function () {
    const user = store.getUser();

    this.username = user.username + '';
    this.email = user.email + '';
    this.id = user.id + '';
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

    submitFile($event) {
      $event.preventDefault();

      const formData = new FormData();
      formData.append('file', this.file);
      axios.post('/api/users/avatar',
        formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        }
      ).then(function () {
        console.log('SUCCESS!!');
      }).catch(function () {
        console.log('FAILURE!!');
      });
    },
  },
};
</script>