<template>
  <v-expansion-panel>
    <v-expansion-panel-header class="py-0 px-1 my-0">
      <toolbar
        title="Recettes disponibles"
        :showAdd="false"
        :elevation="2"
        v-model="search"
      ></toolbar>
    </v-expansion-panel-header>
    <v-expansion-panel-content>
      <v-list dense class="overflow-y-auto py-0" :style="{ height: height }">
        <v-list-item-group>
          <v-list-item
            v-for="recette in recettes"
            :key="recette.id"
            :value="recette"
            :inactive="false"
          >
            <v-list-item-content draggable @dragstart="onDrag($event, recette)">
              <v-list-item-title>{{ recette.nom }}</v-list-item-title>
              <v-list-item-subtitle>
                <i>{{ formatRecetteProprietaire(recette) }}</i>
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
import { Recette, RecetteComplet } from "@/logic/api";
import { DragKind, setDragData } from "../../../utils/utils_drag";
import Toolbar from "../../../utils/Toolbar.vue";
import { searchFunction } from "../../../utils/utils";

const ChoixRecettesProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    height: String,
  },
});

@Component({
  components: { Toolbar },
})
export default class ChoixRecettes extends ChoixRecettesProps {
  search = "";

  formatRecetteProprietaire = this.C.formatter.formatMenuOrRecetteProprietaire;

  get recettes() {
    const sf = searchFunction(this.search);
    return Object.values(this.C.api.recettes).filter((recette) =>
      sf(recette.nom)
    );
  }

  onDrag(event: DragEvent, recette: RecetteComplet) {
    if (event == null || event.dataTransfer == null) return;
    setDragData(event.dataTransfer, DragKind.IdRecette, recette.id);
    event.dataTransfer.effectAllowed = "copyLink";
  }
}
</script>

<style scoped></style>
