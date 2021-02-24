<template>
  <v-app-bar
      app
      color="indigo"
      dark
  >
    <div class="d-flex align-center">
      <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
          transition="scale-transition"
          width="40"
      />

      GoChat
    </div>

    <v-spacer></v-spacer>

    <v-btn
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
        text
    >
      <span class="mr-2">Ivan Kukuruzko</span>
      <v-icon>mdi-account-circle</v-icon>
    </v-btn>
  </v-app-bar>
</template>