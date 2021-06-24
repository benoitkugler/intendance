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
        :C="C"
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

    <v-list-item v-if="treeItems.length === 0">
      <v-list-item-content><i> Aucun fournisseur.</i> </v-list-item-content>
    </v-list-item>
    <v-treeview :items="treeItems" dense class="my-2" open-on-click>
      <template v-slot:label="props">
        <!-- Fournisseur -->
        <v-row v-if="asTI(props.item).isFournisseur" no-gutters>
          <v-col class="align-self-center">
            {{ asTIF(props.item).fournisseur.nom }}
            -
            <i>{{ asTIF(props.item).fournisseur.lieu }}</i>
          </v-col>
          <v-col class="align-self-center text-right">
            <tooltip-btn
              tooltip="Ajouter une contrainte de livraison..."
              mdi-icon="plus"
              color="green"
              @click="startCreateLivraison(asTIF(props.item).fournisseur)"
            ></tooltip-btn>
            <tooltip-btn
              mdi-icon="pencil"
              tooltip="Modifier ce fournisseur..."
              color="secondary"
              @click="startEditFournisseur(asTIF(props.item).fournisseur)"
            ></tooltip-btn>

            <tooltip-btn
              mdi-icon="close"
              color="red"
              tooltip="Supprimer le fournisseur et les produits associés"
              @click="confirmeDeleteFournisseur(asTIF(props.item).fournisseur)"
            ></tooltip-btn>
          </v-col>
        </v-row>
        <!-- Livraisons -->
        <v-row v-else no-gutters>
          <v-col cols="2" class="align-self-center">
            <span
              v-html="formatLivraisonNom(asTIL(props.item).livraison)"
            ></span>
          </v-col>
          <v-col cols="8" class="align-self-center">
            <v-chip
              v-for="jour in filterJoursLivraison(asTIL(props.item).livraison)"
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
              @click="startEditLivraison(asTIL(props.item).livraison)"
            ></tooltip-btn>

            <tooltip-btn
              mdi-icon="close"
              color="red"
              tooltip="Supprimer la contrainte de livraison..."
              @click="confirmeDeleteLivraison(asTIL(props.item).livraison)"
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
import TooltipBtn from "../utils/TooltipBtn.vue";
import DetailsFournisseur from "./DetailsFournisseur.vue";
import DetailsLivraison from "./DetailsLivraison.vue";

import { EditMode, defaultLivraison } from "@/logic/types";
import { Fournisseur, Livraison, New } from "@/logic/api";
import { Days } from "../utils/utils";
import { Controller } from "@/logic/controller";

const ListeFournisseursProps = Vue.extend({
  props: {
    C: Object as () => Controller,
  },
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
  components: { TooltipBtn, DetailsFournisseur, DetailsLivraison },
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
    const fournisseurs = Object.values(this.C.api.fournisseurs);
    const livraisons = Object.values(this.C.api.livraisons);
    return fournisseurs.map((fournisseur) => {
      const lvs = livraisons.filter((l) => l.id_fournisseur == fournisseur.id);
      const children = lvs.map((l) => {
        return { isFournisseur: false, livraison: l };
      });
      return {
        fournisseur: fournisseur,
        children: children,
        isFournisseur: true,
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

  onEditFournisseurDone(fournisseur: Fournisseur) {
    this.showEditFournisseur = false;
    if (this.editMode == "new") {
      this.C.api.CreateFournisseur(fournisseur);
    } else {
      this.C.api.UpdateFournisseur(fournisseur);
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
    await this.C.api.DeleteFournisseur({ id: this.currentFournisseur.id });
    await this.C.api.GetSejours(); // les fournisseurs associés on pu changer
  }

  formatFournisseur(idFournisseur: number) {
    return this.C.getFournisseur(idFournisseur).nom;
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

  onEditLivraisonDone(livraison: Livraison) {
    this.showEditLivraison = false;
    if (this.editMode == "new") {
      this.C.api.CreateLivraison(livraison);
    } else {
      this.C.api.UpdateLivraison(livraison);
    }
  }

  confirmeDeleteLivraison(livraison: Livraison) {
    this.currentLivraison = livraison;
    this.showConfirmeSupprimeLivraison = true;
  }

  deleteLivraison() {
    if (this.currentLivraison == null || this.currentLivraison.id == undefined)
      return;
    this.showConfirmeSupprimeLivraison = false;
    this.C.api.DeleteLivraison({ id: this.currentLivraison.id });
  }
}
</script>

<style>
.v-treeview-node__level {
  width: 15px;
}
</style>
