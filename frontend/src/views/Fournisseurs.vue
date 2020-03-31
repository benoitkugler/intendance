<template>
  <div>
    <v-dialog max-width="500"> </v-dialog>

    <v-container>
      <v-toolbar dense>
        <v-toolbar-title>Fournisseurs et produits</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items> </v-toolbar-items>
      </v-toolbar>
      <v-row justify="center">
        <v-col class="col-6 align-self-center">
          <liste-fournisseurs></liste-fournisseurs>
        </v-col>
        <v-col class="col-2 align-self-center"> </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import { C } from "../logic/controller";
import ListeFournisseurs from "../components/fournisseurs/ListeFournisseurs.vue";

const FournisseursProps = Vue.extend({
  props: {}
});

@Component({
  components: {
    ListeFournisseurs
  }
})
export default class Fournisseurs extends FournisseursProps {
  async mounted() {
    await C.data.loadFournisseurs();
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Fournisseurs chargés.");
    }
  }
}
</script>

<style scoped></style>
