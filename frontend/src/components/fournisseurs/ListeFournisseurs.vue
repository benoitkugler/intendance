<template>
  <div>
    <v-dialog v-model="showEditFournisseur" max-width="500px">
      <details-fournisseur
        :fournisseur="currentFournisseur"
        :editMode="editMode"
        @accept="onEditDone"
      ></details-fournisseur>
    </v-dialog>

    <v-dialog v-model="showConfirmeSupprime" max-width="500px">
      <v-card>
        <v-card-title>Confirmer la suppression</v-card-title>
        <v-card-text>
          Le fournisseur et tout les produits associés seront supprimés. <br />
          Attention, cette opération est <b>irréversible</b>.
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" @click="deleteFournisseur"
            >Supprimer définitivement</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-simple-table fixed-header>
      <thead>
        <tr>
          <th class="text-left">Nom</th>
          <th class="text-center">Lieu</th>
          <th class="text-right">
            <tooltip-btn
              tooltip="Ajouter un fournisseur..."
              mdi-icon="plus"
              color="green"
              @click="startCreateFournisseur()"
            ></tooltip-btn>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="fournisseur in fournisseurs" :key="fournisseur.id">
          <td>
            {{ fournisseur.nom }}
          </td>
          <td class="text-center">
            {{ fournisseur.lieu }}
          </td>
          <td class="text-right">
            <tooltip-btn
              mdi-icon="pencil"
              tooltip="Modifier ce fournisseur..."
              color="secondary"
              @click="startEditFournisseur(fournisseur)"
            ></tooltip-btn>

            <tooltip-btn
              mdi-icon="close"
              color="red"
              tooltip="Supprimer le fournisseur et les produits associés"
              @click="confirmeDeleteFournisseur(fournisseur)"
            ></tooltip-btn>
          </td>
        </tr>
      </tbody>
    </v-simple-table>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { C } from "../../logic/controller";
import TooltipBtn from "../utils/TooltipBtn.vue";
import DetailsFournisseur from "./DetailsFournisseur.vue";
import { EditMode, New } from "../../logic/types2";
import { Fournisseur } from "../../logic/types";

const ListeFournisseursProps = Vue.extend({
  props: {}
});

@Component({
  components: { TooltipBtn, DetailsFournisseur }
})
export default class ListeFournisseurs extends ListeFournisseursProps {
  showConfirmeSupprime = false;

  showEditFournisseur = false;
  editMode: EditMode = "new";
  currentFournisseur: New<Fournisseur> | null = null;

  get fournisseurs() {
    return Object.values(C.data.fournisseurs || {});
  }

  startCreateFournisseur() {
    this.currentFournisseur = { nom: "", lieu: "" };
    this.editMode = "new";
    this.showEditFournisseur = true;
  }

  startEditFournisseur(fournisseur: Fournisseur) {
    this.currentFournisseur = fournisseur;
    this.editMode = "edit";
    this.showEditFournisseur = true;
  }

  async onEditDone(fournisseur: Fournisseur) {
    this.showEditFournisseur = false;
    let message: string;
    if (this.editMode == "new") {
      await C.data.createFournisseur(fournisseur);
      message = "Fournisseur ajouté avec succès.";
    } else {
      await C.data.updateFournisseur(fournisseur);
      message = "Fournisseur édité avec succès.";
    }
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(message);
    }
  }

  confirmeDeleteFournisseur(fournisseur: Fournisseur) {
    this.currentFournisseur = fournisseur;
    this.showConfirmeSupprime = true;
  }

  async deleteFournisseur() {
    if (
      this.currentFournisseur == null ||
      this.currentFournisseur.id == undefined
    )
      return;
    this.showConfirmeSupprime = false;
    await C.data.deleteFournisseur(this.currentFournisseur.id);
    await C.data.loadSejours(); // les fournisseurs associés on pu changer
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Fournisseur supprimé avec succès.");
    }
  }
}
</script>

<style scoped></style>
