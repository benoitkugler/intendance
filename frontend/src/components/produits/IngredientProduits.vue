<template>
  <v-card>
    <v-card-title primary-title>
      TODO
    </v-card-title>
    <v-card-actions>
      <v-btn flat color="primary">text</v-btn>
      <v-btn flat color="primary">text</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

const IngredientProduitsProps = Vue.extend({
  props: {}
});

@Component({})
export default class IngredientProduits extends IngredientProduitsProps {}
</script>

<style scoped></style>
