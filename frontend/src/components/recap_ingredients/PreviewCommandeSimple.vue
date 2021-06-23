<template>
  <v-card style="min-height: 40vh">
    <v-card-title class="secondary py-2 px-3">
      Commande par fournisseurs
      <v-spacer></v-spacer>
      <tooltip-btn
        tooltip="Modifier les associations entre ingrédients et fournisseurs"
        mdi-icon="cogs"
        small
        @click="editAssociations"
      ></tooltip-btn>
    </v-card-title>
    <v-progress-linear indeterminate :active="loading"></v-progress-linear>
    <v-expansion-panels
      :value="commandes.map((_, i) => i)"
      multiple
      accordion
      class="overflow-y-auto"
      style="max-height: 71vh"
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
                <th class="text-left">Fournisseur</th>
                <th class="text-left">Ingrédient</th>
                <th class="text-center">Quantité</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="(item, j) in expandIngredientsJour(commandeJour)"
                :key="j"
              >
                <td v-html="formatLivraison(item.livraison)"></td>
                <td>
                  <v-tooltip left>
                    <template v-slot:activator="props">
                      <a v-on="props.on" @click="showOrigines(item)">
                        {{ item.ingredient.nom }}
                      </a>
                    </template>
                    {{ tooltipOrigin(item) }}
                  </v-tooltip>
                </td>
                <td class="text-center">
                  {{ formatQuantite(item.quantite) }}
                  <i class="grey--text">{{ item.ingredient.unite }}</i>
                </td>
              </tr>
            </tbody>
          </v-simple-table>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <div v-if="commandes.length == 0" class="py-8 pa-2 font-italic text-center">
      <v-btn color="success" outlined @click="editAssociations"
        >Etablir la commande</v-btn
      >
    </div>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../utils/TooltipBtn.vue";
import ListeAmbiguites from "./ListeAmbiguites.vue";

import {
  Time,
  DateIngredientQuantites,
  Produit,
  CommandeSimpleItem,
  CommandeContraintes,
  Date_,
  Livraison,
  Ingredient,
  TimedIngredientQuantite,
} from "@/logic/api";
import { Watch } from "vue-property-decorator";
import { Controller } from "@/logic/controller";
import { Formatter } from "@/logic/formatter";
import { ContraintesProduits } from "./types";

const PreviewCommandeSimpleProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    commande: Array as () => CommandeSimpleItem[],
  },
});

interface commandeJour {
  date: Date_;
  produits: CommandeSimpleItem[];
}

interface commandeItem {
  livraison: Livraison;
  ingredient: Ingredient;
  quantite: number;
  origines: TimedIngredientQuantite[];
}

@Component({
  components: { TooltipBtn, ListeAmbiguites },
})
export default class PreviewCommandeSimple extends PreviewCommandeSimpleProps {
  loading = false;

  formatQuantite = Formatter.formatQuantite;
  formatDate = Formatter.formatDate;

  formatLivraison(liv: Livraison) {
    return this.C.formatter.formatLivraison(liv);
  }

  // livraison par jour
  get commandes() {
    const tmp: { [key: string]: CommandeSimpleItem[] } = {};
    this.commande.forEach((c) => {
      const current = tmp[c.jour_commande] || [];
      current.push(c);
      tmp[c.jour_commande] = current;
    });
    const out: commandeJour[] = [];
    for (const time in tmp) {
      out.push({ date: time as Date_, produits: tmp[time] });
    }
    return out.sort(
      (a, b) => new Date(a.date).valueOf() - new Date(b.date).valueOf()
    );
  }

  tooltipOrigin(item: commandeItem) {
    const or = item.origines || [];
    if (or.length == 1) {
      return "Afficher l'ingrédient d'origine";
    }
    return `Affiches les ${or.length} ingrédients d'origine`;
  }

  showOrigines(item: commandeItem) {
    this.$emit("showOrigines", item.origines);
  }

  editAssociations() {
    this.$emit("editAssociations");
  }

  expandIngredientsJour(commande: commandeJour): commandeItem[] {
    const out: commandeItem[] = [];
    commande.produits.forEach((p) => {
      (p.ingredients || []).forEach((ing) => {
        out.push({
          livraison: p.livraison,
          ingredient: ing.ingredient,
          quantite: ing.quantite,
          origines: ing.origines || [],
        });
      });
    });
    return out;
  }
}
</script>

<style scoped></style>
