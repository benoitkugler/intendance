<template>
  <div>
    <v-snackbar v-model="show" :timeout="-1" bottom right color="success">
      <v-row no-gutters>
        <v-col cols="10" align-self="center" class="pr-2">
          <span v-html="message"></span>
        </v-col>
        <v-col cols="2" align-self="center">
          <v-btn color="black" text icon @click="show = false" class="mx-0">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </v-snackbar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Notifications } from "../logic/notifications";

const SuccessSnackbarProps = Vue.extend({
  props: {
    N: Object as () => Notifications
  }
});

@Component
export default class SuccessSnackbar extends SuccessSnackbarProps {
  get show() {
    return this.N.messages.length != 0;
  }

  set show(b: boolean) {
    if (!b) {
      this.N.clearMessages();
    }
  }

  get message() {
    return this.N.messages
      .map(m => this.formatTime(m.time) + m.text)
      .join("<br/>");
  }

  formatTime(datetime: Date) {
    const min = ("00" + datetime.getMinutes()).substr(-2, 2);
    const ti = datetime.toLocaleTimeString(undefined, {
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit"
    });
    return `<small><i>[${ti}]</i></small>  `;
  }
}
</script>
