<template>
  <v-select
    :disabled="disabled"
    :items="horaires"
    :value="horaire"
    @change="args => $emit('change', args)"
    :label="label"
  >
    <template v-slot:item="{ item }">
      <div :style="{ color: item.color }">{{ item.text }}</div>
    </template></v-select
  >
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { HorairesColors } from "./utils";
import { Horaires } from "../../logic/enums";

const HoraireFieldProps = Vue.extend({
  props: {
    horaire: String,
    label: {
      type: String,
      default: "Horaire"
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  model: {
    prop: "horaire",
    event: "change"
  }
});
@Component({})
export default class HoraireField extends HoraireFieldProps {
  horaires = Horaires.map(h => {
    return { ...h, color: HorairesColors[h.value] };
  });
}
</script>

<style scoped></style>
