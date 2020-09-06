<template>
  <div>
    <v-dialog v-model="showEditIngredient" max-width="800px">
      <edit-ingredient
        :initialIngredient="editedIngredient"
        :mode="editMode"
        @edit="editIngredientDone"
      ></edit-ingredient>
    </v-dialog>

    <v-dialog v-model="showEditProduits" scrollable width="90%">
      <association-ingredient
        :C="C"
        :ingredient="editedIngredientProduit"
        :activated="showEditProduits"
      ></association-ingredient>
    </v-dialog>

    <v-dialog v-model="confirmeSupprime" max-width="800px">
      <v-card>
        <v-card-title primary-title color="warning">
          Confirmer la suppression
        </v-card-title>
        <v-card-text>
          Vous avez demandé la suppression de cet ingrédient. <br />
          Souhaitez-vous vérifiez qu'aucun produit n'y est associé ? Dans le cas
          contraire, les éventuelles associations ingrédient/produit seront
          supprimées.
          <div class="my-3">
            <small>
              Si vous souhaitez seulement enlever l'ingrédient d'un menu ou
              d'une recette, passez plutôt en
              <b>mode édition</b>.
            </small>
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn tile color="warning" @click="supprime(false)">
            Supprimer l'ingrédient et les liens produits
          </v-btn>
          <v-btn tile color="warning" @click="supprime(true)">
            Vérifier les liens avant de supprimer
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <toolbar
      v-model="search"
      tooltipAdd="Ajouter un ingrédient..."
      :title="title"
      :showAdd="
        state.mode == 'visu' &&
          state.selection.idMenu == null &&
          state.selection.idRecette == null
      "
      @add="startCreateIngredient"
    ></toolbar>
    <v-list dense :max-height="height" class="overflow-y-auto" ref="list">
      <v-list-item-group
        :value="state.selection.idIngredient"
        @change="args => $emit('change', args)"
      >
        <v-list-item
          v-for="ingredient in ingredients"
          :key="ingredient.ingredient.id"
          :value="ingredient.ingredient.id"
          :class="classItem(ingredient.ingredient.id)"
        >
          <template v-slot:default="props">
            <v-list-item-content
              draggable="true"
              @dragstart="ev => onDragStart(ev, ingredient.ingredient)"
            >
              <v-list-item-title>
                {{ ingredient.ingredient.nom }}
              </v-list-item-title>
              <v-list-item-subtitle v-html="subtitle(ingredient)">
              </v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action v-if="showActions(props.active)">
              <v-row no-gutters>
                <v-col>
                  <tooltip-btn
                    mdi-icon="cart"
                    tooltip="Produits liés à cet ingrédient..."
                    color="accent"
                    @click.stop="startEditProduits(ingredient)"
                  >
                  </tooltip-btn>
                </v-col>
                <v-col>
                  <tooltip-btn
                    mdi-icon="pencil"
                    tooltip="Modifier cet ingrédient..."
                    color="secondary"
                    @click.stop="startEditIngredient(ingredient)"
                  >
                  </tooltip-btn>
                </v-col>
                <v-col>
                  <tooltip-btn
                    mdi-icon="close"
                    tooltip="Supprimer cet ingrédient"
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
import Vuetify from "vuetify";
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";

import EditIngredient from "./EditIngredient.vue";
import AssociationIngredient from "../produits/AssociationIngredient.vue";
import Toolbar from "../utils/Toolbar.vue";

import { Controller } from "@/logic/controller";
import { Ingredient, New } from "@/logic/api";
import { IngredientOptions, EditMode } from "@/logic/types";
import TooltipBtn from "../utils/TooltipBtn.vue";
import levenshtein from "js-levenshtein";
import { StateMenus, DefautIngredient } from "./types";
import { searchFunction } from "../utils/utils";
import { DragKind, setDragData } from "../utils/utils_drag";
import { BaseList, ListKind } from "./shared";

