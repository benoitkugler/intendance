import Vue from "vue";
import Component from "vue-class-component";

import { StateMenus, SelectionMenu } from "./types";
import { Controller } from "@/logic/controller";

export type ListKind = keyof SelectionMenu;
const Props = Vue.extend({
  props: {
    height: String,
    state: Object as () => StateMenus,
    kind: String as () => ListKind,
    C: Object as () => Controller,
  },
});

@Component({})
export class BaseList extends Props {
  // à implémenter par les composants
  $refs!: {
    list: HTMLElement;
  };

  private classItem(idItem: number) {
    return this.kind + idItem;
  }

  // change l'item courant et scroll
  goToItem(idItem: number) {
    this.state.selection[this.kind] = idItem;
    try {
      this.$vuetify.goTo("." + this.classItem(idItem), {
        container: this.$refs.list,
      });
    } catch {
      // élément inconnu, on ignore
    }
  }
}
