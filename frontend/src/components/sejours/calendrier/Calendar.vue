<template>
  <div class="two-weeks-calendar" ref="weeks">
    <v-dialog v-model="showFormCalcul" max-width="600px">
      <form-calcul :initial-sejour="sejour"></form-calcul>
    </v-dialog>

    <week
      :sejour="sejour"
      :weekdays="weekdays"
      :start="startWeek1"
      :dayHeight="dayHeight"
      :events="events"
      :mode="mode"
      :currentDay="currentDay"
      :hoverDay="hoverDay"
      @editRepas="r => $emit('editRepas', r)"
      @addRepas="r => $emit('addRepas', r)"
      @change="o => $emit('change', o)"
      @hover="d => (hoverDay = d)"
    ></week>
    <week
      :sejour="sejour"
      :weekdays="weekdays"
      :start="startWeek2"
      :dayHeight="dayHeight"
      :events="events"
      :mode="mode"
      :currentDay="currentDay"
      :hoverDay="hoverDay"
      @editRepas="r => $emit('editRepas', r)"
      @addRepas="r => $emit('addRepas', r)"
      @change="o => $emit('change', o)"
      @hover="d => (hoverDay = d)"
    ></week>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import ListeRepas from "./ListeRepas.vue";
import FormCalcul from "../FormCalcul.vue";
import TooltipBtn from "../../utils/TooltipBtn.vue";
import Week from "./Week.vue";

import { Sejour, SejourRepas, RepasWithGroupe } from "../../../logic/types";
import { C } from "../../../logic/controller";
import {
  DetailsSejour,
  New,
  DetailsRepas,
  PreferencesAgenda,
  CalendarMode,
  NullId
} from "../../../logic/types2";
import { fmtHoraire, Horaires } from "../../../logic/enums";
import { Formatter } from "../../../logic/formatter";
import { toDateVuetify } from "./utils";
import { HorairesColors } from "../../utils/utils";

const _days = [0, 1, 2, 3, 4, 5, 6];

type DragKind = "journee" | "repas";

// renvoie l'ordre des jours pour que `start` soit
// affiché en premier
function weekdaysFromStart(start: Date) {
  const d0 = start.getDay();
  return _days.map(d => (d0 + d) % 7);
}

const Props = Vue.extend({
  props: {
    sejour: Object as () => SejourRepas | null,
    preferences: Object as () => PreferencesAgenda,
    mode: String as () => CalendarMode,
    activeJourOffset: Number as () => number | null
  }
});

@Component({
  components: {
    TooltipBtn,
    FormCalcul,
    ListeRepas,
    Week
  }
})
export default class Calendar extends Props {
  hoverDay = "";

  showFormCalcul = false;

  private dayHeight = "35vh";

  dayTitle(date: string) {
    return new Date(date).toLocaleDateString("fr-FR", {
      weekday: "short",
      day: "numeric"
    });
  }

  get currentDay(): string | null {
    if (this.sejour == null || this.activeJourOffset == null) return null;
    return toDateVuetify(C.offsetToDate(this.sejour.id, this.activeJourOffset));
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
    const out = this.startDate;
    out.setDate(this.startDate.getDate() + 7);
    return toDateVuetify(out);
  }

  get events(): { [key: string]: RepasWithGroupe[] } {
    const sejour = this.sejour;
    if (sejour == null) return {};
    let out: { [key: string]: RepasWithGroupe[] } = {};
    C.iterateAllRepas((_, repas) => {
      if (repas.id_sejour != sejour.id) return;
      const d = toDateVuetify(C.offsetToDate(sejour.id, repas.jour_offset));
      const l = out[d] || [];
      l.push(repas);
      out[d] = l;
    });
    const horaires = Horaires.map(v => v.value);
    for (const date in out) {
      const element = out[date];
      out[date] = element.sort((a, b) => {
        const v = horaires.indexOf(a.horaire) - horaires.indexOf(b.horaire);
        return v == 0 ? a.id - b.id : v; // tri déterministe
      });
    }
    return out;
  }
}
</script>

<style></style>
