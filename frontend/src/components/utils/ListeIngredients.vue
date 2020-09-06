<template>
  <v-simple-table dense fixed-header>
    <thead>
      <tr>
        <th class="text-left">Ingrédient</th>
        <th class="text-center">Quantité</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="ingredient in ingredients"
        :key="ingredient.ingredient.id"
        :class="getRowClass(ingredient)"
      >
        <td>
          <v-tooltip left v-if="!hideLinks">
            <template v-slot:activator="props">
              <a @click="$emit('go', ingredient.ingredient.id)" v-on="props.on">
                {{ ingredient.ingredient.nom }}
              </a>
            </template>
            Aller à l'ingrédient...
          </v-tooltip>
          <span v-else>{{ ingredient.ingredient.nom }}</span>
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
import { IngredientQuantite } from "@/logic/api";
import { Formatter } from "@/logic/formatter";
import { Crible } from "./utils";

const ListeIngredientsProps = Vue.extend({
  props: {
    ingredients: Array as () => IngredientQuantite[],
    highlights: Object as () => Crible | undefined, // ids
    hideLinks: Boolean
  }
});
@Component({})
export default class ListeIngredients extends ListeIngredientsProps {
  formatQuantite = Formatter.formatQuantite;

  getRowClass(ingredient: IngredientQuantite) {
    if (this.highlights && this.highlights[ingredient.ingredient.id]) {
      return "lime lighten-3";
    }
    return "";
  }
}
</script>

<style scoped></style>
