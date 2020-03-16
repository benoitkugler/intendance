<template>
  <div>
    <v-dialog v-model="showEditIngredient" max-width="800px">
      <edit-ingredient
        :initialIngredient="state.selection.ingredient"
        :mode="editMode"
        @edit="editIngredientDone"
      ></edit-ingredient>
    </v-dialog>

    <v-dialog v-model="showEditProduits" scrollable max-width="1000px">
      <association-ingredient
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
          state.selection.menu == null &&
          state.selection.recette == null
      "
      @add="startCreateIngredient"
    ></toolbar>
    <v-list dense :max-height="height" class="overflow-y-auto">
      <v-list-item-group
        :value="state.selection.ingredient"
        @change="args => $emit('change', args)"
      >
        <v-list-item
          v-for="ingredient in ingredients"
          :key="ingredient.ingredient.id"
          :value="ingredient"
        >
          <template v-slot:default="{ active }">
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
            <v-list-item-action v-if="showActions(active)">
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
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";

import EditIngredient from "./EditIngredient.vue";
import AssociationIngredient from "../produits/AssociationIngredient.vue";
import Toolbar from "../utils/Toolbar.vue";

import { C } from "../../logic/controller";
import { Ingredient, RecetteIngredient } from "../../logic/types";
import { IngredientOptions, EditMode, New } from "../../logic/types2";
import TooltipBtn from "../utils/TooltipBtn.vue";
import levenshtein from "js-levenshtein";
import { StateMenus, DefautIngredient } from "./types";
import { searchFunction } from "../utils/utils";

const Props = Vue.extend({
  props: {
    height: String,
    state: Object as () => StateMenus
  }
});

const MAX_DIST_LEVENSHTEIN = 4;

@Component({
  components: { TooltipBtn, EditIngredient, AssociationIngredient, Toolbar }
})
export default class ListeIngredients extends Props {
  confirmeSupprime = false;

  showEditIngredient = false;
  editMode: EditMode = "new";

  showEditProduits = false;

  search = "";

  subtitle(ingredient: IngredientOptions) {
    if (ingredient.options) {
      return `${ingredient.options.quantite} <i>${
        ingredient.ingredient.unite
      }</i> - Cuisson : ${ingredient.options.cuisson || "-"}`;
    }
    return `<i>${ingredient.ingredient.unite}</i>`;
  }

  get title() {
    if (this.state.mode == "editMenu") {
      return "Choisir un ingrédient";
    }
    if (this.state.selection.recette != null) {
      return "Ingrédients liés à la recette";
    } else if (this.state.selection.menu != null) {
      return "Ingrédients liés au menu";
    }
    return "Tous les ingrédients";
  }

  get editedIngredientProduit() {
    if (this.state.selection.ingredient == null) return null;
    return this.state.selection.ingredient.ingredient;
  }

  // filtre suivant la recherche
  private searchIngredients(ingredients: IngredientOptions[], search: string) {
    const predicat = searchFunction(search);
    // cas spécial pour l'unité
    return ingredients.filter(ing => {
      return search == ing.ingredient.unite || predicat(ing.ingredient.nom);
    });
  }

  get ingredients() {
    let baseIngredients: IngredientOptions[];
    if (this.state.mode == "editMenu" || this.state.mode == "editRecette") {
      baseIngredients = C.getAllIngredients();
    } else if (this.state.selection.recette != null) {
      baseIngredients = C.getRecetteIngredients(this.state.selection.recette);
    } else if (this.state.selection.menu != null) {
      baseIngredients = C.getMenuIngredients(this.state.selection.menu);
    } else {
      baseIngredients = C.getAllIngredients();
    }
    return this.searchIngredients(baseIngredients, this.search);
  }

  showActions(active: boolean) {
    return (
      this.state.selection.menu == null &&
      this.state.selection.recette == null &&
      active
    );
  }

  async supprime(checkProduits: boolean) {
    this.confirmeSupprime = false;
    if (this.state.selection.ingredient == null) return;
    await C.data.deleteIngredient(
      this.state.selection.ingredient.ingredient,
      checkProduits
    );
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Ingrédient supprimé avec succès.");
    }
  }

  onDragStart(event: DragEvent, ingredient: Ingredient) {
    if (!event.dataTransfer) return;
    event.dataTransfer.setData("id-ingredient", String(ingredient.id));
    event.dataTransfer.effectAllowed = "copy";
  }

  startEditIngredient(ing: IngredientOptions) {
    this.$emit("change", ing);
    this.editMode = "edit";
    this.showEditIngredient = true;
  }

  async editIngredientDone(ing: Ingredient) {
    this.showEditIngredient = false;
    let message = "";
    if (this.editMode == "edit") {
      await C.data.updateIngredient(ing);
      message = "L'ingrédient a été modifié avec succès.";
    } else {
      await C.data.createIngredient(ing);
      message = "L'ingrédient a été ajouté avec succès.";
    }
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(message);
    }
    this.$emit("change", null);
  }

  startCreateIngredient() {
    this.$emit("change", { ingredient: DefautIngredient });
    this.editMode = "new";
    this.showEditIngredient = true;
  }

  startEditProduits(ingredient: IngredientOptions) {
    this.showEditProduits = true;
  }
}
</script>

<style></style>
