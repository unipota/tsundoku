<template lang="pug">
  div
    mobile-top-bar
    router-view
    mobile-tab-bar
    floating-add-tsundoku-button
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MobileTabBar from '@/components/atoms/MobileTabBar.vue'
import MobileTopBar from '@/components/molecules/MobileTopBar.vue'
import FloatingAddTsundokuButton from '@/components/atoms/FloatingAddTsundokuButton.vue'

@Component({
  components: {
    MobileTabBar,
    MobileTopBar,
    FloatingAddTsundokuButton
  }
})
export default class MobileTemplate extends Vue {}
</script>

<style lang="sass"></style>
