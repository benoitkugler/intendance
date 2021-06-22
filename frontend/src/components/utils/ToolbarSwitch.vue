<template>
  <div class="toolbar-switch">
    <v-tooltip bottom>
      <template v-slot:activator="{ on }">
        <div v-on="on" class="mx-1" style="height: 100%">
          <v-switch
            :label="label"
            hide-details
            :value="value"
            :disabled="disabled"
            @change="$emit('input', $event ? true : false)"
          >
          </v-switch>
        </div>
      </template>
      <span v-html="tooltip"></span>
    </v-tooltip>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

const Props = Vue.extend({
  props: {
    tooltipOn: String,
    tooltipOff: String,
    value: Boolean,
    label: String,
    disabled: Boolean,
  },
  model: {
    prop: "value",
    event: "input",
  },
});

@Component
export default class ToolbarSwitch extends Props {
  get tooltip() {
    if (this.value) {
      return this.tooltipOn;
    }
    return this.tooltipOff ? this.tooltipOff : this.tooltipOn;
  }
}
</script>

<style>
.toolbar-switch .v-input__control {
  margin: auto;
}
.toolbar-switch .v-input__slot {
  height: 100%;
}
.toolbar-switch .v-input {
  height: 100%;
}
</style>
