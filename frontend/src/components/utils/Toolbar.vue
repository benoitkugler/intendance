<template>
  <v-toolbar dense color="secondary" class="my-1" :elevation="elevation">
    <v-toolbar-title class="px-2">
      <v-row no-gutters class="mt-1">
        <v-col>
          <span v-show="!showSearch">
            {{ title }}
          </span>
          <v-text-field
            v-show="showSearch"
            label="Rechercher"
            placeholder="Tappez..."
            :value="search"
            @input="s => $emit('change', s)"
            hide-details
            dense
            class="mt-3 mb-1"
            @keyup.esc="showSearch = false"
            @click.stop
            ref="search"
          ></v-text-field>
        </v-col>
      </v-row>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-toolbar-items>
      <tooltip-btn
        mdi-icon="magnify"
        tooltip="Rechercher..."
        @click.stop="showSearch = !showSearch"
      />
      <tooltip-btn
        mdi-icon="plus-thick"
        :tooltip="tooltipAdd"
        @click="$emit('add')"
        color="green"
        v-if="showAdd"
      />
    </v-toolbar-items>
  </v-toolbar>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Watch } from "vue-property-decorator";

import TooltipBtn from "../utils/TooltipBtn.vue";

const ToolbarProps = Vue.extend({
  props: {
    title: String,
    tooltipAdd: String,
    showAdd: Boolean,
    search: String,
    elevation: {
      type: Number,
      default: undefined
    }
  },
  model: {
    prop: "search",
    event: "change"
  }
});

@Component({
  components: { TooltipBtn }
})
export default class Toolbar extends ToolbarProps {
  showSearch = false;

  $refs!: {
    search: Vue;
  };

  // donne le focus au champ de recherche
  @Watch("showSearch")
  onShowSearch(b: boolean) {
    if (!b) {
      this.$emit("change", ""); // annule la recherche
    }
    setTimeout(() => {
      const input = this.$refs.search.$el.querySelector("input");
      if (input != null) input.select();
    }, 50);
  }
}
</script>

<style scoped></style>
