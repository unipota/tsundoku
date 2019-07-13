<template lang="pug">
  div
    slot(name="topBar")
    slot(name="tabBar")
    slot(name="routerView")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class DesktopTemplate extends Vue {}
</script>

<style lang="sass"></style>
