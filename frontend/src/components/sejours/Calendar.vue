<template>
  <div class="two-weeks-calendar" ref="weeks">
    <v-dialog v-model="showPreferences" max-width="800">
      <form-preferences v-model="preferences"></form-preferences>
    </v-dialog>

    <v-dialog v-model="showEditFormRepas" max-width="500">
      <form-repas
        :initialRepas="editedRepas"
        :mode="editMode"
        @accept="onEditRepasDone"
        @delete="deleteRepas"
      ></form-repas>
    </v-dialog>

    <v-dialog v-model="showFormCalcul" max-width="600px">
      <form-calcul :initial-sejour="currentSejour"></form-calcul>
    </v-dialog>

    <v-toolbar class="calendar-toolbar mb-1">
      <v-toolbar-title>Séjours</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          tooltip="Calculer les ingrédients nécessaires..."
          mdi-icon="food-variant"
          @click="showFormCalcul = true"
        ></tooltip-btn>
        <v-divider vertical></v-divider>
        <tooltip-btn
          tooltip="Ajouter un séjour..."
          mdi-icon="shape-rectangle-plus"
          @click="
            editMode = 'new';
            showEditFormSejour = true;
          "
        ></tooltip-btn>
        <v-divider vertical></v-divider>
        <select-sejour
          label="Séjour actif"
          v-model.number="currentSejour"
        ></select-sejour>
        <tooltip-btn
          :disabled="!currentSejour"
          tooltip="Modifier les paramètres du séjour..."
          mdi-icon="pencil"
          @click="
            editMode = 'edit';
            showEditFormSejour = true;
          "
        ></tooltip-btn>
        <tooltip-btn
          :disabled="!currentSejour"
          tooltip="Supprimer le séjour..."
          mdi-icon="close"
          color="red"
          @click="showConfirmeSupprime = true"
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
      :event-height="22"
      :short-weekdays="false"
      :first-interval="firstInterval"
      :interval-count="intervalCount"
      :interval-minutes="intervalMinutes"
      :interval-height="intervalHeight"
      :interval-style="intervalStyle"
      :start="startWeek1"
      :weekdays="weekdays"
      :events="events"
      :event-color="getEventColor"
      @mousedown:time="registerTime"
      @click:time="startAddRepas"
    >
      <template v-slot:event="{ event }">
        <div
          :data-repas="event.dataRepas"
          @click.stop="startEditRepas(event.repas)"
          :style="{
            fontWeight: event.repas.id_sejour == currentSejour ? 'bold' : ''
          }"
        >
          {{ event.name }}
        </div>
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
      class="mt-1"
      type="week"
      locale="fr"
      :event-height="22"
      :short-weekdays="false"
      :first-interval="firstInterval"
      :interval-count="intervalCount"
      :interval-minutes="intervalMinutes"
      :interval-height="intervalHeight"
      :interval-style="intervalStyle"
      :start="startWeek2"
      :weekdays="weekdays"
      :events="events"
      :event-color="getEventColor"
      @mousedown:time="registerTime"
      @click:time="startAddRepas"
    >
      <template v-slot:event="{ event }">
        <div
          :data-repas="event.dataRepas"
          @click.stop="startEditRepas(event.repas)"
          :style="{
            fontWeight: event.repas.id_sejour == currentSejour ? 'bold' : ''
          }"
        >
          {{ event.name }}
        </div>
      </template>
      <template v-slot:day-header="{ date }">
        <div :data-day="date"></div>
      </template>
      <template v-slot:interval
        ><div
          @dragover="e => onDragover(e, 'repas')"
          @drop="e => onDrop(e, 'repas')"
          class="dragover"
        ></div>
      </template>
    </v-calendar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Repas, Sejour, Horaire } from "../../logic/types";
import FormSejour from "./FormSejour.vue";
import FormRepas from "./FormRepas.vue";
import FormPreferences from "./FormPreferences.vue";
import FormCalcul from "./FormCalcul.vue";
import SelectSejour from "./SelectSejour.vue";
import { C } from "../../logic/controller";
import TooltipBtn from "../utils/TooltipBtn.vue";
import {
  DetailsSejour,
  New,
  DetailsRepas,
  PreferencesAgenda
} from "../../logic/types2";
import { Formatter } from "../../logic/formatter";

const _days = [0, 1, 2, 3, 4, 5, 6];

const _colors = [
  "red",
  "pink",
  "purple",
  "indigo",
  "blue",
  "cyan",
  "teal",
  "green",
  "lime",
  "yellow",
  "amber",
  "orange",
  "brown",
  "grey"
];

type DragKind = "journee" | "repas";

interface DataEvent {
  name: string;
  start: Time;
  repas: Repas;
  dataRepas: string; // JSON of Repas
}

// renvoie l'ordre des jours pour que `start` soit
// affiché en premier
function weekdaysFromStart(start: Date) {
  const d0 = start.getDay();
  return _days.map(d => (d0 + d) % 7);
}

function formatDate(d: Date) {
  return d.toISOString().substr(0, 10);
}

function timeToHoraire(time: string): Horaire {
  return {
    heure: Number(time.substr(0, 2)),
    minute: Number(time.substr(3, 2))
  };
}

function getEventStart(r: Repas) {
  const dateDebut = C.formatter.offsetToDate(r.id_sejour, r.jour_offset);
  return formatDate(dateDebut) + " " + Formatter.horaireToTime(r.horaire);
}

