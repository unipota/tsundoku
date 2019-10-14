<template lang="pug">
  .book-list-item
    router-link.book-list-item__body(:to="`${$route.matched[0].path}/book/${book.id}`")
      .book-list-item__cover
        book-cover(:url="book.coverImageUrl" :hasShadow="true")
      .book-list-item__info(:class="`${$store.getters.viewTypeClass}`")
        .book-list-item__icon
          icon(name="right-arrow")
        .book-list-item__detail
          book-major-info(:title="book.title" :authors="book.author")
        .book-list-item__price(v-if="kidoku")
          span.book-list-item__total-price.book-list-item__currency-symbol.kidoku
            | {{ book.price.toLocaleString('ja-JP') }}
          span.book-list-item__time-ago
            relative-time(:from="book.updatedAt" :locale="$store.state.locale")
        .book-list-item__price(v-else)
          span
            span.book-list-item__remaining-price.tsundoku
              span.small
                | 残り
              tweened-number.book-list-item__currency-symbol(:num="remainingPrice" formatLocal)
            span.book-list-item__total-price.tsundoku
              | /{{ book.price.toLocaleString('ja-JP') }}
          span.book-list-item__time-ago
            relative-time(:from="book.updatedAt" :locale="$store.state.locale")
    .book-list-item__progress(v-if="!kidoku")
      book-list-item-progress-controller(:book="book")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import BookCover from '@/components/atoms/BookCover.vue'
import BookMajorInfo from '@/components/atoms/BookMajorInfo.vue'
import TweenedNumber from '@/components/atoms/TweenedNumber.vue'
import BookListItemProgressController from '@/components/molecules/BookListItemProgressController.vue'
import Icon from '@/components/assets/Icon.vue'
import BookProgressPopover from '@/components/organs/BookProgressPopover.vue'
import RelativeTime from '@/components/atoms/RelativeTime.vue'

import { BookRecord } from '../../types/Book'
import { tsundokuPrice } from '@/utils/tsundoku'

@Component({
  components: {
    BookCover,
    BookMajorInfo,
    BookListItemProgressController,
    Icon,
    BookProgressPopover,
    TweenedNumber,
    RelativeTime
  }
})
export default class BookListItem extends Vue {
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

  get remainingPrice(): number {
    return tsundokuPrice(
      this.book.readPages,
      this.book.totalPages,
      this.book.price
    )
  }
}
</script>

<style lang="sass">
.book-list-item
  max-width: 700px
  margin: auto

.book-list-item__body
  display: block
  position: relative
  align-items: center
  width: 100%
  height: 160px
  color: var(--text-black)
  user-select: none

.book-list-item__cover
  position: absolute
  top: 0
  left: 0
  display: flex
  align-items: center
  height: 100%
  margin-right: 16px

.book-list-item__info
  position: relative
  z-index: -1
  margin-left: 80px

  display: flex
  flex-flow: column
  align-items: flex-start
  justify-content: space-between

  width: calc(100% - 80px)
  height: 100%

  padding: 16px
  border-radius: 8px
  background-color: $bg-suppressed-gray
  transition: box-shadow .3s

  .book-list-item__body:hover &
    box-shadow: 0 0 0 4px rgba(0,0,0,0.1)

  &.is-desktop
    padding:
      top: 12px
      right: 12px
      bottom: 12px
      left: calc(12px + 5%)

  &.is-mobile
    padding:
      left: calc(20px + 5%)

.book-list-item__detail
  width: 100%
  margin-right: 4px
  flex:
    grow: 1
    shrink: 1

.book-list-item__title
  font:
    size: 1.2rem
    weight: bold
  width: 100%
  overflow: hidden

  display: -webkit-box
  -webkit-box-orient: vertical

  .is-mobile &
    -webkit-line-clamp: 1
  .is-desktop &
    -webkit-line-clamp: 2

  max-height: 3.6rem

.book-list-item__author
  height: 1.5rem
  width: 100%
  overflow: hidden
  white-space: nowrap
  text-overflow: ellipsis

.book-list-item__progress
  align-self: flex-end

  .is-mobile &
    width: 100%

  .is-desktop &
    flex:
      basis: 280px
      grow: 0
      shrink: 1

.book-list-item__price
  display: flex
  flex-flow: column

.book-list-item__remaining-price
  font:
    weight: bold
    size: 1.4rem
  color: var(--tsundoku-red)

  .small
    font:
      size: 1rem
    color: var(--tsundoku-red-fade60)
    margin:
      right: 8px

.book-list-item__total-price
  line-height: 1.4rem

  &.tsundoku
    font:
      weight: bold
      size: 1rem
    color: var(--tsundoku-red-fade60)

  &.kidoku
    font:
      weight: bold
      size: 1.4rem
    color: var(--kidoku-blue)

.book-list-item__currency-symbol::before
  content: '¥'
  margin:
    right: 0.25rem

.book-list-item__time-ago
  color: var(--text-gray)

.book-list-item__icon
  position: absolute
  right: 16px
  bottom: 14px
  color: var(--border-gray)
</style>
