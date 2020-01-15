<template>
  <div>
    <v-calendar
      v-if="day != null"
      type="day"
      locale="fr"
      :first-interval="4"
      :interval-count="8"
      :interval-minutes="120"
      :interval-height="55"
      :event-height="30"
      :value="dayString"
      :events="events"
      event-color="#6a7d6f"
    >
      <template v-slot:event="{ event }">
        <v-row no-gutters class="fill-height">
          <v-col class="px-1 align-self-center overflow-x-auto">
            <v-chip
              label
              v-for="groupe in getGroupes(event)"
              :key="groupe.id"
              class="mr-1 px-1 align-self-center"
              :color="groupe.couleur"
              small
              :style="{ borderWidth: ' 1.5px' }"
              outlined
              draggable
            >
              {{ groupe.nom }}</v-chip
            >
          </v-col>
        </v-row>
      </template>
    </v-calendar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import {
  toDateVuetify,
  DataEvent,
  getEventStart,
  repartitRepas
} from "./utils";
import { C } from "../../../logic/controller";
import { RepasWithGroupe } from "../../../logic/types";

const DayProps = Vue.extend({
  props: {
    day: Date as new () => Date
  }
});
@Component({})
export default class Day extends DayProps {
  get dayString() {
    if (this.day == null) return "";
    return toDateVuetify(this.day);
  }

  get events(): DataEvent[] {
    const repass: RepasWithGroupe[] = [];
    C.iterateAllRepas((sejour, repas) => {
      if (sejour.id != C.state.idSejour) return;
      const date = C.offsetToDate(sejour.id, repas.jour_offset);
      if (date.getTime() == this.day.getTime()) {
        repass.push(repas);
      }
    });
    const ajusted = repartitRepas(repass, { heure: 8, minute: 0 }, 15 * 60, 63);
    console.log(ajusted);

    return ajusted.map(rep => {
      return { repas: rep, start: getEventStart(rep) };
    });
  }

  getGroupes(event: DataEvent) {
    const gr = C.getRepasGroupes(event.repas);
    const maxChar = 5;
    gr.forEach(groupe => (groupe.nom = groupe.nom.substr(0, maxChar) + "."));
    return gr;
  }
}
</script>

<style scoped></style>