const Props = Vue.extend({
  props: {}
});

interface DateTime {
  date: string;
  time: string;
}

@Component({
  components: {
    TooltipBtn,
    FormSejour,
    FormRepas,
    FormPreferences,
    SelectSejour,
    FormCalcul
  }
})
export default class Calendar extends Props {
  private lastClickedTime: DateTime | null = null;
  private currentSejour: number | null = null;

  private preferences: PreferencesAgenda = {
    restrictSejourCourant: true,
    startPremierJour: true
  };

  private firstInterval = 4;
  private intervalMinutes = 120;
  private intervalCount = 8;
  private intervalHeight = 25;

  protected showPreferences = false;
  protected showEditFormSejour = false;
  protected showEditFormRepas = false;
  protected showConfirmeSupprime = false;
  protected showFormCalcul = false;

  private editMode: "new" | "edit" = "new";

  private editedRepas: Repas = {
    id: -1,
    id_sejour: -1,
    horaire: { heure: 0, minute: 0 },
    id_menu: -1,
    offset_personnes: 0,
    jour_offset: 0
  };

  get startDate(): Date {
    const sejour = this.getCurrentSejour();
    if (!sejour) return new Date();
    return new Date(sejour.date_debut);
  }

  get weekdays() {
    if (this.preferences.startPremierJour) {
      return weekdaysFromStart(this.startDate);
    } else {
      return [1, 2, 3, 4, 5, 6, 0];
    }
  }

  get startWeek1() {
    return formatDate(this.startDate);
  }

  get startWeek2() {
    const out = this.startDate;
    out.setDate(this.startDate.getDate() + 7);
    return formatDate(out);
  }

  get events(): DataEvent[] {
    let out: DataEvent[] = [];
    const restrict = this.preferences.restrictSejourCourant;
    C.iterateAllRepas((sejour, repas) => {
      if (restrict && sejour.id != this.currentSejour) return;
      const data = {
        repas: repas,
        name: C.formatter.formatRepasName(repas),
        start: getEventStart(repas),
        dataRepas: JSON.stringify(repas)
      };
      out.push(data);
    });

    return out;
  }

  private get sortedSejours() {
    const sejours = Object.values(C.data.sejours.sejours);
    sejours.sort((a, b) => {
      return a.date_debut < b.date_debut ? 1 : -1;
    });
    return sejours;
  }

  private sejourColor(idSejour: number) {
    const L = _colors.length;
    const index = this.sortedSejours.findIndex(sej => sej.id == idSejour);
    if (index != -1) {
      return _colors[index % L];
    }
    return "black";
  }

  intervalStyle(interval: DateTime) {
    const day = new Date(interval.date);
    const sej = this.getCurrentSejour();
    if (!sej) return {};
    const debutSejour = new Date(sej.date_debut);
    if (debutSejour <= day) {
      return { backgroundColor: this.sejourColor(sej.id), opacity: 0.15 };
    }
  }

  getEventColor = (event: DataEvent) => {
    return this.sejourColor(event.repas.id_sejour);
  };

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
    if (this.currentSejour == null) return;
    const customDiv = target.querySelector("[data-day]") as HTMLElement;
    const dateFrom = new Date(data);
    if (!customDiv || !customDiv.dataset.day) return;
    const dateTo = new Date(customDiv.dataset.day);
    if (isNaN(dateFrom.getTime()) || isNaN(dateTo.getTime())) return;
    C.data.switchDays(this.currentSejour, dateFrom, dateTo);
  }

  private getCurrentSejour(): Sejour | undefined {
    if (this.currentSejour == null) return;
    return C.data.sejours.sejours[this.currentSejour];
  }

  getInitialCurrentSejour() {
    const newSejour = {
      id_proprietaire: C.idUtilisateur,
      nom: "",
      date_debut: ""
    };
    return this.getCurrentSejour() || newSejour;
  }

  startAddRepas(dt: DateTime) {
    const sejour = this.getCurrentSejour();
    if (sejour == undefined) return;
    const offset = C.data.getOffset(sejour, new Date(dt.date));
    if (offset == undefined) return;
    this.editedRepas = {
      horaire: timeToHoraire(dt.time),
      jour_offset: offset,
      offset_personnes: 0,
      id_menu: -1,
      id_sejour: sejour.id,
      id: -1
    };
    this.editMode = "new";
    this.showEditFormRepas = true;
  }

  startEditRepas(repas: Repas) {
    this.editedRepas = repas;
    this.editMode = "edit";
    this.showEditFormRepas = true;
  }

  async onEditRepasDone(repas: DetailsRepas) {
    this.showEditFormRepas = false;
    if (this.currentSejour == null) return;
    let message = "";
    if (this.editMode == "new") {
      const newRepas: New<Repas> = { ...repas, id_sejour: this.currentSejour };
      await C.data.createRepas(newRepas);
      message = "Le repas a bien été ajouté.";
    } else {
      const repasFull = {
        ...repas,
        id_sejour: this.editedRepas.id_sejour,
        id: this.editedRepas.id
      };
      await C.data.updateManyRepas([repasFull]);
      message = "Le repas a bien été mis à jour.";
    }
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(message);
    }
  }

  async deleteRepas(repas: Repas) {
    this.showEditFormRepas = false;
    await C.data.deleteRepas(repas);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Le repas a été retiré avec succès.");
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