@Component({
  components: { TooltipBtn, EditIngredient, AssociationIngredient, Toolbar },
  props: {
    kind: {
      type: String as () => ListKind,
      default: "idIngredient"
    }
  }
})
export default class ListeIngredients extends BaseList {
  confirmeSupprime = false;

  showEditIngredient = false;
  editMode: EditMode = "new";
  editedIngredient: IngredientOptions = {
    ingredient: { ...DefautIngredient, id: -1 }
  };

  showEditProduits = false;

  search = "";

  $refs!: {
    list: HTMLElement;
  };

  private subtitle(ingredient: IngredientOptions) {
    if (ingredient.options) {
      return `${ingredient.options.quantite} <i>${
        ingredient.ingredient.unite
      }</i> - Cuisson : ${ingredient.options.cuisson || "-"}`;
    }
    return `<i>${ingredient.ingredient.unite}</i>`;
  }

  private get title() {
    if (this.state.mode == "editMenu" || this.state.mode == "editRecette") {
      return "Choisir un ingrédient";
    }
    if (this.state.selection.idRecette != null) {
      return "Ingrédients liés à la recette";
    } else if (this.state.selection.idMenu != null) {
      return "Ingrédients liés au menu";
    }
    return "Tous les ingrédients";
  }

  private get editedIngredientProduit() {
    if (this.state.selection.idIngredient == null) return null;
    return this.C.getIngredient(this.state.selection.idIngredient);
  }

  // filtre suivant la recherche
  private searchIngredients(ingredients: IngredientOptions[], search: string) {
    const predicat = searchFunction(search);
    // cas spécial pour l'unité
    return ingredients
      .filter(ing => {
        return search == ing.ingredient.unite || predicat(ing.ingredient.nom);
      })
      .sort((a, b) => Number(a.ingredient.nom < b.ingredient.nom));
  }

  private get ingredients() {
    let baseIngredients: IngredientOptions[];
    if (this.state.mode == "editMenu" || this.state.mode == "editRecette") {
      baseIngredients = this.C.getAllIngredients();
    } else if (this.state.selection.idRecette != null) {
      baseIngredients = this.C.getRecetteIngredients(
        this.state.selection.idRecette
      );
    } else if (this.state.selection.idMenu != null) {
      baseIngredients = this.C.getMenuIngredients(this.state.selection.idMenu);
    } else {
      baseIngredients = this.C.getAllIngredients();
    }
    return this.searchIngredients(baseIngredients, this.search);
  }

  private showActions(active: boolean) {
    return (
      this.state.selection.idMenu == null &&
      this.state.selection.idRecette == null &&
      active
    );
  }

  private supprime(checkProduits: boolean) {
    this.confirmeSupprime = false;
    if (this.state.selection.idIngredient == null) return;
    this.C.api.DeleteIngredient({
      id: this.state.selection.idIngredient,
      check_produits: checkProduits
    });
  }

  private onDragStart(event: DragEvent, ingredient: Ingredient) {
    if (!event.dataTransfer) return;
    setDragData(event.dataTransfer, DragKind.IdIngredient, ingredient.id);
    event.dataTransfer.effectAllowed = "copy";
  }

  private startEditIngredient(ing: IngredientOptions) {
    this.$emit("change", ing);
    this.editedIngredient = ing;
    this.editMode = "edit";
    this.showEditIngredient = true;
  }

  private async editIngredientDone(ing: Ingredient) {
    this.showEditIngredient = false;
    if (this.editMode == "edit") {
      ing = (await this.C.api.UpdateIngredient(ing)) as Ingredient;
    } else {
      ing = (await this.C.api.CreateIngredient(ing)) as Ingredient;
    }
    if (this.C.notifications.getError() == null) {
      this.state.selection.idIngredient = ing.id;
    }
    this.$emit("change", null);
  }

  private startCreateIngredient() {
    this.$emit("change", { ingredient: DefautIngredient });
    this.editMode = "new";
    this.showEditIngredient = true;
  }

  private startEditProduits(ingredient: IngredientOptions) {
    this.showEditProduits = true;
  }
}
</script>

<style></style>
