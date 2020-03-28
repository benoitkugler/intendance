<template>
  <div>
    <v-dialog v-model="showPrevisuIngredients" max-width="600px">
      <v-skeleton-loader type="card" :loading="loadingIngredients">
        <v-card class="py-2">
          <v-card-title primary-title>
            <h3 class="headline mb-0">
              Ingrédients pour {{ repasNbPersonnes }} personne{{
                repasNbPersonnes > 1 ? "s" : ""
              }}
            </h3>
          </v-card-title>
          <div height="50vh">
            <liste-ingredients
              :ingredients="listeIngredients"
            ></liste-ingredients>
          </div>
        </v-card>
      </v-skeleton-loader>
    </v-dialog>

    <v-row no-gutters>
      <v-col cols="8">
        <div class="overflow-y-auto" style="height: 77vh;">
          <v-list dense>
            <v-list-item-group color="primary">
              <div v-for="(horaire, i) in horaires" :key="horaire.value">
                <v-divider v-if="i > 0"></v-divider>

                <v-hover v-slot="{ hover }">
                  <v-subheader
                    :style="{ color: getHoraireColor(horaire.value) }"
                    @dragover="onDragoverHoraireHeader($event)"
                    @drop="onDropHoraireHeader($event, horaire.value)"
                  >
                    {{ horaire.text }}
                    <v-spacer></v-spacer>
                    <tooltip-btn
                      v-if="hover"
                      mdi-icon="plus"
                      color="green"
                      tooltip="Ajouter un repas..."
                      @click="$emit('addRepas', horaire.value)"
                    ></tooltip-btn>
                  </v-subheader>
                </v-hover>
                <template v-for="repas in events[horaire.value]">
                  <v-hover :key="repas.id" v-slot="{ hover }">
                    <v-list-item
                      @dragover="onDragoverRepas($event)"
                      @drop="onDropRepas($event, repas)"
                      @click="$emit('editRepas', repas)"
                    >
                      <v-list-item-content>
                        <v-row no-gutters class="fill-height">
                          <v-col class="px-1 align-self-center overflow-x-auto">
                            <v-chip
                              label
                              v-for="groupe in getGroupes(repas)"
                              :key="groupe.id"
                              class="mr-1 px-1 align-self-center"
                              :color="groupe.couleur"
                              small
                              :style="{ borderWidth: ' 1.5px' }"
                              outlined
                              draggable
                              @dragstart="onDragStart($event, repas, groupe)"
                            >
                              {{ groupe.nom }}
                            </v-chip>
                            <small
                              v-if="getGroupes(repas).length == 0"
                              class="font-italic mr-1"
                              >Aucun groupe.
                            </small>
                            <v-chip
                              v-if="repas.offset_personnes != 0"
                              label
                              class="mr-1 px-1 align-self-center"
                              small
                              :style="{ borderWidth: ' 1.5px' }"
                              outlined
                            >
                              {{ formatNbOffset(repas) }}
                            </v-chip>
                          </v-col>
                          <v-col class="align-self-center">
                            <case-recettes
                              :recettes="repas.recettes"
                            ></case-recettes>
                          </v-col>
                          <v-col>
                            ingredients
                          </v-col>
                        </v-row>
                      </v-list-item-content>
                      <v-list-item-action class="my-1">
                        <v-row no-gutters>
                          <v-col
                            ><tooltip-btn
                              v-if="hover"
                              mdi-icon="food-variant"
                              small
                              tooltip="Calculer les <b>ingrédients</b> nécessaires au repas..."
                              @click.stop="resoudIngredients(repas)"
                            ></tooltip-btn
                          ></v-col>
                          <v-col
                            ><tooltip-btn
                              v-if="hover"
                              mdi-icon="close"
                              color="red"
                              small
                              tooltip="Supprimer ce repas..."
                              @click.stop="deleteRepas(repas)"
                            ></tooltip-btn
                          ></v-col>
                        </v-row>
                      </v-list-item-action>
                    </v-list-item>
                  </v-hover>
                </template>
              </div>
            </v-list-item-group>
          </v-list></div
      ></v-col>
      <v-col class="px-2">
        <choix-menus height="67vh"></choix-menus>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../../../utils/TooltipBtn.vue";
import ListeIngredients from "../../../utils/ListeIngredients.vue";
import ChoixMenus from "./ChoixMenus.vue";
import CaseRecettes from "./CaseRecettes.vue";

import { toDateVuetify, formatNbOffset } from "../utils";
import { C } from "../../../../logic/controller";
import {
  RepasComplet,
  Groupe,
  IngredientQuantite
} from "../../../../logic/types";
import { New, NullId, deepcopy } from "../../../../logic/types2";
import { Horaires } from "../../../../logic/enums";
import { HorairesColors } from "../../../utils/utils";
import { DragKind, getDragData, setDragData } from "../../../utils/utils_drag";

