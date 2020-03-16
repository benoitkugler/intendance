<template>
  <v-card>
    <v-card-title primary-title>
      Produits associés à l'ingrédient {{ ingredient.nom }}
    </v-card-title>
    <v-card-subtitle>
      Vous pouvez choisir ici quels produits concrets sont associés à
      l'ingrédient générique.
    </v-card-subtitle>
    <v-card-text>
      <v-skeleton-loader type="paragraph" :loading="loading">
        {{ ingredientProduits }}
      </v-skeleton-loader>
    </v-card-text>
    <v-card-actions>
      <v-btn text color="primary">text</v-btn>
      <v-btn text color="primary">text</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Ingredient, IngredientProduits } from "../../logic/types";
import { C } from "../../logic/controller";
import { Watch } from "vue-property-decorator";

const AssociationIngredientProps = Vue.extend({
  props: {
    ingredient: Object as () => Ingredient,
    activated: Boolean
  }
});

@Component({})
export default class AssociationIngredient extends AssociationIngredientProps {
  get loading() {
    return C.notifications.getSpin();
  }

  ingredientProduits: IngredientProduits | null = null;

  private async loadProduits() {
    const res = await C.data.getIngredientProduits(this.ingredient.id);
    if (res == undefined) return;
    this.ingredientProduits = res;
  }

  mounted() {
    this.loadProduits();
  }

  @Watch("activated")
  onChangeIngredient(b: boolean) {
    if (!b) return;
    this.loadProduits();
  }
}
</script>

<style scoped></style>
