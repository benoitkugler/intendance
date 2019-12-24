<template>
  <v-list dense :max-height="height" class="overflow-y-auto">
    <v-subheader>Ingrédients {{ bonusTitle }}</v-subheader>
    <v-list-item-group v-model="currentIngredient">
      <v-list-item
        v-for="ingredient in ingredients"
        :key="ingredient.ingredient.id"
        :value="ingredient.ingredient"
      >
        <v-list-item-content>
          <v-list-item-title>
            <v-row>
              <v-col>
                {{ ingredient.ingredient.nom }} ({{
                  ingredient.ingredient.unite
                }})
              </v-col>
              <v-spacer></v-spacer>
              <v-col v-if="ingredient.options"
                >{{ ingredient.options.quantite }}
                {{ ingredient.options.cuisson }}</v-col
              >
            </v-row>
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list-item-group>
  </v-list>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";
import { D } from "../../logic/controller";
import { Ingredient, RecetteIngredient } from "../../logic/types";
import { IngredientOptions } from "../../logic/types2";

const Props = Vue.extend({
  props: {
    height: String,
    ingredients: Array as () => IngredientOptions[],
    bonusTitle: {
      type: String,
      default: ""
    }
  }
});

@Component
export default class ListeIngredients extends Props {
  currentIngredient: Ingredient | null = null;

  @Watch("currentIngredient")
  onChange() {
    this.$emit("change", this.currentIngredient);
  }

  clearSelection() {
    this.currentIngredient = null;
  }
}
</script>
