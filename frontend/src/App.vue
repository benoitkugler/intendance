<template>
  <v-app>
    <v-app-bar app color="primary" :dense="isLoggedIn">
      <v-toolbar-title class="headline text-uppercase">
        <v-tooltip>
          <template v-slot:activator="props">
            <span v-on="props.on">{{ mainTitle }}</span>
          </template>
          Version {{ version }}
        </v-tooltip>
      </v-toolbar-title>
    </v-app-bar>
    <navigation-bar
      :is-logged-in="isLoggedIn"
      @logout="logout"
    ></navigation-bar>
    <v-main>
      <keep-alive v-if="isLoggedIn">
        <router-view :C="controller"></router-view>
      </keep-alive>
      <loggin v-else :N="notifications" @loggin="onLoggin"></loggin>

      <spinner-snackbar :N="notifications"></spinner-snackbar>
      <error-dialog :N="notifications"></error-dialog>
      <success-snackbar :N="notifications"></success-snackbar>
    </v-main>
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
import Loggin from "./views/Loggin.vue";

import { Controller } from "./logic/controller";
import { devMode, LogginController, metaDev } from "./logic/server";
import { routes, RouteType } from "./router";
import { Notifications } from "./logic/notifications";
import { OutLoggin } from "./logic/api";

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
    NavigationBar,
    Loggin
  }
})
export default class App extends Vue {
  // a créer après le loggin
  private controller: Controller | null = null;
  private notifications = new Notifications();

  get version() {
    return process.env.VUE_APP_VERSION;
  }

  get isLoggedIn() {
    return this.controller != null;
  }

  get mainTitle() {
    if (!this.isLoggedIn) {
      return "Bienvenue sur votre platforme d'intendance";
    }
    return this.getPageTitle();
  }

  created() {
    // point d'entrée : le controller n'est pas encore disponible

    // on commence par vérifier les cookies
    const meta = LogginController.checkCookies();
    if (meta !== null) {
      this.controller = new Controller(meta, this.notifications);
      this.notifications.setMessage("Bon retour parmi nous !");
    } else if (devMode) {
      this.controller = new Controller(metaDev, this.notifications);
      this.notifications.setMessage("Connecté en mode développement");
    } else {
      // pas de loggin automatique : par défaut, affiche la page de loggin
    }
  }

  onLoggin(out: OutLoggin) {
    // loggin réussi
    this.controller = new Controller(
      { idUtilisateur: out.utilisateur.id, token: out.token },
      this.notifications
    );
    const currentPath = this.$router.currentRoute.path;
    if (routes.map(r => r.path).indexOf(currentPath) == -1) {
      this.$router.push("/sejours");
    }
  }

  logout() {
    LogginController.logout();
    this.controller = null;
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

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active below version 2.1.8 */ {
  transform: translateX(10px);
  opacity: 0;
}
</style>
