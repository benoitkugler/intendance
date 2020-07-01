<template>
  <v-card @keyup.native.enter="$emit('edit', current)">
    <v-card-title primary-title>
      Détails de l'ingrédient
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-row>
          <v-col>
            <v-text-field
              label="Nom"
              v-model="current.nom"
              ref="nom"
            ></v-text-field>
          </v-col>
          <v-col>
            <unite-field v-model="current.unite"></unite-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="4">
            <v-switch
              label="Conditionnement requis"
              v-model="conditionnement"
            ></v-switch>
          </v-col>
          <v-col cols="8">
            <conditionnement-field
              v-model="current.conditionnement"
              :disabled="!conditionnement"
              :allowedUnites="allowedUnitesConditionnement"
            ></conditionnement-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-select
              :items="[]"
              v-model="current.categorie"
              label="Catégorie"
            ></v-select>
          </v-col>
          <v-col>
            <v-text-field
              label="Callories"
              v-model="current.callories"
              required
            ></v-text-field>
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="success" @click="$emit('edit', current)">{{
        btnTitle
      }}</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Watch } from "vue-property-decorator";

import UniteField from "../utils/UniteField.vue";
import ConditionnementField from "../utils/ConditionnementField.vue";

import { Ingredient, Unite, UniteLabels } from "../../logic/types";
import {
  IngredientOptions,
  EditMode,
  deepcopy,
  enumStringToOptions
} from "../../logic/types2";
import { DefautIngredient } from "./types";

const EditIngredientProps = Vue.extend({
  props: {
    initialIngredient: Object as () => IngredientOptions | null,
    mode: String as () => EditMode
  }
});

@Component({
  components: { UniteField, ConditionnementField }
})
export default class EditIngredient extends EditIngredientProps {
  current = this.copy(this.initialIngredient);

  private copy(initial: IngredientOptions | null): Ingredient {
    const ing =
      initial == null ? { ...DefautIngredient, id: -1 } : initial.ingredient;
    return deepcopy(ing);
  }

  get btnTitle() {
    return this.mode == "edit" ? "Valider" : "Créer";
  }

  get conditionnement() {
    return this.current.conditionnement.unite != Unite.Zero;
  }
  set conditionnement(b: boolean) {
    if (!b) {
      this.current.conditionnement = { unite: Unite.Zero, quantite: 0 };
    } else {
      this.current.conditionnement.unite =
        this.current.unite ||
        (this.initialIngredient && this.initialIngredient.ingredient.unite) ||
        Unite.Litres;
    }
  }

  get allowedUnitesConditionnement() {
    const items = enumStringToOptions(UniteLabels);
    if (this.current.unite == Unite.Piece) {
      // Pour le conditionnement, l'unité Pièce n'apporte aucune information
      return items.filter(u => u.value != Unite.Piece);
    }
    // le conditionnement doit être compatible avec l'unité.
    return items.filter(u => u.value == this.current.unite);
  }

  @Watch("initialIngredient")
  onPropChange() {
    this.current = this.copy(this.initialIngredient);
  }

  $refs!: {
    nom: Vue;
  };

  focus() {
    const input = this.$refs.nom.$el.querySelector("input");
    if (input != null) input.select();
  }
}
</script>

<style scoped></style>
