<template>
  <v-expansion-panel>
    <v-expansion-panel-header class="py-0 px-1 my-0">
      <toolbar
        title="Ingredients disponibles"
        :showAdd="false"
        :elevation="2"
        v-model="search"
      ></toolbar>
    </v-expansion-panel-header>
    <v-expansion-panel-content>
      <v-list dense class="overflow-y-auto py-0" :style="{ height: height }">
        <v-list-item-group>
          <v-list-item
            v-for="ingredient in ingredients"
            :key="ingredient.id"
            :value="ingredient"
            :inactive="false"
          >
            <v-list-item-content
              draggable
              @dragstart="onDrag($event, ingredient)"
            >
              <v-list-item-title>{{ ingredient.nom }}</v-list-item-title>
              <v-list-item-subtitle>
                <i>{{ ingredient.unite }}</i>
              </v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-expansion-panel-content>
  </v-expansion-panel>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import { Controller } from "@/logic/controller";
import { Ingredient } from "@/logic/api";
import { DragKind, setDragData } from "../../../utils/utils_drag";
import Toolbar from "../../../utils/Toolbar.vue";
import { searchFunction } from "../../../utils/utils";

const ChoixIngredientsProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    height: String,
  },
});

@Component({
  components: { Toolbar },
})
export default class ChoixIngredients extends ChoixIngredientsProps {
  search = "";

  get ingredients() {
    const sf = searchFunction(this.search);
    return Object.values(this.C.api.ingredients).filter((ingredient) =>
      sf(ingredient.nom)
    );
  }

  onDrag(event: DragEvent, ingredient: Ingredient) {
    if (event == null || event.dataTransfer == null) return;
    setDragData(event.dataTransfer, DragKind.IdIngredient, ingredient.id);
    event.dataTransfer.effectAllowed = "copyLink";
  }
}
</script>

<style scoped></style>
