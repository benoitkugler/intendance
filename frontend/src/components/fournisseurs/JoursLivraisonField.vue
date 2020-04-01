<template>
  <v-select
    multiple
    chips
    :items="items"
    v-model="innerJours"
    label="Jours de livraison"
    hint="Jours possible de livraison."
  ></v-select>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { JoursLivraison } from "../../logic/types";
import { EnumItem } from "../../logic/enums";
import { Days } from "../utils/utils";

const JoursLivraisonFieldProps = Vue.extend({
  props: {
    jours: Array as () => JoursLivraison
  },
  model: {
    prop: "jours",
    event: "change"
  }
});

@Component({})
export default class JoursLivraisonField extends JoursLivraisonFieldProps {
  items: EnumItem<number>[] = Days.map((s, i) => {
    return { text: s, value: i };
  });

  // custom v-model
  get innerJours() {
    const out: number[] = [];
    this.jours.forEach((b, i) => {
      if (b) {
        out.push(i);
      }
    });
    return out;
  }
  set innerJours(v: number[]) {
    const out: JoursLivraison = this.items.map(_ => false);
    v.forEach(i => (out[i] = true));
    this.$emit("change", out);
  }
}
</script>

<style scoped></style>
