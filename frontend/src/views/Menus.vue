<template>
  <v-container fluid class="pa-1">
    <v-row wrap>
      <v-col>
        <liste-menus
          ref="menus"
          height="85vh"
          :menus="menus"
          @change="onMenuChange"
        />
      </v-col>
      <v-col>
        <liste-recettes
          ref="recettes"
          height="85vh"
          :recettes="recettes"
          @change="onRecetteChange"
          @reset="onRecettesReset"
          :bonus-title="recettesBonusTitle"
        />
      </v-col>
      <v-col>
        <liste-ingredients
          ref="ingredients"
          height="85vh"
          :ingredients="ingredients"
          @change="onIngredientChange"
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
  kind: "menu" | "recette" | "ingredient";
  item: Menu | Recette | Ingredient;
}

@Component({
  components: { ListeMenus, ListeRecettes, ListeIngredients }
})
export default class Menus extends Vue {
  mode: "visu" | "edit" = "visu";
  selection: Selection | null = null;

  ingredients: IngredientOptions[] = [];
  menus: Menu[] = [];
  recettes: Recette[] = [];

  // annotate refs type
  $refs!: {
    menus: ListeMenus;
    recettes: ListeRecettes;
    ingredients: ListeIngredients;
  };

  get ingredientsBonusTitle() {
    if (this.selection == null) return "";
    if (this.selection.kind == "menu") {
      return " - Menu courant";
    } else if (this.selection.kind == "recette") {
      return " - Recette courante";
    }
    return "";
  }

  get recettesBonusTitle() {
    if (this.selection == null) return "";
    if (this.selection.kind == "menu") {
      return " - Menu courant";
    }
    return "";
  }

  onIngredientChange(ing: Ingredient | null) {}

  onMenuChange(menu: Menu | null) {
    if (menu == null) return;
    this.selection = { kind: "menu", item: menu };
    this.recettes = G.getMenuRecettes(menu);
    this.ingredients = G.getMenuIngredients(menu);
  }

  onRecetteChange(recette: Recette | null) {
    if (recette == null) return;
    this.selection = { kind: "recette", item: recette };
    this.ingredients = G.getRecetteIngredients(recette);
    this.$refs.menus.clearSelection();
  }

  onRecettesReset() {
    this.recettes = Object.values(D.recettes);
    this.$refs.recettes.clearSelection();
    this.ingredients = G.getAllIngredients();
    this.selection = null;
  }

  async mounted() {
    await Promise.all([D.loadMenus(), D.loadIngredients(), D.loadRecettes()]);
    if (NS.getError() == null) {
      NS.setMessage("Les menus, recettes et ingrédients ont bien été chargés.");
    }
    this.menus = Object.values(D.menus);
    this.recettes = Object.values(D.recettes);
  }
}
</script>
