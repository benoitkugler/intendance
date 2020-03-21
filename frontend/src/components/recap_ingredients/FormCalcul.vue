<template>
  <v-card>
    <v-card-title primary-title>
      Choix des journées
      <v-spacer></v-spacer>
      <tooltip-btn
        tooltip="Rafraichir le calcul"
        mdi-icon="refresh"
        small
        @click="emitEvent"
      ></tooltip-btn>
    </v-card-title>
    <v-card-subtitle>Ingrédients nécessaires aux journées :</v-card-subtitle>
    <v-card-text>
      <div class="px-3">
        <v-checkbox
          label="Tout le séjour"
          v-model="selectAll"
          :indeterminate="indeterminate"
        ></v-checkbox>
        <v-checkbox
          v-model="critere"
          v-for="offset in choixJournees"
          :key="offset"
          :label="offsetToDate(offset)"
          :value="offset"
          hide-details
        >
        </v-checkbox>
      </div>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../utils/TooltipBtn.vue";

import { Watch } from "vue-property-decorator";
import { C } from "../../logic/controller";
import { Formatter } from "../../logic/formatter";
import {
  DateIngredientQuantites,
  OutResoudIngredients,
  SejourRepas
} from "../../logic/types";
import { compareArrays, Debounce } from "../utils/utils";

const FormCalculProps = Vue.extend({
  props: {
    sejour: Object as () => SejourRepas | null
  }
});

@Component({
  components: { TooltipBtn }
})
export default class FormCalcul extends FormCalculProps {
  debounce = new Debounce(this.emitEvent, 700);

  get selectAll() {
    return compareArrays(this.choixJournees, this.critere);
  }

  set selectAll(b: boolean) {
    if (b) {
      this.critere = this.choixJournees.map(v => v);
    } else {
      this.critere = [];
    }
  }

  get indeterminate() {
    return this.critere.length > 0 && !this.selectAll;
  }

  critere: number[] = [];

  get choixJournees(): number[] {
    if (this.sejour == null) return [];
    const s = new Set((this.sejour.repass || []).map(rep => rep.jour_offset));
    const offsets = [...s];
    return offsets.sort();
  }

  offsetToDate(offset: number) {
    if (this.sejour == null) return "";
    const d = C.offsetToDate(this.sejour.id, offset);
    return Formatter.formatDate(d.toISOString());
  }

  emitEvent() {
    this.$emit("change", this.critere);
  }

  @Watch("critere")
  onChange() {
    this.debounce.delayJob();
  }
}
</script>

<style scoped></style>
