<template>
  <v-autocomplete
    hide-no-data
    no-filter
    label="Produit"
    placeholder="Taper pour rechercher un autre produit..."
    prepend-icon="mdi-cart"
    :items="items"
    :value="idProduit"
    @change="(v) => $emit('change', v)"
    :loading="isLoading"
    :search-input.sync="search"
  ></v-autocomplete>
</template>

<script lang="ts">
import { Produit } from "@/logic/api";
import { Controller } from "@/logic/controller";
import Vue from "vue";
import Component from "vue-class-component";
import { Watch } from "vue-property-decorator";

const ChoixProduitProps = Vue.extend({
  model: {
    prop: "idProduit",
    event: "change",
  },
  props: {
    idProduit: Number,
    C: Object as () => Controller,
    hints: Array as () => Produit[] | undefined,
  },
});

type item = { text: string; value: number } | { divider: boolean };

@Component({
  components: {},
})
export default class ChoixProduit extends ChoixProduitProps {
  isLoading = false;
  search: string | null = null;

  searchResults: Produit[] = [];

  get items() {
    const hints: item[] = (this.hints || []).map((prod) => ({
      value: prod.id,
      text: prod.nom,
    }));
    const searchResults: item[] = this.searchResults.map((prod) => ({
      value: prod.id,
      text: prod.nom,
    }));
    if (hints.length != 0 && searchResults.length != 0) {
      return hints.concat([{ divider: true }]).concat(searchResults);
    }
    return hints.concat(...searchResults);
  }

  @Watch("search")
  async _(val: string) {
    // Items have already been loaded
    if (this.idProduit != null) return;

    // Items have already been requested
    if (this.isLoading) return;

    this.isLoading = true;

    // Lazily load input items
    const res = await this.C.api.RechercheProduit({ recherche: val });
    this.isLoading = false;
    if (res === undefined) {
      return;
    }

    this.searchResults = res || [];
  }
}
</script>

<style scoped></style>
