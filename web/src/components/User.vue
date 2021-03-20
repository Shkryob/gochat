<template>
  <v-container fluid>
    <v-row>
      <v-col v-if="user">
        <h1 class="text-center">{{ user.username }}</h1>
        <img :src="user.avatar" :alt="user.username">

        <v-row>
          <v-col class="text-right">
            <v-icon @click="startChat(user)">
              mdi-message-outline
            </v-icon>
          </v-col>
          <v-col class="text-center">
            <v-icon @click="toggleFriends(user)" :color="user.friends ? 'red' : ''">
              mdi-heart-outline
            </v-icon>
          </v-col>
          <v-col>
            <v-icon @click="toggleBlacklist(user)" :color="user.blacklisted ? 'red' : ''">
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
    getUser() {
      (new api()).getUser(this.$route.params.id).then((response) => {
        this.user = response.data;
      });
    },

    startChat() {
      (new api()).createChat([this.user.id]).then(() => {});
    },

    toggleFriends() {
      if (this.user.friends) {
        (new api()).removeFriend(this.$route.params.id).then(() => {
          this.user.friends = false;
        });
      } else {
        (new api()).addFriend(this.$route.params.id).then(() => {
          this.user.friends = true;
        });
      }
    },

    toggleBlacklist() {
      if (this.user.blacklisted) {
        (new api()).removeBlacklist(this.$route.params.id).then(() => {
          this.user.blacklisted = false;
        });
      } else {
        (new api()).addBlacklist(this.$route.params.id).then(() => {
          this.user.blacklisted = true;
        });
      }
    },
  },
};
</script>