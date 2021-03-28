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
              <v-list-item-action class="show-on-hover">
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
        <v-row v-for="message in messages" :key="message.id" class="pa-1 relative">
          <v-spacer v-if="message.isOwn"></v-spacer>
            <v-chip :color="getMessageColor(message)"
                    @click="showContextMenu($event, message)">
              {{ message.createdAt | formatDate }}
              <router-link :to="{name: 'user', params: {id: message.user.id}}" class="ml-1">
                {{ message.user.username }}
              </router-link>:
              {{ message.body }}
            </v-chip>
        </v-row>
      </v-col>
    </v-row>
    <v-menu
        v-model="contextMenu.display"
        bottom
        offset-y
        offset-x
        :position-x="contextMenu.x"
        :position-y="contextMenu.y"
        :max-width="150">
      <v-list>
        <v-list-item @click="startEdit">
          <v-list-item-title>Edit</v-list-item-title>
        </v-list-item>
        <v-list-item @click="confirmDelete">
          <v-list-item-title>Delete</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <v-dialog
        v-model="showDeleteDialog"
        max-width="550"
    >
      <v-card>
        <v-card-title class="headline grey lighten-2">
          Are you sure you want to delete message?
        </v-card-title>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
              color="primary"
              text
              @click="sendDeleteMessage">
            Delete
          </v-btn>

          <v-btn
              color="primary"
              text
              @click="cancelDelete">
            Cancel
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

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
            <v-col :cols="editMode ? 7 : 10">
              <v-text-field label="Message" v-model="messageText"/>
            </v-col>
            <v-col :cols="editMode ? 5 : 2" class="pt-5 text-right">
              <v-btn block color="indigo" dark @click="sendMessage" class="hidden-sm-and-down" v-if="!editMode">
                <v-icon left>
                  mdi-send
                </v-icon>
                <span>Send</span>
              </v-btn>

              <v-btn color="indigo" dark @click="sendUpdateMessage" class="hidden-sm-and-down mr-2" v-if="editMode">
                <v-icon left>
                  mdi-check
                </v-icon>
                <span>Save</span>
              </v-btn>

              <v-btn color="indigo" dark @click="cancelUpdateMessage" class="hidden-sm-and-down" v-if="editMode">
                <v-icon left>
                  mdi-close
                </v-icon>
                <span>Cancel</span>
              </v-btn>

              <v-btn color="indigo"
                     fab
                     dark
                     small
                     @click="sendMessage"
                     class="hidden-md-and-up text-center"
                     v-if="!editMode">
                <v-icon>
                  mdi-send
                </v-icon>
              </v-btn>

              <v-btn color="indigo"
                     fab
                     dark
                     small
                     left
                     @click="sendUpdateMessage"
                     class="hidden-md-and-up text-center"
                     v-if="editMode">
                <v-icon>
                  mdi-check
                </v-icon>
              </v-btn>

              <v-btn color="indigo"
                     fab
                     dark
                     small
                     right
                     @click="cancelUpdateMessage"
                     class="hidden-md-and-up text-center"
                     v-if="editMode">
                <v-icon>
                  mdi-close
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
    contextMenu: {
      display: false,
      x: 0,
      y: 0,
    },
    contextMenuMessage: null,
    editMode: false,
    showDeleteDialog: false,
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

    store.eventBus.$on('message_created', (data) => {
      this.addMessage(data.message);
    });

    store.eventBus.$on('chat_created', (data) => {
      this.addChat(data.chat);
    });

    store.eventBus.$on('message_updated', (data) => {
      this.updateMessage(data.message);
    });

    store.eventBus.$on('message_deleted', (data) => {
      this.deleteMessage(data.message);
    });

    store.eventBus.$on('chat_updated', (data) => {
      this.updateChat(data.chat);
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
          this.messages = response.data.messages.map((message) => {
            message.isOwn = this.isOwnMessage(message);

            return message;
          });
          this.scrollToBottom();
          store.hideLeftMenu();
        });
      } else {
        this.messages = [];
      }
    },

    showContextMenu(event, message) {
      event.preventDefault()
      this.contextMenu.display = false;
      this.contextMenu.x = event.clientX;
      this.contextMenu.y = event.clientY;
      this.contextMenuMessage = message;
      this.$nextTick(() => {
        this.contextMenu.display = true;
      })
    },

    startEdit() {
      this.messageText = this.contextMenuMessage.body + '';
      this.editMode = true;
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
          if (this.sharedState.user) {
            return user.username !== this.sharedState.user.username;
          } else {
            return true;
          }
        }).map((user) => {
          return user.username;
        }).join(', ');
      }
    },

    addMessage(message) {
      if (message.chat.id !== this.selectedItem) {
        return;
      }
      message.isOwn = this.isOwnMessage(message);

      this.messages.push(message);
    },

    isOwnMessage(message) {
      return this.sharedState.user && this.sharedState.user.username === message.user.username
    },

    addChat(chat) {
      this.chats.push(chat);
    },

    sendUpdateMessage() {
      (new api()).updateMessage(this.selectedItem, this.contextMenuMessage.id, this.messageText).then(() => {
        this.messageText = '';
        this.editMode = false;
        this.contextMenuMessage = null;
      });
    },

    cancelUpdateMessage() {
      this.messageText = '';
      this.editMode = false;
      this.contextMenuMessage = null;
    },

    updateMessage(message) {
      if (message.chat.id !== this.selectedItem) {
        return;
      }

      for (const curMessage of this.messages) {
        if (curMessage.id === message.id) {
          Object.assign(curMessage, message);
          return;
        }
      }

      this.addMessage(message);
    },

    updateChat(chat) {
      for (const curChat of this.chats) {
        if (curChat.id === chat.id) {
          Object.assign(curChat, chat);

          return;
        }
      }

      this.addChat(chat);
    },

    confirmDelete() {
      this.showDeleteDialog = true;
    },

    cancelDelete() {
      this.showDeleteDialog = false;
    },

    sendDeleteMessage() {
      (new api()).removeMessage(this.contextMenuMessage.chat.id, this.contextMenuMessage.id).then(() => {
        this.showDeleteDialog = false;
      });
    },

    deleteMessage(message) {
      this.messages = this.messages.filter((curMessage) => {
        return curMessage.id !== message.id;
      });
    },

    getMessageColor(message) {
      if (this.sharedState.user && message.user.username === this.sharedState.user.username) {
        return 'indigo';
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