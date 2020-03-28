<template>
  <div @dragover.stop="onDragoverIngredients" @drop.stop="onDropIngredient">
    <v-chip
      small
      v-for="ingredient in ingredients"
      :key="ingredient.id_ingredient"
      close
      @click.stop
      @click:close="$emit('remove', ingredient.id_ingredient)"
    >
      {{ formatIngredient(ingredient.id_ingredient) }}
    </v-chip>
    <small class="font-italic" v-if="(ingredients || []).length == 0"
      >Déposez un ingredient ici...</small
    >
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { C } from "../../../../logic/controller";
import { DragKind, getDragData } from "../../../utils/utils_drag";
import { LienIngredient } from "../../../../logic/types";
const CaseIngredientsProps = Vue.extend({
  props: {
    ingredients: Array as () => LienIngredient[] | null
  }
});

@Component({})
export default class CaseIngredients extends CaseIngredientsProps {
  formatIngredient(idIngredient: number) {
    return C.getIngredient(idIngredient).nom;
  }

  onDragoverIngredients(event: DragEvent) {
    if (!event.dataTransfer) return;
    const isIngredient = event.dataTransfer.types.includes(
      DragKind.IdIngredient
    );
    if (isIngredient) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "copy";
    }
  }

  onDropIngredient(event: DragEvent) {
    if (!event.dataTransfer) return;
    event.preventDefault();
    const idIngredient = getDragData(event.dataTransfer, DragKind.IdIngredient);
    this.$emit("add", idIngredient);
  }
}
</script>

<style scoped></style>
