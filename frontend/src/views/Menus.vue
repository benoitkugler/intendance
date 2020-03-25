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
          :state="state"
          @change="idRecette => (state.selection.idRecette = idRecette)"
          @edit="startEditRecette"
          @new="startCreateRecette"
          ref="listeRecettes"
        />
        <edit-recette
          v-show="state.mode == 'editRecette'"
          key="edit"
          :mode="editMode"
          :initialRecette="editedRecette"
          @undo="editRecetteCancel"
          @done="editRecetteDone"
        ></edit-recette>
      </v-col>
      <v-col md="4" sm="12">
        <liste-ingredients
          height="75vh"
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
import { C } from "../logic/controller";
import { Menu, Recette, Ingredient } from "../logic/types";
import { IngredientOptions, EditMode, New, deepcopy } from "../logic/types2";
import {
  StateMenus,
  DefautRecette,
  DefautMenu
} from "../components/menus/types";
import { Watch } from "vue-property-decorator";

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
    await C.data.loadAllMenus();
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(
        "Les menus, recettes et ingrédients ont bien été chargés."
      );
    }
    this.selectIngredient();
  }

  activated() {
    if (C.notifications.getSpin()) {
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
    if (C.idUtilisateur == null) return;
    const newMenu: Menu = { ...deepcopy(DefautMenu), id: -1 };
    newMenu.id_utilisateur.Int64 = C.idUtilisateur;
    this.editedMenu = newMenu;
    this.editMode = "new";
    this.state.mode = "editMenu";
  }

  editMenuCancel() {
    this.state.mode = "visu";
  }

  async editMenuDone(menu: Menu) {
    let message = "";
    if (this.editMode == "edit") {
      menu = (await C.data.updateMenu(menu)) as Menu;
      message = "Le menu a bien été mis à jour.";
    } else {
      menu = (await C.data.createMenu(menu)) as Menu;
      message = "Le menu a bien été ajouté.";
    }
    this.state.mode = "visu";
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(message);
      this.$nextTick(() => {
        this.$refs.listeMenus.goToItem(menu.id);
      });
    }
  }

  startEditRecette(recette: Recette) {
    this.editedRecette = recette;
    // this.state.selection.idRecette = recette.id;
    this.editMode = "edit";
    this.state.mode = "editRecette";
  }

  startCreateRecette() {
    if (C.idUtilisateur == null) return;
    const newRecette: Recette = { ...deepcopy(DefautRecette), id: -1 };
    newRecette.id_utilisateur.Int64 = C.idUtilisateur;
    this.editedRecette = newRecette;
    this.editMode = "new";
    this.state.mode = "editRecette";
  }
  editRecetteCancel() {
    this.state.mode = "visu";
  }

  async editRecetteDone(recette: Recette) {
    let message = "";
    if (this.editMode == "edit") {
      recette = (await C.data.updateRecette(recette)) as Recette;
      message = "La recette a bien été mise à jour.";
    } else {
      recette = (await C.data.createRecette(recette)) as Recette;
      message = "La recette a bien été ajoutée.";
    }
    this.state.mode = "visu";
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(message);
      this.$refs.listeRecettes.goToItem(recette.id);
    }
  }
}
</script>

<style></style>
