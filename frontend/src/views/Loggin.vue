<template>
  <v-container fluid class="fill-height align-self-center">
    <v-row>
      <v-col xs="0" md="2"></v-col>
      <v-col xs="6" md="4">
        <v-form v-model="formValid" @keyup.native.enter="loggin">
          <v-text-field
            label="Mail"
            v-model="params.mail"
            :rules="[rules.required, rules.mail]"
            required
            name="email"
            autocomplete="email"
          ></v-text-field>
          <v-text-field
            name="password"
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="showPassword = !showPassword"
            v-model="params.password"
            :type="showPassword ? 'password' : 'text'"
            label="Mot de passe"
            autocomplete="password"
          ></v-text-field>
          <v-fade-transition>
            <v-row v-if="error">
              <div v-html="error" class="red--text px-3"></div>
            </v-row>
          </v-fade-transition>
          <v-row class="mt-3">
            <v-spacer></v-spacer>
            <v-btn
              color="success"
              @click="loggin"
              :loading="loading"
              :disabled="!formValid"
            >
              Se connecter
            </v-btn>
          </v-row>
        </v-form>
      </v-col>
      <v-col class="xs-6 md-4">TODO : créer un compte</v-col>
      <v-col></v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import { InLoggin } from "../logic/api";
import { LogginController } from "@/logic/server";
import { Notifications } from "@/logic/notifications";

const patternMail =
  /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

const LogginProps = Vue.extend({
  props: {
    N: Object as () => Notifications,
  },
});

@Component({})
export default class Loggin extends LogginProps {
  params: InLoggin = { mail: "", password: "" };
  showPassword = false;
  error: string | null = null;
  rules = {
    required: (s: string) => {
      return !!s || "Ce champ est requis.";
    },
    mail: (s: string) => {
      return patternMail.test(s) || "L'adresse mail semble invalide.";
    },
  };
  loading = false;
  formValid = false;

  async loggin() {
    if (!this.formValid) return;
    this.loading = true;

    const out = await LogginController.loggin(this.N, this.params);
    this.loading = false;
    if (out === undefined) return; // erreur déjà gérée
    const err = out.erreur;
    if (err == "" || !err) {
      this.error = null;
      this.$emit("loggin", out);
    } else {
      this.error = err;
    }
  }
}
</script>
