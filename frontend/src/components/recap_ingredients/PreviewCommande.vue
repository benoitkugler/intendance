<template>
  <div>
    <v-dialog v-model="showAmbiguites" max-width="800">
      <liste-ambiguites
        :C="C"
        :ambiguites="ambiguites"
        @apply="applyContraintes"
      ></liste-ambiguites>
    </v-dialog>

    <v-card>
      <v-card-title primary-title class="secondary py-2 px-3">
        <h3 class="headline mb-0 ">
          Commande
        </h3>
        <v-spacer></v-spacer>
        <tooltip-btn
          tooltip="Résoudre les associations ingrédients / produits ambigües..."
          :disabled="nbAmbiguites == 0"
          @click="showAmbiguites = true"
        >
          <transition name="slide-fade" mode="out-in" duration="200">
            <v-avatar
              v-if="nbAmbiguites > 0"
              color="accent"
              :size="30"
              class="mr-2"
            >
              {{ nbAmbiguites }}
            </v-avatar>
            <v-icon v-else>mdi-check</v-icon>
          </transition>
          Ambiguités
        </tooltip-btn>
        <v-divider vertical class="mr-2"></v-divider>
        <v-switch
          label="Regroupe"
          v-model="contraintes.regroupe"
          @change="applyRegroupe"
          hide-details
          class="my-auto pt-0"
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
                      <template v-slot:activator="props">
                        <a
                          v-on="props.on"
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
  </div>
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
  Ambiguites,
  CommandeCompleteItem,
  CommandeCompleteContraintes
} from "@/logic/api";
import { Watch } from "vue-property-decorator";
import { Controller } from "@/logic/controller";
import { Formatter } from "@/logic/formatter";
import { ContraintesProduits } from "./types";

const PreviewCommandeProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    dateIngredients: Array as () => DateIngredientQuantites[]
  }
});

interface commandeJour {
  date: Time;
  produits: CommandeCompleteItem[];
}

@Component({
  components: { TooltipBtn, ListeAmbiguites }
})
export default class PreviewCommande extends PreviewCommandeProps {
  data: CommandeCompleteItem[] = [];
  loading = false;

  formatDate = Formatter.formatDate;
  formatQuantite = Formatter.formatQuantite;

  contraintes: CommandeCompleteContraintes = {
    contrainte_produits: {},
    regroupe: false
  };

  ambiguites: Ambiguites = {};
  showAmbiguites = false;

  get nbAmbiguites() {
    return Object.keys(this.ambiguites || {}).length;
  }

  @Watch("dateIngredients")
  onIngredientsChange() {
    this.contraintes = { contrainte_produits: {}, regroupe: false };
    this.computeCommande();
  }

  private async computeCommande() {
    this.loading = true;
    const res = await this.C.api.EtablitCommandeComplete({
      ingredients: this.dateIngredients,
      contraintes: this.contraintes
    });
    this.loading = false;
    if (res == undefined) {
      return;
    }
    this.data = res.commande || [];
    this.ambiguites = res.ambiguites || {};
  }

  applyRegroupe() {
    this.computeCommande();
  }

  applyContraintes(contraintes: ContraintesProduits) {
    this.contraintes.contrainte_produits = contraintes;
    this.showAmbiguites = false;
    this.computeCommande();
  }

  // produit par jour
  get commandes() {
    const tmp: { [key: string]: CommandeCompleteItem[] } = {};
    this.data.forEach(c => {
      const current = tmp[c.jour_commande] || [];
      current.push(c);
      tmp[c.jour_commande] = current;
    });
    const out: commandeJour[] = [];
    for (const time in tmp) {
      out.push({ date: time as Time, produits: tmp[time] });
    }
    return out.sort(
      (a, b) => new Date(a.date).valueOf() - new Date(b.date).valueOf()
    );
  }

  formatProduit = this.C.formatter.formatProduit;

  tooltipOrigin(item: CommandeCompleteItem) {
    const or = item.origines || [];
    if (or.length == 1) {
      return "Afficher l'ingrédient d'origine";
    }
    return `Affiches les ${or.length} ingrédients d'origine`;
  }

  showOrigines(item: CommandeCompleteItem) {
    this.$emit("showOrigines", item.origines);
  }
}
</script>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active below version 2.1.8 */ {
  transform: translateX(10px);
  opacity: 0;
}
</style>
