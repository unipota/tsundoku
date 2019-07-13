<template lang="pug">
  div
    TabBar(selectedTab="tsundoku")
    router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import TabBar from '@/components/atoms/TabBar.vue'

@Component({
  components: {
    TabBar
  }
})
export default class Tsundoku extends Vue {}
</script>

<style lang="sass"></style>
