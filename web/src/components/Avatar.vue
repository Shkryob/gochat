<template>
  <div>
    <v-avatar>
      <img
          :src="src"
          alt="John"
      >
    </v-avatar>
  </div>
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
  props: ['id'],

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