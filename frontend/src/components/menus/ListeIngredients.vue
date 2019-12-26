<template>
  <div>
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

    <v-toolbar color="secondary" class="toolbar-ingredients my-1">
      <v-toolbar-title class="px-2">
        <v-row no-gutters class="mt-1">
          <v-col>
            Ingrédients
          </v-col>
        </v-row>
        <v-row no-gutters
          ><v-col>
            <small>
              <i>{{ bonusTitle }}</i></small
            >
          </v-col></v-row
        >
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          mdi-icon="magnify"
          tooltip="Filtrer les ingrédients..."
          @click="showSearch = !showSearch"
        />
      </v-toolbar-items>
    </v-toolbar>
    <v-text-field
      outlined
      label="Rechercher"
      placeholder="Tappez pour lancer la recherche"
      v-model="search"
      hide-details
      v-if="showSearch"
      class="mt-2"
      ref="search"
    ></v-text-field>
    <v-list dense :max-height="height" class="overflow-y-auto">
      <v-list-item-group
        :value="ingredient"
        @change="args => $emit('change', args)"
      >
        <v-list-item
          v-for="ingredient in ingredientsWithSearch"
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
            <v-list-item-action v-if="active">
              <tooltip-btn
                mdi-icon="close"
                tooltip="Supprimer cet ingrédient"
                color="red"
                @click.stop="confirmeSupprime = true"
              ></tooltip-btn>
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
import { D } from "../../logic/controller";
import { Ingredient, RecetteIngredient } from "../../logic/types";
import { IngredientOptions } from "../../logic/types2";
import TooltipBtn from "../utils/TooltipBtn.vue";
import { NS } from "../../logic/notifications";
import levenshtein from "js-levenshtein";

const Props = Vue.extend({
  props: {
    height: String,
    ingredients: Array as () => IngredientOptions[],
    ingredient: Object as () => IngredientOptions | null,
    bonusTitle: {
      type: String,
      default: ""
    }
  },
  model: {
    prop: "ingredient",
    event: "change"
  }
});

const MAX_DIST_LEVENSHTEIN = 4;

@Component({ components: { TooltipBtn } })
export default class ListeIngredients extends Props {
  confirmeSupprime = false;
  search = "";
  showSearch = false;

  $refs!: {
    search: Vue;
  };

  subtitle(ingredient: IngredientOptions) {
    if (ingredient.options) {
      return `${ingredient.options.quantite} <i>${
        ingredient.ingredient.unite
      }</i> - Cuisson : ${ingredient.options.cuisson || "-"}`;
    }
    return `<i>${ingredient.ingredient.unite}</i>`;
  }

  // filtre suivant la recherche
  get ingredientsWithSearch() {
    if (!this.search || !this.showSearch) return this.ingredients;
    let filterNom: (nom: string) => boolean;
    try {
      const s = new RegExp(this.search, "i");
      filterNom = nom => s.test(nom);
    } catch {
      const sl = this.search.toLowerCase();
      filterNom = (nom: string) => nom.includes(sl);
    }
    return this.ingredients.filter(ing => {
      if (this.search == ing.ingredient.unite) return true;
      const nom = ing.ingredient.nom.toLowerCase();
      if (filterNom(nom)) return true;
      return levenshtein(nom, this.search) <= MAX_DIST_LEVENSHTEIN;
    });
  }

  async supprime(checkProduits: boolean) {
    this.confirmeSupprime = false;
    if (this.ingredient == null) return;
    await D.deleteIngredient(this.ingredient.ingredient, checkProduits);
    if (NS.getError() == null) {
      NS.setMessage("Ingrédient supprimé avec succès.");
    }
  }

  onDragStart(event: DragEvent, ingredient: Ingredient) {
    if (!event.dataTransfer) return;
    event.dataTransfer.setData("id-ingredient", String(ingredient.id));
    event.dataTransfer.effectAllowed = "copy";
  }

  @Watch("showSearch")
  onShowSearch(b: boolean) {
    if (!b) return;
    setTimeout(() => {
      const input = this.$refs.search.$el.querySelector("input");
      console.log(input);
      if (input != null) input.select();
    }, 50);
  }
}
</script>

<style>
.toolbar-ingredients .v-input__control {
  margin: auto;
}
.toolbar-ingredients .v-input__slot {
  height: 100%;
}
.toolbar-ingredients .v-input {
  height: 100%;
  width: 50%;
}
</style>
