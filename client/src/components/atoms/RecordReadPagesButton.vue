<template lang="pug">
  .record-read-pages-button(@click="handleClick" :class="{active: active}")
    .button-inner(v-if="active" key="active")
      .icon
        icon(name="check" color="currentColor" :width="14" :height="14")
      span.label
        | 完了
    .button-inner(v-else key="non-active")
      .icon
        icon(name="bookmark" color="currentColor" :width="18" :height="18")
      span.label
        | {{$t('record')}}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class RecordReadPagesButton extends Vue {
  @Prop({ type: Boolean, default: false })
  private active!: boolean

  handleClick(e: MouseEvent) {
    this.$emit('click', e)
  }
}
</script>

<style lang="sass" scoped>
$width: 72px
$height: 36px

.record-read-pages-button
  position: relative
  cursor: pointer
  width: $width
  height: $height
  border:
    radius: 9999vw
  padding:
    left: 8px
    right: 8px

  &:before
    position: absolute
    z-index: -1
    left: 0
    top: 0
    content: ''
    display: block
    width: $width
    height: $height
    border:
      radius: 9999vw
      style: solid
      width: $height / 2
      color: var(--kidoku-blue)
    box-sizing: border-box
    transition: border-width .3s $easeInOutQuint, border-color .3s

  &.active:before
    border:
      style: solid
      width: 4px
      color: var(--kidoku-blue)

  &:hover::before
    border:
      color: var(--kidoku-blue-hovered)

.button-inner
  color: white
  width: 100%
  height: 100%
  display: flex
  align-items: center
  justify-content: center
  transition: color .3s

  .active &
    color: var(--kidoku-blue)

.icon
  flex-shrink: 0
  width: 18px
  height: 18px
  display: flex
  align-items: center
  justify-content: center

.label
  margin:
    left: 2px
  font:
    weight: bold
    size: 1rem
  white-space: nowrap
  user-select: none

.transition-button
  &-enter-active, &-leave-active
</style>
