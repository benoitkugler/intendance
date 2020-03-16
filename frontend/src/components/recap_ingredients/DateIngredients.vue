<template>
  <div>
    <v-expansion-panels
      :value="dateIngredients.map((_, i) => i)"
      multiple
      accordion
    >
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

import { DateIngredientQuantites, IngredientQuantite } from "../../logic/types";
import { Formatter } from "../../logic/formatter";

const DateIngredientsProps = Vue.extend({
  props: {
    dateIngredients: Array as () => DateIngredientQuantites[]
  }
});

@Component({
  components: { ListeIngredients }
})
export default class DateIngredients extends DateIngredientsProps {
  formatDate = Formatter.formatDate;
}
</script>

<style scoped></style>
