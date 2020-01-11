<template>
  <div>
    <v-dialog v-model="showPrevisuIngredients" max-width="700px">
      <v-skeleton-loader type="card" :loading="loadingIngredients">
        <v-card>
          <v-card-title primary-title class="secondary">
            <h3 class="headline mb-0 ">
              Ingrédients nécessaires
            </h3>
          </v-card-title>
          <date-ingredients
            :date-ingredients="dateIngredients"
          ></date-ingredients>
        </v-card>
      </v-skeleton-loader>
    </v-dialog>

    <v-card>
      <v-card-title primary-title>
        Calcul des ingrédients
      </v-card-title>
      <v-card-text>
        <select-sejour label="Séjour" v-model="sejour"></select-sejour>
        <v-row no-gutters class="mt-4 px-2"><v-col>Journées</v-col></v-row>
        <v-row class="px-2" no-gutters>
          <v-col class="align-self-center">
            <v-switch label="Tout le séjour" v-model="all"></v-switch>
          </v-col>
          <v-col>
            <div class="overflow-y-auto px-3" :style="{ height: '40vh' }">
              <v-checkbox
                v-model="jourOffsets"
                v-for="offset in choixJournees"
                :key="offset"
                :label="offsetToDate(offset)"
                :value="offset"
                hide-details
                :disabled="all"
              >
              </v-checkbox>
            </div>
          </v-col>
        </v-row>

        <div></div>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="secondary" :disabled="!valid" @click="calcul">
          Calculer les ingrédients
        </v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import SelectSejour from "./SelectSejour.vue";
import DateIngredients from "./DateIngredients.vue";
import { Watch } from "vue-property-decorator";
import { C } from "../../logic/controller";
import { Formatter } from "../../logic/formatter";
import {
  DateIngredientQuantites,
  OutResoudIngredients
} from "../../logic/types";

const FormCalculProps = Vue.extend({
  props: {
    initialSejour: Number
  }
});

@Component({
  components: { SelectSejour, DateIngredients }
})
export default class FormCalcul extends FormCalculProps {
  sejour = this.initialSejour;
  all = true;
  jourOffsets: number[] = [];

  showPrevisuIngredients = false;
  loadingIngredients = false;
  dateIngredients: DateIngredientQuantites[] = [];

  @Watch("initialSejour")
  onPropChange() {
    this.sejour = this.initialSejour;
  }

  get valid() {
    return this.all || this.jourOffsets.length > 0;
  }

  get choixJournees(): number[] {
    const s = new Set(
      (C.data.sejours.sejours[this.sejour].repass || []).map(
        rep => rep.jour_offset
      )
    );
    const offsets = [...s];
    return offsets.sort();
  }

  offsetToDate(offset: number) {
    const d = C.formatter.offsetToDate(this.sejour, offset);
    return Formatter.formatDate(d.toISOString());
  }

  async calcul() {
    this.showPrevisuIngredients = true;
    this.loadingIngredients = true;
    let res: OutResoudIngredients | undefined;
    if (this.all) {
      res = await C.calculs.resoudIngredientsSejour(this.sejour);
    } else {
      res = await C.calculs.resoudIngredientsJournees(
        this.sejour,
        this.jourOffsets
      );
    }
    this.loadingIngredients = false;
    if (res == undefined) {
      this.showPrevisuIngredients = false;
      return;
    }
    this.dateIngredients = res.date_ingredients || [];
  }
}
</script>

<style scoped></style>
