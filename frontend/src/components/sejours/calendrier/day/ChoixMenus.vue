<template>
  <v-expansion-panel>
    <v-expansion-panel-header class="py-0 px-1">
      <toolbar
        title="Menus disponibles"
        :showAdd="false"
        :elevation="2"
        v-model="search"
      ></toolbar>
    </v-expansion-panel-header>
    <v-expansion-panel-content>
      <v-list dense class="overflow-y-auto py-0" :style="{ height: height }">
        <v-list-item-group>
          <v-list-item
            v-for="menu in menus"
            :key="menu.id"
            :value="menu"
            :inactive="false"
          >
            <v-list-item-content draggable @dragstart="onDrag($event, menu)">
              <v-list-item-title>{{ formatMenuName(menu) }}</v-list-item-title>
              <v-list-item-subtitle>
                <i>{{ formatMenuProprietaire(menu) }}</i>
              </v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-expansion-panel-content>
  </v-expansion-panel>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import { C } from "../../../../logic/controller";
import { Menu, MenuComplet } from "../../../../logic/api";
import { DragKind, setDragData } from "../../../utils/utils_drag";
import Toolbar from "../../../utils/Toolbar.vue";

const ChoixMenusProps = Vue.extend({
  props: {
    height: String
  }
});

@Component({
  components: { Toolbar }
})
export default class ChoixMenus extends ChoixMenusProps {
  search = "";
  formatMenuName = C.formatter.formatMenuName;
  formatMenuProprietaire = C.formatter.formatMenuOrRecetteProprietaire;

  get menus() {
    return C.searchMenu(this.search);
  }

  onDrag(event: DragEvent, menu: MenuComplet) {
    if (event == null || event.dataTransfer == null) return;
    setDragData(event.dataTransfer, DragKind.Menu, menu);
    event.dataTransfer.effectAllowed = "link";
  }
}
</script>

<style scoped></style>
