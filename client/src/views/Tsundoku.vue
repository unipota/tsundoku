<template lang="pug">
  .tsundoku
    portal(to="priceDisplay")
      price-display(key="price-display" tsundoku :price="tsundokuPrice")
    .view-header-container(v-if="books.length !== 0")
      list-controller(:filterText.sync="filterText")
    transition-group.view(
      tag="div" 
      name="transition-item")
      .empty(v-if="books.length === 0")
        books-empty(name="tsundoku")
      .list-item-container(
        v-else 
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
import { BookRecord } from '../types/Book'

import PriceDisplay from '@/components/atoms/PriceDisplay.vue'
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
    BookListItem
  }
})
export default class Tsundoku extends Vue {
  public $store!: ExStore
  filterText: string = ''

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
    return this.filteredIds.map((id: string) =>
      this.books.find(book => book.id === id)
    )
  }
}
</script>

<style lang="sass" scoped>
.tsundoku
  width: 100%

.view
  position: relative
  max-width: 100%
  margin:
    left: 5%
    right: 5%

.list-item-container
  margin:
    top: 1.5rem
  padding:
    bottom: 1.5rem
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
