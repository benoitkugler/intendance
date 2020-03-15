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
          <v-btn color="red" @click="supprime">
            Supprimer
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-row>
      <v-col>
        <v-toolbar dense color="secondary" class="my-1">
          <v-toolbar-title class="px-2">
            Groupes
          </v-toolbar-title>
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
          <v-list-item-group v-model="groupe" mandatory>
            <v-list-item
              v-for="groupe in groupes"
              :key="groupe.id"
              :value="groupe"
              @click="editMode = 'edit'"
            >
              <v-list-item-content>
                <v-list-item-title>
                  {{ groupe.nom }}
                </v-list-item-title>
                <v-list-item-subtitle>
                  {{ groupe.nb_personnes }} personne(s)
                </v-list-item-subtitle>
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
      </v-col>
      <v-col class="align-self-center">
        <div v-if="tmpGroupe == null">
          <i>Sélectionnez un groupe ou créez un nouveau groupe...</i>
        </div>
        <v-card v-else>
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
                        <template v-slot:activator="{ on }">
                          <div
                            class="color-preview"
                            :style="{ backgroundColor: tmpGroupe.couleur }"
                            v-on="on"
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
              {{
                editMode == "new" ? "Créer" : "Enregistrer les modifications"
              }}
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { SejourRepas, Sejour, Groupe } from "../../../logic/types";
import { C } from "../../../logic/controller";

import TooltipBtn from "../../utils/TooltipBtn.vue";
import { EditMode, New, deepcopy } from "../../../logic/types2";
import { Watch } from "vue-property-decorator";

const ListeGroupesProps = Vue.extend({
  props: {
    sejour: Object as () => Sejour | null
  }
});

@Component({
  components: { TooltipBtn }
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

  get groupes() {
    const sej = this.sejour;
    if (sej === null) return [];
    return Object.values(C.data.sejours.groupes || {}).filter(
      groupe => groupe.id_sejour == sej.id
    );
  }

  startCreateGroupe() {
    if (this.sejour === null) return;
    this.editMode = "new";
    this.groupe = null;
    this.tmpGroupe = {
      id_sejour: this.sejour.id,
      nom: "",
      nb_personnes: 0,
      couleur: "#D1CA3D"
    };
  }

  startEditGroupe(groupe: Groupe) {}

  async editDone() {
    if (this.tmpGroupe === null) return;
    let message: string;
    if (this.editMode == "new") {
      message = "Groupe ajouté avec succès.";
      this.groupe = (await C.data.createGroupe(this.tmpGroupe)) || null;
      this.editMode = "edit";
    } else {
      message = "Groupe modifié avec succès.";
      await C.data.updateGroupe(this.tmpGroupe as Groupe);
    }
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(message);
    }
  }

  askConfirmeSupprime(groupe: Groupe) {
    this.groupe = groupe; // le groupe peut être mis à null en cliquant 2 fois
    this.confirmeSupprime = true;
  }
  async supprime() {
    this.confirmeSupprime = false;
    if (this.groupe == null || this.groupe.id == undefined) return;
    const nbRepas = await C.data.deleteGroupe(this.groupe as Groupe);
    if (C.notifications.getError() == null) {
      C.notifications.setMessage(
        `Groupe bien retiré. ${nbRepas} repas ont été modifié(s).`
      );
    }
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
