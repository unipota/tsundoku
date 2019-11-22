<template lang="pug">
  .book-details-action-button(@click="handleClick" :class="{ expanded }")
    .icon
      icon(:name="icon" color="white" :width="iconSize" :height="iconSize")
    span.label
      | {{ label }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class BookDetailsActionButton extends Vue {
  @Prop({ type: Boolean, default: false })
  private expanded!: boolean

  @Prop({ type: String, default: '' })
  private label!: boolean

  @Prop({ type: String, default: '' })
  private icon!: boolean

  @Prop({ type: Number, default: 18 })
  private iconSize!: boolean

  handleClick(e: MouseEvent) {
    this.$emit('click', e)
  }
}
</script>

<style lang="sass" scoped>
$width: 80px
$height: 36px
$width-sm: 40px

.book-details-action-button
  position: relative
  cursor: pointer
  height: $height
  width: $width
  border:
    radius: 9999vw
  padding:
    left: 12px
    right: 12px
  background: $text-black-fade30
  display: flex
  align-items: center
  justify-content: space-between
  color: white
  transition: all 0.3s ease
  overflow: hidden

  &:hover
    background: $text-black-fade60

  @media (max-width: 340px)
    width: $width-sm
    justify-content: center

  &:not(.expanded)
    width: $width-sm
    justify-content: center

.icon
  flex-shrink: 0
  width: 20px
  height: 20px
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
  transition: all 0.3s ease

  @media (max-width: 340px)
    width: 0
    opacity: 0

  .book-details-action-button:not(.expanded) &
    position: absolute
    width: 0
    opacity: 0
    transform: translateX(10px)
</style>
