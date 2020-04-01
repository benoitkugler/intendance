<template>
  <div>
    <v-dialog v-model="showEditFournisseur" max-width="500px">
      <details-fournisseur
        :fournisseur="currentFournisseur"
        :editMode="editMode"
        @accept="onEditFournisseurDone"
      ></details-fournisseur>
    </v-dialog>

    <v-dialog v-model="showConfirmeSupprimeFournisseur" max-width="500px">
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

    <v-dialog v-model="showEditLivraison" max-width="800px">
      <details-livraison
        :livraison="currentLivraison"
        :editMode="editMode"
        @accept="onEditLivraisonDone"
      ></details-livraison>
    </v-dialog>

    <v-dialog v-model="showConfirmeSupprimeLivraison" max-width="500px">
      <v-card>
        <v-card-title>Confirmer la suppression</v-card-title>
        <v-card-text>
          La contrainte de livraison sera supprimée et retirée de tous les
          produits associés.
          <br />
          Attention, cette opération est <b>irréversible</b>.
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" @click="deleteLivraison"
            >Supprimer définitivement</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-toolbar dense>
      <v-toolbar-title
        >Fournisseurs et contraintes de livraisons</v-toolbar-title
      >
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          tooltip="Ajouter un fournisseur..."
          mdi-icon="plus"
          color="green"
          @click="startCreateFournisseur()"
        ></tooltip-btn>
      </v-toolbar-items>
    </v-toolbar>

    <v-treeview :items="treeItems" dense class="my-2" open-on-click>
      <template v-slot:label="{ item }">
        <!-- Fournisseur -->
        <v-row v-if="asTI(item).isFournisseur" no-gutters>
          <v-col class="align-self-center">
            {{ asTIF(item).fournisseur.nom }}
            -
            <i>{{ asTIF(item).fournisseur.lieu }}</i>
          </v-col>
          <v-col class="align-self-center text-right">
            <tooltip-btn
              tooltip="Ajouter une contrainte de livraison..."
              mdi-icon="plus"
              color="green"
              @click="startCreateLivraison(asTIF(item).fournisseur)"
            ></tooltip-btn>
            <tooltip-btn
              mdi-icon="pencil"
              tooltip="Modifier ce fournisseur..."
              color="secondary"
              @click="startEditFournisseur(asTIF(item).fournisseur)"
            ></tooltip-btn>

            <tooltip-btn
              mdi-icon="close"
              color="red"
              tooltip="Supprimer le fournisseur et les produits associés"
              @click="confirmeDeleteFournisseur(asTIF(item).fournisseur)"
            ></tooltip-btn>
          </v-col>
        </v-row>
        <!-- Livraisons -->
        <v-row v-else no-gutters>
          <v-col cols="2" class="align-self-center">
            <span v-html="formatLivraisonNom(asTIL(item).livraison)"></span>
          </v-col>
          <v-col cols="8" class="align-self-center">
            <v-chip
              v-for="jour in filterJoursLivraison(asTIL(item).livraison)"
              :key="jour"
              small
            >
              {{ jour }}
            </v-chip>
          </v-col>

          <v-col class="align-self-center text-right">
            <tooltip-btn
              mdi-icon="pencil"
              tooltip="Modifier cette contrainte de livraison..."
              color="secondary"
              @click="startEditLivraison(asTIL(item).livraison)"
            ></tooltip-btn>

            <tooltip-btn
              mdi-icon="close"
              color="red"
              tooltip="Supprimer la contrainte de livraison..."
              @click="confirmeDeleteLivraison(asTIL(item).livraison)"
            ></tooltip-btn>
          </v-col>
        </v-row>
      </template>
    </v-treeview>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { C } from "../../logic/controller";
import TooltipBtn from "../utils/TooltipBtn.vue";
import DetailsFournisseur from "./DetailsFournisseur.vue";
import DetailsLivraison from "./DetailsLivraison.vue";

import { EditMode, New, defaultLivraison } from "../../logic/types2";
import { Fournisseur, Livraison } from "../../logic/types";
import { Days } from "../utils/utils";

