<template>
  <v-tooltip bottom>
    <template v-slot:activator="props">
      <v-btn
        :disabled="disabled"
        v-on="props.on"
        :icon="!!mdiIcon"
        :text="!mdiIcon"
        :color="color"
        @click="e => $emit('click', e)"
        :small="small"
      >
        <v-icon v-if="mdiIcon" :color="color">mdi-{{ mdiIcon }}</v-icon>
        <slot></slot>
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
    tooltip: String,
    mdiIcon: String,
    color: {
      type: String,
      default: "black"
    },
    disabled: {
      type: Boolean,
      default: false
    },
    small: {
      type: Boolean,
      default: false
    }
  }
});

@Component
export default class TooltipBtn extends Props {}
</script>
