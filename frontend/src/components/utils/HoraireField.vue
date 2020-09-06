<template>
  <v-select
    :disabled="disabled"
    :items="horaires"
    :value="horaire"
    @change="args => $emit('change', args)"
    :label="label"
  >
    <template v-slot:item="props">
      <div :style="{ color: HorairesColors[asI(props.item).value] }">
        {{ asI(props.item).text }}
      </div>
    </template></v-select
  >
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { HorairesColors } from "./utils";
import { enumIntToOptions, EnumItem } from "@/logic/types";
import { HoraireLabels, Horaire } from "@/logic/api";

const HoraireFieldProps = Vue.extend({
  props: {
    horaire: Number as () => Horaire,
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
  horaires = enumIntToOptions(HoraireLabels);
  HorairesColors = HorairesColors;

  asI = (i: EnumItem<Horaire>) => i;
}
</script>

<style scoped></style>
