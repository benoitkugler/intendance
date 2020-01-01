<template>
  <v-tooltip bottom>
    <template v-slot:activator="{ on }">
      <v-btn
        :disabled="disabled"
        v-on="on"
        icon
        @click="e => $emit('click', e)"
      >
        <v-icon :color="color">mdi-{{ mdiIcon }}</v-icon>
      </v-btn>
    </template>
    <div v-html="tooltip"></div>
  </v-tooltip>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

const Props = Vue.extend({
  props: {
    // Date as string
    tooltip: String,
    mdiIcon: String,
    color: {
      type: String,
      default: "black"
    },
    disabled: {
      type: Boolean,
      default: false
    }
  }
});

@Component
export default class TooltipBtn extends Props {}
</script>
