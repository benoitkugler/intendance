<template>
  <div class="two-weeks-calendar" ref="weeks">
    <week
      :C="C"
      :sejour="sejour"
      :weekdays="weekdays"
      :start="startWeek1"
      :dayHeight="dayHeight"
      :events="events"
      :currentDay="currentDay"
      :hoverDay="hoverDay"
      @edit-repas="(r) => $emit('edit-repas', r)"
      @add-repas="(r) => $emit('add-repas', r)"
      @change="(o) => $emit('change', o)"
      @hover="(d) => (hoverDay = d)"
    ></week>
    <week
      :C="C"
      :sejour="sejour"
      :weekdays="weekdays"
      :start="startWeek2"
      :dayHeight="dayHeight"
      :events="events"
      :currentDay="currentDay"
      :hoverDay="hoverDay"
      @edit-repas="(r) => $emit('edit-repas', r)"
      @add-repas="(r) => $emit('add-repas', r)"
      @change="(o) => $emit('change', o)"
      @hover="(d) => (hoverDay = d)"
    ></week>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import ListeRepas from "./ListeRepas.vue";
import TooltipBtn from "../../utils/TooltipBtn.vue";
import Week from "./Week.vue";

import { Sejour, SejourRepas, RepasComplet, New } from "@/logic/api";
import { Controller } from "@/logic/controller";
import { DetailsSejour, DetailsRepas, PreferencesAgenda } from "@/logic/types";
import { Formatter } from "@/logic/formatter";
import { toDateVuetify } from "./utils";
import { HorairesColors } from "../../utils/utils";
import { Watch } from "vue-property-decorator";

const _days = [0, 1, 2, 3, 4, 5, 6];

type DragKind = "journee" | "repas";

/** renvoie l'ordre des jours pour que `start` soit
affiché en premier */
function weekdaysFromStart(start: Date) {
  const d0 = start.getDay();
  return _days.map((d) => (d0 + d) % 7);
}

const Props = Vue.extend({
  props: {
    C: Object as () => Controller,
    sejour: Object as () => SejourRepas | null,
    preferences: Object as () => PreferencesAgenda,
    activeJourOffset: Number as () => number | null,
  },
});

@Component({
  components: {
    TooltipBtn,
    ListeRepas,
    Week,
  },
})
export default class Calendar extends Props {
  hoverDay = "";

  dayHeight = "35vh";

  dayTitle(date: string) {
    return new Date(date).toLocaleDateString("fr-FR", {
      weekday: "short",
      day: "numeric",
    });
  }

  get currentDay(): string | null {
    if (this.sejour == null || this.activeJourOffset == null) return null;
    return toDateVuetify(
      this.C.offsetToDate(this.sejour.id, this.activeJourOffset)
    );
  }

  get startDate(): Date {
    if (this.sejour == null) return new Date();
    return new Date(this.sejour.date_debut);
  }

  get weekdays() {
    if (this.preferences.startPremierJour) {
      return weekdaysFromStart(this.startDate);
    } else {
      return [1, 2, 3, 4, 5, 6, 0];
    }
  }

  get startWeek1() {
    return toDateVuetify(this.startDate);
  }

  get startWeek2() {
    const out = new Date(this.startDate);
    out.setDate(this.startDate.getDate() + 7);
    return toDateVuetify(out);
  }

  get events(): { [key: string]: RepasComplet[] } {
    const sejour = this.sejour;
    if (sejour == null) return {};
    let out: { [key: string]: RepasComplet[] } = {};
    (sejour.repass || []).forEach((repas) => {
      const d = toDateVuetify(
        this.C.offsetToDate(sejour.id, repas.jour_offset)
      );
      const l = out[d] || [];
      l.push(repas);
      out[d] = l;
    });
    for (const date in out) {
      const element = out[date];
      out[date] = element.sort((a, b) => {
        const v = a.horaire - b.horaire;
        return v == 0 ? a.id - b.id : v; // tri déterministe
      });
    }
    return out;
  }
}
</script>

<style></style>
