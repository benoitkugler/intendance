<template>
  <v-container fluid class="pa-1">
    <v-row wrap>
      <v-col>
        <liste-menus
          ref="menus"
          height="85vh"
          :menus="menus"
          v-model="selection.menu"
        />
      </v-col>
      <v-col>
        <liste-recettes
          ref="recettes"
          height="85vh"
          :recettes="recettes"
          :bonus-title="recettesBonusTitle"
          v-model="selection.recette"
        />
      </v-col>
      <v-col>
        <liste-ingredients
          ref="ingredients"
          height="85vh"
          :ingredients="ingredients"
          :bonus-title="ingredientsBonusTitle"
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
import { D } from "../logic/controller";
import { Menu, Recette, Ingredient } from "../logic/types";
import { G } from "../logic/getters";
import { IngredientOptions } from "../logic/types2";
import { NS } from "../logic/notifications";

interface Selection {
  menu: Menu | null;
  recette: Recette | null;
  ingredient: Ingredient | null;
}

@Component({
  components: { ListeMenus, ListeRecettes, ListeIngredients }
})
export default class Menus extends Vue {
  mode: "visu" | "edit" = "visu";
  selection: Selection = { menu: null, recette: null, ingredient: null };

  // annotate refs type
  $refs!: {
    menus: ListeMenus;
    recettes: ListeRecettes;
    ingredients: ListeIngredients;
  };

  get ingredientsBonusTitle() {
    if (this.selection.recette != null) {
      return " - Recette courante";
    } else if (this.selection.menu != null) {
      return " - Menu courant";
    }
    return "";
  }

  get recettesBonusTitle() {
    if (this.selection.menu != null) {
      return " - Menu courant";
    }
    return "";
  }

  get ingredients() {
    if (this.selection.recette != null) {
      return G.getRecetteIngredients(this.selection.recette);
    } else if (this.selection.menu != null) {
      return G.getMenuIngredients(this.selection.menu);
    } else {
      return G.getAllIngredients();
    }
  }

  get recettes() {
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
}
</script>
