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
          @click="$emit('done', recette)"
        ></tooltip-btn>
        <tooltip-btn
          mdi-icon="undo"
          tooltip="Abandonner les modifications"
          @click="$emit('undo')"
        ></tooltip-btn>
      </v-toolbar-items>
    </v-toolbar>
    <v-row class="px-2">
      <v-col>
        <v-text-field
          label="Nom de la recette"
          v-model="recette.nom"
          required
        ></v-text-field>
        <v-textarea
          class=""
          label="Mode d'emploi"
          v-model="recette.mode_emploi"
          hide-details
          rows="2"
        ></v-textarea>
      </v-col>
      <v-col>
        <v-list dense max-height="60vh" class="mt-0 pt-0 overflow-y-auto">
          <div @dragover="onDragoverIngredients" @drop="onDropIngredient">
            <v-subheader>Ingrédients</v-subheader>
            <v-list-item v-if="(recette.ingredients || []).length == 0">
              <v-list-item-subtitle>
                <i>Cliquer-déplacer pour ajouter un ingrédient...</i>
              </v-list-item-subtitle>
            </v-list-item>
            <v-list-item
              v-for="ingredient in recette.ingredients"
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
                        tooltip="Retirer cet ingrédient de la recette"
                        color="red"
                        @click.stop="removeIngredient(ingredient)"
                      ></tooltip-btn>
                    </v-col>
                  </v-row>
                </v-list-item-action>
              </template>
            </v-list-item>
          </div>
        </v-list>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import TooltipBtn from "../utils/TooltipBtn.vue";
import MenuOrRecetteIngredient from "./MenuOrRecetteIngredient.vue";

import { C } from "../../logic/controller";
import { Recette, RecetteIngredient } from "../../logic/types";
import { New, EditMode, deepcopy } from "../../logic/types2";
import { Watch } from "vue-property-decorator";
const EditRecetteProps = Vue.extend({
  props: {
    mode: String as () => EditMode,
    initialRecette: Object as () => Recette
  }
});

@Component({
  components: { TooltipBtn, MenuOrRecetteIngredient }
})
export default class EditRecette extends EditRecetteProps {
  // recette actuellement éditée
  recette: Recette = deepcopy(this.initialRecette);

  showEditIngredient = false;
  editedIngredient: RecetteIngredient | null = null;

  $refs!: {
    editIngredient: MenuOrRecetteIngredient;
  };
  get title() {
    if (this.mode == "edit") {
      return "Modifier la recette";
    }
    return "Ajouter une recette";
  }

  formatRecetteProprietaire = C.formatter.formatMenuOrRecetteProprietaire;

  getIngredient(ing: RecetteIngredient) {
    return (C.data.ingredients || {})[ing.id_ingredient];
  }

  editIngredient(ing: RecetteIngredient) {
    this.editedIngredient = ing;
    this.showEditIngredient = true;
  }

  editIngredientDone(edited: RecetteIngredient) {
    const ings = this.recette.ingredients;
    if (ings == null) return;
    ings.forEach((ing, index) => {
      if (ing.id_ingredient == edited.id_ingredient) {
        this.$set(ings, index, edited);
      }
    });
    this.showEditIngredient = false;
  }

  removeIngredient(toRemove: RecetteIngredient) {
    this.recette.ingredients = (this.recette.ingredients || []).filter(
      ing => ing.id_ingredient != toRemove.id_ingredient
    );
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
    const ingredients = this.recette.ingredients || [];
    const matchingIngredients = ingredients.filter(
      r => r.id_ingredient == idIngredient
    );
    let newIngredient: RecetteIngredient;
    if (matchingIngredients.length > 0) {
      newIngredient = matchingIngredients[0];
    } else {
      newIngredient = {
        id_recette: this.recette.id,
        id_ingredient: idIngredient,
        quantite: 0,
        cuisson: ""
      };
      ingredients.push(newIngredient);
      this.recette.ingredients = ingredients;
    }
    this.editedIngredient = newIngredient;
    this.showEditIngredient = true;
  }

  // permet de donner le focus automatiquement
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
