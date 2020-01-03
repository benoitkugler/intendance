<template>
  <v-simple-table dense fixed-header>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-left">Ingrédient</th>
          <th class="text-center">Quantité</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="ingredient in ingredients" :key="ingredient.ingredient.id">
          <td>{{ ingredient.ingredient.nom }}</td>
          <td class="text-center">
            {{ formatQuantite(ingredient.quantite) }}
            <i class="grey--text">{{ ingredient.ingredient.unite }}</i>
          </td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { IngredientQuantite } from "../../logic/types";
import { Formatter } from "../../logic/formatter";

const ListeIngredientsProps = Vue.extend({
  props: {
    ingredients: Array as () => IngredientQuantite[]
  }
});
@Component({})
export default class ListeIngredients extends ListeIngredientsProps {
  formatQuantite = Formatter.formatQuantite;
}
</script>

<style scoped></style>
