<template>
  <v-container fluid>
    <v-row>
      <v-col v-if="userData">
        <h1 class="text-center">{{ userData.username }}<Avatar :id="userData.id" /></h1>
        <v-row>
          <v-col class="text-right">
            <v-icon @click="startChat(userData)">
              mdi-message-outline
            </v-icon>
          </v-col>
          <v-col class="text-center">
            <v-icon @click="toggleFriends(userData)" :color="userData.friends ? 'red' : ''">
              mdi-heart-outline
            </v-icon>
          </v-col>
          <v-col>
            <v-icon @click="toggleBlacklist(userData)" :color="userData.blacklisted ? 'red' : ''">
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
import Avatar from "./Avatar";

export default {
  name: 'User',

  components: {
    Avatar,
  },

  props: {user: Object},

  data() {
    return {
      userData: {},
    };
  },

  created() {
    if (this.user) {
      this.userData = this.user;
    } else {
      this.getUser();
    }
  },

  methods: {
    getUser() {
      (new api()).getUser(this.$route.params.id).then((response) => {
        this.userData = response.data;
      });
    },

    startChat() {
      (new api()).createChat([this.userData.id]).then(() => {});
    },

    toggleFriends() {
      if (this.userData.friends) {
        (new api()).removeFriend(this.userData.id).then(() => {
          this.userData.friends = false;
        });
      } else {
        (new api()).addFriend(this.userData.id).then(() => {
          this.userData.friends = true;
        });
      }
    },

    toggleBlacklist() {
      if (this.userData.blacklisted) {
        (new api()).removeBlacklist(this.userData.id).then(() => {
          this.userData.blacklisted = false;
        });
      } else {
        (new api()).addBlacklist(this.userData.id).then(() => {
          this.userData.blacklisted = true;
        });
      }
    },
  },
};
</script>