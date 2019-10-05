<template lang="pug">
  .book-list-item-progress
    .book-list-item-progress__progress
      .book-list-item-progress__progress-bar
        progress-bar(
          :progress="progressRatio" 
          :edit="edit" 
          :editedProgress="editedProgressRatio" 
          @mounted="progressBarWidth = $event.width - 4")
        .book-list-item-progress__progress-knob
          transition(name="slide-in")
            progress-knob(
              v-if="edit"
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
    return this.totalPages === 0 ? 0 : this.readPages / this.totalPages
  }
  get editedProgressRatio(): number {
    return this.syncedEditedReadPages / this.totalPages
  }
  get progressPercent(): number {
    return this.totalPages === 0
      ? 0
      : Math.round((this.readPages / this.totalPages) * 100)
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

.slide-in
  &-enter, &-leave-to
    transform: translateY(10px)
    opacity: 0
  &-enter-active, &-leave-active
    transition: transform .3s $easeOutBack, opacity .3s
</style>
