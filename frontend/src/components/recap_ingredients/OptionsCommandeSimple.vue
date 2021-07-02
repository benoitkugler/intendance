<template>
  <v-card>
    <v-card-title primary-title>
      Options de la commande
      <v-spacer></v-spacer>
      <tooltip-btn
        tooltip="Charger une proposition d'association utilisant vos préférences"
        @click="fetchHints"
        color="accent"
        >Choix automatique des fournisseurs
      </tooltip-btn>
    </v-card-title>
    <v-card-subtitle>
      Cliquez-déposez pour associer un ingrédient à un fournisseur.
    </v-card-subtitle>
    <v-card-text>
      <v-row>
        <v-col>
          <livraison-ingredients
            :C="C"
            :livraison="undefined"
            :ingredients="getIngredientsByLivraison(undefined)"
            @swap-ingredient="(id) => swapIngredient(undefined, id)"
          ></livraison-ingredients>
        </v-col>
        <v-col v-for="livraison in livraisons" :key="livraison.id">
          <livraison-ingredients
            :C="C"
            :livraison="livraison"
            :ingredients="getIngredientsByLivraison(livraison.id)"
            @swap-ingredient="(id) => swapIngredient(livraison, id)"
          ></livraison-ingredients>
        </v-col>
      </v-row>
      <v-row>
        <v-switch
          label="Regroupe au premier jour de commande"
          v-model="options.regroupe"
          class="my-auto pt-0"
          persistent-hint
          hint="Regroupe toutes les commandes sur le premier jour, au lieu d'étaler au mieux."
        ></v-switch>
      </v-row>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn :disabled="!isAssociationsFilled" color="success" @click="valide"
        >Editer la commande</v-btn
      >
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import {
  CommandeContraintes,
  DateIngredientQuantites,
  Ingredient,
  Livraison,
} from "@/logic/api";
import { Controller } from "@/logic/controller";
import Vue from "vue";
import Component from "vue-class-component";
import TooltipBtn from "../utils/TooltipBtn.vue";
import LivraisonIngredients from "./LivraisonIngredients.vue";

const OptionsCommandeSimpleProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    dateIngredients: Array as () => DateIngredientQuantites[],
  },
});

@Component({
  components: { LivraisonIngredients, TooltipBtn },
})
export default class OptionsCommandeSimple extends OptionsCommandeSimpleProps {
  loading = true;

  private options = {
    associations: {} as { [key: number]: number },
    regroupe: false,
  };

  get livraisons() {
    return this.C.getSejourLivraisons();
  }

  get ingredients() {
    const tmp: { [key: number]: boolean } = {};
    this.dateIngredients.forEach((ings) => {
      (ings.ingredients || []).forEach((ing) => {
        tmp[ing.ingredient.id] = true;
      });
    });
    return Object.keys(tmp).map(
      (idIng) => this.C.api.ingredients[Number(idIng)]
    );
  }

  get isAssociationsFilled() {
    return (
      this.ingredients.filter(
        (ing) => this.options.associations[ing.id] === undefined
      ).length === 0
    );
  }

  getIngredientsByLivraison(idLivraison: number | undefined) {
    return this.ingredients.filter(
      (ing) => this.options.associations[ing.id] === idLivraison
    );
  }

  async fetchHints() {
    if (this.C.state.idSejour == null) return;
    const data = await this.C.api.ProposeLienIngredientLivraison({
      id_sejour: this.C.state.idSejour,
      ingredients: this.dateIngredients,
    });
    this.loading = false;
    if (data === undefined) return;
    for (const id in data || {}) {
      // merge into current
      Vue.set(this.options.associations, id, (data || {})[id]); //VRC
    }
  }

  swapIngredient(livraisonCible: Livraison | undefined, idIngredient: number) {
    if (livraisonCible === undefined) {
      Vue.delete(this.options.associations, idIngredient); // VRC
    } else {
      Vue.set(this.options.associations, idIngredient, livraisonCible.id); // VRC
    }
  }

  valide() {
    this.$emit("valide", this.options);
  }
}
</script>

<style scoped></style>
