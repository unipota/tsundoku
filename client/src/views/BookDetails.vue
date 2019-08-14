<template lang="pug">
  modal-frame(path="../../" closeColor="white" no-padding)
    .book-details
      .header
        .header-bg(:style="headerBgStyle")
          .info-bg
            book-major-info(:title="book.title" :authors="book.author")
        .header-fade
        .info
          book-major-info(:title="book.title" :authors="book.author")
        .actions
          book-details-action-button.action(
            :expanded="showSlimHeader"
            icon="pen"
            :label="$t('edit')" :iconSize="20"
            v-tooltip="'本の情報を編集する'")
          book-details-action-button.action(
            :expanded="showSlimHeader"
            icon="remove"
            :label="$t('delete')"
            @click="handleDeleteClick"
            v-tooltip="'この本を削除する'")
      .cover-wrap
        book-cover(:url="book.coverImageUrl" :hasShadow="true")
      .body-wrap
        .body
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
              textarea
              name="メモ"
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
            book-details-item.item(
              :name="$t('isbn')"
              :value="book.isbn")
            book-details-item.item(
              :name="$t('isbn')"
              :value="book.isbn")
            book-details-item.item(
              :name="$t('isbn')"
              :value="book.isbn")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { BookRecord } from '../types/Book'
import BookDetailsActionButton from '@/components/atoms/BookDetailsActionButton.vue'
import BookDetailsItem from '@/components/atoms/BookDetailsItem.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import BookCover from '@/components/atoms/BookCover.vue'
import BookMajorInfo from '@/components/atoms/BookMajorInfo.vue'
import BookListItemProgressController from '@/components/molecules/BookListItemProgressController.vue'

const headerHeight = 240
const slimHeaderHeight = 120

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
export default class BookDetails extends Vue {
  public $store!: ExStore
  private lastBook!: BookRecord // モーダル閉じるときのエラー対策
  private showSlimHeader = false

  public async mounted() {
    this.lastBook = this.book
  }

  public async handleDeleteClick() {
    await this.$store.dispatch('deleteBook', { id: this.book.id })
    this.$router.push('../../')
  }

  get book(): BookRecord {
    const bookId = this.$route.params['id']
    const book = this.$store.getters.getBookById(bookId)
    return book || this.lastBook
  }

  get headerBgStyle() {
    return {
      backgroundImage: this.book ? `url(${this.book.coverImageUrl})` : 'white'
    }
  }

  get price() {
    return this.book && this.book.price
  }

  get totalPages() {
    return this.book && this.book.totalPages
  }

  get remainingPrice(): number {
    return Math.round(
      (1 - this.book.readPages / this.book.totalPages) * this.book.price
    )
  }
}
</script>

<style lang="sass" scoped>
.book-details
  display: flex
  flex:
    direction: column
    wrap: nowrap

  position: relative
  width: 100%
  height: 100%

  overflow-y: hidden

.header
  position: relative
  top: 0
  left: 0

  width: 100%
  height: 240px
  padding: 24px

  overflow: hidden

  flex:
    grow: 0
    shrink: 0

.body-wrap
  height: 100%
  flex:
    grow: 1
    shrink: 1
  overflow-y: scroll
  padding:
    top: 40px

.header-bg
  $blur-radius: 20px

  position: absolute
  top: -$blur-radius
  left: -$blur-radius
  width: calc(100% + #{$blur-radius * 2})
  height: calc(100% + #{$blur-radius * 2})

  background:
    size: cover
    repeat: no-repeat
    position: center center

  filter: blur(20px)

.header-fade
  position: absolute
  top: 0
  left: 0
  width: 100%
  height: 100%

  background: linear-gradient(to bottom, rgba(0, 0, 0, 1) 0%, rgba(0, 0, 0, 0.7) 30%, rgba(0, 0, 0, 0) 70%)
  opacity: 0.5

.close
  display: flex
  justify-content: flex-end

  width: 100%

.info
  margin:
    top: 12px
    left: 8px
    right: 32px
    bottom: 16px
  position: relative

  color: white

.actions
  position: absolute
  bottom: 0
  right: 0
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
  top: 136px
  left: 24px

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
