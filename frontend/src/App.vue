<template>
  <v-app>
    <v-app-bar app color="primary" dark> </v-app-bar>
    <v-content>
      <v-container>
        <calendar :start="start" :forceFirstDay="true" />
        <spinner-snackbar></spinner-snackbar>
        <error-dialog></error-dialog>
        <success-snackbar></success-snackbar>
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import Calendar from "./components/sejours/Calendar.vue";
import ErrorDialog from "./components/ErrorDialog.vue";
import SuccessSnackbar from "./components/SuccessSnackbar.vue";
import SpinnerSnackbar from "./components/SpinnerSnackbar.vue";
import { D } from "./logic/controller";
import { NS } from "./logic/notifications";

export default Vue.extend({
  name: "App",

  components: {
    Calendar,
    ErrorDialog,
    SuccessSnackbar,
    SpinnerSnackbar
  },

  data: () => ({
    start: "2019-12-19",
    storage: D,
    notifications: NS
  })
});
</script>
