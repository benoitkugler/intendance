<template>
  <v-card>
    <v-card-title primary-title>
      Détails du repas
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-text-field
          label="Nombre de personnes"
          v-model.number="repas.nb_personnes"
          type="number"
        ></v-text-field>
        <v-autocomplete
          label="Menu"
          :items="menus"
          v-model="repas.id_menu"
        ></v-autocomplete>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="success" @click="$emit('accept', repas)">Enregistrer</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Repas } from "../../logic/types";
import { New, DetailsRepas } from "../../logic/types2";
import DateField from "../utils/DateField.vue";
import { D } from "../../logic/controller";
import { formatMenuName } from "../../logic/format";

const Props = Vue.extend({
  props: {
    initialRepas: Object as () => Repas
  }
});

@Component({
  components: { DateField }
})
export default class FormRepas extends Props {
  repas: DetailsRepas = JSON.parse(JSON.stringify(this.initialRepas));

  get menus() {
    return Object.values(D.menus).map(menu => {
      return { text: formatMenuName(menu), value: menu.id };
    });
  }
}
</script>

<style></style>
