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
              hideLinks
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
                <v-subheader
                  :style="{ color: getHoraireColor(horaire.value) }"
                  @dragover="onDragoverHoraireHeader($event)"
                  @drop="onDropHoraireHeader($event, horaire.value)"
                >
                  <v-icon
                    small
                    :color="getHoraireColor(horaire.value)"
                    class="mr-2"
                    >mdi-{{ getHoraireIcon(horaire.value) }}</v-icon
                  >
                  {{ horaire.text }}
                  <v-spacer></v-spacer>
                  <tooltip-btn
                    mdi-icon="plus"
                    color="green"
                    tooltip="Ajouter un repas..."
                    @click="$emit('addRepas', horaire.value)"
                  ></tooltip-btn>
                </v-subheader>
                <v-list-item
                  v-for="repas in events[horaire.value]"
                  @dragover="onDragoverRepas($event)"
                  @drop="onDropRepas($event, repas)"
                  @click="$emit('editRepas', repas)"
                  :key="repas.id"
                  :class="repas.anticipation == 0 ? '' : colorAnticipation"
                >
                  <v-list-item-content>
                    <v-row no-gutters class="fill-height">
                      <v-col
                        md="3"
                        class="px-1 align-self-center overflow-x-auto"
                      >
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
                      <v-col md="6" class="align-self-center">
                        <case-recettes
                          :recettes="repas.recettes"
                          @add="idRecette => addRecette(repas, idRecette)"
                          @remove="idRecette => removeRecette(repas, idRecette)"
                        ></case-recettes>
                      </v-col>
                      <v-col class="align-self-center">
                        <liste-lien-ingredients
                          chips
                          v-model="repas.ingredients"
                          @change="updateIngredients(repas)"
                        ></liste-lien-ingredients>
                      </v-col>
                    </v-row>
                  </v-list-item-content>
                  <v-list-item-action class="my-1">
                    <v-row no-gutters>
                      <v-col>
                        <tooltip-btn
                          mdi-icon="food-variant"
                          small
                          tooltip="Calculer les <b>ingrédients</b> nécessaires au repas..."
                          @click.stop="resoudIngredients(repas)"
                        ></tooltip-btn
                      ></v-col>
                      <v-col>
                        <tooltip-btn
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
              </div>
            </v-list-item-group>
          </v-list></div
      ></v-col>
      <v-col class="px-2">
        <v-expansion-panels accordion :value="0">
          <choix-menus height="45vh"></choix-menus>
          <choix-recettes height="45vh"></choix-recettes>
          <choix-ingredients height="45vh"></choix-ingredients>
          <v-expansion-panel></v-expansion-panel>
        </v-expansion-panels>
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
import ChoixRecettes from "./ChoixRecettes.vue";
import CaseRecettes from "./CaseRecettes.vue";

import {
  toDateVuetify,
  formatNbOffset,
  compareRecettesIngredient,
  ColorAnticipation
} from "../utils";
import { C } from "../../../../logic/controller";
import {
  RepasComplet,
  Groupe,
  IngredientQuantite,
  Horaire,
  HoraireLabels
} from "../../../../logic/types";
import { New, deepcopy, enumIntToOptions } from "../../../../logic/types2";
import { HorairesColors, HorairesIcons } from "../../../utils/utils";
import { DragKind, getDragData, setDragData } from "../../../utils/utils_drag";
import ChoixIngredients from "./ChoixIngredients.vue";
import ListeLienIngredients from "../../../utils/ListeLienIngredients.vue";

