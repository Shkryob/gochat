<template>
  <v-container fluid>
    <v-row>
      <v-col>
        <h1>{{ user.username }}</h1>
        <img :src="user.avatar">
      </v-col>
    </v-row>

  </v-container>
</template>

<script>
import api from "../api";

export default {
  name: 'User',

  data: function () {
    return {
      user: {},
    };
  },

  created() {
    this.getUser();
  },

  methods: {
    getUser: function () {
      (new api()).getUser(this.$route.params.id).then((response) => {
        this.user = response.data.user;
      });
    },

    startChat: function () {
      (new api()).createChat([this.$route.params.id]).then((response) => {
        this.user = response.chat.id;
      });
    },
  },
};
</script>