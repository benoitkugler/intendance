<template>
  <div class="fill-height">
    <v-dialog v-model="showEditParametresSimple" max-width="800">
      <options-commande-simple
        :C="C"
        :dateIngredients="dateIngredients"
        @valide="etablitCommandeSimple"
      ></options-commande-simple>
    </v-dialog>

    <v-dialog v-model="showEditParametresComplete" max-width="800">
      <options-commande-complete
        :C="C"
        :dateIngredients="dateIngredients"
        @valide="etablitCommandeComplete"
      ></options-commande-complete>
    </v-dialog>

    <v-row v-if="sejour !== null" class="fill-height px-2 mt-0">
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
            <preview-commande-complete
              :C="C"
              :commande="commandeComplete"
              @showOrigines="(o) => (origineIngredients = o)"
              @editParametres="showEditParametresComplete = true"
            ></preview-commande-complete>
          </v-tab-item>
        </v-tabs-items>
      </v-col>
    </v-row>
    <v-container v-else fluid class="fill-height">
      <v-row no-gutters class="fill-height">
        <v-col v-if="sejour == null" class="align-self-center">
          <v-alert
            class="align-self-center"
            color="secondary"
            :value="true"
            transition="fade-transition"
          >
            Veuillez choisir un séjour via l'onglet <b>Séjours</b>
          </v-alert>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import FormCalcul from "../components/recap_ingredients/FormCalcul.vue";
import ResultIngredients from "../components/recap_ingredients/ResultIngredients.vue";
import PreviewCommandeComplete from "../components/recap_ingredients/PreviewCommandeComplete.vue";
import PreviewCommandeSimple from "../components/recap_ingredients/PreviewCommandeSimple.vue";
import OptionsCommandeSimple from "@/components/recap_ingredients/OptionsCommandeSimple.vue";
import OptionsCommandeComplete from "@/components/recap_ingredients/OptionsCommandeComplete.vue";

import { Controller } from "../logic/controller";
import {
  CommandeCompleteItem,
  CommandeContraintes,
  CommandeSimpleItem,
  DateIngredientQuantites,
  TimedIngredientQuantite,
} from "../logic/api";
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
    OptionsCommandeComplete,
  },
})
export default class RecapIngredients extends RecapIngredientsProps {
  showEditParametresSimple = false;
  showEditParametresComplete = false;

  showFormCalcul = false;
  loadingIngredients = false;
  dateIngredients: DateIngredientQuantites[] = [];

  origineIngredients: TimedIngredientQuantite[] = [];
  // pour pouvoir raffraichir la requête
  critere: number[] = [];

  commandeSimple: CommandeSimpleItem[] = [];
  commandeComplete: CommandeCompleteItem[] = [];

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
    if (this.sejour == null) return;
    this.showEditParametresSimple = false;
    const res = await this.C.api.EtablitCommandeSimple({
      id_sejour: this.sejour.id,
      ingredients: this.dateIngredients,
      contraintes: options,
    });
    if (!res) {
      return;
    }
    this.commandeSimple = res.commande || [];
  }

  async etablitCommandeComplete(options: CommandeContraintes) {
    if (this.sejour == null) return;

    this.showEditParametresComplete = false;
    const res = await this.C.api.EtablitCommandeComplete({
      id_sejour: this.sejour.id,
      ingredients: this.dateIngredients,
      contraintes: options,
    });
    if (!res) {
      return;
    }
    this.commandeComplete = res.commande || [];
  }
}
</script>

<style scoped></style>
