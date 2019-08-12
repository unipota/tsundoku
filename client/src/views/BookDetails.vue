<template lang="pug">
  modal-frame(path="../../" closeColor="white")
    .book-details
      .header
        .header-bg(:style="headerBgStyle")
          .info-bg
            book-major-info(:title="book.title" :authors="book.author")
        .header-fade
        .info
          book-major-info(:title="book.title" :authors="book.author")
      .cover-wrap
        book-cover(:url="book.coverImageUrl" :hasShadow="true")
      .controller
        book-list-item-progress-controller(:book="book")
      .body
        book-details-item.item(:name="$t('price')" :value="`¥ ${price}`")
        book-details-item.item(:name="$t('totalPages')" :value="totalPages")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { BookRecord } from '../types/Book'
import BookDetailsItem from '@/components/atoms/BookDetailsItem.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import BookCover from '@/components/atoms/BookCover.vue'
import BookMajorInfo from '@/components/atoms/BookMajorInfo.vue'
import BookListItemProgressController from '@/components/molecules/BookListItemProgressController.vue'

@Component({
  components: {
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

  public async mounted() {
    this.lastBook = this.book
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
    return this.book && this.book.price.toLocaleString()
  }

  get totalPages() {
    return this.book && this.book.totalPages.toString()
  }
}
</script>

<style lang="sass" scoped>
.book-details
  position: absolute
  top: 0
  left: 0
  width: 100%

.header
  position: relative
  top: 0
  left: 0

  width: 100%
  height: 240px
  padding: 24px

  overflow: hidden

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
    left: 16px
    right: 32px
    bottom: 16px
  position: relative

  color: white

.cover-wrap
  position: absolute
  top: 160px
  left: 24px

.controller
  margin-top: 64px
  padding:
    top: 12px
    right: 24px
    left: 24px

.body
  padding:
    top: 12px
    right: 32px
    left: 32px

.item
  margin-bottom: 8px
</style>