const DayProps = Vue.extend({
  props: {
    jourOffset: Number as () => number | null
  }
});
@Component({
  components: { TooltipBtn, ListeIngredients, ChoixMenus, CaseRecettes }
})
export default class Day extends DayProps {
  showPrevisuIngredients = false;
  loadingIngredients = true;
  listeIngredients: IngredientQuantite[] = [];
  repasNbPersonnes = 0;
  get day(): Date | null {
    if (this.jourOffset == null) return null;
    return C.offsetToDate(C.state.idSejour!, this.jourOffset);
  }

  get dayString() {
    const day = this.day;
    return day ? toDateVuetify(day) : "";
  }

  get horaires() {
    return Horaires;
  }

  getHoraireColor(h: string) {
    return HorairesColors[h];
  }

  formatNbOffset = formatNbOffset;

  get events(): { [key: string]: RepasComplet[] } {
    const out: { [key: string]: RepasComplet[] } = {};
    C.iterateAllRepas((sejour, repas) => {
      if (sejour.id != C.state.idSejour) return;
      if (this.jourOffset == repas.jour_offset) {
        const l = out[repas.horaire] || [];
        l.push(repas);
        out[repas.horaire] = l;
      }
    });
    for (const h in out) {
      // tri par id pour être déterministe
      const repass = out[h];
      out[h] = repass.sort((a, b) => {
        return a.id < b.id ? -1 : +1;
      });
    }
    return out;
  }

  getGroupes(repas: RepasComplet) {
    const grs = C.getRepasGroupes(repas);
    const maxChar = 8;
    return grs.map(groupe => {
      const g: Groupe = deepcopy(groupe);
      g.nom = g.nom.substr(0, maxChar) + (g.nom.length <= maxChar ? "" : ".");
      return groupe;
    });
  }

  // Drag and drop

  onDragStart(event: DragEvent, repas: RepasComplet, groupe: Groupe) {
    if (event == null || event.dataTransfer == null) return;
    const data = { repas: repas, idGroupe: groupe.id };
    setDragData(event.dataTransfer, DragKind.Groupe, data);
    event.dataTransfer.effectAllowed = "copyMove";
  }

  onDragoverHoraireHeader(event: DragEvent) {
    if (!event.dataTransfer) return;
    if (event.dataTransfer.types.includes(DragKind.Groupe)) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "copy";
    }
  }
  // on enlève le groupe du repas de départ
  // et on crée un nouveau repas à l'horaire choisi avec le groupe
  // en question
  async onDropHoraireHeader(event: DragEvent, horaire: string) {
    if (!event.dataTransfer || this.jourOffset == null) return;
    event.preventDefault();
    const data = getDragData(event.dataTransfer, DragKind.Groupe);

    await C.data.createRepas({
      id_sejour: C.state.idSejour!,
      offset_personnes: 0,
      horaire: horaire,
      jour_offset: this.jourOffset,
      groupes: [{ id_groupe: data.idGroupe, id_repas: -1 }],
      recettes: [],
      ingredients: []
    });

    if (C.notifications.getError() != null) return;
    data.repas.groupes = (data.repas.groupes || []).filter(
      g => g.id_groupe != data.idGroupe
    );
    await C.data.updateManyRepas([data.repas]);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Nouveau repas ajouté.");
    }
  }

  onDragoverRepas(event: DragEvent) {
    if (!event.dataTransfer) return;
    if (event.dataTransfer.types.includes(DragKind.Groupe)) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "move";
    }
  }
  // on enlève le groupe du repas de départ et on l'ajoute
  // au repas cible.
  async onDropRepas(event: DragEvent, target: RepasComplet) {
    if (!event.dataTransfer || this.jourOffset == null) return;
    const data = getDragData(event.dataTransfer, DragKind.Groupe);
    if (target.id == data.repas.id) return; // on déplace vers soi-même
    event.preventDefault();

    data.repas.groupes = (data.repas.groupes || []).filter(
      g => g.id_groupe != data.idGroupe
    );
    target = deepcopy(target); // force deepcopy
    target.groupes = (target.groupes || []).concat({
      id_repas: target.id,
      id_groupe: data.idGroupe
    });
    await C.data.updateManyRepas([data.repas, target]);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Groupe déplacé avec succès.");
    }
  }

  async deleteRepas(repas: RepasComplet) {
    await C.data.deleteRepas(repas);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Repas supprimé avec succès.");
    }
  }

  async resoudIngredients(repas: RepasComplet) {
    this.loadingIngredients = true;
    this.showPrevisuIngredients = true;
    const data = await C.calculs.resoudIngredientsRepas(repas.id);
    if (data == undefined || data.date_ingredients == null) return;
    this.repasNbPersonnes = C.getRepasNbPersonnes(repas);
    this.listeIngredients = data.date_ingredients[0].ingredients || [];
    this.loadingIngredients = false;
  }
}
</script>

<style></style>
