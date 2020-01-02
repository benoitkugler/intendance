import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";

Vue.config.productionTip = false;

// Workaround annoying warnings
const ignoreWarnMessage =
  "The .native modifier for v-on is only valid on components but it was used on <div>.";
Vue.config.warnHandler = function(msg, vm, trace) {
  // `trace` is the component hierarchy trace
  if (msg === ignoreWarnMessage) {
    return;
  }
  console.error("[Vue warn]: " + msg + trace);
};

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount("#app");
