import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
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
    path: "/",
    name: "agenda",
    component: Agenda,
    meta: {
      title: "Agenda",
      tooltip: "Vue d'ensemble des séjours et des repas.",
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
      title: "Recettes et ingrédients",
      tooltip: "Accès aux menus, recettes et ingrédients.",
      icon: "mdi-food-variant"
    }
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
