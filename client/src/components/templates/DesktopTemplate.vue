<template lang="pug">
  div
    desktop-tab-bar
    routerView
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import DesktopTabBar from '@/components/molecules/DesktopTabBar.vue'

@Component({
  components: { DesktopTabBar }
})
export default class DesktopTemplate extends Vue {}
</script>

<style lang="sass"></style>
