<template>
  <v-card>
    <v-card-text>
      <v-row>
        <v-col>
          <livraison-ingredients :livraison="null"></livraison-ingredients>
        </v-col>
        <v-col v-for="livraison in livraisons" :key="livraison.id">
          <livraison-ingredients :livraison="livraison"></livraison-ingredients>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import {
  CommandeSimpleContraintes,
  DateIngredientQuantites,
  Ingredient
} from "@/logic/api";
import { Controller } from "@/logic/controller";
import Vue from "vue";
import Component from "vue-class-component";
import LivraisonIngredients from "./LivraisonIngredients.vue";

const AssocieLivraisonsProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    dateIngredients: Array as () => DateIngredientQuantites[]
  }
});

@Component({
  components: { LivraisonIngredients }
})
export default class AssocieLivraisons extends AssocieLivraisonsProps {
  get livraisons() {
    return Object.values(this.C.api.livraisons);
  }

  get ingredients() {
    const tmp: { [key: number]: boolean } = {};
    this.dateIngredients.forEach(ings => {
      (ings.ingredients || []).forEach(ing => {
        tmp[ing.ingredient.id] = true;
      });
    });
    return Object.keys(tmp).map(idIng => this.C.api.ingredients[Number(idIng)]);
  }

  valide() {
    // TODO: construire out
    const out: { [key: number]: number } = {};
    this.$emit("valide", out);
  }
}
</script>

<style scoped></style>
