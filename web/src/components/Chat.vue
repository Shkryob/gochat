<template>
  <v-container fluid>
    <v-row>
      <v-col sm="12" md="4" lg="3" class="left-sidebar" :class="{'hidden-sm-and-down': !showLeftMenu}">
        <v-text-field
            label="Search"
            prepend-icon="mdi-magnify"
            v-model="searchText"
            class="pl-3 pr-3"
        ></v-text-field>

        <div class="chat-list">
        <v-list>
          <v-list-item-group color="primary">
            <v-list-item v-for="chat in filteredChats"
                         :key="chat.id"
                         :value="chat.id"
                         :class="{'v-list-item--active': chat.id === selectedItem}"
                         @click="setChat(chat)">
              <v-list-item-content>
                <v-list-item-title v-text="getChatTitle(chat)"></v-list-item-title>
              </v-list-item-content>
              <v-list-item-action>
                <v-icon color="darken-3" @click="removeChat(chat)">
                  mdi-delete
                </v-icon>
              </v-list-item-action>
            </v-list-item>
          </v-list-item-group>
        </v-list>
        </div>
      </v-col>
      <v-col sm="12" md="8" lg="9" class="message-list" ref="messages" :class="{'hidden-sm-and-down': showLeftMenu}">
        <v-row v-for="message in messages" :key="message.id" class="pa-1">
          <v-spacer v-if="sharedState.user.username === message.user.username"></v-spacer>
          <v-chip :color="getMessageColor(message)">
            {{ message.createdAt | formatDate }}
            <router-link :to="{name: 'user', params: {id: message.user.id}}" class="ml-1">
              {{ message.user.username }}
            </router-link>:
            {{ message.body }}
          </v-chip>
        </v-row>
      </v-col>
    </v-row>

    <v-app-bar
        app
        bottom
    >
      <v-row>
        <v-col sm="12" md="4" lg="3" class="text-center" :class="{'hidden-sm-and-down': !showLeftMenu}">
          <v-btn color="indigo" fab dark small :to="{name: 'users'}">
            <v-icon>
              mdi-plus
            </v-icon>
          </v-btn>
        </v-col>
        <v-col sm="12" md="8" lg="9" :class="{'hidden-sm-and-down': showLeftMenu}">
          <v-row>
            <v-col cols="10">
              <v-text-field label="Message" v-model="messageText"/>
            </v-col>
            <v-col cols="2" class="pt-5 text-right">
              <v-btn block color="indigo" dark @click="sendMessage" class="hidden-sm-and-down">
                <v-icon left>
                  mdi-send
                </v-icon>
                <span>Send</span>
              </v-btn>

              <v-btn color="indigo" fab dark small @click="sendMessage" class="hidden-md-and-up text-center">
                <v-icon>
                  mdi-send
                </v-icon>
              </v-btn>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </v-app-bar>
  </v-container>
</template>

<style scoped>
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
import api from "../api";
import store from "../store";

export default {
  data: () => ({
    selectedItem: null,
    chats: [],
    messages: [],
    messageText: '',
    sharedState: store.state,
    searchText: '',
    showLeftMenu: false,
  }),

  computed: {
    filteredChats() {
      if (!this.searchText) {
        return this.chats;
      }

      return this.chats.filter((chat) => {
        const title = this.getChatTitle(chat).toLowerCase();
        return title.includes(this.searchText.toLocaleLowerCase());
      });
    },
  },

  created() {
    this.getChats();

    store.eventBus.$on('message-received', (message) => {
      this.addMessage(message);
    });

    store.eventBus.$on('toggle-left-menu', (state) => {
      this.showLeftMenu = state;
    });

    if (this.$route.params.id) {
      this.selectedItem = parseInt(this.$route.params.id);
    }
  },

  mounted() {
    this.scrollToBottom();
  },

  methods: {
    getChats() {
      (new api()).getChats().then((response) => {
        this.chats = response.data.chats;
        if (this.chats.length > 0 && !this.selectedItem) {
          this.selectedItem = this.chats[0].id;
          this.getMessages();
        }
      });
    },

    getMessages() {
      if (this.selectedItem) {
        (new api()).getMessages(this.selectedItem).then((response) => {
          this.messages = response.data.messages;
          this.scrollToBottom();
          store.hideLeftMenu();
        });
      } else {
        this.messages = [];
      }
    },

    sendMessage() {
      (new api()).sendMessage(this.selectedItem, this.messageText).then(() => {
        this.messageText = '';
      });
    },

    getChatTitle(chat) {
      if (chat.title) {
        return chat.title;
      } else {
        return chat.participants.filter((user) => {
          return user.username !== this.sharedState.user.username;
        }).map((user) => {
          return user.username;
        }).join(', ');
      }
    },

    addMessage(message) {
      this.messages.push(message);
    },

    getMessageColor(message) {
      if (message.user.username === this.sharedState.user.username) {
        return 'blue';
      }
    },

    scrollToBottom() {
      setTimeout(() => { //wait for next tick
        let container = this.$refs.messages;
        if (container) {
          container.scrollTop = container.scrollHeight;
        }
      });
    },

    removeChat(chat) {
      (new api()).removeChat(chat.id).then(() => {
        this.chats = this.chats.filter((item) => {
          return item.id !== chat.id;
        });
      });
    },

    setChat(chat) {
      this.selectedItem = chat.id;
    },
  },

  watch: {
    selectedItem() {
      if (this.$route.params.id !== this.selectedItem) {
        this.$router.push({
          name: 'chat',
          params: {'id': this.selectedItem},
        });
      }
      this.getMessages();
    },
  },
}
</script>