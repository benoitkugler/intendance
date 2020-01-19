<template>
  <v-list dense style="border-radius: 0;" v-if="repass">
    <v-list-item-group>
      <v-list-item
        v-for="repas in repass"
        :key="repas.id"
        class="px-1"
        @click.stop="$emit('edit', repas)"
        draggable
        @dragstart="onDragstart($event, repas)"
        @dragover.stop="onDragover($event, repas)"
        @drop.stop="onDrop($event, repas)"
      >
        <v-list-item-icon class="mx-0"
          ><v-chip label small :color="getColorRepas(repas)" class="px-1">
            {{ getHoraireInitiale(repas) }}
          </v-chip>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>
            {{ repasTitle(repas) }}
          </v-list-item-title>
          <v-list-item-subtitle>
            {{ repasSubTitle(repas) }}
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
    </v-list-item-group>
  </v-list>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { RepasWithGroupe } from "../../../logic/types";
import { CalendarMode, deepcopy } from "../../../logic/types2";
import { C } from "../../../logic/controller";
import { HorairesColors } from "../../utils/utils";
import { fmtHoraire } from "../../../logic/enums";
import { formatNbOffset } from "./utils";

const ListeRepasProps = Vue.extend({
  props: {
    repass: Array as () => RepasWithGroupe[],
    mode: String as () => CalendarMode
  }
});

@Component({})
export default class ListeRepas extends ListeRepasProps {
  repasTitle(repas: RepasWithGroupe) {
    if (this.mode == "groupes") {
      const nbGroupes = C.getRepasGroupes(repas).length;
      return `${nbGroupes} gr.`;
    } else {
      return C.formatter.formatRepasName(repas);
    }
  }

  repasSubTitle(repas: RepasWithGroupe) {
    return formatNbOffset(repas);
  }

  getColorRepas(repas: RepasWithGroupe) {
    return HorairesColors[repas.horaire];
  }

  getHoraireInitiale(repas: RepasWithGroupe) {
    const horaire = fmtHoraire(repas.horaire);
    return horaire.substr(0, 2);
  }

  onDragstart(event: DragEvent, repas: RepasWithGroupe) {
    if (event == null || event.dataTransfer == null) return;
    event.dataTransfer.setData("repas", JSON.stringify(repas));
    event.dataTransfer.effectAllowed = "linkMove";
  }

  onDragover(event: DragEvent, target: RepasWithGroupe) {
    if (!event.dataTransfer) return;
    if (event.dataTransfer.types.includes("repas")) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "link";
    }
  }

  // on échange les deux repas
  async onDrop(event: DragEvent, target: RepasWithGroupe) {
    if (!event.dataTransfer) return;
    event.preventDefault();
    const origin: RepasWithGroupe = JSON.parse(
      event.dataTransfer.getData("repas")
    );
    target = deepcopy(target);
    [origin.jour_offset, target.jour_offset] = [
      target.jour_offset,
      origin.jour_offset
    ];
    [origin.horaire, target.horaire] = [target.horaire, origin.horaire];
    await C.data.updateManyRepas([target, origin]);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Repas échangés avec succès.");
    }
  }
}
</script>

<style scoped></style>
