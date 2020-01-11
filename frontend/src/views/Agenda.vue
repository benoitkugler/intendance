<template>
  <v-container fluid>
    <calendar ref="calendar" />
  </v-container>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import Calendar from "../components/sejours/Calendar.vue";
import { C } from "../logic/controller";

@Component({
  components: { Calendar }
})
export default class Agenda extends Vue {
  $refs!: {
    calendar: Calendar;
  };

  async mounted() {
    await C.data.loadAllMenus();
    await C.data.loadSejours();
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("L'agenda a bien été chargé.");
    }
  }
}
</script>
