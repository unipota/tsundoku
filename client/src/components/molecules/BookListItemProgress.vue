<template lang="pug">
  .book-list-item-progress
    .book-list-item-progress__progress
      .book-list-item-progress__progress-bar
        progress-bar(
          :progress="progressRatio" 
          :edit="edit" 
          :editedProgress="editedProgressRatio" 
          @mounted="progressBarWidth = $event.width")
        .book-list-item-progress__progress-knob(v-if="edit")
          progress-knob(
            :editedReadPages.sync="syncedEditedReadPages" 
            :totalPages="totalPages"
            :progressBarWidth="progressBarWidth")
      tweened-number(:num="progressPercent")
      | %
</template>

<script lang="ts">
import { Vue, Component, Prop, PropSync } from 'vue-property-decorator'

import ProgressBar from '@/components/atoms/ProgressBar.vue'
import ProgressKnob from '@/components/atoms/ProgressKnob.vue'
import TweenedNumber from '@/components/atoms/TweenedNumber.vue'

@Component({
  components: {
    ProgressBar,
    ProgressKnob,
    TweenedNumber
  }
})
export default class BookListItemProgress extends Vue {
  @Prop({ type: Number, required: true })
  private readPages!: number

  @Prop({ type: Number, required: true })
  private totalPages!: number

  @Prop({ type: Boolean, default: false })
  private edit!: false

  @PropSync('editedReadPages', { type: Number })
  private syncedEditedReadPages!: number

  progressBarWidth = 0

  get progressRatio(): number {
    return this.readPages / this.totalPages
  }
  get editedProgressRatio(): number {
    return this.syncedEditedReadPages / this.totalPages
  }
  get progressPercent(): number {
    return Math.round((this.readPages / this.totalPages) * 100)
  }
}
</script>

<style lang="sass">
.book-list-item-progress
  width: 100%

  &__progress
    position: relative
    padding:
      top: 4px
    color: $kidoku-blue
    font:
      weight: bold

  &__progress-knob
    position: absolute
    z-index: 10
    left: -50%
    bottom: 100%
    width: 200%

// .book-list-item-progress__price
//   display: flex
//   align-items: baseline
//   justify-content: flex-end
//   padding: 0 0.25rem

// .book-list-item-progress__price_remaining_label
//   color: $tsundoku-red-fade60
//   font:
//     weight: bold
//     size: 0.9rem

// .book-list-item-progress__price_total
//   color: $tsundoku-red-fade60
//   font:
//     weight: bold
//     size: 0.9rem
//   &::before
//     content: '/'
//     margin: 0 0.25rem

// .book-list-item-progress__price_remaining
//   color: $tsundoku-red
//   font:
//     weight: bold
//     size: 1.1rem
//   &::before
//     content: '¥'
//     margin: 0 0.25rem
</style>
