<template>
  <v-snackbar v-model="show" bottom right :timeout="4000" color="success">
    <v-row no-gutters>
      <v-col cols="2" align-self="center">
        <i>({{ currentTime() }})</i>
      </v-col>
      <v-col cols="9" align-self="center">
        <span v-html="message"></span>
      </v-col>
      <v-col cols="1">
        <v-btn color="black" text icon @click="show = false" class="mx-0">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-col>
    </v-row>
  </v-snackbar>
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
    return this.N.getMessage() != null;
  }

  set show(b: boolean) {
    if (!b) {
      this.N.setMessage(null);
    }
  }

  get message() {
    return this.N.getMessage();
  }

  currentTime() {
    const datetime = new Date();
    const min = ("00" + datetime.getMinutes()).substr(-2, 2);
    return `${datetime.getHours()}:${min}`;
  }
}
</script>
