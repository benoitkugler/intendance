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
      :showAdd="state.mode == 'visu' && state.selection.idMenu == null"
      @add="$emit('new')"
    ></toolbar>
    <v-list dense :max-height="height" class="overflow-y-auto" ref="list">
      <v-list-item-group
        :value="state.selection.idRecette"
        @change="args => $emit('change', args)"
      >
        <v-list-item
          v-for="recette in recettes"
          :key="recette.id"
          :value="recette.id"
          :class="classItem(recette.id)"
        >
          <template v-slot:default="props">
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
            <v-list-item-action v-if="showActions(props.active, recette)">
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

import { Controller } from "@/logic/controller";
import { Recette } from "@/logic/api";
import { StateMenus } from "./types";
import levenshtein from "js-levenshtein";
import { searchFunction } from "../utils/utils";
import { DragKind, setDragData } from "../utils/utils_drag";
import { ListKind, BaseList } from "./shared";

@Component({
  components: { TooltipBtn, Toolbar },
  props: {
    kind: {
      type: String as () => ListKind,
      default: "idRecette"
    }
  }
})
export default class ListeRecettes extends BaseList {
  confirmeSupprime = false;

  search = "";

  private searchRecettes(recettes: Recette[], search: string) {
    const predicat = searchFunction(search);
    return recettes.filter(recette => predicat(recette.nom));
  }

  get recettes() {
    let baseRecettes: Recette[];
    if (this.state.mode == "editMenu") {
      baseRecettes = Object.values(this.C.api.recettes);
    } else if (this.state.selection.idMenu != null) {
      baseRecettes = this.C.getMenuRecettes(
        this.C.getMenu(this.state.selection.idMenu)
      );
    } else {
      baseRecettes = Object.values(this.C.api.recettes);
    }
    return this.searchRecettes(baseRecettes, this.search);
  }

  get title() {
    if (this.state.mode == "editMenu") {
      return "Choisir une recette";
    }
    if (this.state.selection.idMenu != null) {
      return "Recettes liées au menu";
    }
    return "Toutes les recettes";
  }
  formatRecetteProprietaire = this.C.formatter.formatMenuOrRecetteProprietaire;

  showActions(active: boolean, recette: Recette) {
    if (this.state.selection.idMenu != null) return false;
    return (
      active &&
      (!recette.id_utilisateur.Valid ||
        recette.id_utilisateur.Int64 == this.C.state.idUtilisateur)
    );
  }

  async supprime() {
    this.confirmeSupprime = false;
    if (this.state.selection.idRecette == null) return;
    await this.C.api.DeleteRecette({ id: this.state.selection.idRecette });
    this.$emit("change", null);
  }

  onDragStart(event: DragEvent, recette: Recette) {
    if (!event.dataTransfer) return;
    setDragData(event.dataTransfer, DragKind.IdRecette, recette.id);
    event.dataTransfer.effectAllowed = "copy";
  }
}
</script>
