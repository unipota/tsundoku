<template lang="pug">
  .view-mode(@click="toggleViewMode")
    icon(v-if="isList" name="list" color="var(--text-gray)")
    icon(v-else name="grid" color="var(--text-gray)")
</template>

<script lang="ts">
import { Vue, Component, Model } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: {
    Icon
  }
})
export default class ViewMode extends Vue {
  @Model('toggle', { type: Boolean })
  readonly isList!: boolean

  toggleViewMode() {
    this.$emit('toggle', !this.isList)
  }
}
</script>

<style lang="sass" scoped>
.view-mode
  display: flex
  justify-content: center
  align-items: center
  position: relative
  font:
    weight: bold
    size: 16px
  color: var(--text-gray)
  cursor: pointer

  &::before
    content: ''
    display: block
    position: absolute
    z-index: -1
    width: 32px
    height: 32px
    background: var(--bg-gray)
    border-radius: 100%
    opacity: 0
    transition: opacity .3s

  &:hover::before
    opacity: 1
</style>
