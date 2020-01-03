<template>
  <v-select
    :items="sejours"
    :label="label"
    hide-details
    max-width="200"
    :value="idSejour"
    @change="args => $emit('change', Number(args))"
    class="mx-2"
  ></v-select>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { C } from "../../logic/controller";

const SelectSejourProps = Vue.extend({
  props: {
    label: String,
    idSejour: Number
  },
  model: {
    prop: "idSejour",
    event: "change"
  }
});

@Component({})
export default class SelectSejour extends SelectSejourProps {
  get sejours() {
    const items = Object.values(C.data.agenda.sejours).map(sejour => {
      return { value: sejour.sejour.id, text: sejour.sejour.nom };
    });
    items.sort((a, b) => Number(a.text < b.text));
    return items;
  }
}
</script>

<style scoped></style>
