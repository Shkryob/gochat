<template>
  <v-container fluid>
    <v-row>
      <v-col v-if="user">
        <h1 class="text-center">{{ user.username }}</h1>
        <img :src="user.avatar">

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
        this.user = response.data;
        console.log('this.user', this.user);
      });
    },

    startChat: function () {
      (new api()).createChat([this.user.id]).then(() => {});
    },
  },
};
</script>