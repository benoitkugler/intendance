<template>
  <v-snackbar v-model="show" bottom right :timeout="0">
    <i>Mise à jour...</i>
    <v-progress-linear color="orange" indeterminate></v-progress-linear>
  </v-snackbar>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { NS } from "../logic/notifications";

@Component
export default class SpinnerSnackbar extends Vue {
  get show() {
    return NS.getSpin();
  }

  set show(b: boolean) {}
}
</script>
