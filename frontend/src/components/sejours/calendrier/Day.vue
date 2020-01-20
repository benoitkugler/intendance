<template>
  <div>
    <v-toolbar dense>
      <v-toolbar-title v-if="day">
        Repas du
        {{
          day.toLocaleDateString("fr-FR", {
            weekday: "long",
            day: "numeric",
            month: "short"
          })
        }}</v-toolbar-title
      >
    </v-toolbar>
    <div class="overflow-y-auto">
      <v-list dense>
        <v-list-item-group color="primary">
          <div v-for="horaire in horaires" :key="horaire.value">
            <v-hover v-slot="{ hover }">
              <v-subheader
                :style="{ color: getHoraireColor(horaire.value) }"
                @click="startCreateRepas(horaire.value)"
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
                    </v-row>
                  </v-list-item-content>
                  <v-list-item-action class="my-1">
                    <tooltip-btn
                      v-if="hover"
                      mdi-icon="close"
                      color="red"
                      small
                      tooltip="Supprimer ce repas..."
                      @click.stop="deleteRepas(repas)"
                    ></tooltip-btn>
                  </v-list-item-action>
                </v-list-item>
              </v-hover>
            </template>
          </div>
        </v-list-item-group>
      </v-list>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import TooltipBtn from "../../utils/TooltipBtn.vue";

import { toDateVuetify, formatNbOffset } from "./utils";
import { C } from "../../../logic/controller";
import { RepasWithGroupe, Groupe } from "../../../logic/types";
import { New, NullId, deepcopy } from "../../../logic/types2";
import { Horaires } from "../../../logic/enums";
import { HorairesColors } from "../../utils/utils";

interface dragData {
  idGroupe: number;
  repas: RepasWithGroupe;
}

const DayProps = Vue.extend({
  props: {
    jourOffset: Number as () => number | null
  }
});
@Component({
  components: { TooltipBtn }
})
export default class Day extends DayProps {
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

  get events(): { [key: string]: RepasWithGroupe[] } {
    const out: { [key: string]: RepasWithGroupe[] } = {};
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

  getGroupes(repas: RepasWithGroupe) {
    const grs = C.getRepasGroupes(repas);
    const maxChar = 8;
    return grs.map(groupe => {
      const g: Groupe = deepcopy(groupe);
      g.nom = g.nom.substr(0, maxChar) + (g.nom.length <= maxChar ? "" : ".");
      return groupe;
    });
  }

  // Drag and drop

  onDragStart(event: DragEvent, repas: RepasWithGroupe, groupe: Groupe) {
    if (event == null || event.dataTransfer == null) return;
    const data: dragData = { repas: repas, idGroupe: groupe.id };
    event.dataTransfer.setData("groupe", JSON.stringify(data));
    event.dataTransfer.effectAllowed = "copyMove";
  }

  onDragoverHoraireHeader(event: DragEvent) {
    if (!event.dataTransfer) return;
    if (event.dataTransfer.types.includes("groupe")) {
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
    const data: dragData = JSON.parse(event.dataTransfer.getData("groupe"));

    await C.data.createRepas({
      id_sejour: C.state.idSejour!,
      id_menu: NullId,
      offset_personnes: 0,
      horaire: horaire,
      jour_offset: this.jourOffset,
      groupes: [{ id_groupe: data.idGroupe, id_repas: -1 }]
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
    if (event.dataTransfer.types.includes("groupe")) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "move";
    }
  }
  // on enlève le groupe du repas de départ et on l'ajoute
  // au repas cible.
  async onDropRepas(event: DragEvent, target: RepasWithGroupe) {
    if (!event.dataTransfer || this.jourOffset == null) return;
    const data: dragData = JSON.parse(event.dataTransfer.getData("groupe"));
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

  async deleteRepas(repas: RepasWithGroupe) {
    await C.data.deleteRepas(repas);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Repas supprimé avec succès.");
    }
  }
}
</script>

<style></style>
