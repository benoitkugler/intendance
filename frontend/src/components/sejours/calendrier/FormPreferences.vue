<template>
  <v-card>
    <v-card-title primary-title>
      Préférences
    </v-card-title>
    <v-card-text>
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
import { PreferencesAgenda, deepcopy } from "../../../logic/api";

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
  current = deepcopy(this.preferences);

  onChange() {
    this.$emit("change", this.current);
  }
}
</script>

<style></style>
