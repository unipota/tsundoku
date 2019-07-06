<template lang="pug">
  div
    TabBar(selectedTab="kidoku")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import TabBar from '@/components/TabBar.vue'

@Component({
  components: {
    TabBar,
  },
})
export default class Kidoku extends Vue {
}
</script>

<style lang="sass">
</style>
