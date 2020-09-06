<template>
  <v-card>
    <v-card-title primary-title>
      Détails du fournisseur
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-text-field
          label="Nom du fournisseur"
          v-model="innerFournisseur.nom"
          required
        ></v-text-field>
        <v-text-field
          label="Lieu"
          v-model="innerFournisseur.lieu"
          hint="Le lieu permet de sélectionner rapidement un groupe de fournisseurs."
          persistent-hint
        ></v-text-field>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="success" @click="$emit('accept', innerFournisseur)">
        {{ editMode == "new" ? "Créer" : "Enregistrer" }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Fournisseur } from "../../logic/api";
import { New, EditMode, deepcopy } from "../../logic/api";
import { Watch } from "vue-property-decorator";

const DetailsFournisseurProps = Vue.extend({
  props: {
    fournisseur: Object as () => New<Fournisseur> | null,
    editMode: String as () => EditMode
  }
});

@Component({})
export default class DetailsFournisseur extends DetailsFournisseurProps {
  innerFournisseur = this.duplique();

  @Watch("fournisseur")
  _() {
    this.innerFournisseur = this.duplique();
  }

  private duplique(): New<Fournisseur> {
    if (this.fournisseur == null) {
      return { nom: "", lieu: "" };
    }
    return deepcopy(this.fournisseur);
  }
}
</script>

<style scoped></style>
