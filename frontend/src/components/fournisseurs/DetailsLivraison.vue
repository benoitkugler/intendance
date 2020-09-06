<template>
  <v-card>
    <v-card-title primary-title>
      Détails de la contrainte de livraison
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-row>
          <v-col cols="5">
            <v-text-field
              label="Label de la contrainte de livraison"
              v-model="innerLivraison.nom"
              required
            ></v-text-field>
            <v-select
              v-model="innerLivraison.id_fournisseur"
              :items="optionsFournisseurs"
              :rules="[rules.idRequired]"
              label="Fournisseur"
            ></v-select>
          </v-col>
          <v-col>
            <jours-livraison-field v-model="innerLivraison.jours_livraison">
            </jours-livraison-field>
            <v-text-field
              label="Delai de livraison"
              v-model.number="innerLivraison.delai_commande"
              type="number"
              hint="La date de commande est avancé de ce nombre de jours ouvrés par rapport à la date de livraison."
            ></v-text-field>
            <v-text-field
              label="Anticipation"
              v-model.number="innerLivraison.anticipation"
              type="number"
              hint="La date de livraison est avancé de ce nombre de jours par rapport à la date d'utilisation."
            ></v-text-field>
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="success" @click="$emit('accept', innerLivraison)">
        {{ editMode == "new" ? "Créer" : "Enregistrer" }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Livraison } from "../../logic/api";
import {
  New,
  EditMode,
  deepcopy,
  defaultLivraison,
  EnumItem
} from "../../logic/api";
import { Watch } from "vue-property-decorator";
import { C } from "../../logic/controller";
import JoursLivraisonField from "./JoursLivraisonField.vue";

const DetailsLivraisonProps = Vue.extend({
  props: {
    livraison: Object as () => New<Livraison> | null,
    editMode: String as () => EditMode
  }
});

@Component({
  components: { JoursLivraisonField }
})
export default class DetailsLivraison extends DetailsLivraisonProps {
  innerLivraison = this.duplique();
  @Watch("livraison")
  _() {
    this.innerLivraison = this.duplique();
  }
  private duplique(): New<Livraison> {
    if (this.livraison == null) {
      return defaultLivraison();
    }
    return deepcopy(this.livraison);
  }

  get optionsFournisseurs(): EnumItem<number>[] {
    if (C.data == null) return [];
    const items = Object.values(C.data.fournisseurs || {}).map(fourn => {
      return { text: fourn.nom, value: fourn.id };
    });
    return items;
  }

  rules = {
    idRequired: (id: number) => id >= 0 || "Champ requis"
  };
}
</script>

<style scoped></style>
