<template>
  <div>
    <v-dialog max-width="500"> </v-dialog>

    <v-container>
      <liste-fournisseurs></liste-fournisseurs>
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
