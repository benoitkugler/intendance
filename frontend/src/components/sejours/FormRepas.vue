<template>
  <div>
    <v-dialog v-model="showPrevisuIngredients" max-width="600px">
      <v-skeleton-loader type="card" :loading="loadingIngredients">
        <v-card class="py-2">
          <v-card-title primary-title>
            <h3 class="headline mb-0">
              Ingrédients pour {{ repas.nb_personnes }} personne{{
                repas.nb_personnes > 1 ? "s" : ""
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
        Détails du repas
        <v-spacer></v-spacer>
        <div>
          <small v-if="sejour">Sejour : {{ sejour.nom }} </small>
        </div>
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
          <v-select
            :items="horaires"
            v-model="repas.horaire"
            label="Horaire"
            :placeholder="horaireFormatted"
          ></v-select>
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
  Repas,
  DateIngredientQuantites,
  IngredientQuantite
} from "../../logic/types";
import { New, DetailsRepas, EditMode } from "../../logic/types2";
import DateField from "../utils/DateField.vue";
import TooltipBtn from "../utils/TooltipBtn.vue";
import ListeIngredients from "./ListeIngredients.vue";
import { C } from "../../logic/controller";
import { Watch } from "vue-property-decorator";
import { Horaires } from "../utils/enums";
import { Formatter } from "../../logic/formatter";

const Props = Vue.extend({
  props: {
    initialRepas: Object as () => Repas,
    mode: String as () => EditMode
  }
});

@Component({
  components: { DateField, TooltipBtn, ListeIngredients }
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

  get sejour() {
    return C.data.sejours.sejours[this.initialRepas.id_sejour];
  }

  get menus() {
    return Object.values(C.data.menus).map(menu => {
      return { text: C.formatter.formatMenuName(menu), value: menu.id };
    });
  }

  get horaireFormatted() {
    const horaire = this.repas.horaire;
    return `(Personnalisé) - ${Formatter.horaireToTime(horaire)}`;
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
