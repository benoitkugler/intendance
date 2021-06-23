<template>
  <div
    @dragover="onDragoverIngredients"
    @drop="onDropIngredient"
    style="height: 100%"
  >
    <div v-html="title"></div>
    <v-list>
      <v-list-item
        v-for="ingredient in ingredients"
        :key="ingredient.id"
        :draggable="true"
        @dragstart="(ev) => onDragStart(ev, ingredient)"
      >
        {{ ingredient.nom }}
      </v-list-item>
    </v-list>
  </div>
</template>

<script lang="ts">
import { Ids, Ingredient, Ingredients, Livraison } from "@/logic/api";
import { Controller } from "@/logic/controller";
import Vue from "vue";
import Component from "vue-class-component";
import { DragKind, getDragData, setDragData } from "../utils/utils_drag";

const LivraisonIngredientsProps = Vue.extend({
  props: {
    livraison: Object as () => Livraison | undefined, // undefined pour les ingrédient non associés
    ingredients: Array as () => Ingredient[],
    C: Object as () => Controller,
  },
  model: {
    prop: "ingredients",
    event: "change",
  },
});

@Component({})
export default class LivraisonIngredients extends LivraisonIngredientsProps {
  get title() {
    if (this.livraison === undefined) {
      return "Ingrédients à attribuer";
    }
    return this.C.formatter.formatLivraison(this.livraison);
  }

  onDragStart(event: DragEvent, ingredient: Ingredient) {
    if (!event.dataTransfer) return;
    setDragData(event.dataTransfer, DragKind.IdIngredient, ingredient.id);
    event.dataTransfer.effectAllowed = "move";
  }

  onDragoverIngredients(event: DragEvent) {
    if (!event.dataTransfer) return;
    const isIngredient = event.dataTransfer.types.includes(
      DragKind.IdIngredient
    );
    if (isIngredient) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "move";
    }
  }

  onDropIngredient(event: DragEvent) {
    if (!event.dataTransfer) return;
    event.preventDefault();
    const idIngredient = getDragData(event.dataTransfer, DragKind.IdIngredient);
    this.$emit("swap-ingredient", idIngredient);
  }
}
</script>

<style scoped></style>
