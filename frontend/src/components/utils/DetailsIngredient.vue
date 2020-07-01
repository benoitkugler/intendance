<template>
  <v-card @keyup.native.enter="$emit('edit', current)">
    <v-card-title primary-title>
      Détails de l'ingrédient
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-row>
          <v-col>
            <quantite-relative
              v-model="current.quantite"
              ref="quantite"
              :unite="uniteHint"
            ></quantite-relative>
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
import { LienIngredient, UniteLabels } from "../../logic/types";
import { Watch } from "vue-property-decorator";
import { deepcopy } from "../../logic/types2";
import QuantiteRelative from "./QuantiteRelative.vue";
import { C } from "../../logic/controller";

const DetailsIngredientProps = Vue.extend({
  props: {
    ingredient: Object as () => LienIngredient
  }
});

@Component({
  components: { QuantiteRelative }
})
export default class DetailsIngredient extends DetailsIngredientProps {
  current: LienIngredient = deepcopy(this.ingredient);

  @Watch("ingredient")
  onPropChange() {
    this.current = deepcopy(this.ingredient);
  }

  $refs!: {
    quantite: QuantiteRelative;
  };

  get uniteHint() {
    return UniteLabels[C.getIngredient(this.ingredient.id_ingredient).unite];
  }

  focus() {
    const input = this.$refs.quantite.$el.querySelector("input");
    if (input != null) input.select();
  }
}
</script>

<style scoped></style>
