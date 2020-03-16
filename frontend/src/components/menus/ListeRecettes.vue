<template>
  <div>
    <v-dialog v-model="confirmeSupprime" max-width="500px">
      <v-card>
        <v-card-title primary-title color="warning">
          Confirmer la suppression
        </v-card-title>
        <v-card-text>
          Confirmez-vous la suppression de cette recette ? <br />
          <div class="my-3">
            <small>
              Si vous souhaitez seulement l'enlever d'un menu, passez plutôt en
              <b>mode édition</b>.
            </small>
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn tile color="warning" @click="supprime">Supprimer</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <toolbar
      v-model="search"
      tooltipAdd="Ajouter une recette..."
      :title="title"
      :showAdd="state.mode == 'visu' && state.selection.menu == null"
      @add="$emit('new')"
    ></toolbar>
    <v-list dense :max-height="height" class="overflow-y-auto">
      <v-list-item-group
        :value="state.selection.recette"
        @change="args => $emit('change', args)"
      >
        <v-list-item
          v-for="recette in recettes"
          :key="recette.id"
          :value="recette"
        >
          <template v-slot:default="{ active }">
            <v-list-item-content
              draggable="true"
              @dragstart="ev => onDragStart(ev, recette)"
            >
              <v-list-item-title>
                <v-list-item-title>{{ recette.nom }}</v-list-item-title>
              </v-list-item-title>
              <v-list-item-subtitle>
                <i> {{ formatRecetteProprietaire(recette) }}</i>
              </v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action v-if="showActions(active, recette)">
              <v-row no-gutters>
                <v-col
                  ><tooltip-btn
                    mdi-icon="pencil"
                    tooltip="Modifier cette recette"
                    color="secondary"
                    @click.stop="$emit('edit', recette)"
                  ></tooltip-btn
                ></v-col>
                <v-col>
                  <tooltip-btn
                    mdi-icon="close"
                    tooltip="Supprimer cette recette"
                    color="red"
                    @click.stop="confirmeSupprime = true"
                  ></tooltip-btn>
                </v-col>
              </v-row>
            </v-list-item-action>
          </template>
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
import Toolbar from "../utils/Toolbar.vue";

import { C } from "../../logic/controller";
import { Recette } from "../../logic/types";
import { StateMenus } from "./types";
import levenshtein from "js-levenshtein";
import { searchFunction } from "../utils/utils";
const MAX_DIST_LEVENSHTEIN = 5;

const Props = Vue.extend({
  props: {
    state: Object as () => StateMenus,
    height: String
  }
});

@Component({
  components: { TooltipBtn, Toolbar }
})
export default class ListeRecettes extends Props {
  confirmeSupprime = false;

  search = "";

  private searchRecettes(recettes: Recette[], search: string) {
    const predicat = searchFunction(search);
    return recettes.filter(recette => predicat(recette.nom));
  }

  get recettes() {
    let baseRecettes: Recette[];
    if (this.state.mode == "editMenu") {
      baseRecettes = Object.values(C.data.recettes);
    } else if (this.state.selection.menu != null) {
      baseRecettes = C.getMenuRecettes(this.state.selection.menu);
    } else {
      baseRecettes = Object.values(C.data.recettes);
    }
    return this.searchRecettes(baseRecettes, this.search);
  }

  get title() {
    if (this.state.mode == "editMenu") {
      return "Choisir une recette";
    }
    if (this.state.selection.menu != null) {
      return "Recettes liées au menu";
    }
    return "Toutes les recettes";
  }
  formatRecetteProprietaire = C.formatter.formatMenuOrRecetteProprietaire;

  showActions(active: boolean, recette: Recette) {
    if (this.state.selection.menu != null) return false;
    return (
      active &&
      (!recette.id_proprietaire.Valid ||
        recette.id_proprietaire.Int64 == C.idUtilisateur)
    );
  }

  async supprime() {
    this.confirmeSupprime = false;
    if (this.state.selection.recette == null) return;
    await C.data.deleteRecette(this.state.selection.recette);
    this.$emit("change", null);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Recette supprimée avec succès.");
    }
  }

  onDragStart(event: DragEvent, recette: Recette) {
    if (!event.dataTransfer) return;
    event.dataTransfer.setData("id-recette", String(recette.id));
    event.dataTransfer.effectAllowed = "copy";
  }
}
</script>
