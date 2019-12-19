<template>
  <v-app>
    <v-app-bar app color="primary" dark> </v-app-bar>
    <v-content>
      <calendar :start="start" :forceFirstDay="true" />
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import Calendar from "./components/Calendar.vue";
import { D } from "./logic/controller";

export default Vue.extend({
  name: "App",

  components: {
    Calendar
  },

  data: () => ({
    storage: D,
    start: new Date().toISOString()
  })
});
</script>
