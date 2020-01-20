<template>
  <div>
    <v-toolbar dense>
      <v-toolbar-title>Menus disponibles</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items> </v-toolbar-items>
    </v-toolbar>
    <v-list dense class="overflow-y-auto">
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
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { C } from "../../../logic/controller";
import { Menu } from "../../../logic/types";

const ChoixMenusProps = Vue.extend({
  props: {}
});

@Component({})
export default class ChoixMenus extends ChoixMenusProps {
  formatMenuName = C.formatter.formatMenuName;
  formatMenuProprietaire = C.formatter.formatMenuOrRecetteProprietaire;

  get menus() {
    return Object.values(C.data.menus);
  }

  onDrag(event: DragEvent, menu: Menu) {
    if (event == null || event.dataTransfer == null) return;
    event.dataTransfer.setData("menu", JSON.stringify(menu));
    event.dataTransfer.effectAllowed = "link";
  }
}
</script>

<style scoped></style>
