<template>
  <span>
    <v-avatar size="128">
      <img
          v-if="src"
          :src="src"
          alt="John"
      >
      <v-icon color="darken-3" size="128" v-if="!src">
        mdi-account-circle-outline
      </v-icon>
    </v-avatar>
  </span>
</template>

<script>
import api from "../api";

export default {
  name: 'Avatar',
  data: function () {
    return {
      src: '',
    };
  },
  props: {id: Number},

  created: function () {
    this.loadAvatar();
  },

  methods: {
    loadAvatar: function () {
      if (!this.id) {
        return;
      }
      (new api()).getAvatar(this.id).then((response) => {
        const base64 = btoa(
            new Uint8Array(response.data).reduce(
                (data, byte) => data + String.fromCharCode(byte),
                '',
            ),
        );
        this.src = 'data:;base64,' + base64;
      });
    },
  },

  watch: {
    id() {
      this.loadAvatar();
    }
  }
};
</script>