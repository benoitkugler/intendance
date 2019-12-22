<template>
  <v-menu :close-on-content-click="false" ref="menuDatePicker">
    <template v-slot:activator="{ on }">
      <v-text-field
        readonly
        v-on="on"
        :value="date(currentDate)"
        :label="label"
        :disabled="disabled"
      ></v-text-field>
    </template>
    <v-date-picker
      locale="fr-fr"
      v-model="isoDate"
      no-title
      ref="datePicker"
      @change="save"
      :max="maxDate"
      full-width
    ></v-date-picker>
  </v-menu>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Watch, Model } from "vue-property-decorator";

const Props = Vue.extend({
  model: {
    prop: "currentDate",
    event: "dateChanged"
  },
  props: {
    // Date as string
    currentDate: String,
    label: String,
    disabled: Boolean
  }
});

const Months = [
  "Janvier",
  "Février",
  "Mars",
  "Avril",
  "Mai",
  "Juin",
  "Juillet",
  "Août",
  "Septembre",
  "Octobre",
  "Novembre",
  "Décembre"
];

function showDate(dateString: string) {
  dateString = dateString || "";
  if (dateString.length < 10 || dateString.substr(0, 10) == "0001-01-01") {
    return null;
  }
  const year = dateString.substr(0, 4);
  const month = Number(dateString.substr(5, 2));
  const day = dateString.substr(8, 2);
  return `${day} ${Months[month - 1]} ${year}`;
}

@Component
export default class DateField extends Props {
  isoDate = "";

  date = showDate;
  toDateObject(date) {
    if (!date) return null;
    return date + "T00:00:00Z";
  }
  save(date) {
    (this.$refs.menuDatePicker as any).save(date);
  }
  get maxDate() {
    var d = new Date();
    d.setFullYear(d.getFullYear() + 3);
    return d.toISOString().substr(0, 10);
  }

  @Watch("isoDate")
  onIsoDateChanged() {
    const currentDate = this.toDateObject(this.isoDate);
    this.$emit("dateChanged", currentDate);
  }
}
</script>

<style></style>
