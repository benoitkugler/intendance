<template>
  <div>
    <v-toolbar dense color="secondary" class="my-1">
      <v-toolbar-title class="px-2">
        Fournisseurs associés au séjour
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          tooltip="Sauvegarder les fournisseurs sélectionnés au séjour"
          mdi-icon="content-save"
          @click="save"
        ></tooltip-btn>
      </v-toolbar-items>
    </v-toolbar>
    <v-list dense max-height="75vh" class="overflow-y-auto">
      <v-list-item-group v-model="idsFournisseurs" multiple>
        <template v-for="(fournisseur, i) in allFournisseurs">
          <v-list-item :key="i" :value="fournisseur.id">
            <template v-slot:default="{ active, toggle }">
              <v-list-item-action>
                <v-checkbox
                  :input-value="active"
                  :true-value="fournisseur.id"
                  @click="toggle"
                ></v-checkbox>
              </v-list-item-action>
              <v-list-item-content>
                <v-row no-gutters>
                  <v-col class="align-self-center"
                    ><v-list-item-title
                      >{{ fournisseur.nom }}
                    </v-list-item-title></v-col
                  >
                  <v-col class="text-right">
                    <tooltip-btn
                      color="primary"
                      @click.stop="select(fournisseur.lieu)"
                      tooltip="Sélectionner tous les fournisseurs du lieu"
                    >
                      {{ fournisseur.lieu }}</tooltip-btn
                    >
                  </v-col>
                </v-row>
              </v-list-item-content>
            </template>
          </v-list-item>
        </template>
      </v-list-item-group>
    </v-list>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../utils/TooltipBtn.vue";
import { C } from "../../logic/controller";
import { SejourRepas } from "../../logic/api";
import { Watch } from "vue-property-decorator";

const ListeFournisseursProps = Vue.extend({
  props: {
    sejour: Object as () => SejourRepas | null
  }
});

@Component({
  components: { TooltipBtn }
})
export default class ListeFournisseurs extends ListeFournisseursProps {
  idsFournisseurs: number[] = [];

  @Watch("sejour")
  onSejourChange() {
    this.idsFournisseurs = this.initialIdsFournisseurs();
  }

  private initialIdsFournisseurs() {
    if (this.sejour == null) return [];
    return (this.sejour.fournisseurs || []).map(f => f.id_fournisseur);
  }

  get allFournisseurs() {
    return Object.values(C.data.fournisseurs || {});
  }

  select(lieu: string) {
    this.idsFournisseurs = this.allFournisseurs
      .filter(f => f.lieu == lieu)
      .map(f => f.id);
  }

  async save() {
    if (this.sejour == null) return;
    await C.data.updateSejourFournisseurs(this.sejour, this.idsFournisseurs);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Fournisseurs mis à jour.");
    }
  }
}
</script>

<style scoped></style>
