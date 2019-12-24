<template>
  <div>
    <v-toolbar color="secondary" dense class="my-1">
      <v-toolbar-title>
        Recettes <i>{{ bonusTitle }}</i>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items> </v-toolbar-items>
    </v-toolbar>
    <v-list dense :max-height="height" class="overflow-y-auto">
      <v-list-item-group
        :value="recette"
        @change="args => $emit('change', args)"
      >
        <v-list-item
          v-for="recette in recettes"
          :key="recette.id"
          :value="recette"
        >
          <v-list-item-content>
            <v-list-item-title>{{ recette.nom }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";

import TooltipBtn from "../utils/TooltipBtn.vue";

import { D } from "../../logic/controller";
import { Recette } from "../../logic/types";

const Props = Vue.extend({
  props: {
    height: String,
    recettes: Array as () => Recette[],
    bonusTitle: {
      type: String,
      default: ""
    },
    recette: Object as () => Recette | null
  },
  model: {
    prop: "recette",
    event: "change"
  }
});

@Component({
  components: { TooltipBtn }
})
export default class ListeRecettes extends Props {}
</script>
