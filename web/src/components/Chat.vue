<template>
  <v-container fluid>
    <v-row>
      <v-col cols="4" class="left-sidebar">
        <v-text-field
            label="Search"
            prepend-icon="mdi-magnify"
        ></v-text-field>

        <div class="chat-list">
        <v-list>
          <v-list-item-group v-model="selectedItem" color="primary">
            <v-list-item v-for="(item, i) in items" :key="i">
              <v-list-item-icon>
                <v-icon v-text="item.icon"></v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title v-text="item.text"></v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>
        </div>
      </v-col>
      <v-col cols="8" class="message-list">
        <v-row v-for="(message, i) in messages" :key="i" class="pa-1">
          <v-chip :color="message.color">{{ message.text }}</v-chip>
        </v-row>
      </v-col>
    </v-row>

    <v-app-bar
        app
        bottom
    >
      <v-row>
        <v-col cols="4" class="text-center">
          <v-btn
              fab
              dark
              color="indigo"
          >
            <v-icon dark>
              mdi-plus
            </v-icon>
          </v-btn>
        </v-col>
        <v-col cols="8">
          <v-row>
            <v-col cols="10">
              <v-text-field label="Message"/>
            </v-col>
            <v-col cols="2" class="pt-5">
              <v-btn block color="indigo" dark>
                <v-icon
                    left

                >
                  mdi-send
                </v-icon>
                Send
              </v-btn>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </v-app-bar>
  </v-container>
</template>

<style>
.container.container--fluid,
.container.container--fluid > .row {
  height: 100%;
  padding: 0;
  margin: 0;
}

.message-list {
  height: 100%;
  padding-top: 20px;
  padding-bottom: 20px;
}

.left-sidebar {
  padding: 0;
  margin: 0;
  position: relative;
  height: 100%;
}

.chat-list {
  position: absolute;
  top: 70px;
  bottom: 0;
  overflow: auto;
  right: 0;
  left: 0;
  width: auto;
}

.message-list {
  height: 100%;
  overflow-y: auto;
}
</style>

<script>
export default {
  data: () => ({
    selectedItem: 1,
    items: [
      {text: 'Lorem ipsum dolor sit amet.', icon: 'mdi-clock'},
      {text: 'Species warp on peace at hyperspace!', icon: 'mdi-account'},
      {text: 'Meet, scotty, solid advice!', icon: 'mdi-flag'},
      {text: 'Lorem ipsum dolor sit amet.', icon: 'mdi-clock'},
      {text: 'Species warp on peace at hyperspace!', icon: 'mdi-account'},
      {text: 'Meet, scotty, solid advice!', icon: 'mdi-flag'},
      {text: 'Lorem ipsum dolor sit amet.', icon: 'mdi-clock'},
      {text: 'Species warp on peace at hyperspace!', icon: 'mdi-account'},
      {text: 'Meet, scotty, solid advice!', icon: 'mdi-flag'},
      {text: 'Lorem ipsum dolor sit amet.', icon: 'mdi-clock'},
      {text: 'Species warp on peace at hyperspace!', icon: 'mdi-account'},
      {text: 'Meet, scotty, solid advice!', icon: 'mdi-flag'},
      {text: 'Lorem ipsum dolor sit amet.', icon: 'mdi-clock'},
      {text: 'Species warp on peace at hyperspace!', icon: 'mdi-account'},
      {text: 'Meet, scotty, solid advice!', icon: 'mdi-flag'},
      {text: 'Lorem ipsum dolor sit amet.', icon: 'mdi-clock'},
      {text: 'Species warp on peace at hyperspace!', icon: 'mdi-account'},
      {text: 'Meet, scotty, solid advice!', icon: 'mdi-flag'},
    ],

    messages: [
      {text: 'Wildly convert a ship.', color: 'yellow'},
      {text: 'Hypnosis, attitude, and paralysis.', color: 'yellow'},
      {
        text: 'When the particle meets for deep space, all captains examine ancient, interstellar phenomenans.',
        color: 'yellow'
      },
      {text: 'Fly without love, and we won’t imitate a phenomenan.', color: 'green'},
      {text: 'Assimilation at the holodeck was the mystery of advice, raised to a final hur\'q.', color: 'green'},
      {text: 'The planet is more ferengi now than species. ugly and patiently final.', color: 'green'},
      {text: 'Why does the teleporter walk?', color: 'green'},
      {text: 'Teleporters go with metamorphosis!', color: 'green'},
      {text: 'Wildly convert a ship.', color: 'yellow'},
      {text: 'Hypnosis, attitude, and paralysis.', color: 'yellow'},
      {
        text: 'When the particle meets for deep space, all captains examine ancient, interstellar phenomenans.',
        color: 'yellow'
      },
      {text: 'Fly without love, and we won’t imitate a phenomenan.', color: 'green'},
      {text: 'Assimilation at the holodeck was the mystery of advice, raised to a final hur\'q.', color: 'green'},
      {text: 'The planet is more ferengi now than species. ugly and patiently final.', color: 'green'},
      {text: 'Why does the teleporter walk?', color: 'green'},
      {text: 'Teleporters go with metamorphosis!', color: 'green'},
      {text: 'Wildly convert a ship.', color: 'yellow'},
      {text: 'Hypnosis, attitude, and paralysis.', color: 'yellow'},
      {
        text: 'When the particle meets for deep space, all captains examine ancient, interstellar phenomenans.',
        color: 'yellow'
      },
      {text: 'Fly without love, and we won’t imitate a phenomenan.', color: 'green'},
      {text: 'Assimilation at the holodeck was the mystery of advice, raised to a final hur\'q.', color: 'green'},
      {text: 'The planet is more ferengi now than species. ugly and patiently final.', color: 'green'},
      {text: 'Why does the teleporter walk?', color: 'green'},
      {text: 'Teleporters go with metamorphosis!', color: 'green'},
      {text: 'Wildly convert a ship.', color: 'yellow'},
      {text: 'Hypnosis, attitude, and paralysis.', color: 'yellow'},
      {
        text: 'When the particle meets for deep space, all captains examine ancient, interstellar phenomenans.',
        color: 'yellow'
      },
      {text: 'Fly without love, and we won’t imitate a phenomenan.', color: 'green'},
      {text: 'Assimilation at the holodeck was the mystery of advice, raised to a final hur\'q.', color: 'green'},
      {text: 'The planet is more ferengi now than species. ugly and patiently final.', color: 'green'},
      {text: 'Why does the teleporter walk?', color: 'green'},
      {text: 'Teleporters go with metamorphosis!', color: 'green', reverse: true},
    ]
  }),
}
</script>