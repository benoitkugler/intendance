<template>
  <div>
    <v-toolbar color="secondary" dense>
      <v-toolbar-title>Recettes {{ bonusTitle }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          mdi-icon="refresh"
          tooltip="Afficher toutes les recettes"
          @click="$emit('reset')"
        ></tooltip-btn>
      </v-toolbar-items>
    </v-toolbar>
    <v-list dense :max-height="height" class="overflow-y-auto">
      <v-list-item-group v-model="currentRecette">
        <v-list-item
          v-for="recette in recettes"
          :key="recette.id"
          :value="recette"
        >
          <v-list-item-content>
            <v-list-item-title>{{ recette.nom }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";

import TooltipBtn from "../utils/TooltipBtn.vue";

import { D } from "../../logic/controller";
import { Recette } from "../../logic/types";

const Props = Vue.extend({
  props: {
    height: String,
    recettes: Array as () => Recette[],
    bonusTitle: {
      type: String,
      default: ""
    }
  }
});

@Component({
  components: { TooltipBtn }
})
export default class ListeRecettes extends Props {
  currentRecette: Recette | null = null;

  @Watch("currentRecette")
  onChange() {
    this.$emit("change", this.currentRecette);
  }

  clearSelection() {
    this.currentRecette = null;
  }
}
</script>
