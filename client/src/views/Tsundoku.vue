<template lang="pug">
  div
    TabBar(selectedTab="tsundoku")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import TabBar from '@/components/TabBar.vue'


@Component({
  components: {
    TabBar,
  },
})
export default class Tsundoku extends Vue {
}
</script>

<style lang="sass">
</style>
