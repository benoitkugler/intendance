<template>
  <div>
    <v-dialog v-model="showEditFormSejour" max-width="500">
      <form-sejour
        :sejour="sejour"
        :editMode="editModeSejour"
        @accept="onEditSejourDone"
      ></form-sejour>
    </v-dialog>

    <v-dialog v-model="showConfirmeSupprime" max-width="500px">
      <v-card>
        <v-card-title primary-title color="warning">
          Confirmer la suppression
        </v-card-title>
        <v-card-text>
          Confirmez-vous la suppression du séjour
          <b>{{ (sejour || {}).nom }}</b> ?
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn tile color="warning" @click="supprimeSejour">Supprimer</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="showAssistantCreate" max-width="1000px">
      <assistant-create-repass></assistant-create-repass>
    </v-dialog>

    <v-container>
      <v-toolbar dense>
        <v-toolbar-title>Séjours</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-btn text v-on="on" @click="showAssistantCreate = true">
                <v-icon class="mx-1" color="green"
                  >mdi-calendar-multiple</v-icon
                >
                Ajouter plusieurs repas
              </v-btn>
            </template>
            Initier plusieurs repas à partir des groupes
          </v-tooltip>
        </v-toolbar-items>
      </v-toolbar>
      <v-row justify="center">
        <v-col class="col-4 align-self-center">
          <select-sejour
            label="Séjour actif"
            v-model.number="idSejour"
          ></select-sejour>
        </v-col>
        <v-col class="col-2 align-self-center">
          <tooltip-btn
            :disabled="!idSejour"
            tooltip="Modifier les paramètres du séjour..."
            mdi-icon="pencil"
            @click="
              editModeSejour = 'edit';
              showEditFormSejour = true;
            "
            color="secondary"
          ></tooltip-btn>
          <tooltip-btn
            :disabled="!idSejour"
            tooltip="Supprimer le séjour..."
            mdi-icon="close"
            color="red"
            @click="showConfirmeSupprime = true"
          ></tooltip-btn>
          <tooltip-btn
            tooltip="Ajouter un séjour..."
            mdi-icon="plus"
            color="green"
            @click="
              editModeSejour = 'new';
              showEditFormSejour = true;
            "
          ></tooltip-btn>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <liste-groupes :sejour="sejour"></liste-groupes>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../components/utils/TooltipBtn.vue";
import SelectSejour from "../components/sejours/SelectSejour.vue";
import FormSejour from "../components/sejours/FormSejour.vue";
import ListeGroupes from "../components/sejours/groupes/ListeGroupes.vue";
import AssistantCreateRepass from "../components/sejours/groupes/AssistantCreateRepass.vue";

import { EditMode, DetailsSejour, New } from "../logic/types2";
import { C } from "../logic/controller";
import { Sejour } from "../logic/types";

const SejoursProps = Vue.extend({
  props: {}
});

@Component({
  components: {
    TooltipBtn,
    FormSejour,
    SelectSejour,
    ListeGroupes,
    AssistantCreateRepass
  }
})
export default class Sejours extends SejoursProps {
  editModeSejour: EditMode = "new";
  showEditFormSejour = false;
  showConfirmeSupprime = false;
  showAssistantCreate = false;

  get idSejour() {
    const sej = this.sejour;
    return sej == null ? null : sej.id;
  }
  set idSejour(idSejour: number | null) {
    C.state.idSejour = idSejour;
  }

  get sejour() {
    return C.state.getSejour();
  }

  async mounted() {
    await C.data.loadSejours();
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Les séjours ont bien été chargés.");
      this.setClosestSejour();
    }
  }

  private setClosestSejour() {
    const sejours = Object.values(C.data.sejours.sejours);
    if (sejours.length == 0) return;
    const now = new Date().valueOf();
    const computeDistance = (sejour: Sejour) => {
      const diff = now - new Date(sejour.date_debut).valueOf();
      return Math.abs(diff);
    };
    const criteres = sejours.map(sej => {
      return { id: sej.id, distance: computeDistance(sej) };
    });
    const best = criteres.sort((a, b) => (a.distance < b.distance ? 1 : -1))[0];
    this.idSejour = best.id;
  }

  onEditSejourDone(modif: DetailsSejour) {
    this.showEditFormSejour = false;
    if (this.editModeSejour == "edit") {
      this.editSejour(modif);
    } else {
      this.addSejour(modif);
    }
  }

  private async addSejour(modif: DetailsSejour) {
    if (C.idUtilisateur == null) return;
    const sejour: New<Sejour> = {
      nom: modif.nom,
      date_debut: modif.date_debut,
      id_proprietaire: C.idUtilisateur
    };
    const newSejour = await C.data.createSejour(sejour);
    if (C.notifications.getError() == null && newSejour != undefined) {
      C.notifications.setMessage("Le séjour a bien été ajouté.");
      C.state.idSejour = newSejour.id;
    }
  }

  private async editSejour(modif: DetailsSejour) {
    const sejour = this.sejour;
    if (sejour === null) return;
    sejour.date_debut = modif.date_debut;
    sejour.nom = modif.nom;
    await C.data.updateSejour(sejour);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Le séjour a bien été modifié.");
    }
  }

  async supprimeSejour() {
    const sej = this.sejour;
    if (sej == null) return;
    this.idSejour = null;
    this.showConfirmeSupprime = false;
    await C.data.deleteSejour(sej);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Le séjour a été supprimé avec succès.");
      this.setClosestSejour();
    }
  }
}
</script>

<style scoped></style>
