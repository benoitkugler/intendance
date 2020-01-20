<template>
  <v-calendar
    type="custom-weekly"
    hide-header
    :start="start"
    :weekdays="weekdays"
  >
    <template v-slot:day-label="{ date }">
      <v-row no-gutters @mouseover="$emit('hover', date)" class="mb-1">
        <v-col>
          <v-btn
            text
            small
            @click="$emit('change', date)"
            :input-value="currentDay === date"
            class="mx-0 px-1"
            @dragover="onDayDragover($event, date)"
            @drop="onDayDrop($event, date)"
            >{{ dayTitle(date) }}
          </v-btn></v-col
        >
        <v-col cols="4">
          <tooltip-btn
            v-if="hoverDay === date"
            small
            mdi-icon="plus"
            color="green"
            tooltip="Ajouter un repas..."
            @click="$emit('addRepas', date)"
          ></tooltip-btn>
        </v-col>
      </v-row>
    </template>
    <template v-slot:day="{ date }">
      <div
        @dragover="onDayDragover($event, date)"
        @drop="onDayDrop($event, date)"
        class="overflow-y-auto"
        :style="{ height: dayHeight }"
        @mouseover="$emit('hover', date)"
      >
        <liste-repas
          :repass="events[date]"
          :mode="mode"
          @edit="args => $emit('editRepas', args)"
        ></liste-repas>
      </div>
    </template>
  </v-calendar>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import ListeRepas from "./ListeRepas.vue";
import TooltipBtn from "../../utils/TooltipBtn.vue";
import { C } from "../../../logic/controller";
import { CalendarMode } from "../../../logic/types2";
import { RepasWithGroupe, Sejour } from "../../../logic/types";

const Props = Vue.extend({
  props: {
    sejour: Object as () => Sejour,
    weekdays: Array,
    start: String,
    dayHeight: String,
    events: Object as () => { [key: string]: RepasWithGroupe[] },
    mode: String as () => CalendarMode,
    currentDay: String as () => string | null,
    hoverDay: String
  }
});

@Component({
  components: {
    ListeRepas,
    TooltipBtn
  }
})
export default class Week extends Props {
  dayTitle(date: string) {
    return new Date(date).toLocaleDateString("fr-FR", {
      weekday: "short",
      day: "numeric"
    });
  }

  onDayDragover(event: DragEvent, date: string) {
    if (!event.dataTransfer || this.sejour == null) return;
    const debut = new Date(this.sejour.date_debut).valueOf();
    const target = new Date(date).valueOf();
    if (target < debut) return; // empêche un offset négatif
    const isRepas = event.dataTransfer.types.includes("repas");
    // const isDay = event.dataTransfer.types.includes("journee");
    if (isRepas) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "move";
    }
  }

  async onDayDrop(event: DragEvent, date: string) {
    if (!event.dataTransfer || C.state.idSejour == null) return;
    event.preventDefault();
    const repas: RepasWithGroupe = JSON.parse(
      event.dataTransfer.getData("repas")
    );
    const targetOffset = C.dateToOffset(C.state.idSejour, new Date(date));
    if (targetOffset == repas.jour_offset) return;
    repas.jour_offset = targetOffset;
    await C.data.updateManyRepas([repas]);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Repas déplacé avec succès.");
    }
  }
}
</script>

<style></style>
