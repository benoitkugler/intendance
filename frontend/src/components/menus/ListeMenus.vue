<template>
  <div>
    <v-dialog v-model="confirmeSupprime" max-width="500px">
      <v-card>
        <v-card-title primary-title color="warning">
          Confirmer la suppression
        </v-card-title>
        <v-card-text>
          Confirmez-vous la suppression de ce menu ?
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn tile color="warning" @click="supprime">Supprimer</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-toolbar color="secondary" class="my-1">
      <v-toolbar-title>Menus</v-toolbar-title>
    </v-toolbar>
    <v-list dense class="overflow-y-auto">
      <v-list-item-group :value="menu" @change="args => $emit('change', args)">
        <v-list-item v-for="menu in menus" :key="menu.id" :value="menu">
          <template v-slot:default="{ active }">
            <v-list-item-content>
              <v-list-item-title>{{ formatMenuName(menu) }}</v-list-item-title>
              <v-list-item-subtitle>
                <i> {{ formatMenuProprietaire(menu) }}</i>
              </v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action v-if="showButtons(active, menu)">
              <v-row no-gutters>
                <v-col
                  ><tooltip-btn
                    mdi-icon="pencil"
                    tooltip="Modifier ce menu"
                    color="secondary"
                    @click.stop="$emit('edit', menu)"
                  ></tooltip-btn
                ></v-col>
                <v-col
                  ><tooltip-btn
                    mdi-icon="close"
                    tooltip="Supprimer ce menu"
                    color="red"
                    @click.stop="confirmeSupprime = true"
                  ></tooltip-btn
                ></v-col>
              </v-row>
            </v-list-item-action>
          </template>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";
import { D } from "../../logic/controller";
import { Menu } from "../../logic/types";
import {
  formatMenuName,
  formatMenuOrRecetteProprietaire
} from "../../logic/format";
import { G } from "../../logic/getters";
import TooltipBtn from "../utils/TooltipBtn.vue";
import { NS } from "../../logic/notifications";

const Props = Vue.extend({
  props: {
    menus: Array as () => Menu[],
    menu: Object as () => Menu | null
  },
  model: {
    prop: "menu",
    event: "change"
  }
});

@Component({
  components: { TooltipBtn }
})
export default class ListeMenus extends Props {
  confirmeSupprime = false;
  formatMenuName = formatMenuName;
  formatMenuProprietaire = formatMenuOrRecetteProprietaire;

  showButtons(active: boolean, menu: Menu) {
    return (
      active &&
      (!menu.id_proprietaire.Valid ||
        menu.id_proprietaire.Int64 == D.idUtilisateur)
    );
  }

  async supprime() {
    this.confirmeSupprime = false;
    if (this.menu == null) return;
    await D.deleteMenu(this.menu);
    if (NS.getError() == null) {
      NS.setMessage("Menu supprimé avec succès.");
    }
  }
}
</script>
