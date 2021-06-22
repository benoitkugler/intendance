<template>
  <v-row class="fill-height px-2">
    <v-col md="3" sm="6" class="align-self-center">
      <form-calcul :C="C" :sejour="sejour" @change="onChange"></form-calcul>
    </v-col>
    <v-col md="4" sm="6" class="align-self-center">
      <result-ingredients
        :loading="loadingIngredients"
        :dateIngredients="dateIngredients"
        :origineIngredients="origineIngredients"
        @go-to-ingredient="goToIngredient"
      >
      </result-ingredients>
    </v-col>
    <v-col class="align-self-center" md="5" sm="12">
      <associe-livraisons
        :C="C"
        :dateIngredients="dateIngredients"
      ></associe-livraisons>
    </v-col>
    <!-- <v-col class="align-self-center" md="5" sm="12">
      <preview-commande
        :C="C"
        :dateIngredients="dateIngredients"
        @showOrigines="o => (origineIngredients = o)"
      ></preview-commande>
    </v-col> -->
  </v-row>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import FormCalcul from "../components/recap_ingredients/FormCalcul.vue";
import ResultIngredients from "../components/recap_ingredients/ResultIngredients.vue";
import PreviewCommande from "../components/recap_ingredients/PreviewCommande.vue";

import { Controller } from "../logic/controller";
import { DateIngredientQuantites, TimedIngredientQuantite } from "../logic/api";
import AssocieLivraisons from "@/components/recap_ingredients/AssocieLivraisons.vue";

const RecapIngredientsProps = Vue.extend({
  props: {
    C: Object as () => Controller,
  },
});

@Component({
  components: {
    FormCalcul,
    ResultIngredients,
    PreviewCommande,
    AssocieLivraisons,
  },
})
export default class RecapIngredients extends RecapIngredientsProps {
  showFormCalcul = false;
  loadingIngredients = false;
  dateIngredients: DateIngredientQuantites[] = [];

  origineIngredients: TimedIngredientQuantite[] = [];

  // pour pouvoir raffraichir la requête
  critere: number[] = [];

  get sejour() {
    return this.C.getSejour();
  }

  async mounted() {
    if (Object.keys(this.C.api.ingredients || {}).length == 0) {
      await this.C.api.GetIngredients();
    }
    if (Object.keys(this.C.api.fournisseurs || {}).length == 0) {
      await this.C.api.GetFournisseurs();
    }
  }

  onChange(critere: number[]) {
    this.critere = critere;
    this.calcul();
  }

  async calcul() {
    if (this.sejour == null) return;

    this.loadingIngredients = true;
    const res = await this.C.resoudIngredientsJournees(
      this.sejour.id,
      this.critere
    );
    this.loadingIngredients = false;
    if (!res) {
      return;
    }
    this.dateIngredients = res;
  }

  goToIngredient(id: number) {
    this.$router.push({ name: "menus", query: { idIngredient: String(id) } });
  }
}
</script>

<style scoped></style>
