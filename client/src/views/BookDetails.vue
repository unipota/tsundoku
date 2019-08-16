<template lang="pug">
  modal-frame(path="../../" closeColor="white" no-padding)
    .book-details
      .header(ref="header")
        .header-bg(:style="headerBgStyle")
          .info-bg
            book-major-info(:title="book.title" :authors="book.author")
        .header-fade
        .info(ref="info")
          book-major-info(:title="book.title" :authors="book.author")
        .actions
          book-details-action-button.action(
            :expanded="isButtonExpanded"
            icon="pen"
            :label="$t('edit')" :iconSize="20"
            v-tooltip="'本の情報を編集する'")
          book-details-action-button.action(
            :expanded="isButtonExpanded"
            icon="remove"
            :label="$t('delete')"
            @click="handleDeleteClick"
            v-tooltip="'この本を削除する'")
      .cover-wrap(ref="coverWrap")
        book-cover(:url="book.coverImageUrl" :hasShadow="true")
      .body-wrap(ref="bodyWrap")
        .body(ref="body")
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

const initialHeaderHeight = 200
const slimHeaderHeight = 120

const initialCoverTop = 120
const slimInfoLeftPadding = 100

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
  private currentHeaderHeight = initialHeaderHeight
  private animationFrameRequestId?: number
  private animationEndTimeoutId?: number

  public async mounted() {
    this.lastBook = this.book
    if (!(this.$refs.bodyWrap instanceof Element)) {
      return
    }
    window.addEventListener('resize', this.handleResize)
    this.handleResize()
    this.updateHeader()
  }

  public destroy() {
    window.removeEventListener('resize', this.handleResize)
    this.handleResize()
    if (this.animationFrameRequestId) {
      cancelAnimationFrame(this.animationFrameRequestId)
    }
    if (this.animationEndTimeoutId) {
      window.clearTimeout(this.animationEndTimeoutId)
    }
  }

  public handleResize() {
    if (
      !(this.$refs.body instanceof Element) ||
      !(this.$refs.bodyWrap instanceof Element)
    ) {
      return
    }
    // 本体がラッパー以上の場合はラッパー + ヘッダの高さ変化分までは伸ばす
    const {
      height: bodyWrapHeight
    } = this.$refs.bodyWrap.getBoundingClientRect()
    const { height: bodyHeight } = this.$refs.body.getBoundingClientRect()
    if (bodyHeight >= bodyWrapHeight) {
      this.$refs.body.style.minHeight = `${bodyWrapHeight +
        initialHeaderHeight -
        slimHeaderHeight}px`
    } else {
      this.$refs.body.style.minHeight = ''
    }
  }

  public updateHeader() {
    if (
      !(this.$refs.body instanceof Element) ||
      !(this.$refs.bodyWrap instanceof Element) ||
      !(this.$refs.header instanceof Element) ||
      !(this.$refs.info instanceof Element) ||
      !(this.$refs.coverWrap instanceof Element)
    ) {
      return
    }

    const {
      top: bodyWrapTop,
      height: bodyWrapHeight
    } = this.$refs.bodyWrap.getBoundingClientRect()
    const {
      top: bodyTop,
      height: bodyHeight
    } = this.$refs.body.getBoundingClientRect()

    const scrollAmount = bodyWrapTop - bodyTop
    let newHeight = initialHeaderHeight - scrollAmount

    if (newHeight < slimHeaderHeight) {
      newHeight = slimHeaderHeight
    } else if (newHeight > initialHeaderHeight + 80) {
      newHeight = initialHeaderHeight - (scrollAmount - 80) * 0.25
    } else if (newHeight > initialHeaderHeight) {
      newHeight = initialHeaderHeight - scrollAmount * 0.5
    }

    if (
      Math.abs(this.currentHeaderHeight - newHeight) >=
        (initialHeaderHeight - slimHeaderHeight) * 1 &&
      scrollAmount <= initialHeaderHeight - slimHeaderHeight &&
      scrollAmount >= 0
    ) {
      this.animationFrameRequestId = requestAnimationFrame(() =>
        this.updateHeader()
      )
      return
    }

    const progress =
      (initialHeaderHeight - this.currentHeaderHeight) /
      (initialHeaderHeight - slimHeaderHeight)

    // リアクティブガン無視
    this.$refs.header.style.height = `${newHeight}px`
    this.$refs.info.style.left = `${Math.max(
      0,
      progress * slimInfoLeftPadding
    )}px`
    this.$refs.coverWrap.style.top = `${(1 - progress) * initialCoverTop}px`
    this.$refs.coverWrap.style.transform = `scale(${1 - progress / 2})`

    this.currentHeaderHeight = newHeight
    this.animationFrameRequestId = requestAnimationFrame(() =>
      this.updateHeader()
    )
  }

  public async handleDeleteClick() {
    await this.$store.dispatch('deleteBook', { id: this.book.id })
    this.$router.push('../../')
  }

  get isButtonExpanded() {
    return (
      this.currentHeaderHeight >= (initialHeaderHeight + slimHeaderHeight) / 2
    )
  }

  get headerTransitionProgress() {
    return (
      (initialHeaderHeight - this.currentHeaderHeight) /
      (initialHeaderHeight - slimHeaderHeight)
    )
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
  position: relative
  width: 100%
  height: 100%

  overflow-y: hidden
  overflow-

.header
  position: absolute
  top: 0
  left: 0
  z-index: 3

  height: 200px
  width: 100%
  padding: 24px

  overflow: hidden

  will-change: height

.body-wrap
  position: absolute
  height: 100%
  width: 100%
  flex:
    grow: 1
    shrink: 1
  overflow-y: scroll
  -webkit-overflow-scrolling: touch

.body
  padding-top: 200px
  // min-height: calc(100% + 100px)

.header-bg
  $blur-radius: 20px

  position: absolute
  top: -$blur-radius
  left: -$blur-radius
  width: calc(100% + #{$blur-radius * 2})
  height: calc(100% + #{$blur-radius * 2})

  background:
    color: rgb(0, 128, 255)
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
  left: 0
  width: 100%
  position: relative
  margin:
    top: 12px
    left: 8px
    right: 32px
    bottom: 16px

  color: white
  will-change: left, width

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
  top: 120px
  left: 24px
  z-index: 4
  transform: scale(1)
  will-change: top, transform

.spacer
  height: 60px

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
