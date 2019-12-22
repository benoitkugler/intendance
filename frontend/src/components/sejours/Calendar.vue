<template>
  <div class="two-weeks-calendar" ref="weeks">
    <v-dialog v-model="showPreferences" max-width="800">
      <v-card>
        <v-card-title primary-title>
          Préférences
        </v-card-title>
        <v-card-text>
          <v-switch
            label="Se restreindre au séjour courant"
            v-model="showOnlyCurrent"
            persistent-hint
            hint="N'afficher que le séjour actuellement sélectionné."
          ></v-switch>
        </v-card-text>
      </v-card>
    </v-dialog>

    <v-dialog v-model="showEditFormSejour" max-width="500">
      <form-sejour :initialSejour="{}" @accept="editSejour"></form-sejour>
    </v-dialog>

    <v-toolbar class="calendar-toolbar mb-1">
      <v-toolbar-title>Séjours</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          tooltip="Ajouter un séjour..."
          mdi-icon="shape-rectangle-plus"
          @click="addSejour"
        ></tooltip-btn>
        <v-divider vertical></v-divider>
        <v-select
          :items="sejours"
          label="Séjour courant"
          hide-details
          max-width="200"
          v-model.number="currentSejour"
          class="mx-2"
        ></v-select>
        <tooltip-btn
          tooltip="Modifier les paramètres du séjour..."
          mdi-icon="pencil"
          @click="showEditFormSejour = true"
        ></tooltip-btn>
        <v-divider vertical></v-divider>
        <tooltip-btn
          tooltip="Préférences d'affichage..."
          mdi-icon="settings"
          @click="showPreferences = true"
        ></tooltip-btn>
      </v-toolbar-items>
    </v-toolbar>
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
        { name: 'jejejej', start: '2019-12-20T09:00', dataRepas: '{ id: 78 }' },
        { name: 'jejejej', start: '2019-12-20T08:00', dataRepas: '{ id: 75 }' }
      ]"
      @click:time="registerTime"
    >
      <template v-slot:event="{ event }">
        <div :data-repas="event.dataRepas">{{ event.name }}</div>
      </template>
      <template v-slot:day-header="{ date }">
        <div :data-day="date"></div>
      </template>
      <template v-slot:interval
        ><div
          @dragover="e => onDragover(e, 'repas')"
          @drop="e => onDrop(e, 'repas')"
          class="dragover"
        ></div
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
      @click:time="registerTime"
    ></v-calendar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Repas, Sejour, Horaire, SejourJournees } from "../../logic/types";
import FormSejour from "./FormSejour.vue";
import { D } from "../../logic/controller";
import TooltipBtn from "../utils/TooltipBtn.vue";
import { DetailsSejour } from "../../logic/types2";

const _days = [0, 1, 2, 3, 4, 5, 6];

type DragKind = "journee" | "repas";

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

const Props = Vue.extend({
  props: {
    // Date as string
    start: String,
    forceFirstDay: Boolean
  }
});

interface DateTime {
  date: string;
  time: string;
}

@Component({
  components: { TooltipBtn, FormSejour }
})
export default class Calendar extends Props {
  private lastClickedTime: DateTime | null = null;
  private currentSejour: number | null = null;
  private showOnlyCurrent = true;
  private firstInterval = 4;
  private intervalCount = 7;
  private intervalMinutes = 120;
  private intervalHeight = 25;

  private showPreferences = false;
  private showEditFormSejour = false;

  get startDate(): Date {
    return new Date(this.start);
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

  get sejours() {
    const items = Object.keys(D.agenda.sejours).map(idSejour => {
      const sejour = D.agenda.sejours[Number(idSejour)];
      return { value: idSejour, text: sejour.sejour.nom };
    });
    items.sort((a, b) => Number(a.text < b.text));
    return items;
  }

  mounted() {
    this.setupDrag();
  }

  updated() {
    this.setupDrag();
  }

  // on stocke le moment correspondant au dernier click,
  // pour contourner le manque d'une méthode getTime(pos)
  registerTime(time: DateTime) {
    this.lastClickedTime = time;
  }

  private setupDrag() {
    if (this.currentSejour == null) return; // désactivé si aucun séjour n'est sélectionné.
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
      const repas: Repas = JSON.parse(item.dataset.repas || "");
      item.draggable = repas.id_sejour == this.currentSejour; // uniquement les repas lié au séjour courant
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
    const repas: Repas = JSON.parse(dataRepas);
    const jour = new Date(targetTime.date);
    const horaire: Horaire = {
      heure: Number(targetTime.time.substr(0, 2)),
      minute: Number(targetTime.time.substr(3, 2))
    };
    D.deplaceRepas(repas, jour, horaire);
  }

  private onDropRepas(
    event: DragEvent,
    dataRepas: string,
    target: HTMLElement
  ) {
    // hack pour contourner les limitations de VCalendar
    const ev = new MouseEvent("click", {
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
    if (this.currentSejour == null) return;
    const customDiv = target.querySelector("[data-day]") as HTMLElement;
    const dateFrom = new Date(data);
    if (!customDiv || !customDiv.dataset.day) return;
    const dateTo = new Date(customDiv.dataset.day);
    if (isNaN(dateFrom.getTime()) || isNaN(dateTo.getTime())) return;
    D.switchDays(this.currentSejour, dateFrom, dateTo);
  }

  addSejour() {
    console.log("add");
  }

  editSejour(sejour: DetailsSejour) {
    this.showEditFormSejour = false;
    console.log(sejour);
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
