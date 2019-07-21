<template lang="pug">
  div
    desktop-nav
    routerView
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import DesktopNav from '@/components/molecules/DesktopNav.vue'

@Component({
  components: { DesktopNav }
})
export default class DesktopTemplate extends Vue {}
</script>

<style lang="sass"></style>
