<template>
  <v-navigation-drawer v-if="isLoggedIn" app :mini-variant="menuMini" permanent>
    <v-layout align-center justify-center row>
      <v-avatar @click="menuMini = !menuMini" size="50">
        <v-img margin="auto" :src="require('@/assets/logo.png')" />
      </v-avatar>
    </v-layout>
    <v-list>
      <v-list-item
        v-for="route in routes"
        router
        :key="route.name"
        :to="route.path"
      >
        <v-tooltip top>
          <template v-slot:activator="{ on }">
            <v-list-item-action v-on="on">
              <v-icon>{{ route.meta.icon }}</v-icon>
            </v-list-item-action>
          </template>
          {{ route.meta.tooltip }}
        </v-tooltip>
        <v-list-item-content>{{ route.meta.title }}</v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>
      <v-list-item @click="$emit('logout')">
        <v-tooltip top>
          <template v-slot:activator="{ on }">
            <v-list-item-action v-on="on">
              <v-icon>mdi-logout</v-icon>
            </v-list-item-action>
          </template>
          Se déconnecter
        </v-tooltip>
        <v-list-item-content>Se déconnecter</v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop } from "vue-property-decorator";
import { routes } from "../router/index";

const Props = Vue.extend({
  props: {
    isLoggedIn: Boolean
  }
});

@Component
export default class NavigationBar extends Props {
  menuMini = true;
  routes = routes;
}
</script>
