<template>
  <v-list dense style="border-radius: 0" v-if="repass">
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
        :class="repas.anticipation == 0 ? '' : colorAnticipation"
      >
        <v-list-item-icon class="ml-0 mr-1 y-0 align-self-center">
          <v-icon :color="getColorRepas(repas)"
            >mdi-{{ getHoraireIcon(repas) }}</v-icon
          >
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
import { RepasComplet, Menu, MenuComplet } from "@/logic/api";
import { deepcopy, toNullableId } from "@/logic/types";
import { Controller } from "@/logic/controller";
import { HorairesColors, HorairesIcons } from "../../utils/utils";
import { DragKind, getDragData, setDragData } from "../../utils/utils_drag";
import {
  formatNbOffset,
  compareRecettesIngredient,
  ColorAnticipation,
} from "./utils";

const ListeRepasProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    repass: Array as () => RepasComplet[],
  },
});
@Component({})
export default class ListeRepas extends ListeRepasProps {
  colorAnticipation = ColorAnticipation;

  repasTitle(repas: RepasComplet) {
    return this.C.formatter.formatRepasName(repas);
  }

  repasSubTitle(repas: RepasComplet) {
    const nbPersonnes = this.C.getRepasNbPersonnes(repas);
    return `${nbPersonnes} personne(s)`;
  }

  getColorRepas(repas: RepasComplet) {
    return HorairesColors[repas.horaire];
  }

  getHoraireIcon(repas: RepasComplet) {
    return HorairesIcons[repas.horaire];
  }

  onDragstart(event: DragEvent, repas: RepasComplet) {
    if (event == null || event.dataTransfer == null) return;
    setDragData(event.dataTransfer, DragKind.Repas, repas);
    event.dataTransfer.effectAllowed = "linkMove";
  }

  // drop un repas pour échanger
  onDragover(event: DragEvent, target: RepasComplet) {
    if (!event.dataTransfer) return;
    if (event.dataTransfer.types.includes(DragKind.Repas)) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "link";
    }
  }

  onDrop(event: DragEvent, target: RepasComplet) {
    if (!event.dataTransfer) return;
    if (event.dataTransfer.types.includes(DragKind.Repas)) {
      event.preventDefault();
      this.onDropRepas(event, target);
    }
  }
  private async onDropRepas(event: DragEvent, target: RepasComplet) {
    if (!event.dataTransfer) return;
    const origin = getDragData(event.dataTransfer, DragKind.Repas);

    if (target.id == origin.id) return; // on évite les échanges inutiles
    target = deepcopy(target);
    [origin.jour_offset, target.jour_offset] = [
      target.jour_offset,
      origin.jour_offset,
    ];
    await this.C.api.UpdateManyRepas([target, origin]);
  }
}
</script>

<style scoped></style>
