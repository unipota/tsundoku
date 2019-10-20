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
    .empty(v-if="isEmpty" key="empty")
      books-empty(name="tsundoku")
    transition-group.view(
      v-else
      tag="div" 
      name="transition-item")
      .list-item-container(
        v-for="book in filteredBooks"
        :key="book.id")
        book-list-item(:book="book")
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

import PriceDisplay from '@/components/atoms/PriceDisplay.vue'
import SortBy from '@/components/atoms/SortBy.vue'
import ListController from '@/components/molecules/ListController.vue'
import BooksEmpty from '@/components/molecules/BooksEmpty.vue'
import BookListItem from '@/components/organs/BookListItem.vue'

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
    SortBy
  }
})
export default class Tsundoku extends Vue {
  public $store!: ExStore
  filterText: string = ''

  sortItems: string[] = ['購入価格', 'ツンドク残額', '総ページ数', '更新日']
  sortItemId: number = 0
  sortIsDesc: boolean = true

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
        return comparePrice(this.sortIsDesc)
      case 1:
        return compareTsundokuPrice(this.sortIsDesc)
      case 2:
        return compareTotalPages(this.sortIsDesc)
      case 3:
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
  justify-content: space-between
  padding:
    top: 12px
    left: calc(5% + 6px)
    right: calc(5% + 6px)

.books-number
  font:
    weight: bold
    size: 16px
  color: var(--text-gray)

.empty
  width: 100%

.view
  position: relative
  width: 100%
  padding:
    left: 5%
    right: 5%

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
