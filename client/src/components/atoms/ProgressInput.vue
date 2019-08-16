<template lang="pug">
  .progress-input
    .price-wrap
      input.current-price(
        type="number"
        autocomplete="off"
        :min="0"
        :max="totalPages"
        :value="readPages"
        @paste.prevent=""
        @keypress="isNumber($event)"
        @input="$emit('input', validateNumber($event.target.value))"
        @keydown.up.prevent="handleUpKey"
        @keydown.down.prevent="handleDownKey")
      span.price-total
        | /{{totalPages}}p
    .close-button(@click.stop="handleCancel")
      .icon-wrap
        icon(name="close" color="white" :width="12" :height="12")
      span キャンセル
</template>

<script lang="ts">
import { Vue, Component, Model, Prop } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class ProgressInput extends Vue {
  @Model('input', { type: Number })
  private readonly readPages!: number

  @Prop({ type: Number, required: true })
  private readonly totalPages!: number

  handleCancel() {
    this.$emit('cancel')
  }

  isNumber(e: KeyboardEvent) {
    const charCode = e.which ? e.which : e.keyCode
    if (
      (charCode > 31 && (charCode < 48 || charCode > 57)) ||
      this.readPages >= this.totalPages
    ) {
      e.preventDefault()
    } else {
      return true
    }
  }

  validateNumber(value: string) {
    return this.limitationNumber(Number(value))
  }

  limitationNumber(value: number): number {
    return Math.min(Math.max(value, 0), this.totalPages)
  }

  handleUpKey() {
    this.$emit('input', this.limitationNumber(this.readPages + 1))
  }

  handleDownKey() {
    this.$emit('input', this.limitationNumber(this.readPages - 1))
  }
}
</script>

<style lang="sass" scoped>
.progress-input
  display: flex
  align-items: center
  padding:
    top: 8px
    right: 4px

.price-wrap
  margin:
    right: 8px

.price-total
  font:
    weight: bold
  color: var(--kidoku-blue)
  margin:
    left: 4px

.current-price
  text-align: right
  height: 28px
  max-width: 120px
  background: var(--bg-gray)
  border:
    radius: 8px
  padding:
    right: 8px
    left: 8px
  font:
    size: 1rem
    weight: bold
  color: var(--text-gray)

.close-button
  display: inline-flex
  align-items: center
  height: 28px
  width: 116px
  padding:
    left: 4px
  border:
    radius: 9999vw
  background:
    color: $text-gray
  cursor: pointer
  transition: background-color .5s

  &:hover
    background:
      color: darken($text-gray, 10%)

  span
    user-select: none
    color: white
    font:
      size: 0.9rem
      weight: bold

.icon-wrap
  display: flex
  align-items: center
  justify-content: center
  width: 24px
  height: 24px
  margin:
    right: 2px
  border:
    radius: 9999vw
</style>
