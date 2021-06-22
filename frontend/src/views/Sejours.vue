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
      <assistant-create-repass
        :C="C"
        @create="assistantCreate"
      ></assistant-create-repass>
    </v-dialog>

    <v-container>
      <v-toolbar dense>
        <v-toolbar-title>Séjours</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-tooltip bottom>
            <template v-slot:activator="props">
              <v-btn text v-on="props.on" @click="showAssistantCreate = true">
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
      <v-row justify="center" class="mt-2">
        <v-col class="col-4 align-self-center">
          <select-sejour
            :C="C"
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
          <liste-groupes :C="C" :sejour="sejour"></liste-groupes>
        </v-col>
        <v-col>
          <liste-fournisseurs :C="C" :sejour="sejour"></liste-fournisseurs>
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
import ListeFournisseurs from "../components/sejours/ListeFournisseurs.vue";
import AssistantCreateRepass from "../components/sejours/groupes/AssistantCreateRepass.vue";

import { EditMode, DetailsSejour } from "../logic/types";
import { Controller } from "../logic/controller";
import { Sejour, OptionsAssistantCreateRepass, New } from "../logic/api";

const SejoursProps = Vue.extend({
  props: {
    C: Object as () => Controller,
  },
});

@Component({
  components: {
    TooltipBtn,
    FormSejour,
    SelectSejour,
    ListeFournisseurs,
    ListeGroupes,
    AssistantCreateRepass,
  },
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
    this.C.state.idSejour = idSejour;
  }

  get sejour() {
    return this.C.getSejour();
  }

  async mounted() {
    await Promise.all([this.C.api.GetSejours(), this.C.api.GetFournisseurs()]);
    if (this.C.notifications.getError() == null) {
      this.setClosestSejour();
    }
  }

  private setClosestSejour() {
    const sejours = Object.values(this.C.api.sejours.sejours || {});
    if (sejours.length == 0) return;
    const now = new Date().valueOf();
    const computeDistance = (sejour: Sejour) => {
      const diff = now - new Date(sejour.date_debut).valueOf();
      return Math.abs(diff);
    };
    const criteres = sejours.map((sej) => {
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
    const sejour: New<Sejour> = {
      nom: modif.nom,
      date_debut: modif.date_debut,
      id_utilisateur: this.C.state.idUtilisateur,
    };
    const newSejour = await this.C.api.CreateSejour(sejour);
    if (this.C.notifications.getError() == null && newSejour != undefined) {
      this.C.state.idSejour = newSejour.id;
    }
  }

  private async editSejour(modif: DetailsSejour) {
    const sejour = this.sejour;
    if (sejour === null) return;
    sejour.date_debut = modif.date_debut;
    sejour.nom = modif.nom;
    await this.C.api.UpdateSejour(sejour);
  }

  async supprimeSejour() {
    const sej = this.sejour;
    if (sej == null) return;
    this.idSejour = null;
    this.showConfirmeSupprime = false;
    await this.C.api.DeleteSejour(sej);
    if (this.C.notifications.getError() == null) {
      this.setClosestSejour();
    }
  }

  async assistantCreate(
    options: OptionsAssistantCreateRepass,
    groupesSorties: { [key: number]: number[] }
  ) {
    if (this.idSejour == null) return;
    await this.C.api.AssistantCreateRepas({
      id_sejour: this.idSejour,
      options,
      groupes_sorties: groupesSorties,
    });
    this.showAssistantCreate = false;
  }
}
</script>

<style scoped></style>
