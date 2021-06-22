<template>
  <div @dragover.stop="onDragoverRecettes" @drop.stop="onDropRecette">
    <v-chip
      small
      v-for="idRecette in recettes || []"
      :key="idRecette"
      close
      @click.stop
      @click:close="$emit('remove', idRecette)"
    >
      {{ formatRecette(idRecette) }}
    </v-chip>
    <small class="font-italic" v-if="(recettes || []).length == 0"
      >Déposez une recette ici...</small
    >
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Controller } from "@/logic/controller";
import { DragKind, getDragData } from "../../../utils/utils_drag";
const CaseRecettesProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    recettes: Array as () => number[] | null,
  },
});

@Component({})
export default class CaseRecettes extends CaseRecettesProps {
  formatRecette(idRecette: number) {
    return this.C.getRecette(idRecette).nom;
  }

  onDragoverRecettes(event: DragEvent) {
    if (!event.dataTransfer) return;
    const isRecette = event.dataTransfer.types.includes(DragKind.IdRecette);
    if (isRecette) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "copy";
    }
  }

  onDropRecette(event: DragEvent) {
    if (!event.dataTransfer) return;
    event.preventDefault();
    const idRecette = getDragData(event.dataTransfer, DragKind.IdRecette);
    this.$emit("add", idRecette);
  }
}
</script>

<style scoped></style>
