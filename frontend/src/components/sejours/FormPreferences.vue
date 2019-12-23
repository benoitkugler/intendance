<template>
  <v-card>
    <v-card-title primary-title>
      Préférences
    </v-card-title>
    <v-card-text>
      <v-switch
        label="Se restreindre au séjour courant"
        v-model="current.restrictSejourCourant"
        @change="onChange"
        persistent-hint
        hint="N'afficher que le séjour actuellement sélectionné."
      ></v-switch>
      <v-switch
        label="Afficher le premier jour en tête de semaine"
        v-model="current.startPremierJour"
        @change="onChange"
        persistent-hint
        hint="Le premier jour affiché est ajusté, au lieu d'être Lundi."
      ></v-switch>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { PreferencesAgenda } from "../../logic/types2";

const Props = Vue.extend({
  model: {
    prop: "preferences",
    event: "change"
  },
  props: {
    // Date as string
    preferences: Object as () => PreferencesAgenda
  }
});

@Component({})
export default class FormPre extends Props {
  current = JSON.parse(JSON.stringify(this.preferences));

  onChange() {
    this.$emit("change", this.current);
  }
}
</script>

<style></style>
