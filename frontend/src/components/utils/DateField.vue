<template>
  <v-menu :close-on-content-click="false" ref="menuDatePicker">
    <template v-slot:activator="props">
      <v-text-field
        readonly
        v-on="props.on"
        :value="showDate(currentDate)"
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
import { Formatter } from "@/logic/formatter";
import { Time } from "@/logic/api";

const Props = Vue.extend({
  model: {
    prop: "currentDate",
    event: "dateChanged"
  },
  props: {
    // Date as string
    currentDate: String as () => Time,
    label: String,
    disabled: Boolean
  }
});

@Component
export default class DateField extends Props {
  isoDate = "";

  showDate = Formatter.formatDate;
  toDateObject(date: string | any) {
    if (!date) return null;
    return date + "T00:00:00Z";
  }
  save(date: string) {
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
