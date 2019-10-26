<template lang="pug">
  .tsundoku
    portal(to="priceDisplay")
      price-display(key="price-display" tsundoku :price="tsundokuPrice")
    .view-header-container(v-if="!isEmpty")
      list-controller(:filterText.sync="filterText")
    .controller-container(v-if="!isEmpty")
      .books-number
        | {{booksNumber}}冊
      sort-by(
        :items="sortItems" 
        @change="handleChangeSortBy"
        @toggle="handleChangeIsDesc")
      view-mode.view-mode-icon(v-model="isList")
    .empty(v-if="isEmpty" key="empty")
      books-empty(name="tsundoku")
    .view(v-else)
      transition-group.inner-view(
        tag="div" 
        name="transition-item"
        :style="innerViewStyle")
        template(v-if="isList")
          .list-item-container(
            v-for="book in filteredBooks"
            :key="book.id")
            book-list-item(:book="book")
        template(v-else)
          .grid-item-container(
            v-for="book in filteredBooks"
            :key="book.id")
            book-grid-item(:book="book")
    portal(to="modalView")
      transition(name="modal-show")
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import Fuse from 'fuse.js'
import { kanaToHira } from '@/utils/string'
import {
  comparePrice,
  compareTsundokuPrice,
  compareUpdatedAt,
  compareTotalPages
} from '@/utils/sort'
import { BookRecord } from '../types/Book'
import { throttle } from 'lodash'

import PriceDisplay from '@/components/atoms/PriceDisplay.vue'
import SortBy from '@/components/atoms/SortBy.vue'
import ViewMode from '@/components/atoms/ViewMode.vue'
import ListController from '@/components/molecules/ListController.vue'
import BooksEmpty from '@/components/molecules/BooksEmpty.vue'
import BookListItem from '@/components/organs/BookListItem.vue'
import BookGridItem from '@/components/organs/BookGridItem.vue'

type FilterTargetBookRecord = Pick<
  BookRecord,
  'id' | 'title' | 'author' | 'caption' | 'publisher' | 'memo'
>

const options: Fuse.FuseOptions<FilterTargetBookRecord> = {
  keys: ['title', 'author', 'caption', 'publisher', 'memo']
}

@Component({
  components: {
    PriceDisplay,
    ListController,
    BooksEmpty,
    BookListItem,
    BookGridItem,
    SortBy,
    ViewMode
  }
})
export default class Tsundoku extends Vue {
  public $store!: ExStore
  filterText: string = ''

  sortItems: string[] = ['ツンドク残額', '購入価格', '総ページ数', '最終更新']
  sortItemId: number = 0
  sortIsDesc: boolean = true

  isList: boolean = true
  gridColumnNum: number = 0

  mounted() {
    window.addEventListener('resize', throttle(this.handleResizeWindow, 100))
    window.addEventListener('orientationchange', this.handleResizeWindow)
    this.$nextTick(() => {
      this.handleResizeWindow()
    })
  }

  handleResizeWindow() {
    this.gridColumnNum = Math.floor(
      (this.$el.getBoundingClientRect().width * 0.9) / 128
    )
  }

  get innerViewStyle() {
    return {
      width: this.isList ? '' : `${128 * this.gridColumnNum}px`
    }
  }
  get books(): BookRecord[] {
    return this.$store.getters.tsundokuBooks
  }
  get tsundokuPrice() {
    return this.$store.getters.tsundokuPrice
  }
  get processedBooks(): FilterTargetBookRecord[] {
    return this.books.map(book => {
      return {
        id: book.id,
        title: kanaToHira(book.title),
        author: book.author.map(author => kanaToHira(author)),
        caption: book.caption && kanaToHira(book.caption),
        publisher: kanaToHira(book.publisher),
        memo: book.memo && kanaToHira(book.memo)
      }
    })
  }
  get fuseInstance() {
    return new Fuse(this.processedBooks, options)
  }
  get filteredIds(): string[] {
    return this.fuseInstance
      .search(kanaToHira(this.filterText))
      .map(book => book.id)
  }
  get filteredBooks() {
    return this.books
      .filter(book => this.filteredIds.includes(book.id))
      .sort(this.compareFunction)
  }
  get booksNumber() {
    return this.filteredBooks.length
  }
  get isEmpty() {
    return this.books.length === 0
  }
  get compareFunction() {
    switch (this.sortItemId) {
      case 0:
        return compareTsundokuPrice(this.sortIsDesc)
      case 1:
        return comparePrice(this.sortIsDesc)
      case 2:
        return compareTotalPages(this.sortIsDesc)
      default:
        return compareUpdatedAt(this.sortIsDesc)
    }
  }

  handleChangeSortBy(id: number) {
    this.sortItemId = id
  }
  handleChangeIsDesc(isDesc: boolean) {
    this.sortIsDesc = isDesc
  }
}
</script>

<style lang="sass" scoped>
.tsundoku
  width: 100%
  padding:
    bottom: 24px

.view-header-container
  .view-desktop &
    position: sticky
    z-index: 100
    top: 0
  background: var(--bg-white)
  width: 100%
  padding:
    top: 8px
    left: 5%
    right: 5%
    bottom: 12px
  // border:
  //   bottom: solid 1px var(--border-gray)

.controller-container
  display: flex
  justify-content: flex-end
  padding:
    top: 4px
    left: calc(5% + 6px)
    right: calc(5% + 6px)

.books-number
  font:
    weight: bold
    size: 16px
  color: var(--text-gray)
  margin:
    right: auto

.view-mode-icon
  margin:
    left: 18px

.empty
  width: 100%

.view
  width: 100%
  padding:
    left: 5%
    right: 5%

.inner-view
  position: relative
  display: flex
  flex-wrap: wrap
  width: 100%
  margin: 0 auto

.list-item-container
  margin:
    top: 15px
  padding:
    bottom: 15px
  width: 100%
  position: relative

  &::after
    content: ''
    display: block
    position: absolute
    z-index: -1
    bottom: 0
    left: 0
    right: 0
    margin: auto
    width: calc(100% - 36px)
    max-width: 680px
    height: 2px
    border:
      radius: 9999vw
    transition: background .3s

  &:not(:last-child)::after
    background: var(--border-gray)

.transition-item
  &-enter, &-leave-to
    transform: translateY(10px)
    opacity: 0

  &-enter-active, &-leave-active
    transition: transform .5s $easeInOutQuint, opacity .5s

  &-leave-active
    position: absolute

  &-move
    transition: transform .5s $easeInOutQuint
</style>
