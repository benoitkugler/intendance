<template>
  <div>
    <v-dialog max-width="500"> </v-dialog>

    <v-container>
      <liste-fournisseurs :C="C"></liste-fournisseurs>
    </v-container>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import { Controller } from "../logic/controller";
import ListeFournisseurs from "../components/fournisseurs/ListeFournisseurs.vue";

const FournisseursProps = Vue.extend({
  props: {
    C: Object as () => Controller
  }
});

@Component({
  components: {
    ListeFournisseurs
  }
})
export default class Fournisseurs extends FournisseursProps {
  mounted() {
    this.C.api.GetFournisseurs();
  }
}
</script>

<style scoped></style>
