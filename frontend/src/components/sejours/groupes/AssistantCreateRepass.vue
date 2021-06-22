<template>
  <div>
    <v-dialog v-model="showConfirme" max-width="800px">
      <v-card>
        <v-card-title primary-title>
          Confirmer la création des repas
        </v-card-title>
        <v-card-text>
          Le séjour <b>{{ sejour }}</b> contient déjà des repas. Souhaitez-vous
          ajouter ces nouveaux repas, ou effacer les repas existants et repartir
          de zéro ?
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="warning" @click="onConfirme(true)"
            >Effacer les repas existants</v-btn
          >
          <v-btn color="success" @click="onConfirme(false)"
            >Conserver les repas existants</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-card>
      <v-card-title primary-title>
        Assistant de création des repas
      </v-card-title>
      <v-card-subtitle
        >Vous pouvez créer rapidement les repas associé au séjour
        <b>{{ sejour }}</b
        >. Les groupes marqués en sortie auront un repas à part.
      </v-card-subtitle>
      <v-card-text>
        <v-row>
          <v-col cols="4">
            <v-row>
              <v-col>
                <v-text-field
                  label="Durée du séjour"
                  type="number"
                  v-model.number="options.duree"
                  suffix="jours"
                ></v-text-field
              ></v-col>
            </v-row>
            <v-row>
              <v-col
                ><v-switch
                  label="Inclure un goûter"
                  v-model="options.with_gouter"
                ></v-switch
              ></v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-select
                  hide-hint
                  :items="groupes"
                  v-model="options.cinquieme"
                  label="Inclure un cinquième"
                  chips
                  multiple
                ></v-select>
              </v-col>
            </v-row>
          </v-col>
          <v-col>
            <div class="overflow-y-auto" style="height: 50vh">
              <v-row v-for="offset in offsets" :key="offset" no-gutters>
                <v-col cols="3" class="align-self-center">
                  {{ formatOffset(offset) }}
                </v-col>
                <v-col>
                  <v-select
                    hide-hint
                    :items="groupes"
                    v-model="groupesSorties[offset]"
                    label="Groupes en sortie"
                    chips
                    multiple
                  ></v-select>
                </v-col>
              </v-row>
            </div>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="green" @click="onCreate">Créer ces repas</v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { RepasGroupe, OptionsAssistantCreateRepass } from "@/logic/api";
import { Controller } from "@/logic/controller";

const AssitantCreateRepassProps = Vue.extend({
  props: {
    C: Object as () => Controller,
  },
});

@Component({})
export default class AssitantCreateRepass extends AssitantCreateRepassProps {
  options: OptionsAssistantCreateRepass = {
    duree: 7,
    cinquieme: [],
    with_gouter: true,
    delete_existing: false,
  };
  groupesSorties: { [key: number]: number[] } = {};

  showConfirme = false;

  get offsets() {
    return Array.from(Array(this.options.duree).keys());
  }

  get groupes() {
    return this.C.getGroupes().map((groupe) => {
      return { text: groupe.nom, value: groupe.id };
    });
  }

  get sejour() {
    const sej = this.C.getSejour();
    return sej ? sej.nom : "";
  }

  formatOffset(offset: number) {
    if (this.C.state.idSejour == null) return "";
    const date = this.C.offsetToDate(this.C.state.idSejour, offset);
    return date.toLocaleDateString("fr-FR", {
      weekday: "long",
      day: "numeric",
      month: "long",
    });
  }

  onCreate() {
    const sej = this.C.getSejour();
    if (sej && sej.repass && sej.repass.length > 0) {
      this.showConfirme = true;
    } else {
      this.create();
    }
  }

  onConfirme(deleteExisting: boolean) {
    this.showConfirme = false;
    this.options.delete_existing = deleteExisting;
    this.create();
  }

  private create() {
    this.$emit("create", this.options, this.groupesSorties);
  }
}
</script>

<style scoped></style>
