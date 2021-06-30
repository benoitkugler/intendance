<template>
  <div>
    <v-card>
      <v-card-title class="secondary py-2 px-3">
        Commande par produits
        <v-spacer></v-spacer>
        <tooltip-btn
          tooltip="Modifier les paramètres de la commande"
          mdi-icon="cogs"
          small
          @click="editParametres"
        ></tooltip-btn>
      </v-card-title>

      <v-expansion-panels
        :value="commandes.map((_, i) => i)"
        multiple
        accordion
        class="overflow-y-auto"
        style="max-height: 75vh"
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
      <div
        v-if="commandes.length == 0"
        class="py-8 pa-2 font-italic text-center"
      >
        <v-btn color="success" outlined @click="editParametres"
          >Etablir la commande</v-btn
        >
      </div>
    </v-card>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../utils/TooltipBtn.vue";

import {
  Time,
  DateIngredientQuantites,
  Produit,
  CommandeCompleteItem,
  CommandeContraintes,
} from "@/logic/api";
import { Watch } from "vue-property-decorator";
import { Controller } from "@/logic/controller";
import { Formatter } from "@/logic/formatter";
import { ContraintesProduits } from "./types";

const PreviewCommandeCompleteProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    dateIngredients: Array as () => DateIngredientQuantites[],
  },
});

interface commandeJour {
  date: Time;
  produits: CommandeCompleteItem[];
}

@Component({
  components: { TooltipBtn },
})
export default class PreviewCommandeComplete extends PreviewCommandeCompleteProps {
  data: CommandeCompleteItem[] = [];

  formatDate = Formatter.formatDate;
  formatQuantite = Formatter.formatQuantite;

  // produit par jour
  get commandes() {
    const tmp: { [key: string]: CommandeCompleteItem[] } = {};
    this.data.forEach((c) => {
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

  editParametres() {
    this.$emit("editParametres");
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
