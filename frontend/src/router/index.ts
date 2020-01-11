import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import Sejours from "../views/Sejours.vue";
import Agenda from "../views/Agenda.vue";

Vue.use(VueRouter);

interface RouteMeta {
  title: string;
  tooltip: string;
  icon: string;
}
export interface RouteType extends RouteConfig {
  meta: RouteMeta;
}

export const routes: RouteType[] = [
  {
    path: "/sejours",
    name: "sejours",
    // route level code-splitting
    // this generates a separate chunk (agenda.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Sejours,
    meta: {
      title: "Sejours",
      tooltip: "Vue d'ensemble des séjours et des groupes.",
      icon: "mdi-account-group"
    }
  },
  {
    path: "/agenda",
    name: "agenda",
    // route level code-splitting
    // this generates a separate chunk (agenda.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "agenda" */ "../views/Agenda.vue"),
    meta: {
      title: "Agenda",
      tooltip: "Organisation du séjour.",
      icon: "mdi-calendar-month"
    }
  },
  {
    path: "/menus",
    name: "menus",
    // route level code-splitting
    // this generates a separate chunk (menus.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "menus" */ "../views/Menus.vue"),
    meta: {
      title: "Menus, recettes et ingrédients",
      tooltip: "Accès aux menus, recettes et ingrédients.",
      icon: "mdi-food-variant"
    }
  }
];

const router = new VueRouter({
  mode: "history",
  routes
});

export default router;
