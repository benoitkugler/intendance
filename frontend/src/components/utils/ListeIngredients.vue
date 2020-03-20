<template>
  <v-simple-table dense fixed-header>
    <thead>
      <tr>
        <th class="text-left">Ingrédient</th>
        <th class="text-center">Quantité</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="ingredient in ingredients" :key="ingredient.ingredient.id">
        <td>
          <v-tooltip left>
            <template v-slot:activator="{ on }">
              <a @click="$emit('go', ingredient.ingredient.id)" v-on="on">
                {{ ingredient.ingredient.nom }}
              </a>
            </template>
            Aller à l'ingrédient...
          </v-tooltip>
        </td>
        <td class="text-center">
          {{ formatQuantite(ingredient.quantite) }}
          <i class="grey--text">{{ ingredient.ingredient.unite }}</i>
        </td>
      </tr>
    </tbody>
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
