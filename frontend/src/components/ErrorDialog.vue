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
import { C } from "../logic/controller";

@Component
export default class ErrorDialog extends Vue {
  get show() {
    return C.notifications.getError() != null;
  }

  set show(b: boolean) {
    if (!b) {
      C.notifications.setError(null);
    }
  }

  get mainTitle() {
    const err = C.notifications.getError();
    if (err == null) return "";
    return err.kind;
  }

  get subTitle() {
    const err = C.notifications.getError();
    if (err && err.code != null) {
      return `code ${err.code}`;
    }
    return "";
  }

  get message() {
    const err = C.notifications.getError() || { messageHtml: "" };
    return err.messageHtml;
  }
}
</script>
