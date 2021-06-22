<template>
  <v-autocomplete
    label="Recettes"
    :value="recettes"
    @change="(v) => $emit('change', v)"
    chips
    multiple
    :items="items"
    :filter="filter"
  >
  </v-autocomplete>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Controller } from "@/logic/controller";
import { searchFunction } from "./utils";
import { EnumItem } from "@/logic/types";
const RecettesFieldsProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    recettes: Array as () => number[],
  },
  model: {
    prop: "recettes",
    event: "change",
  },
});

@Component({})
export default class RecettesFields extends RecettesFieldsProps {
  searchFunctionCache: { [key: string]: (_: string) => boolean } = {};

  filter(recette: EnumItem<number>, search: string, _: string) {
    let sf = this.searchFunctionCache[search];
    if (sf === undefined) {
      sf = searchFunction(search);
      // mise en cache
      this.searchFunctionCache[search] = sf;
    }
    return sf(recette.text);
  }

  get items() {
    return Object.values(this.C.api.recettes).map((rec) => {
      return { text: rec.nom, value: rec.id };
    });
  }
}
</script>

<style scoped></style>
