<template>
  <v-row class="fill-height" justify="space-around">
    <v-col cols="auto" class="align-self-center">
      <form-calcul :sejour="sejour" @change="onChange"></form-calcul>
    </v-col>
    <v-col cols="6" class="align-self-center">
      <result-ingredients
        :loading="loadingIngredients"
        :dateIngredients="dateIngredients"
      >
      </result-ingredients>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import FormCalcul from "../components/recap_ingredients/FormCalcul.vue";
import ResultIngredients from "../components/recap_ingredients/ResultIngredients.vue";

import { C } from "../logic/controller";
import { DateIngredientQuantites, OutResoudIngredients } from "../logic/types";

const RecapIngredientsProps = Vue.extend({
  props: {}
});

@Component({
  components: { FormCalcul, ResultIngredients }
})
export default class RecapIngredients extends RecapIngredientsProps {
  showFormCalcul = false;

  loadingIngredients = false;
  dateIngredients: DateIngredientQuantites[] = [];

  // pour pouvoir raffraichir la requête
  critere: number[] = [];

  get sejour() {
    return C.state.getSejour();
  }

  onChange(critere: number[]) {
    this.critere = critere;
    this.calcul();
  }

  async calcul() {
    if (this.sejour == null) return;

    this.loadingIngredients = true;
    const res = await C.calculs.resoudIngredientsJournees(
      this.sejour.id,
      this.critere
    );
    this.loadingIngredients = false;
    if (!res) {
      return;
    }
    this.dateIngredients = res.date_ingredients || [];
  }

  activated() {
    // les repas on put changer,
    // donc on relance une requête de calcul
    this.calcul();
  }
}
</script>

<style scoped></style>
