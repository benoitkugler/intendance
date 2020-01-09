<template>
  <v-card>
    <v-card-title primary-title>
      Produits associés à l'ingrédient {{ ingredient.nom }}
    </v-card-title>
    <v-card-text> </v-card-text>
    <v-card-actions>
      <v-btn flat color="primary">text</v-btn>
      <v-btn flat color="primary">text</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Ingredient } from "../../logic/types";

const IngredientProduitsProps = Vue.extend({
  props: {
    ingredient: Object as () => Ingredient
  }
});

@Component({})
export default class IngredientProduits extends IngredientProduitsProps {}
</script>

<style scoped></style>
