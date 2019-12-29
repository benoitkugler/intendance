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
      <v-btn color="success" @click="$emit('edit', current)">Valider</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { RecetteIngredient, Ingredient } from "../../logic/types";
import { Watch } from "vue-property-decorator";
import { IngredientOptions } from "../../logic/types2";
import UniteField from "../utils/UniteField.vue";
import { Unites, UnitePiece } from "../utils/enums";

const EditIngredientProps = Vue.extend({
  props: {
    initialIngredient: Object as () => IngredientOptions
  }
});

@Component({
  components: { UniteField }
})
export default class EditIngredient extends EditIngredientProps {
  current: Ingredient = JSON.parse(
    JSON.stringify(this.initialIngredient.ingredient)
  );

  get conditionnement() {
    return this.current.conditionnement.unite != "";
  }
  set conditionnement(b: boolean) {
    if (!b) {
      this.current.conditionnement = { unite: "", quantite: 0 };
    } else {
      this.current.conditionnement.unite =
        this.current.unite || this.initialIngredient.ingredient.unite || "L";
    }
  }

  get allowedUnitesConditionnement() {
    if (this.current.unite == UnitePiece) {
      // Pour le conditionnement, l'unité Pièce n'apporte aucune information
      return Unites.filter(u => u.value != UnitePiece);
    }
    // le conditionnement doit être compatible avec l'unité.
    return Unites.filter(u => u.value == this.current.unite);
  }

  @Watch("initialIngredient")
  onPropChange() {
    this.current = JSON.parse(
      JSON.stringify(this.initialIngredient.ingredient)
    );
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
