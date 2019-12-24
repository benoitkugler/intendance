<template>
  <v-app>
    <v-app-bar app color="primary" dense> </v-app-bar>
    <navigation-bar
      :is-logged-in="isLoggedIn"
      @logout="logout"
    ></navigation-bar>
    <v-content>
      <v-container>
        <keep-alive v-if="isLoggedIn">
          <router-view></router-view>
        </keep-alive>
      </v-container>

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

import { D } from "./logic/controller";
import { NS } from "./logic/notifications";
import { RouteType } from "./router";

@Component({
  components: {
    ErrorDialog,
    SuccessSnackbar,
    SpinnerSnackbar,
    NavigationBar
  }
})
export default class App extends Vue {
  private storage = D;
  private notifications = NS;

  isLoggedIn = false;

  logout() {
    console.log("TODO");
  }

  private updatePageTitle() {
    let title;
    if (!this.isLoggedIn) {
      title = "Intendance - Connexion";
    } else {
      const route = this.$route as RouteType;
      title = "Intendance - " + route.meta.title;
    }
    document.title = title;
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
