<template>
  <v-list dense :height="height" class="overflow-y-auto">
    <v-subheader>Menus</v-subheader>
    <v-list-item-group v-model="currentMenu">
      <v-list-item v-for="menu in menus" :key="menu.id" :value="menu">
        <v-list-item-content>
          <v-list-item-title>{{ formatMenuName(menu) }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list-item-group>
  </v-list>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";
import { D } from "../../logic/controller";
import { Menu } from "../../logic/types";
import { formatMenuName } from "../../logic/format";

const Props = Vue.extend({
  props: {
    height: String,
    menus: Array as () => Menu[]
  }
});

@Component
export default class ListeMenus extends Props {
  currentMenu: Menu | null = null;

  formatMenuName = formatMenuName;

  @Watch("currentMenu")
  onChange() {
    this.$emit("change", this.currentMenu);
  }

  clearSelection() {
    this.currentMenu = null;
  }
}
</script>
