import Vue from "vue";
import Vuetify from "vuetify/lib";

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: "#e0a912",
        secondary: "#a5d0d9",
        accent: "#e8d52c"
      }
    }
  }
});
