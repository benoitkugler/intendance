<template>
  <div class="fill-height">
    <v-dialog v-model="showPreferences" max-width="800">
      <form-preferences v-model="preferences"></form-preferences>
    </v-dialog>

    <v-dialog v-model="showEditFormRepas" size="lg">
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
            <tooltip-btn
              mdi-icon="information-outline"
              :tooltip="tooltip"
            ></tooltip-btn>
            <v-toolbar-title>
              Organisation du séjour <b>{{ sejour.nom }}</b>
              <span v-if="viewMode == 'day'">
                - Journée du
                {{
                  activeDay.toLocaleDateString("fr-FR", {
                    weekday: "long",
                    day: "numeric",
                    month: "short"
                  })
                }}
              </span>
            </v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items>
              <tooltip-btn
                tooltip="Reculer"
                mdi-icon="arrow-left"
                small
                @click="activeJourOffset -= 1"
                v-if="viewMode == 'day'"
              ></tooltip-btn>
              <tooltip-btn
                tooltip="Avancer"
                mdi-icon="arrow-right"
                small
                @click="activeJourOffset += 1"
                v-if="viewMode == 'day'"
              ></tooltip-btn>
              <v-divider vertical></v-divider>
              <tooltip-btn
                mdi-icon="calendar-month"
                tooltip="Revenir à la vue générale"
                :disabled="viewMode == 'month'"
                @click="viewMode = 'month'"
              ></tooltip-btn>
              <v-divider vertical></v-divider>
              <tooltip-btn
                tooltip="Préférences d'affichage..."
                mdi-icon="cogs"
                @click="showPreferences = true"
              ></tooltip-btn>
            </v-toolbar-items>
          </v-toolbar>

          <calendar
            v-if="viewMode == 'month'"
            ref="calendar"
            :sejour="sejour"
            :preferences="preferences"
            :activeJourOffset="activeJourOffset"
            @change="onChangeDay"
            @addRepas="startAddRepasFromDate"
            @editRepas="startEditRepas"
          />
          <day
            v-else
            :jourOffset="activeJourOffset"
            @addRepas="startAddRepasFromHoraire"
            @editRepas="startEditRepas"
          ></day>
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
  EditMode,
  New,
  DetailsRepas,
  ViewMode
} from "../logic/api";

import Calendar from "../components/sejours/calendrier/Calendar.vue";
import Day from "../components/sejours/calendrier/day/Day.vue";
import TooltipBtn from "../components/utils/TooltipBtn.vue";
import ToolbarSwitch from "../components/utils/ToolbarSwitch.vue";
import FormPreferences from "../components/sejours/calendrier/FormPreferences.vue";
import FormRepas from "../components/sejours/FormRepas.vue";

import { RepasGroupe, RepasComplet, Horaire } from "../logic/api";

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

  private editedRepas: RepasComplet = {
    id: -1,
    id_sejour: -1,
    offset_personnes: 0,
    jour_offset: 0,
    horaire: Horaire.Midi,
    anticipation: 0,
    groupes: [],
    recettes: [],
    ingredients: []
  };

  viewMode: ViewMode = "month";
  activeJourOffset: number | null = null; // mode "groupes" only

  $refs!: {
    calendar: Calendar;
  };

  get sejour() {
    return C.state.getSejour();
  }

  get activeDay(): Date | null {
    if (this.activeJourOffset == null || C.state.idSejour == null) return null;
    return C.offsetToDate(C.state.idSejour, this.activeJourOffset);
  }

  get tooltip() {
    const sej = C.state.getSejour();
    let nbRepas = sej ? (sej.repass || []).length : 0;
    return `<b>${nbRepas}</b> repas`;
  }

  async mounted() {
    await C.data.loadAllMenus();
    await C.data.loadSejours();
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("L'agenda a bien été chargé.");
    }
  }

  onChangeDay(date: string) {
    if (C.state.idSejour == null) {
      this.activeJourOffset = null;
    } else {
      const target = new Date(date);
      this.activeJourOffset = C.dateToOffset(C.state.idSejour, target);
      this.viewMode = "day";
    }
  }

  // Edition des repas
  startAddRepasFromDate(date: string) {
    if (C.state.idSejour == null) return;
    const offset = C.dateToOffset(C.state.idSejour, new Date(date));
    this.startAddRepas(offset, Horaire.Midi);
  }

  startAddRepasFromHoraire(horaire: Horaire) {
    if (this.activeJourOffset == null) return;
    this.startAddRepas(this.activeJourOffset, horaire);
  }

  private startAddRepas(jourOffset: number, horaire: Horaire) {
    const sejour = this.sejour;
    if (sejour == null) return;
    this.editedRepas = {
      id: -1,
      id_sejour: sejour.id,
      jour_offset: jourOffset,
      offset_personnes: 0,
      horaire: horaire,
      anticipation: 0,
      groupes: [],
      recettes: [],
      ingredients: []
    };
    this.editMode = "new";
    this.showEditFormRepas = true;
  }

  startEditRepas(repas: RepasComplet) {
    this.editedRepas = repas;
    this.editMode = "edit";
    this.showEditFormRepas = true;
  }

  async onEditRepasDone(repas: DetailsRepas) {
    this.showEditFormRepas = false;
    if (this.sejour == null) return;
    let message = "";
    if (this.editMode == "new") {
      const newRepas: New<RepasComplet> = {
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

  async deleteRepas(repas: RepasComplet) {
    this.showEditFormRepas = false;
    await C.data.deleteRepas(repas);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Le repas a été retiré avec succès.");
    }
  }
}
</script>
