<template>
  <div>
    <v-dialog v-model="confirmeSupprime" max-width="500px">
      <v-card>
        <v-card-title primary-title color="warning">
          Confirmer la suppression
        </v-card-title>
        <v-card-text>
          Confirmez-vous la suppression de cette recette ? <br />
          <div class="my-3">
            <small>
              Si vous souhaitez seulement l'enlever d'un menu, passez plutôt en
              <b>mode édition</b>.
            </small>
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn tile color="warning" @click="supprime">Supprimer</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-toolbar color="secondary" dense class="my-1">
      <v-toolbar-title>
        Recettes <i>{{ bonusTitle }}</i>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items> </v-toolbar-items>
    </v-toolbar>
    <v-list dense :max-height="height" class="overflow-y-auto">
      <v-list-item-group
        :value="recette"
        @change="args => $emit('change', args)"
      >
        <v-list-item
          v-for="recette in recettes"
          :key="recette.id"
          :value="recette"
        >
          <template v-slot:default="{ active }">
            <v-list-item-content>
              <v-list-item-title>
                <v-list-item-title>{{ recette.nom }}</v-list-item-title>
              </v-list-item-title>
              <v-list-item-subtitle>
                <i> {{ formatRecetteProprietaire(recette) }}</i>
              </v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action v-if="showDelete(active, recette)">
              <tooltip-btn
                mdi-icon="close"
                tooltip="Supprimer cette recette"
                color="red"
                @click.stop="confirmeSupprime = true"
              ></tooltip-btn>
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

import { D } from "../../logic/controller";
import { Recette } from "../../logic/types";
import { formatMenuOrRecetteProprietaire } from "../../logic/format";
import { NS } from "../../logic/notifications";

const Props = Vue.extend({
  props: {
    height: String,
    recettes: Array as () => Recette[],
    bonusTitle: {
      type: String,
      default: ""
    },
    recette: Object as () => Recette | null
  },
  model: {
    prop: "recette",
    event: "change"
  }
});

@Component({
  components: { TooltipBtn }
})
export default class ListeRecettes extends Props {
  confirmeSupprime = false;

  formatRecetteProprietaire = formatMenuOrRecetteProprietaire;

  showDelete(active: boolean, recette: Recette) {
    return (
      active &&
      (!recette.id_proprietaire.Valid ||
        recette.id_proprietaire.Int64 == D.idUtilisateur)
    );
  }

  async supprime() {
    this.confirmeSupprime = false;
    if (this.recette == null) return;
    await D.deleteRecette(this.recette);
    if (NS.getError() == null) {
      NS.setMessage("Recette supprimée avec succès.");
    }
  }
}
</script>
