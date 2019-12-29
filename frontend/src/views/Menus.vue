<template>
  <v-container fluid class="pa-1">
    <v-row wrap>
      <v-fade-transition group hide-on-leave :style="{ maxWidth: '35%' }">
        <v-col v-if="state.mode == 'visu'" key="liste">
          <liste-menus
            :style="{ height: '80vh' }"
            :state="state"
            @edit="startEditMenu"
            @change="menu => (state.selection.menu = menu)"
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
import { D } from "../logic/controller";
import { Menu, Recette, Ingredient } from "../logic/types";
import { G } from "../logic/getters";
import { IngredientOptions, EditMode } from "../logic/types2";
import { NS } from "../logic/notifications";
import { StateMenus } from "../components/menus/types";

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
    await Promise.all([D.loadIngredients(), D.loadUtilisateurs()]);
    await D.loadRecettes(); // recettes dépend des ingrédients
    await D.loadMenus(); // menus dépends des recettes, ingrédients et utilisateurs
    if (NS.getError() == null) {
      NS.setMessage("Les menus, recettes et ingrédients ont bien été chargés.");
    }
  }

  startEditMenu(menu: Menu) {
    this.state.selection.menu = menu;
    this.state.mode = "editMenu";
    this.editMode = "edit";
  }

  async editMenuDone(menu: Menu) {
    await D.updateMenu(menu);
    if (NS.getError() == null) {
      NS.setMessage("Le menu a bien été mis à jour.");
    }
    this.state.mode = "visu";
    this.state.selection.menu = null;
  }

  startEditRecette(recette: Recette) {
    this.state.selection.recette = recette;
    this.state.mode = "editRecette";
    this.editMode = "edit";
  }

  async editRecetteDone(recette: Recette) {
    await D.updateRecette(recette);
    if (NS.getError() == null) {
      NS.setMessage("La recette a bien été mise à jour.");
    }
    this.state.mode = "visu";
    this.state.selection.recette = null;
  }
}
</script>

<style></style>
