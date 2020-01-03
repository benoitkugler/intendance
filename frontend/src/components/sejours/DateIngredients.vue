<template>
  <v-expansion-panels :value="0">
    <v-expansion-panel
      v-for="(jour, i) in dateIngredients"
      :key="i"
      :disabled="dateIngredients.length <= 1"
    >
      <v-expansion-panel-header :hide-actions="dateIngredients.length <= 1">
        <b>{{ formatDate(jour.date) }}</b>
      </v-expansion-panel-header>
      <v-expansion-panel-content>
        <liste-ingredients :ingredients="jour.ingredients"></liste-ingredients>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import ListeIngredients from "./ListeIngredients.vue";

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
