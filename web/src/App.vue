<template>
  <v-app>
    <Header/>
    <v-main>
      <Chat/>
    </v-main>
  </v-app>
</template>

<script>
import Chat from "./components/Chat";
import Header from "./components/Header";

export default {
  name: 'App',

  components: {
    Chat,
    Header,
  },

  data: () => ({
    //
  }),
};
</script>
