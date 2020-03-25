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
        <v-list-item-icon class="ml-0 mr-1 y-0 align-self-center"
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
import { RepasComplet, Menu } from "../../../logic/types";
import { CalendarMode, deepcopy, toNullableId } from "../../../logic/types2";
import { C } from "../../../logic/controller";
import { HorairesColors } from "../../utils/utils";
import { fmtHoraire } from "../../../logic/enums";
import {
  formatNbOffset,
  asRepasRecette,
  asRepasIngredient,
  compareRecettesIngredient
} from "./utils";

const ListeRepasProps = Vue.extend({
  props: {
    repass: Array as () => RepasComplet[],
    mode: String as () => CalendarMode
  }
});

@Component({})
export default class ListeRepas extends ListeRepasProps {
  repasTitle(repas: RepasComplet) {
    if (this.mode == "groupes") {
      const nbGroupes = C.getRepasGroupes(repas).length;
      return `${nbGroupes} gr.`;
    } else {
      return C.formatter.formatRepasName(repas);
    }
  }

  repasSubTitle(repas: RepasComplet) {
    if (this.mode == "groupes") {
      return formatNbOffset(repas);
    } else {
      const nbPersonnes = C.getRepasNbPersonnes(repas);
      return `${nbPersonnes} pers.`;
    }
  }

  getColorRepas(repas: RepasComplet) {
    return HorairesColors[repas.horaire];
  }

  getHoraireInitiale(repas: RepasComplet) {
    const horaire = fmtHoraire(repas.horaire);
    return horaire.substr(0, 2);
  }

  onDragstart(event: DragEvent, repas: RepasComplet) {
    if (event == null || event.dataTransfer == null) return;
    event.dataTransfer.setData("repas", JSON.stringify(repas));
    event.dataTransfer.effectAllowed = "linkMove";
  }

  // deux types de drop sont possibles :
  // - un autre repas pour échange
  // - un menu
  onDragover(event: DragEvent, target: RepasComplet) {
    if (!event.dataTransfer) return;
    if (event.dataTransfer.types.includes("repas")) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "link";
    } else if (event.dataTransfer.types.includes("menu")) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "link";
    }
  }

  // cf onDragover
  onDrop(event: DragEvent, target: RepasComplet) {
    if (!event.dataTransfer) return;
    if (event.dataTransfer.types.includes("repas")) {
      event.preventDefault();
      this.onDropRepas(event, target);
    } else if (event.dataTransfer.types.includes("menu")) {
      event.preventDefault();
      this.onDropMenu(event, target);
    }
  }
  private async onDropRepas(event: DragEvent, target: RepasComplet) {
    if (!event.dataTransfer) return;
    const origin: RepasComplet = JSON.parse(
      event.dataTransfer.getData("repas")
    );
    if (target.id == origin.id) return; // on évite les échanges inutiles
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

  private async onDropMenu(event: DragEvent, target: RepasComplet) {
    if (!event.dataTransfer) return;
    const menu: Menu = JSON.parse(event.dataTransfer.getData("menu"));
    if (compareRecettesIngredient(menu, target)) return; // on évite les requettes inutiles

    target = deepcopy(target); // on évite la modification locale
    // on copie le contenu du menu sur le repas
    target.recettes = (menu.recettes || []).map(r =>
      asRepasRecette(r, target.id)
    );
    target.ingredients = (menu.ingredients || []).map(r =>
      asRepasIngredient(r, target.id)
    );
    await C.data.updateManyRepas([target]);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage("Menu associé avec succès.");
    }
  }
}
</script>

<style scoped></style>
