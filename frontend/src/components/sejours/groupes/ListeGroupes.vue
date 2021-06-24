<template>
  <div>
    <v-dialog v-model="confirmeSupprime" max-width="600px">
      <v-card>
        <v-card-title primary-title color="warning">
          Confirmer la suppression
        </v-card-title>
        <v-card-text>
          Etes vous sûr de vouloir retirer ce groupe ? <br />
          Il sera retiré des repas associés.
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="red" @click="supprime"> Supprimer </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-if="tmpGroupe != null" v-model="showEdit" max-width="600px">
      <v-card>
        <v-card-title primary-title>
          {{
            editMode == "edit" ? "Groupe " + tmpGroupe.nom : "Nouveau groupe"
          }}
        </v-card-title>
        <v-card-text>
          <v-form>
            <v-row>
              <v-col>
                <v-text-field
                  label="Nom du groupe"
                  v-model="tmpGroupe.nom"
                ></v-text-field>
                <v-text-field
                  type="number"
                  label="Taille du groupe"
                  v-model.number="tmpGroupe.nb_personnes"
                ></v-text-field>
                <v-text-field
                  label="Couleur"
                  v-model="tmpGroupe.couleur"
                  hide-details
                  class="ma-0 pa-0"
                >
                  <template v-slot:append-outer>
                    <v-menu
                      v-model="showColorPicker"
                      top
                      nudge-bottom="105"
                      nudge-left="16"
                      :close-on-content-click="false"
                    >
                      <template v-slot:activator="props">
                        <div
                          class="color-preview"
                          :style="{
                            backgroundColor: (tmpGroupe || {}).couleur,
                          }"
                          v-on="props.on"
                        />
                      </template>
                      <v-card>
                        <v-card-text class="pa-0">
                          <v-color-picker
                            v-model="tmpGroupe.couleur"
                            flat
                            mode="hexa"
                            hide-inputs
                          />
                        </v-card-text>
                      </v-card>
                    </v-menu>
                  </template>
                </v-text-field>
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green" @click="editDone">
            {{ editMode == "new" ? "Créer" : "Enregistrer les modifications" }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-toolbar dense color="secondary" class="my-1">
      <v-toolbar-title class="px-2"> Groupes </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <tooltip-btn
          mdi-icon="plus-thick"
          tooltip="Ajouter un groupe..."
          @click="startCreateGroupe"
          color="green"
          :disabled="!sejour"
        />
      </v-toolbar-items>
    </v-toolbar>
    <v-list dense max-height="75vh" class="overflow-y-auto">
      <v-list-item v-if="groupes.length === 0">
        <v-list-item-content><i> Aucun groupe.</i> </v-list-item-content>
      </v-list-item>
      <v-list-item-group v-model="groupe" mandatory>
        <v-list-item
          v-for="groupe in groupes"
          :key="groupe.id"
          :value="groupe"
          @click="startEditGroupe"
        >
          <v-list-item-content>
            <v-row no-gutters
              ><v-col>
                <v-list-item-title>
                  <v-chip
                    label
                    class="mr-1 px-1 align-self-center"
                    :color="groupe.couleur"
                    :style="{ borderWidth: ' 1.5px' }"
                    outlined
                  >
                    {{ groupe.nom }}
                  </v-chip>
                </v-list-item-title>
              </v-col>
              <v-col class="text-right align-self-center">
                <v-list-item-subtitle>
                  {{ groupe.nb_personnes }} personne(s)
                </v-list-item-subtitle>
              </v-col></v-row
            >
          </v-list-item-content>
          <v-list-item-action>
            <v-row no-gutters>
              <v-col>
                <tooltip-btn
                  mdi-icon="close"
                  tooltip="Supprimer ce groupe"
                  color="red"
                  @click.stop="askConfirmeSupprime(groupe)"
                ></tooltip-btn>
              </v-col>
            </v-row>
          </v-list-item-action>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { SejourRepas, Sejour, Groupe, New } from "@/logic/api";
import { Controller } from "@/logic/controller";

import TooltipBtn from "../../utils/TooltipBtn.vue";
import { EditMode, deepcopy } from "@/logic/types";
import { Watch } from "vue-property-decorator";

const ListeGroupesProps = Vue.extend({
  props: {
    C: Object as () => Controller,
    sejour: Object as () => Sejour | null,
  },
});

@Component({
  components: { TooltipBtn },
})
export default class ListeGroupes extends ListeGroupesProps {
  groupe: Groupe | null = null;

  tmpGroupe: New<Groupe> | null = null;

  @Watch("groupe")
  onGroupeChange(groupe: Groupe | undefined) {
    if (this.editMode == "new") return;
    if (groupe == undefined) {
      this.tmpGroupe = null;
    } else {
      this.tmpGroupe = deepcopy(groupe);
    }
  }

  showColorPicker = false;

  confirmeSupprime = false;

  editMode: EditMode = "edit";
  showEdit = false;

  get groupes() {
    const sej = this.sejour;
    if (sej === null) return [];
    return Object.values(this.C.api.sejours.groupes).filter(
      (groupe) => groupe.id_sejour == sej.id
    );
  }

  startCreateGroupe() {
    if (this.sejour === null) return;
    this.groupe = null;
    this.tmpGroupe = {
      id_sejour: this.sejour.id,
      nom: "",
      nb_personnes: 0,
      couleur: "#D1CA3D",
    };
    this.editMode = "new";
    this.showEdit = true;
  }

  startEditGroupe() {
    this.tmpGroupe = deepcopy(this.groupe);
    this.editMode = "edit";
    this.showEdit = true;
  }

  async editDone() {
    if (this.tmpGroupe === null) return;
    if (this.editMode == "new") {
      await this.C.api.CreateGroupe(this.tmpGroupe);
    } else {
      await this.C.api.UpdateGroupe(this.tmpGroupe as Groupe);
    }
    this.tmpGroupe = null;
    this.showEdit = false;
  }

  askConfirmeSupprime(groupe: Groupe) {
    this.groupe = groupe; // le groupe peut être mis à null en cliquant 2 fois
    this.confirmeSupprime = true;
  }
  async supprime() {
    this.confirmeSupprime = false;
    if (this.groupe == null || this.groupe.id == undefined) return;
    await this.C.api.DeleteGroupe(this.groupe);
  }
}
</script>

<style scoped>
div.color-preview {
  height: 30px;
  width: 60px;
  cursor: pointer;
}
</style>
