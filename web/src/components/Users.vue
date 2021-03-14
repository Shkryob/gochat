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
        <v-row>
          <v-col class="text-center">
            {{ user.username }}
          </v-col>
        </v-row>

        <v-row>
          <v-col class="text-right">
            <v-icon @click="startChat(user)">
              mdi-message-outline
            </v-icon>
          </v-col>
          <v-col class="text-center">
            <v-icon>
              mdi-heart-outline
            </v-icon>
          </v-col>
          <v-col>
            <v-icon>
              mdi-block-helper
            </v-icon>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import api from "../api";

export default {
  name: 'Users',

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

    startChat: function (user) {
      (new api()).createChat([user.id]).then(() => {});
    },
  },
};
</script>