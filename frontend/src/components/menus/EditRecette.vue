<template>
  <div>
    <v-toolbar color="secondary" dense class="my-1">
      <v-toolbar-title>{{ title }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          mdi-icon="check-outline"
          tooltip="Valider"
          color="success"
          @click="$emit('done', recette)"
        ></tooltip-btn>
        <tooltip-btn
          mdi-icon="undo"
          tooltip="Abandonner les modifications"
          @click="$emit('undo')"
        ></tooltip-btn>
      </v-toolbar-items>
    </v-toolbar>
    <v-row class="px-2">
      <v-col>
        <v-text-field
          label="Nom de la recette"
          v-model="recette.nom"
          required
        ></v-text-field>
        <v-textarea
          class=""
          label="Mode d'emploi"
          v-model="recette.mode_emploi"
          hide-details
          rows="2"
        ></v-textarea>
      </v-col>
      <v-col>
        <liste-lien-ingredients
          style="max-height: 60vh;"
          class="mt-0 pt-0 overflow-y-auto"
          subheader="Ingrédients"
          v-model="recette.ingredients"
        ></liste-lien-ingredients>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import TooltipBtn from "../utils/TooltipBtn.vue";

import { C } from "../../logic/controller";
import { Recette, RecetteComplet, LienIngredient } from "../../logic/types";
import { New, EditMode, deepcopy } from "../../logic/types2";
import { Watch } from "vue-property-decorator";

import ListeLienIngredients from "../utils/ListeLienIngredients.vue";

const EditRecetteProps = Vue.extend({
  props: {
    mode: String as () => EditMode,
    initialRecette: Object as () => RecetteComplet
  }
});

@Component({
  components: { TooltipBtn, ListeLienIngredients }
})
export default class EditRecette extends EditRecetteProps {
  // recette actuellement éditée
  recette: RecetteComplet = deepcopy(this.initialRecette);

  @Watch("initialRecette")
  onRecetteChange() {
    this.recette = deepcopy(this.initialRecette);
  }

  get title() {
    if (this.mode == "edit") {
      return "Modifier la recette";
    }
    return "Ajouter une recette";
  }

  formatRecetteProprietaire = C.formatter.formatMenuOrRecetteProprietaire;
}
</script>

<style scoped></style>
