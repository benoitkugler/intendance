<template>
  <div class="fill-height">
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

    <v-container fluid class="fill-height">
      <v-row no-gutters class="fill-height">
        <v-col v-if="sejour == null" class="align-self-center">
          <v-alert
            class="align-self-center"
            color="secondary"
            :value="true"
            transition="fade-transition"
          >
            Veuillez choisir un séjour via l'onglet <b>Séjours</b>
          </v-alert>
        </v-col>
        <v-col v-else>
          <v-toolbar dense class="calendar-toolbar mb-1">
            <v-toolbar-title>
              Organisation du séjour <b>{{ sejour.nom }}</b>
            </v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items>
              <tooltip-btn
                tooltip="Calculer les ingrédients nécessaires..."
                mdi-icon="food-variant"
                @click="showFormCalcul = true"
              ></tooltip-btn>
              <v-divider vertical></v-divider>
              <toolbar-switch
                v-model="calendarModeGroupe"
                label="Afficher les groupes"
                tooltip-on="Passer vers le choix des <b>menus</b>"
                tooltip-off="Passer vers la répartition des <b>groupes</b>"
              ></toolbar-switch>
              <v-divider vertical></v-divider>
              <tooltip-btn
                tooltip="Préférences d'affichage..."
                mdi-icon="settings"
                @click="showPreferences = true"
              ></tooltip-btn>
            </v-toolbar-items>
          </v-toolbar>
          <v-row>
            <v-col v-if="calendarMode == 'menus'"></v-col>
            <v-col cols="8">
              <calendar
                ref="calendar"
                :sejour="sejour"
                :preferences="preferences"
                :mode="calendarMode"
                @change="onChangeDay"
                @addRepas="startAddRepasFromDate"
                @editRepas="startEditRepas"
              />
            </v-col>
            <v-col cols="4" class="align-self-center">
              <day
                :jourOffset="activeJourOffset"
                v-if="activeJourOffset != null"
                @addRepas="startAddRepasFromHoraire"
              ></day>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import { C } from "../logic/controller";
import {
  PreferencesAgenda,
  CalendarMode,
  NullId,
  EditMode,
  New,
  DetailsRepas
} from "../logic/types2";

import Calendar from "../components/sejours/calendrier/Calendar.vue";
import Day from "../components/sejours/calendrier/Day.vue";
import TooltipBtn from "../components/utils/TooltipBtn.vue";
import ToolbarSwitch from "../components/utils/ToolbarSwitch.vue";
import FormPreferences from "../components/sejours/calendrier/FormPreferences.vue";
import FormRepas from "../components/sejours/FormRepas.vue";
import { DateTime } from "../components/sejours/calendrier/utils";
import { RepasGroupe, RepasWithGroupe } from "../logic/types";

@Component({
  components: {
    Calendar,
    TooltipBtn,
    FormPreferences,
    ToolbarSwitch,
    Day,
    FormRepas
  }
})
export default class Agenda extends Vue {
  showPreferences = false;
  preferences: PreferencesAgenda = {
    startPremierJour: true
  };

  showEditFormRepas = false;
  editMode: EditMode = "new";

  private editedRepas: RepasWithGroupe = {
    id: -1,
    id_sejour: -1,
    horaire: "",
    id_menu: NullId,
    offset_personnes: 0,
    jour_offset: 0,
    groupes: []
  };

  calendarMode: CalendarMode = "menus";
  activeJourOffset: number | null = null; // mode "groupes" only

  $refs!: {
    calendar: Calendar;
  };

  get sejour() {
    return C.state.getSejour();
  }

  get calendarModeGroupe() {
    return this.calendarMode == "groupes";
  }
  set calendarModeGroupe(b: boolean) {
    this.calendarMode = b ? "groupes" : "menus";
  }

  async mounted() {
    await C.data.loadAllMenus();
    await C.data.loadSejours();
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("L'agenda a bien été chargé.");
    }
  }

  onChangeDay(event: DateTime) {
    if (C.state.idSejour == null) {
      this.activeJourOffset = null;
    } else {
      const target = new Date(event.date);
      this.activeJourOffset = C.dateToOffset(C.state.idSejour, target);
    }
  }

  // Edition des repas
  startAddRepasFromDate(date: string) {
    if (C.state.idSejour == null) return;
    const offset = C.dateToOffset(C.state.idSejour, new Date(date));
    this.startAddRepas(offset, "");
  }

  startAddRepasFromHoraire(horaire: string) {
    if (this.activeJourOffset == null) return;
    this.startAddRepas(this.activeJourOffset, horaire);
  }

  private startAddRepas(jourOffset: number, horaire: string) {
    const sejour = this.sejour;
    if (sejour == null) return;
    this.editedRepas = {
      horaire: horaire,
      jour_offset: jourOffset,
      offset_personnes: 0,
      id_menu: NullId,
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