const DayProps = Vue.extend({
  props: {
    jourOffset: Number as () => number | null
  }
});
@Component({
  components: {
    TooltipBtn,
    ListeIngredients,
    ChoixMenus,
    ChoixRecettes,
    ChoixIngredients,
    CaseRecettes,
    ListeLienIngredients
  }
})
export default class Day extends DayProps {
  colorAnticipation = ColorAnticipation;
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
    return enumIntToOptions(HoraireLabels);
  }

  getHoraireColor(h: Horaire) {
    return HorairesColors[h];
  }
  getHoraireIcon(h: Horaire) {
    return HorairesIcons[h];
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
  async onDropHoraireHeader(event: DragEvent, horaire: Horaire) {
    if (!event.dataTransfer || this.jourOffset == null) return;
    event.preventDefault();
    const data = getDragData(event.dataTransfer, DragKind.Groupe);

    await C.data.createRepas({
      id_sejour: C.state.idSejour!,
      offset_personnes: 0,
      horaire: horaire,
      jour_offset: this.jourOffset,
      anticipation: data.repas.anticipation,
      groupes: [{ id_groupe: data.idGroupe, id_repas: -1 }],
      recettes: data.repas.recettes,
      ingredients: data.repas.ingredients
    });

    if (C.notifications.getError() != null) return;
    data.repas.groupes = (data.repas.groupes || []).filter(
      g => g.id_groupe != data.idGroupe
    );
    if (data.repas.groupes.length == 0 && data.repas.offset_personnes == 0) {
      // le repas est maintenant vide, on le supprime
      await C.data.deleteRepas(data.repas);
    } else {
      await C.data.updateManyRepas([data.repas]);
    }
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Nouveau repas ajouté.");
    }
  }

  // deux drops possibles:
  // - groupe
  // - menu
  onDragoverRepas(event: DragEvent) {
    if (!event.dataTransfer) return;
    if (event.dataTransfer.types.includes(DragKind.Groupe)) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "move";
    } else if (event.dataTransfer.types.includes(DragKind.Menu)) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "link";
    }
  }

  onDropRepas(event: DragEvent, target: RepasComplet) {
    if (!event.dataTransfer || this.jourOffset == null) return;
    if (event.dataTransfer.types.includes(DragKind.Groupe)) {
      event.preventDefault();
      this.onDropGroupe(event.dataTransfer, target);
    } else if (event.dataTransfer.types.includes(DragKind.Menu)) {
      event.preventDefault();
      this.onDropMenu(event.dataTransfer, target);
    }
  }

  // on enlève le groupe du repas de départ et on l'ajoute
  // au repas cible.
  private async onDropGroupe(dataTransfer: DataTransfer, target: RepasComplet) {
    const data = getDragData(dataTransfer, DragKind.Groupe);
    if (target.id == data.repas.id) return; // on déplace vers soi-même

    // on enlève le groupe
    data.repas.groupes = (data.repas.groupes || []).filter(
      g => g.id_groupe != data.idGroupe
    );

    // on l'ajoute à la cible
    target = deepcopy(target); // force deepcopy
    target.groupes = (target.groupes || []).concat({
      id_repas: target.id,
      id_groupe: data.idGroupe
    });

    if (data.repas.groupes.length == 0 && data.repas.offset_personnes == 0) {
      // le repas est maintenant vide, on le supprime
      await Promise.all([
        C.data.deleteRepas(data.repas),
        C.data.updateManyRepas([target])
      ]);
    } else {
      await C.data.updateManyRepas([data.repas, target]);
    }
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Groupe déplacé avec succès.");
    }
  }

  private async onDropMenu(dataTransfer: DataTransfer, target: RepasComplet) {
    const menu = getDragData(dataTransfer, DragKind.Menu);

    if (compareRecettesIngredient(menu, target)) return; // on évite les requettes inutiles

    target = deepcopy(target); // on évite la modification locale
    // on copie le contenu du menu sur le repas
    target.recettes = menu.recettes;
    target.ingredients = menu.ingredients;
    await C.data.updateManyRepas([target]);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Menu associé avec succès.");
    }
  }

  async addRecette(repas: RepasComplet, idRecette: number) {
    // check pour éviter une requête inutile
    if ((repas.recettes || []).includes(idRecette)) return;

    repas.recettes = (repas.recettes || []).concat(idRecette);
    await C.data.updateManyRepas([repas]);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Recette ajoutée avec succès.");
    }
  }

  async removeRecette(repas: RepasComplet, toRemove: number) {
    repas.recettes = (repas.recettes || []).filter(idR => idR != toRemove);
    await C.data.updateManyRepas([repas]);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Recette enlevée avec succès.");
    }
  }

  async updateIngredients(repas: RepasComplet) {
    await C.data.updateManyRepas([repas]);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Ingredients mis à jour avec succès.");
    }
  }

  // async addIngredient(repas: RepasComplet, idIngredient: number) {
  //   // check pour éviter une requête inutile
  //   if (
  //     (repas.ingredients || [])
  //       .map(ing => ing.id_ingredient)
  //       .includes(idIngredient)
  //   )
  //     return;

  //   repas.ingredients = (repas.ingredients || []).concat({
  //     id_ingredient: idIngredient,
  //     quantite: 0,
  //     cuisson: ""
  //   });
  //   await C.data.updateManyRepas([repas]);
  //   if (C.notifications.getError() == null) {
  //     C.notifications.setMessage("Ingredient ajouté avec succès.");
  //   }
  // }

  // async removeIngredient(repas: RepasComplet, toRemove: number) {
  //   repas.ingredients = (repas.ingredients || []).filter(
  //     ing => ing.id_ingredient != toRemove
  //   );
  //   await C.data.updateManyRepas([repas]);
  //   if (C.notifications.getError() == null) {
  //     C.notifications.setMessage("Ingredient enlevé avec succès.");
  //   }
  // }

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
