<template>
  <div class="two-weeks-calendar" ref="weeks">
    <v-calendar
      type="week"
      locale="fr"
      :first-interval="firstInterval"
      :interval-count="intervalCount"
      :interval-minutes="intervalMinutes"
      :interval-height="intervalHeight"
      :start="startWeek1"
      :weekdays="weekdays"
      :events="[
        { name: 'jejejej', start: '2019-12-20T09:00', id_repas: 78 },
        { name: 'jejejej', start: '2019-12-20', id_repas: 75 }
      ]"
      @click:interval="log"
      @click:time="log"
    >
      <template v-slot:event="{ event }">
        <div :data-id-repas="event.id_repas">{{ event.name }}</div>
      </template>
      <template v-slot:day-header="{ date }">
        <div :data-day="date"></div>
      </template>
      <template v-slot:interval
        ><div @dragover="onDragover" class="dragover"></div
      ></template>
    </v-calendar>
    <v-calendar
      type="week"
      locale="fr"
      :first-interval="firstInterval"
      :interval-count="intervalCount"
      :interval-minutes="intervalMinutes"
      :interval-height="intervalHeight"
      :start="startWeek2"
      :weekdays="weekdays"
    ></v-calendar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Repas, Sejour } from "../logic/types";
import { D } from "../logic/controller";

const _days = [0, 1, 2, 3, 4, 5, 6];

// renvoie l'ordre des jours pour que `start` soit
// affiché en premier
function weekdaysFromStart(start: Date) {
  const d0 = start.getDay();
  return _days.map(d => (d0 + d) % 7);
}

function formatDate(d: Date) {
  return d.toISOString().substr(0, 10);
}

function getEventStart(r: Repas, sejour: Sejour) {
  const dateDebut = new Date(sejour.date_debut);
  dateDebut.setDate(dateDebut.getDate() + r.jour_offset);
  return (
    formatDate(dateDebut) +
    `T${"00" + r.horaire.heure}:${"00" + r.horaire.minute}`
  );
}

function onDragStart(
  event: DragEvent,
  kind: "repas" | "journee",
  arg: string | undefined
) {
  if (!event.dataTransfer || !arg) return;
  event.dataTransfer.setData("text/plain", arg);
  event.dataTransfer.effectAllowed = "move";
}

const Props = Vue.extend({
  props: {
    // Date as string
    start: String,
    forceFirstDay: Boolean
  }
});

@Component
export default class Calendar extends Props {
  firstInterval = 2;
  intervalCount = 8;
  intervalMinutes = 120;
  intervalHeight = 25;

  get startDate(): Date {
    return new Date(this.start);
  }

  log(args: any) {
    console.log(args);
  }

  get weekdays() {
    if (this.forceFirstDay) {
      return weekdaysFromStart(this.startDate);
    }
    return _days;
  }

  get startWeek1() {
    return formatDate(this.startDate);
  }

  get startWeek2() {
    const out = this.startDate;
    out.setDate(this.startDate.getDate() + 7);
    return formatDate(out);
  }

  private setupDrag() {
    const htmlEl = this.$refs.weeks as HTMLDivElement;
    htmlEl.querySelectorAll<HTMLElement>("[data-day]").forEach(item => {
      if (!item.parentElement) return;
      item.parentElement.draggable = true;
      item.parentElement.ondragstart = e =>
        onDragStart(e, "journee", item.dataset["day"]);
    });
    htmlEl.querySelectorAll<HTMLElement>("[data-id-repas]").forEach(item => {
      item.draggable = true;
      item.ondragstart = e => onDragStart(e, "repas", item.dataset["id-repas"]);
    });
  }

  mounted() {
    this.setupDrag();
  }

  updated() {
    this.setupDrag();
  }

  onDragover(event: DragEvent) {
    if (!event.dataTransfer) return;
    const isRepas = event.dataTransfer.types.includes("repas");
    const isDay = event.dataTransfer.types.includes("journee");
    console.log(event);
    if (isRepas || isDay) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "move";
    }
  }
}
</script>

<style>
.two-weeks-calendar .v-calendar-daily__day-interval {
  cursor: pointer;
}
.two-weeks-calendar .v-calendar-daily__day-interval:hover {
  background-color: rgba(34, 182, 187, 0.267);
}
.dragover {
  height: 100%;
}
</style>
