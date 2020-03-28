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
          @click="$emit('done', menu)"
        ></tooltip-btn>
        <tooltip-btn
          mdi-icon="undo"
          tooltip="Abandonner les modifications"
          @click="$emit('undo')"
        ></tooltip-btn>
      </v-toolbar-items>
    </v-toolbar>
    <v-textarea
      class="pa-2 mt-4 my-3 mb-0"
      label="Commentaire"
      v-model="menu.commentaire"
      hide-details
      rows="2"
    ></v-textarea>

    <div @dragover="onDragoverRecettes" @drop="onDropRecette">
      <v-list dense max-height="60vh" class="mt-0 pt-0 overflow-y-auto">
        <v-subheader>Recettes du menu</v-subheader>
        <v-list-item v-if="recettes.length == 0">
          <v-list-item-subtitle>
            <i>Cliquer-déplacer pour ajouter une recette...</i>
          </v-list-item-subtitle>
        </v-list-item>
        <v-list-item v-for="recette in recettes" :key="recette.id">
          <template v-slot:default="{}">
            <v-list-item-content>
              <v-list-item-title>
                <v-list-item-title>{{ recette.nom }}</v-list-item-title>
              </v-list-item-title>
              <v-list-item-subtitle>
                <i> {{ formatRecetteProprietaire(recette) }}</i>
              </v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action>
              <tooltip-btn
                mdi-icon="close"
                tooltip="Retirer cette recette du menu"
                color="red"
                @click.stop="removeRecette(recette)"
              ></tooltip-btn>
            </v-list-item-action>
          </template>
        </v-list-item>
      </v-list>
    </div>
    <liste-lien-ingredients
      subheader="Ingrédients additionnels"
      v-model="menu.ingredients"
    ></liste-lien-ingredients>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import TooltipBtn from "../utils/TooltipBtn.vue";

import { C } from "../../logic/controller";
import { Menu, Recette, LienIngredient, MenuComplet } from "../../logic/types";
import { New, EditMode, deepcopy } from "../../logic/types2";
import { Watch } from "vue-property-decorator";
import ListeLienIngredients from "../utils/ListeLienIngredients.vue";

const EditMenuProps = Vue.extend({
  props: {
    mode: String as () => EditMode,
    initialMenu: Object as () => New<MenuComplet>
  }
});

@Component({
  components: { TooltipBtn, ListeLienIngredients }
})
export default class EditMenu extends EditMenuProps {
  // menu actuellement édité
  menu: New<MenuComplet> = deepcopy(this.initialMenu);

  @Watch("initialMenu")
  onMenuChange() {
    this.menu = deepcopy(this.initialMenu);
  }

  get title() {
    if (this.mode == "edit") {
      return "Modifier le menu";
    }
    return "Ajouter un menu";
  }

  // résoud les recettes à partir des ids
  get recettes() {
    return C.getMenuRecettes(this.menu);
  }

  formatRecetteProprietaire = C.formatter.formatMenuOrRecetteProprietaire;

  removeRecette(toRemove: Recette) {
    this.menu.recettes = (this.menu.recettes || []).filter(
      idRecette => idRecette != toRemove.id
    );
  }

  onDragoverRecettes(event: DragEvent) {
    if (!event.dataTransfer) return;
    const isRecette = event.dataTransfer.types.includes("id-recette");
    if (isRecette) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "copy";
    }
  }

  onDropRecette(event: DragEvent) {
    if (!event.dataTransfer) return;
    event.preventDefault();
    const idRecette = Number(event.dataTransfer.getData("id-recette"));
    const recettes = this.menu.recettes || [];
    const hasRecette = recettes.includes(idRecette);
    if (hasRecette) return;
    recettes.push(idRecette);
    this.menu.recettes = recettes;
  }
}
</script>

<style scoped></style>
