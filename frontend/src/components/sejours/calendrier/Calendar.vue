<template>
  <div class="two-weeks-calendar" ref="weeks">
    <v-dialog v-model="showEditFormRepas" max-width="500">
      <form-repas
        :initialRepas="editedRepas"
        :mode="editMode"
        @accept="onEditRepasDone"
        @delete="deleteRepas"
      ></form-repas>
    </v-dialog>

    <v-dialog v-model="showFormCalcul" max-width="600px">
      <form-calcul :initial-sejour="sejour"></form-calcul>
    </v-dialog>

    <v-calendar
      type="week"
      locale="fr"
      :event-height="22"
      :short-weekdays="false"
      :first-interval="firstInterval"
      :interval-count="intervalCount"
      :interval-minutes="intervalMinutes"
      :interval-height="intervalHeight"
      :start="startWeek1"
      :weekdays="weekdays"
      :events="events"
      @mousedown:time="registerTime"
      @click:time="startAddRepas"
      @click:date="args => $emit('change', args)"
    >
      <template v-slot:event="{ event }">
        <div
          :data-repas="JSON.stringify(event.repas)"
          @click.stop="startEditRepas(event.repas)"
          :style="{ fontWeight: 'bold' }"
          class="px-1"
        >
          {{ formatEventName(event) }}
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
      :start="startWeek2"
      :weekdays="weekdays"
      :events="events"
      @mousedown:time="registerTime"
      @click:time="startAddRepas"
      @click:date="args => $emit('change', args)"
    >
      <template v-slot:event="{ event }">
        <div
          :data-repas="JSON.stringify(event.repas)"
          @click.stop="startEditRepas(event.repas)"
          :style="{ fontWeight: 'bold' }"
          class="px-1"
        >
          {{ formatEventName(event) }}
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
import {
  Sejour,
  Horaire,
  SejourRepas,
  RepasWithGroupe
} from "../../../logic/types";
import FormRepas from "../FormRepas.vue";
import FormCalcul from "../FormCalcul.vue";
import { C } from "../../../logic/controller";
import TooltipBtn from "../../utils/TooltipBtn.vue";
import {
  DetailsSejour,
  New,
  DetailsRepas,
  PreferencesAgenda,
  CalendarMode
} from "../../../logic/types2";
import { Formatter } from "../../../logic/formatter";
import { DateTime } from "./types";
import { toDateVuetify } from "./utils";

const _days = [0, 1, 2, 3, 4, 5, 6];

type DragKind = "journee" | "repas";

interface DataEvent {
  start: Time;
  repas: RepasWithGroupe;
}

// renvoie l'ordre des jours pour que `start` soit
// affiché en premier
function weekdaysFromStart(start: Date) {
  const d0 = start.getDay();
  return _days.map(d => (d0 + d) % 7);
}

function timeToHoraire(time: string): Horaire {
  return {
    heure: Number(time.substr(0, 2)),
    minute: Number(time.substr(3, 2))
  };
}

function getEventStart(r: RepasWithGroupe) {
  const dateDebut = C.formatter.offsetToDate(r.id_sejour, r.jour_offset);
  return toDateVuetify(dateDebut) + " " + Formatter.horaireToTime(r.horaire);
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
    FormRepas,
    FormCalcul
  }
})
export default class Calendar extends Props {
  private lastClickedTime: DateTime | null = null;

  private firstInterval = 4;
  private intervalCount = 8;
  private intervalMinutes = 120;
  private intervalHeight = 25;

  protected showEditFormRepas = false;
  protected showFormCalcul = false;

  private editMode: "new" | "edit" = "new";

  private editedRepas: RepasWithGroupe = {
    id: -1,
    id_sejour: -1,
    horaire: { heure: 0, minute: 0 },
    id_menu: -1,
    offset_personnes: 0,
    jour_offset: 0,
    groupes: []
  };

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

  get events(): DataEvent[] {
    const sejour = this.sejour;
    if (sejour == null) return [];
    let out: DataEvent[] = [];
    C.iterateAllRepas((_, repas) => {
      if (repas.id_sejour != sejour.id) return;
      const data = {
        repas: repas,
        start: getEventStart(repas)
      };
      out.push(data);
    });
    return out;
  }

  formatEventName(event: DataEvent) {
    if (this.mode == "groupes") {
      return C.getRepasGroupes(event.repas)
        .map(g => g.nom)
        .join(" / ");
    } else {
      return C.formatter.formatRepasName(event.repas);
    }
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
    if (this.sejour == null) return;
    const customDiv = target.querySelector("[data-day]") as HTMLElement;
    const dateFrom = new Date(data);
    if (!customDiv || !customDiv.dataset.day) return;
    const dateTo = new Date(customDiv.dataset.day);
    if (isNaN(dateFrom.getTime()) || isNaN(dateTo.getTime())) return;
    C.data.switchDays(this.sejour.id, dateFrom, dateTo);
  }

  startAddRepas(dt: DateTime) {
    const sejour = this.sejour;
    if (sejour == null) return;
    const offset = C.data.getOffset(sejour, new Date(dt.date));
    if (offset == undefined) return;
    this.editedRepas = {
      horaire: timeToHoraire(dt.time),
      jour_offset: offset,
      offset_personnes: 0,
      id_menu: -1,
      id_sejour: sejour.id,
      id: -1,
      groupes: []
    };
    this.editMode = "new";
    this.showEditFormRepas = true;
  }

  startEditRepas(repas: RepasWithGroupe) {
    this.editedRepas = repas;
    this.editMode = "edit";
    this.showEditFormRepas = true;
  }

  async onEditRepasDone(repas: DetailsRepas) {
    this.showEditFormRepas = false;
    if (this.sejour == null) return;
    let message = "";
    if (this.editMode == "new") {
      const newRepas: New<RepasWithGroupe> = {
        ...repas,
        id_sejour: this.sejour.id
      };
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

  async deleteRepas(repas: RepasWithGroupe) {
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
