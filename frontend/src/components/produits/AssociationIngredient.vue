<template>
  <v-card v-if="ingredient != null">
    <v-card-title primary-title>
      Produits associés à l'ingrédient {{ ingredient.nom }}
    </v-card-title>
    <v-card-subtitle>
      Vous pouvez choisir ici quels produits concrets sont associés à
      l'ingrédient générique.
    </v-card-subtitle>
    <v-card-text>
      <v-row>
        <v-col>
          <v-skeleton-loader type="paragraph" :loading="loading" class="h-100">
            <v-simple-table dense fixed-header>
              <thead>
                <tr>
                  <th></th>
                  <th class="text-left">Fournisseur</th>
                  <th class="text-left">Nom</th>
                  <th class="text-center">Prix</th>
                  <th class="text-center">Conditionnement</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="produit in produits" :key="produit.id">
                  <td class="px-1">
                    <tooltip-btn
                      small
                      mdi-icon="star"
                      :color="isDefault(produit) ? 'yellow' : ''"
                      :tooltip="
                        isDefault(produit)
                          ? 'Retirer ce produit des favoris'
                          : 'Définir ce produit par défaut'
                      "
                      @click="setDefault(produit)"
                    ></tooltip-btn>
                  </td>
                  <td v-html="formatFournisseur(produit)"></td>
                  <td>{{ produit.nom }}</td>
                  <td class="text-center">{{ produit.prix }} €</td>
                  <td class="text-center">
                    {{ formatConditionnement(produit.conditionnement) }}
                  </td>
                  <td class="px-1">
                    <tooltip-btn
                      small
                      mdi-icon="close"
                      color="red"
                      tooltip="Supprimer ce produit"
                      @click="deleteProduit(produit)"
                    ></tooltip-btn>
                  </td>
                </tr>
              </tbody>
            </v-simple-table>
          </v-skeleton-loader>
        </v-col>
        <v-col>
          <details-produit :C="C" @add="addProduit" :produit="newProduit">
          </details-produit>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import DetailsProduit from "./DetailsProduit.vue";
import TooltipBtn from "../utils/TooltipBtn.vue";

import { Ingredient, IngredientProduits, Produit, Unite } from "@/logic/api";
import { Controller } from "@/logic/controller";
import { Watch } from "vue-property-decorator";
import { New } from "@/logic/api";
import { Formatter } from "@/logic/formatter";

const AssociationIngredientProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    ingredient: Object as () => Ingredient | null,
    activated: Boolean
  }
});

@Component({
  components: { DetailsProduit, TooltipBtn }
})
export default class AssociationIngredient extends AssociationIngredientProps {
  formatConditionnement = Formatter.formatConditionnement;

  get newProduit(): New<Produit> {
    let cond = { quantite: 0, unite: Unite.Litres };
    if (this.ingredient != null) {
      cond = this.ingredient.conditionnement;
      if (this.ingredient.unite != Unite.Piece) {
        // les unités doivent être indentiques
        cond.unite = this.ingredient.unite;
      }
    }
    return {
      id_livraison: -1,
      nom: "",
      conditionnement: cond,
      prix: 0,
      reference_fournisseur: "",
      colisage: 1
    };
  }

  get loading() {
    return this.C.notifications.getSpin();
  }

  ingredientProduits: IngredientProduits | null = null;

  get produits() {
    if (this.ingredientProduits != null) {
      return this.ingredientProduits.produits || [];
    }
    return [];
  }

  formatFournisseur(produit: Produit) {
    return this.C.formatter.formatFournisseur(produit);
  }

  isDefault(produit: Produit) {
    if (this.ingredientProduits == null) return false;
    return (this.ingredientProduits.defaults || {})[produit.id];
  }

  private async loadProduits() {
    if (this.ingredient == null) return;
    const res = await this.C.api.GetIngredientProduits({
      id: this.ingredient.id
    });
    if (res == undefined) return;
    this.ingredientProduits = res;
  }

  mounted() {
    this.loadProduits();
  }

  @Watch("activated")
  onChangeIngredient(b: boolean) {
    if (!b) return;
    this.loadProduits();
  }

  async addProduit(produit: Produit) {
    if (this.ingredient == null) return;
    const res = await this.C.api.AjouteIngredientProduit({
      produit: produit,
      id_ingredient: this.ingredient.id
    });
    if (res) {
      this.ingredientProduits = res;
    }
  }

  async deleteProduit(produit: Produit) {
    if (this.ingredient == null) return;
    await this.C.api.DeleteProduit({ id: produit.id });
    await this.loadProduits();
  }

  async setDefault(produit: Produit) {
    const newState = !this.isDefault(produit);
    if (this.ingredient == null) return;
    const res = await this.C.api.SetDefautProduit({
      id_ingredient: this.ingredient.id,
      id_produit: produit.id,
      on: newState
    });
    if (res) {
      this.ingredientProduits = res;
    }
  }
}
</script>

<style scoped></style>
