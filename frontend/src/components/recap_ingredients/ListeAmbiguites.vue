<template>
  <v-card>
    <v-card-title>Résoudre les ambiguités</v-card-title>
    <v-card-subtitle>
      Plusieurs produits sont disponible pour les ingrédients demandés. <br />
      Pour résoudre définitivement une ambiguité, vous pouvez marquer un produit
      <b>par défaut</b>.
    </v-card-subtitle>
    <v-card-text>
      <v-list class="overflow-y-auto" style="max-height: 55vh">
        <div v-for="(a, i) in items" :key="i">
          <v-list-item>
            <v-row>
              <v-col cols="3" class="align-self-center">
                <b>{{ a.ingredient.nom }}</b>
              </v-col>
              <v-col class="align-self-center">
                <v-radio-group v-model="contraintes[a.ingredient.id]">
                  <v-radio
                    v-for="produit in a.produits"
                    :key="produit.id"
                    :value="produit.id"
                  >
                    <template v-slot:label>
                      <span v-html="formatProduit(produit)"></span>
                    </template>
                  </v-radio> </v-radio-group
              ></v-col>
            </v-row>
          </v-list-item>
          <v-divider></v-divider>
        </div>
      </v-list>
    </v-card-text>
    <v-card-actions>
      <v-spacer> </v-spacer>
      <v-btn color="primary" @click="$emit('apply', contraintes)">
        Appliquer à la commande
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Ingredient, Produit } from "@/logic/api";
import { Controller } from "@/logic/controller";
import { ContraintesProduits } from "./types";

const ListeAmbiguitesProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    // ambiguites: Object as () => Ambiguites
  },
});

interface Amb {
  ingredient: Ingredient;
  produits: Produit[];
}

@Component({})
export default class ListeAmbiguites extends ListeAmbiguitesProps {
  contraintes: ContraintesProduits = {};

  get items(): Amb[] {
    return []; // TODO:
    // const ambs = this.ambiguites || {};
    // return Object.keys(ambs).map(idIngredient => {
    //   const ing = this.C.getIngredient(Number(idIngredient));
    //   return { ingredient: ing, produits: ambs[Number(idIngredient)] || [] };
    // });
  }

  formatProduit = this.C.formatter.formatProduit;
}
</script>

<style scoped></style>
