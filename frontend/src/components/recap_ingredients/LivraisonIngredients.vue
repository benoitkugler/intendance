<template>
  <div>
    {{ title }}
    <v-list>
      <v-list-group>
        <v-list-item
          v-for="ingredient in ingredients"
          :key="ingredient.id"
          draggable="true"
          @dragstart="onDragStart(ingredient)"
        >
          {{ ingredient.nom }}
        </v-list-item>
      </v-list-group>
    </v-list>
  </div>
</template>

<script lang="ts">
import { Ids, Ingredient, Ingredients, Livraison } from "@/logic/api";
import { Controller } from "@/logic/controller";
import Vue from "vue";
import Component from "vue-class-component";

const LivraisonIngredientsProps = Vue.extend({
  props: {
    livraison: Object as () => Livraison | undefined, // undefined pour les ingrédient non associés
    ingredients: Array as () => Ingredient[],
    C: Object as () => Controller
  },
  model: {
    prop: "ingredients",
    event: "change"
  }
});

@Component({})
export default class LivraisonIngredients extends LivraisonIngredientsProps {
  get title() {
    if (this.livraison === undefined) {
      return "A attribuer";
    }
    return this.C.formatter.formatLivraison(this.livraison);
  }

  onDragStart(ingredient: Ingredient) {
    console.log(ingredient);
  }
}
</script>

<style scoped></style>
