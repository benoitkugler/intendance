<template>
  <v-row no-gutters>
    <v-col>
      <v-text-field
        :disabled="disabled"
        label="Conditionnement - Quantité"
        type="number"
        v-model.number="conditionnement.quantite"
        @change="onChange"
        hint="Conditionnement d'un exemplaire"
      ></v-text-field>
    </v-col>
    <v-col class="ml-3">
      <unite-field
        :disabled="disabled"
        v-model="conditionnement.unite"
        label="Conditionnement - Unité"
        :unites="allowedUnites"
        @change="onChange"
      ></unite-field>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import UniteField from "./UniteField.vue";
import { Conditionnement } from "../../logic/types";
import { EnumItem } from "../../logic/types2";

const ConditionnementFieldProps = Vue.extend({
  props: {
    conditionnement: Object as () => Conditionnement,
    disabled: Boolean,
    allowedUnites: Array as () => EnumItem[]
  },
  model: {
    prop: "conditionnement",
    event: "change"
  }
});

@Component({
  components: { UniteField }
})
export default class ConditionnementField extends ConditionnementFieldProps {
  onChange() {
    this.$emit("change", this.conditionnement);
  }
}
</script>

<style scoped></style>
