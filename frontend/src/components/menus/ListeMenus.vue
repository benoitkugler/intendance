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

    <toolbar
      v-model="search"
      tooltipAdd="Ajouter un menu..."
      title="Menus"
      :showAdd="state.mode == 'visu'"
      @add="$emit('new')"
    ></toolbar>
    <v-list dense :max-height="height" class="overflow-y-auto" ref="list">
      <v-list-item-group
        :value="state.selection.idMenu"
        @change="args => $emit('change', args)"
      >
        <v-list-item
          v-for="menu in menus"
          :key="menu.id"
          :value="menu.id"
          :class="classItem(menu.id)"
        >
          <template v-slot:default="{ active }">
            <v-list-item-content>
              <v-list-item-title>{{ formatMenuName(menu) }}</v-list-item-title>
              <v-list-item-subtitle>
                <i> {{ formatMenuProprietaire(menu) }}</i>
              </v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action v-if="showButtons(active, menu)">
              <v-row no-gutters>
                <v-col>
                  <tooltip-btn
                    mdi-icon="pencil"
                    tooltip="Modifier ce menu"
                    color="secondary"
                    @click.stop="$emit('edit', menu)"
                  ></tooltip-btn>
                </v-col>
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

import TooltipBtn from "../utils/TooltipBtn.vue";
import Toolbar from "../utils/Toolbar.vue";

import { C } from "../../logic/controller";
import { Menu } from "../../logic/api";
import { StateMenus } from "./types";
import { searchFunction } from "../utils/utils";
import { BaseList, ListKind } from "./shared";

@Component({
  components: { TooltipBtn, Toolbar },
  props: {
    kind: {
      type: String as () => ListKind,
      default: "idMenu"
    }
  }
})
export default class ListeMenus extends BaseList {
  confirmeSupprime = false;
  formatMenuName = C.formatter.formatMenuName;
  formatMenuProprietaire = C.formatter.formatMenuOrRecetteProprietaire;

  search = "";

  get menus() {
    return C.searchMenu(this.search);
  }

  showButtons(active: boolean, menu: Menu) {
    return (
      active &&
      (!menu.id_utilisateur.Valid ||
        menu.id_utilisateur.Int64 == C.idUtilisateur)
    );
  }

  async supprime() {
    this.confirmeSupprime = false;
    if (this.state.selection.idMenu == null) return;
    await C.data.deleteMenu(this.state.selection.idMenu);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Menu supprimé avec succès.");
    }
  }
}
</script>
