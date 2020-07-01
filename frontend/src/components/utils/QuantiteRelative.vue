<template>
  <v-row>
    <v-col cols="8">
      <v-text-field
        label="Quantité"
        :hint="hint"
        persistent-hint
        type="number"
        :value="displayedQuantite"
        @input="onInput"
        :suffix="unite"
      ></v-text-field>
    </v-col>
    <v-col>
      <v-text-field
        label="Nombre de personnes"
        type="number"
        v-model.number="nbPersonnes"
        :min="1"
        prefix="Pour"
      >
      </v-text-field>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

const QuantiteRelativeProps = Vue.extend({
  props: {
    quantite: Number,
    unite: String
  },
  model: {
    prop: "quantite",
    event: "change"
  }
});

// Permet de modifier une quantité relative
// La valeur effective est toujours ramenée à une personne,
// mais l'utilisateur peut choisir le nombre de personnes
// pour éviter de faire la division manuellement

@Component({})
export default class QuantiteRelative extends QuantiteRelativeProps {
  nbPersonnes = 8;

  get hint() {
    const s = this.nbPersonnes > 1 ? "s" : "";
    return `Quantité désirée pour ${this.nbPersonnes} personne${s}.`;
  }

  get displayedQuantite() {
    return this.nbPersonnes * this.quantite;
  }

  onInput(displayedQuantite: number) {
    if (this.nbPersonnes <= 0) {
      this.nbPersonnes = 1;
    }
    const quantite = displayedQuantite / this.nbPersonnes;
    this.$emit("change", quantite);
  }
}
</script>

<style scoped></style>
