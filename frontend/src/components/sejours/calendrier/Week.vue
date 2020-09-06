<template>
  <v-calendar
    type="custom-weekly"
    hide-header
    :start="start"
    :weekdays="weekdays"
  >
    <template v-slot:day-label="props">
      <v-row no-gutters @mouseover="$emit('hover', props.date)" class="mb-1">
        <v-col>
          <v-btn
            text
            small
            @click="$emit('change', props.date)"
            :input-value="currentDay === props.date"
            class="mx-0 px-1"
            @dragover="onDayDragover($event, props.date)"
            @drop="onDayDrop($event, props.date)"
            >{{ dayTitle(props.date) }}
          </v-btn></v-col
        >
        <v-col cols="4">
          <tooltip-btn
            v-if="hoverDay === props.date"
            small
            mdi-icon="plus"
            color="green"
            tooltip="Ajouter un repas..."
            @click="$emit('add-repas', props.date)"
          ></tooltip-btn>
        </v-col>
      </v-row>
    </template>
    <template v-slot:day="props">
      <div
        @dragover="onDayDragover($event, props.date)"
        @drop="onDayDrop($event, props.date)"
        class="overflow-y-auto"
        :style="{ height: dayHeight }"
        @mouseover="$emit('hover', props.date)"
      >
        <liste-repas
          :C="C"
          :repass="events[props.date]"
          @edit="args => $emit('edit-repas', args)"
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
import { Controller } from "@/logic/controller";
import { RepasComplet, Sejour } from "@/logic/api";
import { DragKind, getDragData } from "../../utils/utils_drag";
import { Formatter } from "@/logic/formatter";

const Props = Vue.extend({
  props: {
    C: Object as () => Controller,
    sejour: Object as () => Sejour,
    weekdays: Array,
    start: String,
    dayHeight: String,
    events: Object as () => { [key: string]: RepasComplet[] },
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
  dayTitle = Formatter.formatDate;

  onDayDragover(event: DragEvent, date: string) {
    if (!event.dataTransfer || this.sejour == null) return;
    const debut = new Date(this.sejour.date_debut).valueOf();
    const target = new Date(date).valueOf();
    if (target < debut) return; // empêche un offset négatif
    const isRepas = event.dataTransfer.types.includes(DragKind.Repas);
    if (isRepas) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "move";
    }
  }

  async onDayDrop(event: DragEvent, date: string) {
    if (!event.dataTransfer || this.C.state.idSejour == null) return;
    event.preventDefault();
    const repas = getDragData(event.dataTransfer, DragKind.Repas);

    const targetOffset = this.C.dateToOffset(
      this.C.state.idSejour,
      new Date(date)
    );
    if (targetOffset == repas.jour_offset) return;
    repas.jour_offset = targetOffset;
    await this.C.api.UpdateManyRepas([repas]);
  }
}
</script>

<style></style>
