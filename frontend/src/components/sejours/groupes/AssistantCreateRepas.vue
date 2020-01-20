<template>
  <v-card>
    <v-card-title primary-title>
      Assistant de création des repas
    </v-card-title>
    <v-card-actions>
      <v-btn flat color="primary">Créer ces repas</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

const AssitantCreateRepasProps = Vue.extend({
  props: {}
});

@Component({})
export default class AssitantCreateRepas extends AssitantCreateRepasProps {
  //TODO:
}
</script>

<style scoped></style>
