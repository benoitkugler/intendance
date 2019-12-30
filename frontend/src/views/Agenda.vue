<template>
  <v-container fluid>
    <calendar ref="calendar" />
  </v-container>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import Calendar from "../components/sejours/Calendar.vue";
import { D } from "../logic/controller";
import { NS } from "../logic/notifications";

@Component({
  components: { Calendar }
})
export default class Agenda extends Vue {
  $refs!: {
    calendar: Calendar;
  };

  async mounted() {
    await D.loadAllMenus();
    await D.loadAgenda();
    if (NS.getError() == null) {
      NS.setMessage("L'agenda a bien été chargé.");
    }
    this.$refs.calendar.setClosestSejour();
  }
}
</script>
