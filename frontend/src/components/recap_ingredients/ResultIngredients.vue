<template>
  <v-card>
    <v-card-title primary-title class="secondary py-2 px-3">
      <h3 class="headline mb-0 ">
        Ingrédients nécessaires
      </h3>
    </v-card-title>
    <v-progress-linear indeterminate :active="loading"></v-progress-linear>
    <date-ingredients
      class="overflow-y-auto"
      style="max-height: 75vh;"
      :date-ingredients="dateIngredients"
      :highlight="origineIngredients"
      @go="id => $emit('goToIngredient', id)"
    ></date-ingredients>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import DateIngredients from "./DateIngredients.vue";

import {
  DateIngredientQuantites,
  TimedIngredientQuantite
} from "../../logic/api";

const ResultIngredientsProps = Vue.extend({
  props: {
    loading: Boolean,
    dateIngredients: Array as () => DateIngredientQuantites[] | null,
    origineIngredients: Array as () => TimedIngredientQuantite[]
  }
});

@Component({
  components: { DateIngredients }
})
export default class ResultIngredients extends ResultIngredientsProps {}
</script>

<style scoped></style>
