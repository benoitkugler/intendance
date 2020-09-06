<template>
  <div>
    <v-expansion-panels :value="openPanels" multiple accordion>
      <v-expansion-panel
        v-for="(jour, i) in dateIngredients"
        :key="i"
        :disabled="dateIngredients.length <= 1"
      >
        <v-expansion-panel-header
          :hide-actions="dateIngredients.length <= 1"
          class="py-0"
        >
          <b>{{ formatDate(jour.date) }}</b>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <liste-ingredients
            :ingredients="jour.ingredients"
            :highlights="getJourHighlights(jour)"
            @go="id => $emit('go', id)"
          ></liste-ingredients>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <div v-if="dateIngredients.length == 0" class="pa-2 font-italic">
      Aucune journée ne nécessite d'ingrédients.
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import ListeIngredients from "../utils/ListeIngredients.vue";

import {
  DateIngredientQuantites,
  IngredientQuantite,
  TimedIngredientQuantite
} from "../../logic/api";
import { Formatter } from "../../logic/formatter";
import { Crible } from "../utils/utils";

const DateIngredientsProps = Vue.extend({
  props: {
    dateIngredients: Array as () => DateIngredientQuantites[],
    highlight: Array as () => TimedIngredientQuantite[]
  }
});

@Component({
  components: { ListeIngredients }
})
export default class DateIngredients extends DateIngredientsProps {
  formatDate = Formatter.formatDate;

  get openPanels() {
    // cas particulier en l'absence de highlight : on affiche tout
    if (this.highlight.length == 0) {
      return this.dateIngredients.map((_, i) => i);
    }

    const out: number[] = [];
    this.dateIngredients.forEach((jour, i) => {
      const isConcerned =
        this.highlight.filter(h => h.date == jour.date).length > 0;
      if (isConcerned) {
        out.push(i);
      }
    });
    return out;
  }

  getJourHighlights(jour: DateIngredientQuantites): Crible {
    const out: Crible = {};
    this.highlight
      .filter(h => h.date == jour.date)
      .forEach(h => {
        out[h.ingredient.id] = true;
      });
    return out;
  }
}
</script>

<style scoped></style>
