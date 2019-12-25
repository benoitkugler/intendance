<template>
  <v-container fluid class="pa-1">
    <v-row wrap>
      <transition name="slide-fade">
        <keep-alive>
          <v-col v-if="mode == 'visu'">
            <liste-menus
              ref="menus"
              height="80vh"
              :menus="menus"
              v-model="selection.menu"
              @edit="startEditMenu"
            />
          </v-col>
        </keep-alive>
      </transition>
      <transition name="slide-fade">
        <keep-alive>
          <v-col v-if="mode == 'editMenu'">
            <edit-menu :mode="menuEditMode" :initialMenu="menuEdit"></edit-menu>
          </v-col>
        </keep-alive>
      </transition>
      <transition name="slide-fade">
        <keep-alive>
          <v-col>
            <liste-recettes
              ref="recettes"
              height="80vh"
              :recettes="recettes"
              :bonus-title="recettesBonusTitle"
              v-model="selection.recette"
            />
          </v-col>
        </keep-alive>
      </transition>
      <transition name="slide-fade">
        <keep-alive>
          <v-col>
            <liste-ingredients
              ref="ingredients"
              height="80vh"
              :ingredients="ingredients"
              :bonus-title="ingredientsBonusTitle"
              v-model="selection.ingredient"
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
import { D } from "../logic/controller";
import { Menu, Recette, Ingredient } from "../logic/types";
import { G } from "../logic/getters";
import { IngredientOptions, EditMode } from "../logic/types2";
import { NS } from "../logic/notifications";

interface Selection {
  menu: Menu | null;
  recette: Recette | null;
  ingredient: IngredientOptions | null;
}

@Component({
  components: { ListeMenus, ListeRecettes, ListeIngredients, EditMenu }
})
export default class Menus extends Vue {
  mode: "visu" | "editMenu" | "editRecette" | "editIngredient" = "visu";
  selection: Selection = { menu: null, recette: null, ingredient: null };

  menuEditMode: EditMode = "new";
  menuEdit: Menu | null = null;

  // annotate refs type
  $refs!: {
    menus: ListeMenus;
    recettes: ListeRecettes;
    ingredients: ListeIngredients;
  };

  get ingredientsBonusTitle() {
    if (this.mode == "editMenu") {
      return " - Tous";
    }
    if (this.selection.recette != null) {
      return " - Recette courante";
    } else if (this.selection.menu != null) {
      return " - Menu courant";
    }
    return "";
  }

  get recettesBonusTitle() {
    if (this.mode == "editMenu") {
      return " - Toutes";
    }
    if (this.selection.menu != null) {
      return " - Menu courant";
    }
    return "";
  }

  get ingredients() {
    if (this.mode == "editMenu") {
      return G.getAllIngredients();
    }
    if (this.selection.recette != null) {
      return G.getRecetteIngredients(this.selection.recette);
    } else if (this.selection.menu != null) {
      return G.getMenuIngredients(this.selection.menu);
    } else {
      return G.getAllIngredients();
    }
  }

  get recettes() {
    if (this.mode == "editMenu") {
      return Object.values(D.recettes);
    }
    if (this.selection.menu != null) {
      return G.getMenuRecettes(this.selection.menu);
    }
    return Object.values(D.recettes);
  }

  get menus() {
    return Object.values(D.menus);
  }

  async mounted() {
    await Promise.all([D.loadIngredients(), D.loadUtilisateurs()]);
    await D.loadRecettes(); // recettes dépend des ingrédients
    await D.loadMenus(); // menus dépends des recettes, ingrédients et utilisateurs
    if (NS.getError() == null) {
      NS.setMessage("Les menus, recettes et ingrédients ont bien été chargés.");
    }
  }

  startEditMenu(menu: Menu) {
    this.mode = "editMenu";
    this.menuEditMode = "edit";
    this.menuEdit = menu;
  }
}
</script>
