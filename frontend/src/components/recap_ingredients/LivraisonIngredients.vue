<template>
  <div>{{ livraison }}</div>
</template>

<script lang="ts">
import { Ids, Livraison } from "@/logic/api";
import Vue from "vue";
import Component from "vue-class-component";

const LivraisonIngredientsProps = Vue.extend({
  props: {
    livraison: Object as () => Livraison | null, // null pour les ingrédient non associés
    ingredients: Array as () => Ids
  },
  model: {
    prop: "ingredients",
    event: "change"
  }
});

@Component({})
export default class LivraisonIngredients extends LivraisonIngredientsProps {}
</script>

<style scoped></style>
