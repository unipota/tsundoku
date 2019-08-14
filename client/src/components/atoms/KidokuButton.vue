<template lang="pug">
  .check-button(@click="handleClick" :class="{'disable': disable}")
    icon(name="kidoku" color="white" :width="19" :height="14")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class KidokuButton extends Vue {
  @Prop({ type: Boolean, default: false })
  private disable!: boolean

  handleClick(e: MouseEvent) {
    if (this.disable) return
    this.$emit('click', e)
  }
}
</script>

<style lang="sass" scoped>
.check-button
  cursor: pointer
  display: flex
  align-items: center
  justify-content: center
  width: 36px
  height: 36px
  border:
    radius: 100%
  background: var(--kidoku-blue)
  transition: background .3s

  &:hover
    background: var(--kidoku-blue-hovered)

  &.disable
    cursor: auto
    background: var(--text-gray)
</style>
