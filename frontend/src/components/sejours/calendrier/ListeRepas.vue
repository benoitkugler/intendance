<template>
  <v-list dense style="border-radius: 0;">
    <v-list-item-group>
      <v-list-item
        v-for="repas in repass"
        :key="repas.id"
        @click.stop="$emit('edit', repas)"
        class="px-1"
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
import { CalendarMode } from "../../../logic/types2";
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
}
</script>

<style scoped></style>
