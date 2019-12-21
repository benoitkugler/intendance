<template>
  <v-dialog v-model="show" max-width="500">
    <v-card>
      <v-card-title class="title error" v-html="title"></v-card-title>
      <v-card-text class="pa-3">
        <div v-html="message" class="subtitle-2"></div>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { NS } from "../logic/notifications";

@Component
export default class ErrorDialog extends Vue {
  get show() {
    return NS.getError() != null;
  }

  set show(b: boolean) {
    if (!b) {
      NS.setError(null);
    }
  }

  get title() {
    const err = NS.getError();
    if (err == null) return "";
    let title = err.kind;
    if (err.code != null) {
      title += ` <i>(code ${err.code})</i>`;
    }
    return title;
  }

  get message() {
    const err = NS.getError() || { messageHtml: "" };
    return err.messageHtml;
  }
}
</script>
