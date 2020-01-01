<template>
  <v-app>
    <v-app-bar app color="primary" :dense="isLoggedIn">
      <v-toolbar-title class="headline text-uppercase">
        <v-tooltip>
          <template v-slot:activator="{ on }">
            <span v-on="on">{{ mainTitle }}</span>
          </template>
          Version {{ version }}
        </v-tooltip>
      </v-toolbar-title>
    </v-app-bar>
    <navigation-bar
      :is-logged-in="isLoggedIn"
      @logout="logout"
    ></navigation-bar>
    <v-content>
      <keep-alive v-if="isLoggedIn">
        <router-view></router-view>
      </keep-alive>

      <spinner-snackbar></spinner-snackbar>
      <error-dialog></error-dialog>
      <success-snackbar></success-snackbar>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Watch } from "vue-property-decorator";

import ErrorDialog from "./components/ErrorDialog.vue";
import SuccessSnackbar from "./components/SuccessSnackbar.vue";
import SpinnerSnackbar from "./components/SpinnerSnackbar.vue";
import NavigationBar from "./components/NavigationBar.vue";

import { C } from "./logic/controller";
import { devMode } from "./logic/data";
import { RouteType } from "./router";

declare var process: {
  env: {
    VUE_APP_VERSION: string;
  };
};

@Component({
  components: {
    ErrorDialog,
    SuccessSnackbar,
    SpinnerSnackbar,
    NavigationBar
  }
})
export default class App extends Vue {
  private controller = C;

  isLoggedIn = devMode;

  get version() {
    return process.env.VUE_APP_VERSION;
  }

  get mainTitle() {
    if (!this.isLoggedIn) {
      return "Bienvenue sur votre platforme d'intendance";
    }
    return this.getPageTitle();
  }

  logout() {
    console.log("TODO");
  }

  private getPageTitle() {
    let title;
    if (!this.isLoggedIn) {
      title = "Intendance - Connexion";
    } else {
      const route = this.$route as RouteType;
      title = "Intendance - " + route.meta.title;
    }
    return title;
  }

  private updatePageTitle() {
    document.title = this.getPageTitle();
  }

  @Watch("$route")
  onRouteChanged() {
    this.updatePageTitle();
  }

  mounted() {
    this.updatePageTitle();
  }
}
</script>
