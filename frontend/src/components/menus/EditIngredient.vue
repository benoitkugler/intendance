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
          <v-col>
            <v-switch
              label="Conditionnement requis"
              v-model="conditionnement"
            ></v-switch>
          </v-col>
          <v-col>
            <v-text-field
              :disabled="!conditionnement"
              label="Conditionnement - Quanité"
              type="number"
              v-model.number="current.conditionnement.quantite"
            ></v-text-field>
          </v-col>
          <v-col>
            <unite-field
              :disabled="!conditionnement"
              v-model="current.conditionnement.unite"
              label="Conditionnement - Unité"
              :unites="allowedUnitesConditionnement"
            ></unite-field>
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
import { RecetteIngredient, Ingredient } from "../../logic/types";
import { Watch } from "vue-property-decorator";
import { IngredientOptions, EditMode } from "../../logic/types2";
import UniteField from "../utils/UniteField.vue";
import { Unites, UniteFields } from "../../logic/enums";
import { DefautIngredient } from "./types";

const EditIngredientProps = Vue.extend({
  props: {
    initialIngredient: Object as () => IngredientOptions | null,
    mode: String as () => EditMode
  }
});

@Component({
  components: { UniteField }
})
export default class EditIngredient extends EditIngredientProps {
  current = this.copy(this.initialIngredient);

  private copy(initial: IngredientOptions | null): Ingredient {
    const ing = initial == null ? DefautIngredient : initial.ingredient;
    return JSON.parse(JSON.stringify(ing));
  }

  get btnTitle() {
    return this.mode == "edit" ? "Valider" : "Créer";
  }

  get conditionnement() {
    return this.current.conditionnement.unite != "";
  }
  set conditionnement(b: boolean) {
    if (!b) {
      this.current.conditionnement = { unite: "", quantite: 0 };
    } else {
      this.current.conditionnement.unite =
        this.current.unite ||
        (this.initialIngredient && this.initialIngredient.ingredient.unite) ||
        "L";
    }
  }

  get allowedUnitesConditionnement() {
    if (this.current.unite == UniteFields.Piece) {
      // Pour le conditionnement, l'unité Pièce n'apporte aucune information
      return Unites.filter(u => u.value != UniteFields.Piece);
    }
    // le conditionnement doit être compatible avec l'unité.
    return Unites.filter(u => u.value == this.current.unite);
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
