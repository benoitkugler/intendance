<template>
  <v-container fluid class="py-1 px-2">
    <v-row wrap>
      <v-col
        md="4"
        sm="6"
        v-if="state.mode == 'visu' || state.mode == 'editMenu'"
      >
        <transition-group name="fade" mode="out-in" duration="200">
          <liste-menus
            v-show="state.mode == 'visu'"
            key="liste"
            :C="C"
            height="75vh"
            :state="state"
            @change="idMenu => (state.selection.idMenu = idMenu)"
            @edit="startEditMenu"
            @new="startCreateMenu"
            ref="listeMenus"
          />
          <edit-menu
            v-show="state.mode == 'editMenu'"
            key="edit"
            :C="C"
            :mode="editMode"
            :initialMenu="editedMenu"
            @undo="editMenuCancel"
            @done="editMenuDone"
          ></edit-menu>
        </transition-group>
      </v-col>
      <v-col>
        <liste-recettes
          v-show="state.mode == 'visu' || state.mode == 'editMenu'"
          key="liste"
          height="75vh"
          :C="C"
          :state="state"
          @change="idRecette => (state.selection.idRecette = idRecette)"
          @edit="startEditRecette"
          @new="startCreateRecette"
          ref="listeRecettes"
        />
        <edit-recette
          v-show="state.mode == 'editRecette'"
          key="edit"
          :C="C"
          :mode="editMode"
          :initialRecette="editedRecette"
          @undo="editRecetteCancel"
          @done="editRecetteDone"
        ></edit-recette>
      </v-col>
      <v-col md="4" sm="12">
        <liste-ingredients
          height="75vh"
          :C="C"
          :state="state"
          @change="
            idIngredient => (state.selection.idIngredient = idIngredient)
          "
          ref="listeIngredients"
        />
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
import { Controller } from "../logic/controller";
import {
  Menu,
  Recette,
  Ingredient,
  RecetteComplet,
  MenuComplet,
  New
} from "../logic/api";
import { IngredientOptions, EditMode, deepcopy } from "../logic/types";
import {
  StateMenus,
  DefautRecette,
  DefautMenu
} from "../components/menus/types";
import { Watch } from "vue-property-decorator";
import { ClickOutside } from "vuetify/lib";

const MenuProps = Vue.extend({
  props: {
    C: Object as () => Controller
  }
});

@Component({
  components: {
    ListeMenus,
    ListeRecettes,
    ListeIngredients,
    EditMenu,
    EditRecette
  }
})
export default class Menus extends MenuProps {
  state: StateMenus = {
    mode: "visu",
    selection: { idMenu: null, idRecette: null, idIngredient: null }
  };

  editMode: EditMode = "new"; // s'applique au menu, recette ou ingrédient

  // ces champs vont être mis à jour avant édition
  editedMenu: New<Menu> = DefautMenu;
  editedRecette: New<Recette> = DefautRecette;

  $refs!: {
    listeMenus: ListeMenus;
    listeRecettes: ListeRecettes;
    listeIngredients: ListeIngredients;
  };

  async mounted() {
    await this.C.api.loadAllMenus();
    this.selectIngredient();
  }

  activated() {
    if (this.C.notifications.spin) {
      return; // données en cours de chargement
    }
    this.selectIngredient();
  }

  private selectIngredient() {
    const idIngredient = Number(this.$route.query["idIngredient"]);
    if (idIngredient) {
      this.state.mode = "visu";
      // pour afficher tout les ingrédients
      this.state.selection.idRecette = null;
      this.state.selection.idMenu = null;
      this.$refs.listeIngredients.goToItem(idIngredient);
    }
  }

  startEditMenu(menu: Menu) {
    this.editedMenu = menu;
    this.state.selection.idMenu = menu.id;
    this.state.mode = "editMenu";
    this.editMode = "edit";
  }

  startCreateMenu() {
    if (this.C.state.idUtilisateur == null) return;
    const newMenu: Menu = { ...deepcopy(DefautMenu), id: -1 };
    newMenu.id_utilisateur.Int64 = this.C.state.idUtilisateur;
    this.editedMenu = newMenu;
    this.editMode = "new";
    this.state.mode = "editMenu";
  }

  editMenuCancel() {
    this.state.mode = "visu";
  }

  async editMenuDone(menu: MenuComplet) {
    let out: MenuComplet | undefined;
    if (this.editMode == "edit") {
      out = await this.C.api.UpdateMenu(menu);
    } else {
      out = await this.C.api.CreateMenu(menu);
    }
    this.state.mode = "visu";
    if (out !== undefined) {
      this.$refs.listeMenus.goToItem(out.id);
    }
  }

  startEditRecette(recette: Recette) {
    this.editedRecette = recette;
    // this.state.selection.idRecette = recette.id;
    this.editMode = "edit";
    this.state.mode = "editRecette";
  }

  startCreateRecette() {
    const newRecette: Recette = { ...deepcopy(DefautRecette), id: -1 };
    newRecette.id_utilisateur.Int64 = this.C.state.idUtilisateur;
    this.editedRecette = newRecette;
    this.editMode = "new";
    this.state.mode = "editRecette";
  }
  editRecetteCancel() {
    this.state.mode = "visu";
  }

  async editRecetteDone(recette: RecetteComplet) {
    let out: RecetteComplet | undefined;
    if (this.editMode == "edit") {
      out = await this.C.api.UpdateRecette(recette);
    } else {
      out = await this.C.api.CreateRecette(recette);
    }
    this.state.mode = "visu";
    if (out !== undefined) {
      this.$refs.listeRecettes.goToItem(out.id);
    }
  }
}
</script>

<style></style>
