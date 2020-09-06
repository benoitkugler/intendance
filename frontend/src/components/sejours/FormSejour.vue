<template>
  <v-card>
    <v-card-title primary-title>
      Détails du séjour
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-text-field
          label="Nom du séjour"
          v-model="tmpSejour.nom"
          required
        ></v-text-field>
        <date-field
          v-model="tmpSejour.date_debut"
          label="Date du premier jour"
        ></date-field>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="success" @click="$emit('accept', tmpSejour)">
        {{ editMode == "new" ? "Créer" : "Enregistrer" }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Sejour, New } from "@/logic/api";
import { DetailsSejour, EditMode, deepcopy } from "@/logic/types";
import DateField from "../utils/DateField.vue";
import { Watch } from "vue-property-decorator";

const Props = Vue.extend({
  props: {
    sejour: Object as () => Sejour,
    editMode: String as () => EditMode
  }
});

@Component({
  components: { DateField }
})
export default class FormSejour extends Props {
  tmpSejour: DetailsSejour = this.copy();

  private copy() {
    return deepcopy(
      this.sejour || { nom: "", date_debut: new Date().toISOString() }
    );
  }

  @Watch("sejour")
  onPropChange() {
    this.tmpSejour = this.copy();
  }
}
</script>

<style></style>
