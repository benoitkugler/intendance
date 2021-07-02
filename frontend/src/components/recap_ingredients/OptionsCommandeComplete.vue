<template>
  <v-card>
    <v-card-title primary-title>
      Options de la commande
      <v-spacer></v-spacer>
      <tooltip-btn
        tooltip="Charger une proposition d'association utilisant vos préférences"
        @click="fetchHints"
        color="accent"
        >Choix automatique des produits
      </tooltip-btn>
    </v-card-title>
    <v-card-subtitle>
      Choisissez le produit à associer un ingrédient.
    </v-card-subtitle>
    <v-card-text>
      <v-row>
        <v-col>
          <v-list class="overflow-y-auto" style="max-height: 53vh">
            <div v-for="(a, i) in ingredients" :key="i">
              <v-list-item>
                <v-row>
                  <v-col cols="3" class="align-self-center">
                    <b>{{ a.nom }}</b>
                  </v-col>
                  <v-col class="align-self-center">
                    <choix-produit
                      :C="C"
                      :idProduit="produitForIngredient(a.id)"
                      @change="(v) => setProduitForIngredient(a.id, v)"
                      :hints="hints[a.id]"
                    >
                    </choix-produit>
                  </v-col>
                </v-row>
              </v-list-item>
              <v-divider></v-divider>
            </div>
          </v-list>
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
  Produit,
} from "@/logic/api";
import { Controller } from "@/logic/controller";
import Vue from "vue";
import Component from "vue-class-component";
import { Watch } from "vue-property-decorator";
import TooltipBtn from "../utils/TooltipBtn.vue";
import ChoixProduit from "./ChoixProduit.vue";
import LivraisonIngredients from "./LivraisonIngredients.vue";

const OptionsCommandeCompleteProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    dateIngredients: Array as () => DateIngredientQuantites[],
  },
});

@Component({
  components: { LivraisonIngredients, TooltipBtn, ChoixProduit },
})
export default class OptionsCommandeComplete extends OptionsCommandeCompleteProps {
  loadingSearch = false;
  search: string | null = null;
  private options = {
    associations: {} as { [key: number]: number },
    regroupe: false,
  };

  hints: { [key: number]: Produit[] } = {};

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
    const data = await this.C.api.ProposeLienIngredientProduit({
      id_sejour: this.C.state.idSejour,
      ingredients: this.dateIngredients,
    });

    if (data === undefined) return;
    for (const id in data || {}) {
      // merge into current
      const produits: Produit[] = (data || {})[id] || [];
      Vue.set(this.hints, id, produits); //VRC

      // choisit la première proposition par défaut
      if (produits.length != 0) {
        this.options.associations[id] = produits[0].id;
      }
    }
  }

  produitForIngredient(idIngredient: number) {
    return this.options.associations[idIngredient];
  }

  setProduitForIngredient(idIngredient: number, idProduit: number | null) {
    Vue.set(this.options.associations, idIngredient, idProduit); //VRC
  }

  valide() {
    this.$emit("valide", this.options);
  }
}
</script>

<style scoped></style>
