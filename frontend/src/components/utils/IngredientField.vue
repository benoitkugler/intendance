<template>
  <v-autocomplete
    color="success"
    label="Ajouter un ingrédient"
    :items="items"
    :filter="filter"
    v-model="id"
    @change="$emit('change', id)"
    @blur="$emit('change', id)"
  >
  </v-autocomplete>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Controller } from "@/logic/controller";
import { searchFunction } from "./utils";
import { EnumItem } from "@/logic/types";

const IngredientFieldProps = Vue.extend({
  props: {
    C: Object as () => Controller
  }
});

@Component({})
export default class IngredientField extends IngredientFieldProps {
  id: number = -1;

  searchFunctionCache: { [key: string]: (_: string) => boolean } = {};

  filter(ingredient: EnumItem<number>, search: string, _: string) {
    let sf = this.searchFunctionCache[search];
    if (sf === undefined) {
      sf = searchFunction(search);
      // mise en cache
      this.searchFunctionCache[search] = sf;
    }
    return sf(ingredient.text);
  }

  get items() {
    return this.C.getAllIngredients().map(ing => {
      return { text: ing.ingredient.nom, value: ing.ingredient.id };
    });
  }
}
</script>

<style scoped></style>
