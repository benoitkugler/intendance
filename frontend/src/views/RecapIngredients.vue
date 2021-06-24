<template>
  <div>
    <v-dialog v-model="showEditParametresSimple" max-width="800">
      <options-commande-simple
        :C="C"
        :dateIngredients="dateIngredients"
        @valide="etablitCommandeSimple"
      ></options-commande-simple>
    </v-dialog>

    <v-row class="fill-height px-2 mt-0">
      <v-col md="3" sm="6" class="align-self-center">
        <form-calcul :C="C" :sejour="sejour" @change="onChange"></form-calcul>
      </v-col>
      <v-col md="4" sm="6" class="align-self-center">
        <result-ingredients
          :loading="loadingIngredients"
          :dateIngredients="dateIngredients"
          :origineIngredients="origineIngredients"
          @go-to-ingredient="goToIngredient"
        >
        </result-ingredients>
      </v-col>
      <v-col class="align-self-center" md="5" sm="12">
        <v-tabs v-model="modeCommande" grow>
          <v-tab> Par fournisseurs </v-tab>
          <v-tab> Par produits </v-tab>
        </v-tabs>

        <v-tabs-items v-model="modeCommande">
          <v-tab-item>
            <preview-commande-simple
              :C="C"
              :commande="commandeSimple"
              @showOrigines="(o) => (origineIngredients = o)"
              @editParametres="showEditParametresSimple = true"
            ></preview-commande-simple>
          </v-tab-item>
          <v-tab-item>
            TODO
            <!-- <preview-commande
            :C="C"
            :dateIngredients="dateIngredients"
            @showOrigines="(o) => (origineIngredients = o)"
          ></preview-commande> -->
          </v-tab-item>
        </v-tabs-items>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import FormCalcul from "../components/recap_ingredients/FormCalcul.vue";
import ResultIngredients from "../components/recap_ingredients/ResultIngredients.vue";
import PreviewCommandeComplete from "../components/recap_ingredients/PreviewCommandeComplete.vue";
import PreviewCommandeSimple from "../components/recap_ingredients/PreviewCommandeSimple.vue";

import { Controller } from "../logic/controller";
import {
  CommandeContraintes,
  CommandeSimpleItem,
  DateIngredientQuantites,
  TimedIngredientQuantite,
} from "../logic/api";
import OptionsCommandeSimple from "@/components/recap_ingredients/OptionsCommandeSimple.vue";
import AssociationIngredient from "@/components/produits/AssociationIngredient.vue";

const RecapIngredientsProps = Vue.extend({
  props: {
    C: Object as () => Controller,
  },
});

@Component({
  components: {
    FormCalcul,
    ResultIngredients,
    PreviewCommandeComplete,
    PreviewCommandeSimple,
    OptionsCommandeSimple,
  },
})
export default class RecapIngredients extends RecapIngredientsProps {
  showEditParametresSimple = false;
  showFormCalcul = false;
  loadingIngredients = false;
  dateIngredients: DateIngredientQuantites[] = [];

  origineIngredients: TimedIngredientQuantite[] = [];

  // pour pouvoir raffraichir la requête
  critere: number[] = [];

  commandeSimple: CommandeSimpleItem[] = [];

  modeCommande = 0; // 0 = simple, 1 = complète

  get sejour() {
    return this.C.getSejour();
  }

  async mounted() {
    if (Object.keys(this.C.api.ingredients || {}).length == 0) {
      await this.C.api.GetIngredients();
    }
    if (Object.keys(this.C.api.fournisseurs || {}).length == 0) {
      await this.C.api.GetFournisseurs();
    }
  }

  onChange(critere: number[]) {
    this.critere = critere;
    this.calcul();
  }

  async calcul() {
    if (this.sejour == null) return;

    this.loadingIngredients = true;
    const res = await this.C.resoudIngredientsJournees(
      this.sejour.id,
      this.critere
    );
    this.loadingIngredients = false;
    if (!res) {
      return;
    }
    this.dateIngredients = res;
    this.origineIngredients = []; // reset the highlights
  }

  goToIngredient(id: number) {
    this.$router.push({ name: "menus", query: { idIngredient: String(id) } });
  }

  async etablitCommandeSimple(options: CommandeContraintes) {
    this.showEditParametresSimple = false;
    const res = await this.C.api.EtablitCommandeSimple({
      ingredients: this.dateIngredients,
      contraintes: options,
    });
    if (!res) {
      return;
    }
    this.commandeSimple = res.commande || [];
  }
}
</script>

<style scoped></style>
