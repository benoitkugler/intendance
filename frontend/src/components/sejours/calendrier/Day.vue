<template>
  <div>
    <v-calendar
      v-if="day != null"
      type="day"
      locale="fr"
      :first-interval="4"
      :interval-count="8"
      :interval-minutes="120"
      :interval-height="55"
      :value="dayString"
    >
    </v-calendar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { toDateVuetify } from "./utils";

const DayProps = Vue.extend({
  props: {
    day: Date as new () => Date
  }
});

@Component({})
export default class Day extends DayProps {
  get dayString() {
    if (this.day == null) return "";
    return toDateVuetify(this.day);
  }
}
</script>

<style scoped></style>
