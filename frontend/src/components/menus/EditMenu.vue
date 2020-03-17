<template>
  <div>
    <v-dialog v-model="showEditIngredient" max-width="500px">
      <menu-or-recette-ingredient
        :ingredient="editedIngredient"
        @edit="editIngredientDone"
        ref="editIngredient"
      ></menu-or-recette-ingredient>
    </v-dialog>

    <v-toolbar color="secondary" dense class="my-1">
      <v-toolbar-title>{{ title }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          mdi-icon="check-outline"
          tooltip="Valider"
          color="success"
          @click="$emit('done', menu)"
        ></tooltip-btn>
        <tooltip-btn
          mdi-icon="undo"
          tooltip="Abandonner les modifications"
          @click="$emit('undo')"
        ></tooltip-btn>
      </v-toolbar-items>
    </v-toolbar>
    <v-textarea
      class="pa-2 mt-4 my-3 mb-0"
      label="Commentaire"
      v-model="menu.commentaire"
      hide-details
      rows="2"
    ></v-textarea>

    <v-list dense max-height="60vh" class="mt-0 pt-0 overflow-y-auto">
      <div @dragover="onDragoverRecettes" @drop="onDropRecette">
        <v-subheader>Recettes du menu</v-subheader>
        <v-list-item v-if="recettes.length == 0">
          <v-list-item-subtitle>
            <i>Cliquer-déplacer pour ajouter une recette...</i>
          </v-list-item-subtitle>
        </v-list-item>
        <v-list-item v-for="recette in recettes" :key="recette.id">
          <template v-slot:default="{}">
            <v-list-item-content>
              <v-list-item-title>
                <v-list-item-title>{{ recette.nom }}</v-list-item-title>
              </v-list-item-title>
              <v-list-item-subtitle>
                <i> {{ formatRecetteProprietaire(recette) }}</i>
              </v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action>
              <tooltip-btn
                mdi-icon="close"
                tooltip="Retirer cette recette du menu"
                color="red"
                @click.stop="removeRecette(recette)"
              ></tooltip-btn>
            </v-list-item-action>
          </template>
        </v-list-item>
      </div>
      <div @dragover="onDragoverIngredients" @drop="onDropIngredient">
        <v-subheader>Ingrédients additionnels</v-subheader>
        <v-list-item v-if="(menu.ingredients || []).length == 0">
          <v-list-item-subtitle>
            <i>Cliquer-déplacer pour ajouter un ingrédient...</i>
          </v-list-item-subtitle>
        </v-list-item>
        <v-list-item
          v-for="ingredient in menu.ingredients"
          :key="ingredient.id_ingredient"
        >
          <template v-slot:default="{}">
            <v-list-item-content>
              <v-list-item-title>
                {{ getIngredient(ingredient).nom }}
              </v-list-item-title>
              <v-list-item-subtitle>
                {{ ingredient.quantite }}
                <i>{{ getIngredient(ingredient).unite }}</i> - Cuisson :
                {{ ingredient.cuisson }}
              </v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-action>
              <v-row no-gutters>
                <v-col
                  ><tooltip-btn
                    mdi-icon="pencil"
                    tooltip="Modifier les détails..."
                    color="secondary"
                    @click.stop="editIngredient(ingredient)"
                  ></tooltip-btn
                ></v-col>
                <v-col
                  ><tooltip-btn
                    mdi-icon="close"
                    tooltip="Retirer cet ingrédient du menu"
                    color="red"
                    @click.stop="removeIngredient(ingredient)"
                  ></tooltip-btn
                ></v-col>
              </v-row>
            </v-list-item-action>
          </template>
        </v-list-item>
      </div>
    </v-list>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import TooltipBtn from "../utils/TooltipBtn.vue";
import MenuOrRecetteIngredient from "./MenuOrRecetteIngredient.vue";

import { C } from "../../logic/controller";
import {
  Menu,
  Recette,
  MenuIngredient,
  RecetteIngredient
} from "../../logic/types";
import { New, EditMode, deepcopy } from "../../logic/types2";
import { Watch } from "vue-property-decorator";
const EditMenuProps = Vue.extend({
  props: {
    mode: String as () => EditMode,
    initialMenu: Object as () => Menu
  }
});

@Component({
  components: { TooltipBtn, MenuOrRecetteIngredient }
})
export default class EditMenu extends EditMenuProps {
  // menu actuellement édité
  menu: Menu = deepcopy(this.initialMenu);

  showEditIngredient = false;
  editedIngredient: MenuIngredient | null = null;

  $refs!: {
    editIngredient: MenuOrRecetteIngredient;
  };
  get title() {
    if (this.mode == "edit") {
      return "Modifier le menu";
    }
    return "Ajouter un menu";
  }

  // résoud les recettes à partir des ids
  get recettes() {
    return C.getMenuRecettes(this.menu);
  }

  formatRecetteProprietaire = C.formatter.formatMenuOrRecetteProprietaire;

  getIngredient(ing: MenuIngredient) {
    return (C.data.ingredients || {})[ing.id_ingredient];
  }

  removeRecette(toRemove: Recette) {
    this.menu.recettes = (this.menu.recettes || []).filter(
      rec => rec.id_recette != toRemove.id
    );
  }

  editIngredient(ing: MenuIngredient) {
    this.editedIngredient = ing;
    this.showEditIngredient = true;
  }

  editIngredientDone(edited: MenuIngredient) {
    const ings = this.menu.ingredients;
    if (ings == null) return;
    ings.forEach((ing, index) => {
      if (ing.id_ingredient == edited.id_ingredient) {
        this.$set(ings, index, edited);
      }
    });
    this.showEditIngredient = false;
  }

  removeIngredient(toRemove: MenuIngredient) {
    this.menu.ingredients = (this.menu.ingredients || []).filter(
      ing => ing.id_ingredient != toRemove.id_ingredient
    );
  }

  onDragoverRecettes(event: DragEvent) {
    if (!event.dataTransfer) return;
    const isRecette = event.dataTransfer.types.includes("id-recette");
    if (isRecette) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "copy";
    }
  }

  onDropRecette(event: DragEvent) {
    if (!event.dataTransfer) return;
    event.preventDefault();
    const idRecette = Number(event.dataTransfer.getData("id-recette"));
    const recettes = this.menu.recettes || [];
    const hasRecette =
      recettes.filter(r => r.id_recette == idRecette).length > 0;
    if (hasRecette) return;
    recettes.push({ id_menu: this.menu.id, id_recette: idRecette });
    this.menu.recettes = recettes;
  }

  onDragoverIngredients(event: DragEvent) {
    if (!event.dataTransfer) return;
    const isIngredient = event.dataTransfer.types.includes("id-ingredient");
    if (isIngredient) {
      event.preventDefault();
      event.dataTransfer.dropEffect = "copy";
    }
  }

  onDropIngredient(event: DragEvent) {
    if (!event.dataTransfer) return;
    event.preventDefault();
    const idIngredient = Number(event.dataTransfer.getData("id-ingredient"));
    const ingredients = this.menu.ingredients || [];
    const matchingIngredients = ingredients.filter(
      r => r.id_ingredient == idIngredient
    );
    let newIngredient: MenuIngredient;
    if (matchingIngredients.length > 0) {
      newIngredient = matchingIngredients[0];
    } else {
      newIngredient = {
        id_menu: this.menu.id,
        id_ingredient: idIngredient,
        quantite: 0,
        cuisson: ""
      };
      ingredients.push(newIngredient);
      this.menu.ingredients = ingredients;
    }
    this.editedIngredient = newIngredient;
    this.showEditIngredient = true;
  }

  @Watch("showEditIngredient")
  onEditIngredient(b: boolean) {
    if (b) {
      setTimeout(() => {
        const ed = this.$refs.editIngredient;
        if (ed) ed.focus();
      }, 50);
    }
  }
}
</script>

<style scoped></style>
