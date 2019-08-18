<template lang="pug">
  .book-list-item-progress
    .book-list-item-progress__progress
      | {{ progressPercentStr }}
      .book-list-item-progress__progress-bar
        progress-bar(:progress="progressRatio")
    //- .book-list-item-progress__price
      //- span.book-list-item-progress__price_remaining_label
      //-   | {{ $t('remaining') }}
      //- span.book-list-item-progress__price_remaining
      //-   | {{ remainingPrice }}
      //- span.book-list-item-progress__price_total
      //-   | {{ book.price.toLocaleString() }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'
import ProgressBar from '@/components/atoms/ProgressBar.vue'

@Component({
  components: {
    ProgressBar
  }
})
export default class BookListItemProgress extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord

  get progressRatio(): number {
    return this.book.readPages / this.book.totalPages
  }
  get progressPercentStr(): string {
    return `${Math.round((this.book.readPages / this.book.totalPages) * 100)}%`
  }
  get remainingPrice(): string {
    return `${Math.round(
      (1 - this.book.readPages / this.book.totalPages) * this.book.price
    ).toLocaleString()}`
  }
}
</script>

<style lang="sass">
.book-list-item-progress
  width: 100%

.book-list-item-progress__progress
  margin: 0.5rem 0
  color: $kidoku-blue
  font:
    weight: bold

.book-list-item-progress__price
  display: flex
  align-items: baseline
  justify-content: flex-end
  padding: 0 0.25rem

.book-list-item-progress__price_remaining_label
  color: $tsundoku-red-fade60
  font:
    weight: bold
    size: 0.9rem

.book-list-item-progress__price_total
  color: $tsundoku-red-fade60
  font:
    weight: bold
    size: 0.9rem
  &::before
    content: '/'
    margin: 0 0.25rem

.book-list-item-progress__price_remaining
  color: $tsundoku-red
  font:
    weight: bold
    size: 1.1rem
  &::before
    content: '¥'
    margin: 0 0.25rem
</style>
