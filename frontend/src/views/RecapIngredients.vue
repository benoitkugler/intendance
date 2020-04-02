<template>
  <v-row class="fill-height px-2">
    <v-col md="3" sm="6" class="align-self-center">
      <form-calcul :sejour="sejour" @change="onChange"></form-calcul>
    </v-col>
    <v-col md="4" sm="6" class="align-self-center">
      <result-ingredients
        :loading="loadingIngredients"
        :dateIngredients="dateIngredients"
        :origineIngredients="origineIngredients"
        @goToIngredient="goToIngredient"
      >
      </result-ingredients>
    </v-col>
    <v-col class="align-self-center" md="5" sm="12">
      <preview-commande
        :dateIngredients="dateIngredients"
        @showOrigines="o => (origineIngredients = o)"
      ></preview-commande>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import FormCalcul from "../components/recap_ingredients/FormCalcul.vue";
import ResultIngredients from "../components/recap_ingredients/ResultIngredients.vue";
import PreviewCommande from "../components/recap_ingredients/PreviewCommande.vue";

import { C } from "../logic/controller";
import {
  DateIngredientQuantites,
  OutResoudIngredients,
  TimedIngredientQuantite
} from "../logic/types";

const RecapIngredientsProps = Vue.extend({
  props: {}
});

@Component({
  components: { FormCalcul, ResultIngredients, PreviewCommande }
})
export default class RecapIngredients extends RecapIngredientsProps {
  showFormCalcul = false;

  loadingIngredients = false;
  dateIngredients: DateIngredientQuantites[] = [];

  origineIngredients: TimedIngredientQuantite[] = [];

  // pour pouvoir raffraichir la requête
  critere: number[] = [];

  get sejour() {
    return C.state.getSejour();
  }

  async mounted() {
    let ps = [];
    if (Object.keys(C.data.ingredients || {}).length == 0) {
      ps.push(C.data.loadIngredients);
    }
    if (Object.keys(C.data.fournisseurs || {}).length == 0) {
      ps.push(C.data.loadFournisseurs);
    }
    await Promise.all(ps);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Fournisseurs et ingrédients chargés.");
    }
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

  goToIngredient(id: number) {
    this.$router.push({ name: "menus", query: { idIngredient: String(id) } });
  }
}
</script>

<style scoped></style>
