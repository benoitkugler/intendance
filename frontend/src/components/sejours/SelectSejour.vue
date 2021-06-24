<template>
  <v-select
    :items="sejours"
    :label="label"
    hide-details
    max-width="200"
    :value="idSejour"
    @change="(args) => $emit('change', Number(args))"
    class="mx-2"
    outlined
  >
    <template v-slot:no-data>
      <v-list-item>
        <v-list-item-content>
          Aucun séjour n'est encore déclaré.
        </v-list-item-content>
      </v-list-item>
    </template>
  </v-select>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Controller } from "@/logic/controller";
import { sortByText } from "../utils/utils";

const SelectSejourProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    label: String,
    idSejour: Number,
  },
  model: {
    prop: "idSejour",
    event: "change",
  },
});

@Component({})
export default class SelectSejour extends SelectSejourProps {
  get sejours() {
    const items = Object.values(this.C.api.sejours.sejours).map((sejour) => {
      return { value: sejour.id, text: sejour.nom };
    });
    return sortByText(items);
  }
}
</script>

<style scoped></style>
