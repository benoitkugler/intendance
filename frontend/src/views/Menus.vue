<template>
  <v-container fluid class="pa-1">
    <v-row wrap>
      <v-fade-transition group hide-on-leave :style="{ maxWidth: '35%' }">
        <v-col v-if="state.mode == 'visu'" key="liste">
          <liste-menus
            :style="{ height: '80vh' }"
            :state="state"
            @change="menu => (state.selection.menu = menu)"
            @edit="startEditMenu"
            @new="startCreateMenu"
          />
        </v-col>
        <v-col v-if="state.mode == 'editMenu'" key="edit">
          <edit-menu
            :mode="editMode"
            :initialMenu="state.selection.menu"
            @undo="state.mode = 'visu'"
            @done="editMenuDone"
          ></edit-menu>
        </v-col>
      </v-fade-transition>
      <v-slide-x-reverse-transition group hide-on-leave>
        <v-col
          v-if="state.mode == 'visu' || state.mode == 'editMenu'"
          key="liste"
        >
          <liste-recettes
            height="80vh"
            :state="state"
            @change="rec => (state.selection.recette = rec)"
            @edit="startEditRecette"
            @new="startCreateRecette"
          />
        </v-col>
        <v-col v-if="state.mode == 'editRecette'" key="edit">
          <edit-recette
            :mode="editMode"
            :initialRecette="state.selection.recette"
            @undo="state.mode = 'visu'"
            @done="editRecetteDone"
          ></edit-recette>
        </v-col>
      </v-slide-x-reverse-transition>
      <transition name="slide-fade">
        <keep-alive>
          <v-col>
            <liste-ingredients
              height="80vh"
              :state="state"
              @change="ing => (state.selection.ingredient = ing)"
            />
          </v-col>
        </keep-alive>
      </transition>
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
import { IngredientOptions, EditMode, New } from "../logic/types2";
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
    const newMenu: Menu = JSON.parse(JSON.stringify(DefautMenu));
    newMenu.id_proprietaire.Int64 = C.idUtilisateur;
    this.state.selection.menu = newMenu;
    this.editMode = "new";
    this.state.mode = "editMenu";
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
    const newRecette: Recette = JSON.parse(JSON.stringify(DefautRecette));
    newRecette.id_proprietaire.Int64 = C.idUtilisateur;
    this.state.selection.recette = newRecette;
    this.editMode = "new";
    this.state.mode = "editRecette";
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