const ListeFournisseursProps = Vue.extend({
  props: {}
});

interface treeItem {
  isFournisseur: boolean;
}

type treeItemFournisseur = treeItem & {
  fournisseur: Fournisseur;
  children: treeItemLivraison[];
};

type treeItemLivraison = treeItem & {
  livraison: Livraison;
};

@Component({
  components: { TooltipBtn, DetailsFournisseur, DetailsLivraison }
})
export default class ListeFournisseurs extends ListeFournisseursProps {
  showConfirmeSupprimeFournisseur = false;
  showEditFournisseur = false;
  currentFournisseur: New<Fournisseur> | null = null;

  editMode: EditMode = "new"; // shared

  showConfirmeSupprimeLivraison = false;
  showEditLivraison = false;
  currentLivraison: New<Livraison> | null = null;

  // to support typing in template
  asTI = (a: treeItem) => a;
  asTIF = (a: treeItemFournisseur) => a;
  asTIL = (a: treeItemLivraison) => a;

  get treeItems(): treeItemFournisseur[] {
    const fournisseurs = Object.values(C.data.fournisseurs || {});
    const livraisons = Object.values(C.data.livraisons || {});
    return fournisseurs.map(fournisseur => {
      const lvs = livraisons.filter(l => l.id_fournisseur == fournisseur.id);
      const children = lvs.map(l => {
        return { isFournisseur: false, livraison: l };
      });
      return {
        fournisseur: fournisseur,
        children: children,
        isFournisseur: true
      };
    });
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

  async onEditFournisseurDone(fournisseur: Fournisseur) {
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
    this.showConfirmeSupprimeFournisseur = true;
  }

  async deleteFournisseur() {
    if (
      this.currentFournisseur == null ||
      this.currentFournisseur.id == undefined
    )
      return;
    this.showConfirmeSupprimeFournisseur = false;
    await C.data.deleteFournisseur(this.currentFournisseur.id);
    await C.data.loadSejours(); // les fournisseurs associés on pu changer
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Fournisseur supprimé avec succès.");
    }
  }

  formatFournisseur(idFournisseur: number) {
    return C.getFournisseur(idFournisseur).nom;
  }

  formatLivraisonNom(livraison: Livraison) {
    const n = livraison.nom;
    if (n == "") return "<i class='grey--text'>par défaut</i>";
    return n;
  }

  filterJoursLivraison(livraison: Livraison) {
    const out: string[] = [];
    livraison.jours_livraison.forEach((v, i) => {
      if (v) {
        out.push(Days[i]);
      }
    });
    return out;
  }

  startCreateLivraison(fournisseur: Fournisseur) {
    this.currentLivraison = defaultLivraison();
    this.currentLivraison.id_fournisseur = fournisseur.id;
    this.editMode = "new";
    this.showEditLivraison = true;
  }

  startEditLivraison(livraison: Livraison) {
    this.currentLivraison = livraison;
    this.editMode = "edit";
    this.showEditLivraison = true;
  }

  async onEditLivraisonDone(livraison: Livraison) {
    this.showEditLivraison = false;
    let message: string;
    if (this.editMode == "new") {
      await C.data.createLivraison(livraison);
      message = "Contrainte de livraison ajoutée avec succès.";
    } else {
      await C.data.updateLivraison(livraison);
      message = "Contrainte de livraison éditée avec succès.";
    }
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(message);
    }
  }

  confirmeDeleteLivraison(livraison: Livraison) {
    this.currentLivraison = livraison;
    this.showConfirmeSupprimeLivraison = true;
  }

  async deleteLivraison() {
    if (this.currentLivraison == null || this.currentLivraison.id == undefined)
      return;
    this.showConfirmeSupprimeLivraison = false;
    await C.data.deleteLivraison(this.currentLivraison.id);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(
        "Contrainte de livraison supprimée avec succès."
      );
    }
  }
}
</script>

<style>
.v-treeview-node__level {
  width: 15px;
}
</style>
