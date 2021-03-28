<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <v-form @submit="getUsers($event)">
          <v-text-field
              label="Search"
              prepend-icon="mdi-magnify"
              v-model="search"
          ></v-text-field>
        </v-form>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="4" v-for="user in users" :key="user.username">
        <User :user="user"/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import api from "../api";
import User from "./User";

export default {
  name: 'Users',

  components: {
    User,
  },

  data: function () {
    return {
      users: [],
      search: '',
    };
  },

  created() {
    this.getUsers();
  },

  methods: {
    getUsers: function ($event) {
      if ($event) {
        $event.preventDefault();
      }

      (new api()).getUsers(this.search).then((response) => {
        this.users = response.data.users;
        console.log('this.users', this.users);
      });
    },
  },
};
</script>