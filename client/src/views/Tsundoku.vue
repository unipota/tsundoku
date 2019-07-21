<template lang="pug">
  div
    router-view
    floating-add-tsundoku-button
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import FloatingAddTsundokuButton from '@/components/atoms/FloatingAddTsundokuButton.vue'

@Component({ components: { FloatingAddTsundokuButton } })
export default class Tsundoku extends Vue {}
</script>

<style lang="sass"></style>
