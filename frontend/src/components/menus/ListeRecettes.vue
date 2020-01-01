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

    <v-toolbar color="secondary" class="my-1">
      <v-toolbar-title>
        <v-row no-gutters class="mt-1">
          <v-col>
            Recettes
          </v-col>
        </v-row>
        <v-row no-gutters>
          <v-col>
            <small>
              <i>{{ bonusTitle }}</i>
            </small>
          </v-col>
        </v-row>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          mdi-icon="magnify"
          tooltip="Filtrer les recettes..."
          @click="showSearch = !showSearch"
        />
        <tooltip-btn
          mdi-icon="plus-thick"
          color="green"
          tooltip="Ajouter une recette..."
          @click="$emit('new')"
          v-if="state.mode == 'visu' && state.selection.menu == null"
        />
      </v-toolbar-items>
    </v-toolbar>
    <v-text-field
      outlined
      label="Rechercher"
      placeholder="Tappez pour lancer la recherche"
      v-model="search"
      hide-details
      v-if="showSearch"
      class="my-2"
      ref="search"
    ></v-text-field>
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

import { C } from "../../logic/controller";
import { Recette } from "../../logic/types";
import { StateMenus } from "./types";
import levenshtein from "js-levenshtein";

const Props = Vue.extend({
  props: {
    state: Object as () => StateMenus,
    height: String
  }
});

const MAX_DIST_LEVENSHTEIN = 5;

@Component({
  components: { TooltipBtn }
})
export default class ListeRecettes extends Props {
  confirmeSupprime = false;

  search = "";
  showSearch = false;

  $refs!: {
    search: Vue;
  };

  private searchRecettes(recettes: Recette[]) {
    if (!this.search || !this.showSearch) return recettes;
    let filterNom: (nom: string) => boolean;
    try {
      const s = new RegExp(this.search, "i");
      filterNom = nom => s.test(nom);
    } catch {
      const sl = this.search.toLowerCase();
      filterNom = (nom: string) => nom.includes(sl);
    }
    return recettes.filter(ing => {
      const nom = ing.nom.toLowerCase();
      if (filterNom(nom)) return true;
      return levenshtein(nom, this.search) <= MAX_DIST_LEVENSHTEIN;
    });
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
    return this.searchRecettes(baseRecettes);
  }

  get bonusTitle() {
    if (this.state.mode == "editMenu") {
      return "Toutes";
    }
    if (this.state.selection.menu != null) {
      return "Menu courant";
    }
    return "";
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

  @Watch("showSearch")
  onShowSearch(b: boolean) {
    if (!b) return;
    setTimeout(() => {
      const input = this.$refs.search.$el.querySelector("input");
      if (input != null) input.select();
    }, 50);
  }
}
</script>
