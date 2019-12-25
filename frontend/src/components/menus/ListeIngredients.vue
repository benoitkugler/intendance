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

    <v-toolbar color="secondary" dense class="my-1">
      <v-toolbar-title class="px-2">
        Ingrédients <i>{{ bonusTitle }}</i>
      </v-toolbar-title>
    </v-toolbar>
    <v-list dense :max-height="height" class="overflow-y-auto">
      <v-list-item-group
        :value="ingredient"
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
                <v-row no-gutters>
                  <v-col>
                    {{ ingredient.ingredient.nom }}
                  </v-col>
                  <v-spacer></v-spacer>
                  <v-col v-if="ingredient.options"
                    >{{ ingredient.options.quantite }}
                    {{ ingredient.options.cuisson }}</v-col
                  >
                </v-row>
              </v-list-item-title>
              <v-list-item-subtitle>
                <i>{{ ingredient.ingredient.unite }}</i>
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

@Component({ components: { TooltipBtn } })
export default class ListeIngredients extends Props {
  confirmeSupprime = false;

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
}
</script>
