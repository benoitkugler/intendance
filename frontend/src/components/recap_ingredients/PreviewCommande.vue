<template>
  <v-card>
    <v-card-title primary-title class="secondary py-2 px-3">
      <h3 class="headline mb-0 ">
        Commande
      </h3>
      <v-spacer></v-spacer>
      <v-switch
        label="Regroupe"
        v-model="regroupe"
        @change="computeCommande"
        hide-details
        class="my-auto"
      ></v-switch>
    </v-card-title>
    <v-progress-linear indeterminate :active="loading"></v-progress-linear>
    <v-expansion-panels
      :value="dateIngredients.map((_, i) => i)"
      multiple
      accordion
      class="overflow-y-auto"
      style="max-height: 75vh;"
    >
      <v-expansion-panel
        v-for="(commandeJour, i) in commandes"
        :key="i"
        :disabled="commandes.length <= 1"
      >
        <v-expansion-panel-header
          :hide-actions="commandes.length <= 1"
          class="py-0"
        >
          <b>{{ formatDate(commandeJour.date) }}</b>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-simple-table dense fixed-header>
            <thead>
              <tr>
                <th class="text-left">Produit</th>
                <th class="text-center">Quantité</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, j) in commandeJour.produits" :key="j">
                <td>
                  <v-tooltip left>
                    <template v-slot:activator="{ on }">
                      <a
                        v-on="on"
                        @click="showOrigines(item)"
                        v-html="formatProduit(item.produit)"
                      >
                      </a>
                    </template>
                    {{ tooltipOrigin(item) }}
                  </v-tooltip>
                </td>
                <td class="text-center">
                  {{ item.quantite }}
                </td>
              </tr>
            </tbody>
          </v-simple-table>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <div v-if="commandes.length == 0" class="pa-2 font-italic">
      La commande est vide.
    </div>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../utils/TooltipBtn.vue";

import {
  CommandeItem,
  Time,
  DateIngredientQuantites,
  Produit
} from "../../logic/types";
import { Watch } from "vue-property-decorator";
import { C } from "../../logic/controller";
import { Formatter } from "../../logic/formatter";

const PreviewCommandeProps = Vue.extend({
  props: {
    dateIngredients: Array as () => DateIngredientQuantites[]
  }
});

interface commandeJour {
  date: Time;
  produits: CommandeItem[];
}

@Component({
  components: { TooltipBtn }
})
export default class PreviewCommande extends PreviewCommandeProps {
  data: CommandeItem[] = [];
  loading = false;

  formatDate = Formatter.formatDate;
  formatQuantite = Formatter.formatQuantite;

  regroupe = false;

  @Watch("dateIngredients")
  onIngredientsChange() {
    this.computeCommande();
  }

  private async computeCommande() {
    this.loading = true;
    const res = await C.calculs.previewCommande({
      ingredients: this.dateIngredients,
      contraintes: { contrainte_produits: {}, regroupe: this.regroupe } //TODO:
    });
    this.loading = false;
    if (res == undefined) {
      return;
    }
    this.data = res.commande || [];
  }

  // produit par jour
  get commandes() {
    const tmp: { [key: string]: CommandeItem[] } = {};
    this.data.forEach(c => {
      const current = tmp[c.jour_commande] || [];
      current.push(c);
      tmp[c.jour_commande] = current;
    });
    const out: commandeJour[] = [];
    for (const time in tmp) {
      out.push({ date: time, produits: tmp[time] });
    }
    return out.sort((a, b) => Number(new Date(a.date) < new Date(b.date)));
  }

  formatProduit(produit: Produit) {
    const fourstring = C.formatter.formatFournisseur(produit);
    return `${fourstring} - ${produit.nom}`;
  }

  tooltipOrigin(item: CommandeItem) {
    const or = item.origines || [];
    if (or.length == 1) {
      return "Afficher l'ingrédient d'origine";
    }
    return `Affiches les ${or.length} ingrédients d'origine`;
  }

  showOrigines(item: CommandeItem) {
    this.$emit("showOrigines", item.origines);
  }
}
</script>

<style scoped></style>
