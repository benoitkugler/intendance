<template>
  <calendar />
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import Calendar from "../components/sejours/Calendar.vue";

@Component({
  components: { Calendar }
})
export default class Agenda extends Vue {}
</script>
