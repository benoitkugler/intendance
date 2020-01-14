<template>
  <div class="fill-height">
    <v-dialog v-model="showPreferences" max-width="800">
      <form-preferences v-model="preferences"></form-preferences>
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
              />
            </v-col>
            <v-col cols="4" class="align-self-center">
              <day :day="activeDay"></day>
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
import { PreferencesAgenda, CalendarMode } from "../logic/types2";

import Calendar from "../components/sejours/calendrier/Calendar.vue";
import Day from "../components/sejours/calendrier/Day.vue";
import TooltipBtn from "../components/utils/TooltipBtn.vue";
import ToolbarSwitch from "../components/utils/ToolbarSwitch.vue";
import FormPreferences from "../components/sejours/calendrier/FormPreferences.vue";
import { DateTime } from "../components/sejours/calendrier/utils";

@Component({
  components: { Calendar, TooltipBtn, FormPreferences, ToolbarSwitch, Day }
})
export default class Agenda extends Vue {
  showPreferences = false;
  preferences: PreferencesAgenda = {
    startPremierJour: true
  };

  calendarMode: CalendarMode = "groupes";
  activeDay: Date | null = null; // mode "groupes" only

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
    this.activeDay = new Date(event.date);
  }
}
</script>
