<template>
  <v-snackbar v-model="show" bottom left :timeout="-1">
    <i>Mise à jour...</i>
    <v-progress-linear color="orange" indeterminate></v-progress-linear>
  </v-snackbar>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Notifications } from "../logic/notifications";

const SpinnerSnackbarProps = Vue.extend({
  props: {
    N: Object as () => Notifications
  }
});

@Component
export default class SpinnerSnackbar extends SpinnerSnackbarProps {
  get show() {
    return this.N.getSpin();
  }

  set show(b: boolean) {}
}
</script>
