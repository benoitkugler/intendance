<template>
  <v-card>
    <v-card-title>Détails du produit</v-card-title>
    <v-card-text>
      <v-form ref="form">
        <v-row>
          <v-col>
            <v-select
              :items="optionsLivraisons"
              label="Fournisseur"
              :rules="[rules.idRequired]"
              v-model="innerProduit.id_livraison"
            >
              <template v-slot:item="props">
                <span v-html="asE(props.item).text"></span>
              </template>
              <template v-slot:selection="props">
                <span v-html="asE(props.item).text"></span>
              </template>
            </v-select>
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
              hint="Prix d'un exemplaire"
            ></v-text-field>
          </v-col>
          <v-col>
            <v-text-field
              label="Référence fournisseur"
              v-model="innerProduit.reference_fournisseur"
              hint="Optionnel"
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
              hint="Nombre minimum d'exemplaires livrables"
              :min="1"
              persistent-hint
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

import { Produit, New } from "@/logic/api";
import { Controller } from "@/logic/controller";
import { sortByText } from "../utils/utils";
import { deepcopy, NullId, toNullableId, EnumItem } from "@/logic/types";
import { Watch } from "vue-property-decorator";

const DetailsProduitProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    // contraint le produit possible
    produit: Object as () => Produit,
  },
});

type VForm = Vue & { validate: () => boolean };

@Component({
  components: { ConditionnementField },
})
export default class DetailsProduit extends DetailsProduitProps {
  innerProduit: Produit = this.duplique();
  $refs!: {
    form: VForm;
  };

  asE = (i: EnumItem) => i;

  @Watch("produit")
  onChange() {
    this.innerProduit = this.duplique();
  }

  duplique() {
    return deepcopy<Produit>(this.produit);
  }

  get optionsLivraisons() {
    const items: EnumItem<number>[] = Object.values(this.C.api.livraisons).map(
      (livraison) => {
        return {
          text: this.C.formatter.formatLivraison(livraison),
          value: livraison.id,
        };
      }
    );
    return items;
  }

  rules = {
    required: (v: any) => !!v || "Champ requis",
    idRequired: (id: number) => id >= 0 || "Champ requis",
  };

  validate() {
    if (this.$refs.form.validate()) {
      this.$emit("add", this.innerProduit);
    }
  }
}
</script>

<style scoped></style>
