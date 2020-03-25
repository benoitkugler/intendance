<template>
  <div>
    <v-toolbar dense color="secondary" class="my-1">
      <v-toolbar-title class="px-2">
        Fournisseurs
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          mdi-icon="plus-thick"
          tooltip="Ajouter un groupe..."
          color="green"
        />
      </v-toolbar-items>
    </v-toolbar>
    <v-list dense max-height="75vh" class="overflow-y-auto"> TODO </v-list>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../utils/TooltipBtn.vue";

const ListeFournisseursProps = Vue.extend({
  props: {}
});

@Component({
  components: { TooltipBtn }
})
export default class ListeFournisseurs extends ListeFournisseursProps {}
</script>

<style scoped></style>
