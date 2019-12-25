<template>
  <v-card>
    <v-card-title primary-title>
      Détails de l'ingrédient
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-row>
          <v-col>
            <v-text-field
              label="Quantité"
              hint="Quantité désirée pour 1 personne."
              type="number"
              v-model.number="current.quantite"
              ref="quantite"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-select
              :items="[]"
              v-model="current.cuisson"
              label="Cuisson"
            ></v-select>
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
import { RecetteIngredient } from "../../logic/types";
import { Watch } from "vue-property-decorator";

const MenuOrRecetteIngredientProps = Vue.extend({
  props: {
    ingredient: Object as () => RecetteIngredient
  }
});

@Component({})
export default class MenuOrRecetteIngredient extends MenuOrRecetteIngredientProps {
  current: RecetteIngredient = JSON.parse(JSON.stringify(this.ingredient));

  @Watch("ingredient")
  onPropChange() {
    this.current = JSON.parse(JSON.stringify(this.ingredient));
  }

  $refs!: {
    quantite: Vue;
  };

  focus() {
    const input = this.$refs.quantite.$el.querySelector("input");
    if (input != null) input.select();
  }
}
</script>

<style scoped></style>
