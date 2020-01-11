<template>
  <div>
    <v-dialog v-model="showPreferences" max-width="800">
      <form-preferences v-model="preferences"></form-preferences>
    </v-dialog>

    <v-container fluid style="height:100%;">
      <v-row style="height:100%;">
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
          <calendar
            ref="calendar"
            :sejour="sejour"
            :preferences="preferences"
            :mode="calendarModeGroupe ? 'groupes' : 'menus'"
          />
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
import TooltipBtn from "../components/utils/TooltipBtn.vue";
import ToolbarSwitch from "../components/utils/ToolbarSwitch.vue";
import FormPreferences from "../components/sejours/calendrier/FormPreferences.vue";

@Component({
  components: { Calendar, TooltipBtn, FormPreferences, ToolbarSwitch }
})
export default class Agenda extends Vue {
  showPreferences = false;
  preferences: PreferencesAgenda = {
    startPremierJour: true
  };

  calendarModeGroupe = true;

  $refs!: {
    calendar: Calendar;
  };

  get sejour() {
    return C.state.getSejour();
  }

  async mounted() {
    await C.data.loadAllMenus();
    await C.data.loadSejours();
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("L'agenda a bien été chargé.");
    }
  }
}
</script>
