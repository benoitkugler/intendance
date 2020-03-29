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
                  <td>{{ formatFournisseur(produit) }}</td>
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
          <details-produit @add="addProduit" :produit="newProduit">
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

import { Ingredient, IngredientProduits, Produit } from "../../logic/types";
import { C } from "../../logic/controller";
import { Watch } from "vue-property-decorator";
import { New } from "../../logic/types2";
import { Formatter } from "../../logic/formatter";
import { UniteFields } from "../../logic/enums";

const AssociationIngredientProps = Vue.extend({
  props: {
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
    let cond = { quantite: 0, unite: "" };
    if (this.ingredient != null) {
      cond = this.ingredient.conditionnement;
      if (this.ingredient.unite != UniteFields.Piece) {
        // les unités doivent être indentiques
        cond.unite = this.ingredient.unite;
      }
    }
    return {
      id_fournisseur: -1,
      nom: "",
      conditionnement: cond,
      prix: 0,
      reference_fournisseur: "",
      colisage: 0
    };
  }

  get loading() {
    return C.notifications.getSpin();
  }

  ingredientProduits: IngredientProduits | null = null;

  get produits() {
    if (this.ingredientProduits != null) {
      return this.ingredientProduits.produits || [];
    }
    return [];
  }

  formatFournisseur(produit: Produit) {
    const f = C.getFournisseur(produit);
    return f.nom;
  }

  isDefault(produit: Produit) {
    if (this.ingredientProduits == null) return false;
    return (this.ingredientProduits.defaults || {})[produit.id];
  }

  private async loadProduits() {
    if (this.ingredient == null) return;
    const res = await C.data.getIngredientProduits(this.ingredient.id);
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
    const res = await C.data.ajouteIngredientProduit({
      produit: produit,
      id_ingredient: this.ingredient.id
    });
    if (res) {
      this.ingredientProduits = res;
      C.notifications.setMessage(
        `Produit créé et lié à ${this.ingredient.nom}`
      );
    }
  }

  async deleteProduit(produit: Produit) {
    if (this.ingredient == null) return;
    await C.data.deleteProduit(produit.id);
    if (C.notifications.getError() != null) return;
    await this.loadProduits();
    if (C.notifications.getError() != null) return;
    C.notifications.setMessage(`Produit ${produit.nom} supprimé avec succès`);
  }

  async setDefault(produit: Produit) {
    const newState = !this.isDefault(produit);

    if (this.ingredient == null) return;
    const res = await C.data.setDefautProduit({
      id_ingredient: this.ingredient.id,
      id_produit: produit.id,
      on: newState
    });
    if (res) {
      this.ingredientProduits = res;
      C.notifications.setMessage("Préférence modifiée avec succès.");
    }
  }
}
</script>

<style scoped></style>
