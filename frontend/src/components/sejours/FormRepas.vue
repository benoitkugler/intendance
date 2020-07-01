<template>
  <div>
    <v-card>
      <v-card-title primary-title>
        {{ mode == "new" ? "Détails du nouveau repas" : "Détails du repas" }}
        <v-spacer></v-spacer>
        <div>
          <small v-if="sejour">Sejour : {{ sejour.nom }} </small>
        </div>
      </v-card-title>
      <v-card-text>
        <v-form>
          <v-row>
            <v-col md="6">
              <v-select
                :items="groupes"
                v-model="repas.groupes"
                label="Groupes"
                chips
                multiple
              ></v-select>
              <v-text-field
                label="Nombre additionnel de personnes "
                v-model.number="repas.offset_personnes"
                type="number"
                hint="S'ajoute aux groupes. Peut être négatif."
              ></v-text-field>
              <horaire-field v-model="repas.horaire"></horaire-field>
              <v-text-field
                label="Anticiper l'utilisation"
                v-model.number="repas.anticipation"
                type="number"
                suffix="jour(s)"
                hint="Un nombre positif avance le repas lors de la résolution des ingrédients."
              ></v-text-field>
            </v-col>
            <v-col md="6">
              <recettes-fields v-model="repas.recettes"></recettes-fields>
              <liste-lien-ingredients
                class="ml-n2"
                subheader="Ingrédients"
                showAdd
                v-model="repas.ingredients"
              ></liste-lien-ingredients>
            </v-col>
          </v-row>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-btn
          v-if="mode == 'edit'"
          color="error"
          @click="$emit('delete', repas)"
          >Supprimer</v-btn
        >
        <v-spacer></v-spacer>
        <v-btn color="success" @click="$emit('accept', repas)">{{
          mode == "edit" ? "Enregistrer" : "Ajouter"
        }}</v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import {
  RepasComplet,
  DateIngredientQuantites,
  IngredientQuantite,
  RepasGroupe
} from "../../logic/types";
import {
  New,
  DetailsRepas,
  EditMode,
  toNullableId,
  deepcopy
} from "../../logic/types2";
import DateField from "../utils/DateField.vue";
import HoraireField from "../utils/HoraireField.vue";
import TooltipBtn from "../utils/TooltipBtn.vue";
import { C } from "../../logic/controller";
import { Watch } from "vue-property-decorator";
import { Formatter } from "../../logic/formatter";
import RecettesFields from "../utils/RecettesFields.vue";
import ListeLienIngredients from "../utils/ListeLienIngredients.vue";

const Props = Vue.extend({
  props: {
    initialRepas: Object as () => RepasComplet,
    mode: String as () => EditMode
  }
});

@Component({
  components: {
    DateField,
    HoraireField,
    TooltipBtn,
    RecettesFields,
    ListeLienIngredients
  }
})
export default class FormRepas extends Props {
  repas: DetailsRepas = deepcopy(this.initialRepas);

  @Watch("initialRepas")
  onPropChange() {
    this.repas = deepcopy(this.initialRepas);
  }

  get sejour() {
    return (C.data.sejours.sejours || {})[this.initialRepas.id_sejour];
  }

  get menus() {
    return Object.values(C.data.menus).map(menu => {
      return { text: C.formatter.formatMenuName(menu), value: menu.id };
    });
  }

  get groupes() {
    return C.state.getGroupes().map(groupe => {
      const rg: RepasGroupe = {
        id_repas: this.initialRepas.id,
        id_groupe: groupe.id
      };
      return { text: groupe.nom, value: rg };
    });
  }
}
</script>

<style></style>
