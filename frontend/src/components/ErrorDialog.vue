<template>
  <v-dialog v-model="show" max-width="500">
    <v-card>
      <v-card-title class="title error">
        <v-row>
          <v-col cols="7">{{ mainTitle }}</v-col>
          <v-col cols="5" class="text-right"
            ><small>{{ subTitle }}</small></v-col
          >
        </v-row>
      </v-card-title>
      <v-card-text class="pa-3">
        <div v-html="message" class="subtitle-2"></div>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Notifications } from "../logic/notifications";

const ErrorDialogProps = Vue.extend({
  props: {
    N: Object as () => Notifications,
  },
});

@Component
export default class ErrorDialog extends ErrorDialogProps {
  get show() {
    return this.N.getError() != null;
  }

  set show(b: boolean) {
    if (!b) {
      this.N.setError(null);
    }
  }

  get mainTitle() {
    const err = this.N.getError();
    if (err == null) return "";
    return err.kind;
  }

  get subTitle() {
    const err = this.N.getError();
    if (err && err.code != null) {
      return `code ${err.code}`;
    }
    return "";
  }

  get message() {
    const err = this.N.getError() || { messageHtml: "" };
    return err.messageHtml;
  }
}
</script>
