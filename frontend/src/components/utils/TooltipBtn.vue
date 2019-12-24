<template>
  <v-tooltip bottom>
    <template v-slot:activator="{ on }">
      <v-btn v-on="on" icon @click="$emit('click')">
        <v-icon :color="color">mdi-{{ mdiIcon }}</v-icon>
      </v-btn>
    </template>
    {{ tooltip }}
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
    }
  }
});

@Component
export default class TooltipBtn extends Props {}
</script>
