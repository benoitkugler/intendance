<template>
  <v-snackbar v-model="show" bottom right :timeout="6000" color="success">
    <i>({{ currentTime() }}) </i>
    &nbsp;
    <span v-html="message"></span>
    <v-btn color="black" text icon @click="show = false">
      <v-icon>mdi-close</v-icon>
    </v-btn>
  </v-snackbar>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { NS } from "../logic/notifications";

@Component
export default class SuccessSnackbar extends Vue {
  get show() {
    return NS.getMessage() != null;
  }

  set show(b: boolean) {
    if (!b) {
      NS.setMessage(null);
    }
  }

  get message() {
    return NS.getMessage();
  }

  currentTime() {
    const datetime = new Date();
    const min = ("00" + datetime.getMinutes()).substr(-2, 2);
    return `${datetime.getHours()}:${min}`;
  }
}
</script>
