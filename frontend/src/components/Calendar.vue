<template>
  <div>
    <p>Test</p>

    <v-calendar
      type="week"
      locale="fr"
      :first-interval="firstInterval"
      :interval-count="intervalCount"
      :interval-minutes="intervalMinutes"
      :interval-height="intervalHeight"
    >
      <template v-slot:day-label>dmlkdmsmskd</template>
      <template v-slot:day-month>dmlkdmsmskd</template>
    </v-calendar>
    <v-calendar
      type="week"
      locale="fr"
      :min-weeks="firstInterval"
      :first-interval="2"
      :interval-count="intervalCount"
      :interval-minutes="intervalMinutes"
      :interval-height="intervalHeight"
    ></v-calendar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "Calendar",

  data: () => ({
    firstInterval: 2,
    intervalCount: 4,
    intervalMinutes: 240,
    intervalHeight: 50
  })
});
</script>
