<template>
  <v-card>
    <v-card-text>
      <v-btn @click="fetchHints">Hint</v-btn>
      <v-row>
        <v-col>
          <livraison-ingredients
            :C="C"
            :livraison="undefined"
            :ingredients="getIngredientsByLivraison(undefined)"
          ></livraison-ingredients>
        </v-col>
        <v-col v-for="livraison in livraisons" :key="livraison.id">
          <livraison-ingredients
            :C="C"
            :livraison="livraison"
            :ingredients="getIngredientsByLivraison(livraison.id)"
          ></livraison-ingredients>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import {
  CommandeContraintes,
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
  loading = true;

  private associations: { [key: number]: number } = {};

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

  getIngredientsByLivraison(idLivraison: number | undefined) {
    return this.ingredients.filter(
      ing => this.associations[ing.id] === idLivraison
    );
  }

  async fetchHints() {
    const data = await this.C.api.ProposeLienIngredientLivraison(
      this.dateIngredients
    );
    this.loading = false;
    if (data === undefined) return;
    for (const id in data || {}) {
      // merge into current
      Vue.set(this.associations, id, (data || {})[id]); //VRC
    }
  }

  valide() {
    // TODO: construire out
    const out: { [key: number]: number } = {};
    this.$emit("valide", out);
  }
}
</script>

<style scoped></style>
