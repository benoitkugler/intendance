<template>
  <div class="two-weeks-calendar" ref="weeks">
    <v-dialog v-model="showFormCalcul" max-width="600px">
      <form-calcul :initial-sejour="sejour"></form-calcul>
    </v-dialog>

    <v-calendar
      type="week"
      locale="fr"
      :event-height="22"
      :short-weekdays="false"
      :interval-count="1"
      :interval-height="dayHeight"
      :interval-width="0"
      :start="startWeek1"
      :weekdays="weekdays"
      @mousedown:time="registerTime"
      @click:time="args => $emit('addRepas', args.date)"
      @click:date="args => $emit('change', args)"
    >
      <template v-slot:interval="{ date }">
        <div
          @dragover="e => onDragover(e, 'repas')"
          @drop="e => onDrop(e, 'repas')"
          class="dragover overflow-y-auto"
        >
          <liste-repas
            :repass="events[date]"
            :mode="mode"
            @edit="args => $emit('editRepas', args)"
          ></liste-repas>
        </div>
      </template>
      <template v-slot:day-header="{ date }">
        <div :data-day="date"></div>
      </template>
    </v-calendar>
    <v-calendar
      class="mt-1"
      type="week"
      locale="fr"
      :interval-count="1"
      :interval-height="dayHeight"
      :interval-width="0"
      :short-weekdays="false"
      :start="startWeek2"
      :weekdays="weekdays"
      @mousedown:time="registerTime"
      @click:time="args => $emit('addRepas', args.date)"
      @click:date="args => $emit('change', args)"
    >
      <template v-slot:interval="{ date }">
        <div
          @dragover="e => onDragover(e, 'repas')"
          @drop="e => onDrop(e, 'repas')"
          class="dragover y-overflow-auto"
        >
          <liste-repas
            :repass="events[date]"
            :mode="mode"
            @edit="args => $emit('editRepas', args)"
          ></liste-repas>
        </div>
      </template>
    </v-calendar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Sejour, SejourRepas, RepasWithGroupe } from "../../../logic/types";
import ListeRepas from "./ListeRepas.vue";
import FormCalcul from "../FormCalcul.vue";
import { C } from "../../../logic/controller";
import TooltipBtn from "../../utils/TooltipBtn.vue";
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
import { toDateVuetify, DateTime, DataEvent } from "./utils";
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
    mode: String as () => CalendarMode
  }
});

@Component({
  components: {
    TooltipBtn,
    FormCalcul,
    ListeRepas
  }
})
export default class Calendar extends Props {
  private lastClickedTime: DateTime | null = null;

  showFormCalcul = false;

  private dayHeight = 250;

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
        return horaires.indexOf(a.horaire) - horaires.indexOf(b.horaire);
      });
    }
    return out;
  }

  mounted() {
    // this.setupDrag();
  }

  updated() {
    // this.setupDrag();
  }

  // on stocke le moment correspondant au dernier click,
  // pour contourner le manque d'une méthode getTime(pos)
  registerTime(time: DateTime) {
    this.lastClickedTime = time;
  }

  private setupDrag() {
    const sejour = this.sejour;
    if (sejour == null) return; // désactivé si aucun séjour n'est sélectionné.
    const htmlEl = this.$refs.weeks as HTMLDivElement;
    htmlEl.querySelectorAll<HTMLElement>("[data-day]").forEach(item => {
      if (!item.parentElement) return;
      item.parentElement.draggable = true;
      item.parentElement.ondragstart = e =>
        this.onDragStart(e, "journee", item.dataset.day);
      item.parentElement.ondragover = e => this.onDragover(e, "journee");
      item.parentElement.ondrop = e => this.onDrop(e, "journee");
    });
    htmlEl.querySelectorAll<HTMLElement>("[data-repas]").forEach(item => {
      const repas: RepasWithGroupe = JSON.parse(item.dataset.repas || "");
      item.draggable = repas.id_sejour == sejour.id; // uniquement les repas lié au séjour courant
      item.ondragstart = e => this.onDragStart(e, "repas", item.dataset.repas);
    });
  }

  private onDragStart = (
    event: DragEvent,
    kind: DragKind,
    arg: string | undefined
  ) => {
    if (!event.dataTransfer || !arg) return;
    event.dataTransfer.setData(kind, arg);
    event.dataTransfer.effectAllowed = "move";
  };

  onDragover(event: DragEvent, kind: DragKind) {
    if (!event.dataTransfer) return;
    const isRepas = event.dataTransfer.types.includes("repas");
    const isDay = event.dataTransfer.types.includes("journee");
    if ((kind == "repas" && isRepas) || (kind == "journee" && isDay)) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "move";
    }
  }

  onDrop(event: DragEvent, kind: DragKind) {
    if (!event.dataTransfer) return;
    event.preventDefault();
    const data = event.dataTransfer.getData(kind);
    const target = event.currentTarget as HTMLElement;
    if (kind == "repas") {
      this.onDropRepas(event, data, target);
    } else {
      this.onDropJournee(data, target);
    }
  }

  private jobDropRepas(targetTime: DateTime, dataRepas: string) {
    const repas: RepasWithGroupe = JSON.parse(dataRepas);
    const jour = new Date(targetTime.date);
    //FIXME:
    const horaire = "";
    C.data.deplaceRepas(repas, jour, horaire);
  }

  private onDropRepas(
    event: DragEvent,
    dataRepas: string,
    target: HTMLElement
  ) {
    // hack pour contourner les limitations de VCalendar
    const ev = new MouseEvent("mousedown", {
      view: window,
      bubbles: true,
      cancelable: true,
      clientX: event.clientX,
      clientY: event.clientY
    });
    this.lastClickedTime = null;
    target.dispatchEvent(ev);
    // on attend la gestion de l'évènement par 'registerTime'
    let currentTry = 0;
    const afterClick = () => {
      if (currentTry >= 200) {
        // on évite une boucle infinie
        return;
      }
      if (!this.lastClickedTime) {
        currentTry += 1;
        setTimeout(afterClick, 50); // on ressaie plus tard
        return;
      }
      this.jobDropRepas(this.lastClickedTime, dataRepas);
    };
    setTimeout(afterClick, 50);
  }

  private onDropJournee(data: string, target: HTMLElement) {
    if (this.sejour == null) return;
    const customDiv = target.querySelector("[data-day]") as HTMLElement;
    const dateFrom = new Date(data);
    if (!customDiv || !customDiv.dataset.day) return;
    const dateTo = new Date(customDiv.dataset.day);
    if (isNaN(dateFrom.getTime()) || isNaN(dateTo.getTime())) return;
    C.data.switchDays(this.sejour.id, dateFrom, dateTo);
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

.calendar-toolbar .v-input__control {
  margin: auto;
}
.calendar-toolbar .v-input__slot {
  height: 100%;
}
.calendar-toolbar .v-input {
  height: 100%;
}
</style>
