<template>
  <v-card>
    <v-card-title primary-title>
      Détails du séjour
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-text-field
          label="Nom du séjour"
          v-model="sejour.nom"
          required
        ></v-text-field>
        <date-field
          v-model="sejour.date_debut"
          label="Date du premier jour"
        ></date-field>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="success" @click="$emit('accept', sejour)"
        >Enregistrer</v-btn
      >
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Sejour, SejourJournees } from "../../logic/types";
import { New, DetailsSejour } from "../../logic/types2";
import DateField from "../utils/DateField.vue";

const Props = Vue.extend({
  props: {
    initialSejour: Object as () => Sejour
  }
});

@Component({
  components: { DateField }
})
export default class FormSejour extends Props {
  sejour: DetailsSejour = JSON.parse(JSON.stringify(this.initialSejour));
}
</script>

<style></style>
