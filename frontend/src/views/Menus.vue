<template>
  <v-container fluid class="py-1 px-2">
    <v-row wrap>
      <v-col cols="4" v-if="state.mode == 'visu' || state.mode == 'editMenu'">
        <v-fade-transition group hide-on-leave>
          <liste-menus
            v-if="state.mode == 'visu'"
            key="liste"
            style="height: 75vh;"
            :state="state"
            @change="menu => (state.selection.menu = menu)"
            @edit="startEditMenu"
            @new="startCreateMenu"
          />
          <edit-menu
            v-if="state.mode == 'editMenu'"
            key="edit"
            :mode="editMode"
            :initialMenu="state.selection.menu"
            @undo="editMenuCancel"
            @done="editMenuDone"
          ></edit-menu>
        </v-fade-transition>
      </v-col>
      <v-col>
        <v-slide-x-reverse-transition group hide-on-leave>
          <liste-recettes
            v-if="state.mode == 'visu' || state.mode == 'editMenu'"
            key="liste"
            height="75vh"
            :state="state"
            @change="rec => (state.selection.recette = rec)"
            @edit="startEditRecette"
            @new="startCreateRecette"
          />
          <edit-recette
            v-if="state.mode == 'editRecette'"
            key="edit"
            :mode="editMode"
            :initialRecette="state.selection.recette"
            @undo="editRecetteCancel"
            @done="editRecetteDone"
          ></edit-recette>
        </v-slide-x-reverse-transition>
      </v-col>
      <v-col cols="4">
        <transition name="slide-fade">
          <keep-alive>
            <liste-ingredients
              height="75vh"
              :state="state"
              @change="ing => (state.selection.ingredient = ing)"
            />
          </keep-alive>
        </transition>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import ListeMenus from "../components/menus/ListeMenus.vue";
import ListeRecettes from "../components/menus/ListeRecettes.vue";
import ListeIngredients from "../components/menus/ListeIngredients.vue";
import EditMenu from "../components/menus/EditMenu.vue";
import EditRecette from "../components/menus/EditRecette.vue";
import { C } from "../logic/controller";
import { Menu, Recette, Ingredient } from "../logic/types";
import { IngredientOptions, EditMode, New, deepcopy } from "../logic/types2";
import {
  StateMenus,
  DefautRecette,
  DefautMenu
} from "../components/menus/types";

@Component({
  components: {
    ListeMenus,
    ListeRecettes,
    ListeIngredients,
    EditMenu,
    EditRecette
  }
})
export default class Menus extends Vue {
  state: StateMenus = {
    mode: "visu",
    selection: { menu: null, recette: null, ingredient: null }
  };

  editMode: EditMode = "new"; // s'applique au menu, recette ou ingrédient

  async mounted() {
    await C.data.loadAllMenus();
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(
        "Les menus, recettes et ingrédients ont bien été chargés."
      );
    }
  }

  startEditMenu(menu: Menu) {
    this.state.selection.menu = menu;
    this.state.mode = "editMenu";
    this.editMode = "edit";
  }

  startCreateMenu() {
    if (C.idUtilisateur == null) return;
    const newMenu: Menu = { ...deepcopy(DefautMenu), id: -1 };
    newMenu.id_proprietaire.Int64 = C.idUtilisateur;
    this.state.selection.menu = newMenu;
    this.editMode = "new";
    this.state.mode = "editMenu";
  }

  editMenuCancel() {
    if (this.editMode == "new") {
      // on remet le menu courant à zéro
      this.state.selection.menu = null;
    }
    this.state.mode = "visu";
  }

  async editMenuDone(menu: Menu) {
    let message = "";
    if (this.editMode == "edit") {
      await C.data.updateMenu(menu);
      message = "Le menu a bien été mis à jour.";
    } else {
      await C.data.createMenu(menu);
      message = "Le menu a bien été ajouté.";
    }
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(message);
    }
    this.state.selection.menu = null;
    this.state.mode = "visu";
  }

  startEditRecette(recette: Recette) {
    this.state.selection.recette = recette;
    this.editMode = "edit";
    this.state.mode = "editRecette";
  }

  startCreateRecette() {
    if (C.idUtilisateur == null) return;
    const newRecette: Recette = { ...deepcopy(DefautRecette), id: -1 };
    newRecette.id_proprietaire.Int64 = C.idUtilisateur;
    this.state.selection.recette = newRecette;
    this.editMode = "new";
    this.state.mode = "editRecette";
  }
  editRecetteCancel() {
    if (this.editMode == "new") {
      // on remet la recette ourante à zéro
      this.state.selection.recette = null;
    }
    this.state.mode = "visu";
  }

  async editRecetteDone(recette: Recette) {
    let message = "";
    if (this.editMode == "edit") {
      await C.data.updateRecette(recette);
      message = "La recette a bien été mise à jour.";
    } else {
      await C.data.createRecette(recette);
      message = "La recette a bien été ajoutée.";
    }
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(message);
    }
    this.state.selection.recette = null;
    this.state.mode = "visu";
  }
}
</script>

<style></style>
