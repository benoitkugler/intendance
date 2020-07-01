<template>
  <v-snackbar v-model="show" bottom right :timeout="-1">
    <i>Mise à jour...</i>
    <v-progress-linear color="orange" indeterminate></v-progress-linear>
  </v-snackbar>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { C } from "../logic/controller";

@Component
export default class SpinnerSnackbar extends Vue {
  get show() {
    return C.notifications.getSpin();
  }

  set show(b: boolean) {}
}
</script>
