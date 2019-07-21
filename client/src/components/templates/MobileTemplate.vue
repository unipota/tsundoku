<template lang="pug">
  div
    slot(name="topBar")
    slot(name="tabBar")
    slot(name="routerView")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MobileTabBar from '@/components/atoms/MobileTabBar.vue'

@Component({
  components: {
    MobileTabBar
  }
})
export default class MobileTemplate extends Vue {}
</script>

<style lang="sass"></style>
