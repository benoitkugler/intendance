<template>
  <div>
    <v-dialog v-model="showPrevisuIngredients" max-width="600px">
      <v-skeleton-loader type="card" :loading="loadingIngredients">
        <v-card class="py-2">
          <v-card-title primary-title>
            <h3 class="headline mb-0">
              Ingrédients pour {{ repas.offset_personnes }} personne{{
                repas.offset_personnes > 1 ? "s" : ""
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
          <v-autocomplete
            label="Menu"
            :items="menus"
            v-model="idMenu"
          ></v-autocomplete>
          <horaire-field v-model="repas.horaire"></horaire-field>
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
        <tooltip-btn
          v-if="mode == 'edit'"
          mdi-icon="food-variant"
          tooltip="Calculer les <b>ingrédients</b> nécessaires au repas..."
          @click="resoudIngredients"
        ></tooltip-btn>
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
  RepasWithGroupe,
  DateIngredientQuantites,
  IngredientQuantite,
  RepasGroupe
} from "../../logic/types";
import { New, DetailsRepas, EditMode, toNullableId } from "../../logic/types2";
import DateField from "../utils/DateField.vue";
import HoraireField from "../utils/HoraireField.vue";
import TooltipBtn from "../utils/TooltipBtn.vue";
import ListeIngredients from "./ListeIngredients.vue";
import { C } from "../../logic/controller";
import { Watch } from "vue-property-decorator";
import { Horaires } from "../../logic/enums";
import { Formatter } from "../../logic/formatter";
import { fmtHoraire } from "../../logic/enums";

const Props = Vue.extend({
  props: {
    initialRepas: Object as () => RepasWithGroupe,
    mode: String as () => EditMode
  }
});

@Component({
  components: { DateField, HoraireField, TooltipBtn, ListeIngredients }
})
export default class FormRepas extends Props {
  repas: DetailsRepas = JSON.parse(JSON.stringify(this.initialRepas));

  horaires = Horaires;

  showPrevisuIngredients = false;
  loadingIngredients = true;
  listeIngredients: IngredientQuantite[] = [];

  @Watch("initialRepas")
  onPropChange() {
    this.repas = JSON.parse(JSON.stringify(this.initialRepas));
  }

  get idMenu() {
    return this.repas.id_menu.Valid ? this.repas.id_menu.Int64 : -1;
  }
  set idMenu(id: number) {
    this.repas.id_menu = toNullableId(id);
  }

  get sejour() {
    return C.data.sejours.sejours[this.initialRepas.id_sejour];
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

  async resoudIngredients() {
    this.loadingIngredients = true;
    this.showPrevisuIngredients = true;
    const data = await C.calculs.resoudIngredientsRepas(
      this.initialRepas.id,
      this.repas.offset_personnes
    );
    if (data == undefined || data.date_ingredients == null) return;
    this.listeIngredients = data.date_ingredients[0].ingredients || [];
    this.loadingIngredients = false;
  }
}
</script>

<style></style>
