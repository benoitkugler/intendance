<template>
  <v-autocomplete
    label="Recettes"
    :value="idsRecettes"
    @change="v => $emit('change', v)"
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
import { C } from "../../logic/controller";
import { EnumItem } from "../../logic/enums";
import { searchFunction } from "./utils";
const RecettesFieldsProps = Vue.extend({
  props: {
    idsRecettes: Array as () => number[]
  },
  model: {
    prop: "idsRecettes",
    event: "change"
  }
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
    return Object.values(C.data.recettes || {}).map(rec => {
      return { text: rec.nom, value: rec.id };
    });
  }
}
</script>

<style scoped></style>
