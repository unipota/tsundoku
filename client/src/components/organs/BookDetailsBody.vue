<template lang="pug">
  .book-details-body
    .spacer
    .controller
      book-list-item-progress-controller(:book="book")
    .info-items
      book-details-item.item(
        name="ツンドク残額"
        :value="`¥ ${remainingPrice.toLocaleString()}`"
        valueColor="var(--tsundoku-red)")
      book-details-item.item(
        :name="$t('price')"
        :value="`¥ ${price.toLocaleString()}`")
      book-details-item.item(
        :name="$t('totalPages')"
        :value="`${totalPages}`")
      book-details-item.item(
        name="最後に読んだ日"
        value="0日前")
      book-details-item.item(
        textarea
        clickable
        name="メモ"
        placeholder="メモがありません"
        :value="book.memo")
      book-details-item.item(
        :name="$t('overview')"
        :value="book.caption")
      book-details-item.item(
        :name="$t('publisher')"
        :value="book.publisher")
      book-details-item.item(
        :name="$t('isbn')"
        :value="book.isbn")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'
import BookDetailsActionButton from '@/components/atoms/BookDetailsActionButton.vue'
import BookDetailsItem from '@/components/molecules/BookDetailsItem.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import BookCover from '@/components/atoms/BookCover.vue'
import BookMajorInfo from '@/components/atoms/BookMajorInfo.vue'
import BookListItemProgressController from '@/components/molecules/BookListItemProgressController.vue'

@Component({
  components: {
    BookDetailsActionButton,
    BookDetailsItem,
    ModalFrame,
    BookCover,
    BookMajorInfo,
    BookListItemProgressController
  }
})
export default class BookDetailsBody extends Vue {
  @Prop({ type: Object, required: true })
  book!: BookRecord

  get remainingPrice(): number {
    return Math.round(
      (1 - this.book.readPages / this.book.totalPages) * this.book.price
    )
  }

  get price() {
    return this.book && this.book.price
  }

  get totalPages() {
    return this.book && this.book.totalPages
  }
}
</script>

<style lang="sass" scoped>
.close
  display: flex
  justify-content: flex-end

  width: 100%

.info
  position: absolute
  width: calc(100% - 192px)
  top: 24px
  left: 32px
  z-index: 5
  position: relative

  color: white
  transform: translateZ(0)
  will-change: transform

.actions
  position: absolute
  top: 72px
  right: 0
  z-index: 5
  padding:
    top: 0
    right: 16px
    bottom: 16px
    left: 0
  display: flex

.action
  margin: 0 4px

.cover-wrap
  position: absolute
  top: 0
  left: 24px
  z-index: 4
  transform: translateY(120px) scale(1)
  will-change: transform

.controller
  padding:
    right: 24px
    left: 24px

.info-items
  padding:
    right: 32px
    left: 32px
  margin:
    top: 16px

.item
  margin:
    bottom: 16px
</style>
