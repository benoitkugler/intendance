<template>
  <v-select
    :disabled="disabled"
    :items="unites"
    :value="unite"
    @change="args => $emit('change', args)"
    :label="label"
  ></v-select>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { UniteLabels } from "../../logic/types";
import { enumStringToOptions } from "../../logic/types2";

const UniteFieldProps = Vue.extend({
  props: {
    unite: String,
    label: {
      type: String,
      default: "Unité"
    },
    disabled: {
      type: Boolean,
      default: false
    },
    unites: {
      type: Array,
      default: function() {
        return enumStringToOptions(UniteLabels);
      }
    }
  },
  model: {
    prop: "unite",
    event: "change"
  }
});

@Component({})
export default class UniteField extends UniteFieldProps {}
</script>

<style scoped></style>
