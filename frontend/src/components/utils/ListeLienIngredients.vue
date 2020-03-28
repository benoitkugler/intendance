<template>
  <div>
    <v-dialog v-model="showEditIngredient" max-width="500px">
      <details-ingredient
        :ingredient="editedIngredient"
        @edit="editIngredientDone"
        ref="editIngredient"
      ></details-ingredient>
    </v-dialog>

    <div @dragover="onDragoverIngredients" @drop="onDropIngredient">
      <v-list dense>
        <v-subheader>
          {{ subheader }}
        </v-subheader>

        <v-list-item v-if="(ingredients || []).length == 0 && !showAdd">
          <v-list-item-subtitle>
            <i>Cliquer-déplacer pour ajouter un ingrédient...</i>
          </v-list-item-subtitle>
        </v-list-item>

        <v-list-item
          v-for="ingredient in ingredients"
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
        <div class="px-3">
          <ingredient-field
            v-if="showAdd"
            @change="addIngredient"
          ></ingredient-field>
        </div>
      </v-list>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Watch } from "vue-property-decorator";

import DetailsIngredient from "./DetailsIngredient.vue";
import TooltipBtn from "./TooltipBtn.vue";
import IngredientField from "./IngredientField.vue";

import { LienIngredient } from "../../logic/types";
import { C } from "../../logic/controller";
import { DragKind, getDragData } from "./utils_drag";

const ListeLienIngredientsProps = Vue.extend({
  props: {
    subheader: String,
    showAdd: Boolean,
    ingredients: Array as () => LienIngredient[] | null
  },
  model: {
    prop: "ingredients",
    event: "change"
  }
});

@Component({
  components: { DetailsIngredient, TooltipBtn, IngredientField }
})
export default class ListeLienIngredients extends ListeLienIngredientsProps {
  showEditIngredient = false;
  editedIngredient: LienIngredient | null = null;

  $refs!: {
    editIngredient: DetailsIngredient;
  };

  getIngredient(ing: LienIngredient) {
    return (C.data.ingredients || {})[ing.id_ingredient];
  }

  editIngredient(ing: LienIngredient) {
    this.editedIngredient = ing;
    this.showEditIngredient = true;
  }

  editIngredientDone(edited: LienIngredient) {
    const ings = this.ingredients;
    if (ings == null) return;
    const index = ings.findIndex(
      ing => ing.id_ingredient == edited.id_ingredient
    );
    ings[index] = edited;
    this.showEditIngredient = false;
    this.$emit("change", ings);
  }

  removeIngredient(toRemove: LienIngredient) {
    const ings = (this.ingredients || []).filter(
      ing => ing.id_ingredient != toRemove.id_ingredient
    );
    this.$emit("change", ings);
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
    this.addIngredient(idIngredient);
  }

  addIngredient(idIngredient: number) {
    const ingredients = this.ingredients || [];
    const matchingIngredients = ingredients.filter(
      r => r.id_ingredient == idIngredient
    );
    let newIngredient: LienIngredient;
    if (matchingIngredients.length > 0) {
      newIngredient = matchingIngredients[0];
    } else {
      newIngredient = {
        id_ingredient: idIngredient,
        quantite: 0,
        cuisson: ""
      };
      ingredients.push(newIngredient);
    }
    this.$emit("change", ingredients);
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
