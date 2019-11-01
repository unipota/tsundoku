<template lang="pug">
  .book-grid-item
    router-link.book-grid-item__body(:to="`${$route.matched[0].path}/book/${book.id}`")
      .book-grid-item__cover
        book-cover(:url="book.coverImageUrl" :hasShadow="true")
    .book-grid-item__progress(v-if="!kidoku")
      progress-bar(:progress="progressRatio")
    .book-grid-item__kidoku-price(v-if="kidoku")
      | ¥{{price.toLocaleString('ja-JP')}}
    .book-grid-item__price(v-else)
      | ¥{{remainingPrice.toLocaleString('ja-JP')}}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { BookRecord } from '../../types/Book'
import { tsundokuPrice } from '@/utils/tsundoku'

import BookCover from '@/components/atoms/BookCover.vue'
import ProgressBar from '@/components/atoms/ProgressBar.vue'

@Component({
  components: {
    BookCover,
    ProgressBar
  }
})
export default class BookGridItem extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord

  @Prop({ type: Boolean, default: false })
  private kidoku!: boolean

  public $store!: ExStore

  isHovered = false

  handleMouseover() {
    this.isHovered = true
  }

  handleMouseleave() {
    this.isHovered = false
  }

  get price(): number {
    return this.book.price
  }

  get remainingPrice(): number {
    return tsundokuPrice(
      this.book.readPages,
      this.book.totalPages,
      this.book.price
    )
  }

  get progressRatio(): number {
    return this.book.totalPages === 0
      ? 0
      : this.book.readPages / this.book.totalPages
  }
}
</script>

<style lang="sass">
.book-grid-item
  margin: 24px 14px

// .book-grid-item__body

.book-grid-item__cover
  display: flex
  align-items: center
  height: 100%

.book-grid-item__progress
  margin:
    top: 8px

.book-grid-item__kidoku-price
  text-align: right
  font:
    weight: bold
    size: 1.2rem
  color: var(--kidoku-blue)

.book-grid-item__price
  text-align: right
  font:
    weight: bold
    size: 1.2rem
  color: var(--tsundoku-red)
</style>
