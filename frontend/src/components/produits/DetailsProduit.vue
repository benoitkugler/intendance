<template>
  <v-card>
    <v-card-title>Détails du produit</v-card-title>
    <v-card-text>
      <v-form ref="form">
        <v-row>
          <v-col>
            <v-select
              :items="optionsFournisseurs"
              label="Fournisseur"
              :rules="[rules.idRequired]"
              v-model="innerProduit.id_fournisseur"
            ></v-select>
          </v-col>
          <v-col>
            <v-select
              :items="optionsLivraisons"
              label="Contrainte de livraison"
              v-model="idLivraison"
            ></v-select>
          </v-col>
          <v-col>
            <v-text-field
              label="Nom"
              :rules="[rules.required]"
              v-model="innerProduit.nom"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-text-field
              label="Prix"
              type="number"
              suffix="€"
              :rules="[rules.required]"
              v-model.number="innerProduit.prix"
            ></v-text-field>
          </v-col>
          <v-col>
            <v-text-field
              label="Référence fournisseur"
              v-model="innerProduit.reference_fournisseur"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="8">
            <conditionnement-field
              v-model="innerProduit.conditionnement"
            ></conditionnement-field>
          </v-col>
          <v-col cols="4">
            <v-text-field
              label="Colisage"
              type="number"
              v-model.number="innerProduit.colisage"
            ></v-text-field>
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="success" @click="validate">Ajouter ce produit</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

import ConditionnementField from "../utils/ConditionnementField.vue";

import { Produit } from "../../logic/types";
import { C } from "../../logic/controller";
import { EnumItem } from "../../logic/enums";
import { sortByText } from "../utils/utils";
import { New, deepcopy, NullId, toNullableId } from "../../logic/types2";
import { Watch } from "vue-property-decorator";

const DetailsProduitProps = Vue.extend({
  props: {
    // contraint le produit possible
    produit: Object as () => Produit
  }
});

type VForm = Vue & { validate: () => boolean };

@Component({
  components: { ConditionnementField }
})
export default class DetailsProduit extends DetailsProduitProps {
  innerProduit: Produit = this.duplique();

  $refs!: {
    form: VForm;
  };

  @Watch("produit")
  onChange() {
    this.innerProduit = this.duplique();
  }

  duplique() {
    return deepcopy<Produit>(this.produit);
  }

  // custom v-model pour NullInt64
  get idLivraison() {
    return this.innerProduit.id_livraison.Valid
      ? this.innerProduit.id_livraison.Int64
      : null;
  }
  set idLivraison(id: number | null) {
    if (id == null) {
      this.innerProduit.id_livraison = NullId();
    } else {
      this.innerProduit.id_livraison = toNullableId(id);
    }
  }

  get optionsFournisseurs(): EnumItem<number>[] {
    if (C.data == null) return [];
    const items = Object.values(C.data.fournisseurs || {}).map(fourn => {
      return { text: fourn.nom, value: fourn.id };
    });
    return sortByText(items);
  }

  get optionsLivraisons() {
    if (C.data == null) return [];
    const items: EnumItem<number | null>[] = Object.values(
      C.data.livraisons || {}
    )
      .filter(
        livraison =>
          !livraison.id_fournisseur.Valid ||
          livraison.id_fournisseur.Int64 == this.innerProduit.id_fournisseur
      )
      .map(livraison => {
        return { text: livraison.nom, value: livraison.id };
      });
    if (items.findIndex(item => item.value == this.idLivraison) == -1) {
      // on met à jour le model sous-jacent
      this.idLivraison = null;
    }
    return sortByText(items).concat({ text: "Aucune", value: null });
  }

  rules = {
    required: (v: any) => !!v || "Champ requis",
    idRequired: (id: number) => id >= 0 || "Champ requis"
  };

  validate() {
    if (this.$refs.form.validate()) {
      this.$emit("add", this.innerProduit);
    }
  }
}
</script>

<style scoped></style>
