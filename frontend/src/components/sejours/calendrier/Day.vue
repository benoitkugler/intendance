<template>
  <div class="day">
    <v-calendar
      v-if="day != null"
      type="custom-daily"
      locale="fr"
      :interval-count="1"
      :interval-height="500"
      :interval-width="0"
      :event-height="30"
      :start="dayString"
      event-color="#6a7d6f"
    >
      <template v-slot:interval>
        <v-list dense>
          <div v-for="horaire in horaires" :key="horaire.value">
            <v-subheader @click="startCreateRepas">
              {{ horaire.text }}
              <v-spacer></v-spacer>
              <tooltip-btn mdi-icon="plus" color="green" tooltip="Ajouter un repas..."></tooltip-btn>
            </v-subheader>
            <v-list-item-group color="primary">
              <v-list-item v-for="repas in events[horaire.value]" :key="repas.id">
                <v-list-item-content>
                  <v-row no-gutters class="fill-height">
                    <v-col class="px-1 align-self-center overflow-x-auto">
                      <v-chip
                        label
                        v-for="groupe in getGroupes(repas)"
                        :key="groupe.id"
                        class="mr-1 px-1 align-self-center"
                        :color="groupe.couleur"
                        small
                        :style="{ borderWidth: ' 1.5px' }"
                        outlined
                        draggable
                      >{{ groupe.nom }}</v-chip>
                    </v-col>
                  </v-row>
                </v-list-item-content>
              </v-list-item>
            </v-list-item-group>
          </div>
        </v-list>
      </template>
    </v-calendar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../../utils/TooltipBtn.vue";

import { toDateVuetify, DataEvent, getEventStart } from "./utils";
import { C } from "../../../logic/controller";
import { RepasWithGroupe, Groupe } from "../../../logic/types";
import { Horaires } from "../../utils/enums";

const DayProps = Vue.extend({
  props: {
    day: Date as new () => Date
  }
});
@Component({
  components: { TooltipBtn }
})
export default class Day extends DayProps {
  get dayString() {
    if (this.day == null) return "";
    return toDateVuetify(this.day);
  }

  get horaires() {
    return Horaires;
  }

  get events(): { [key: string]: RepasWithGroupe[] } {
    const repass: RepasWithGroupe[] = [];
    C.iterateAllRepas((sejour, repas) => {
      if (sejour.id != C.state.idSejour) return;
      const date = C.offsetToDate(sejour.id, repas.jour_offset);
      if (date.getTime() == this.day.getTime()) {
        repass.push(repas);
      }
    });
    //TODO: et tri dans chaque catégorie
    return { midi: repass };
  }

  getGroupes(repas: RepasWithGroupe) {
    const grs = C.getRepasGroupes(repas);
    const maxChar = 8;
    return grs.map(groupe => {
      const g: Groupe = JSON.parse(JSON.stringify(groupe));
      g.nom = g.nom.substr(0, maxChar) + (g.nom.length <= maxChar ? "" : ".");
      return groupe;
    });
  }

  startCreateRepas() {
    console.log("oj");
  }
}
</script>

<style>
</style>
