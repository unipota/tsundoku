<template lang="pug">
  .kidoku
    portal(to="priceDisplay")
      price-display(key="price-display" kidoku :price="kidokuPrice")
    .view-header-container
    .view
      .empty(v-if="books.length === 0")
        books-empty(name="kidoku")
      .list-item-container(v-else v-for="book in books")
        book-list-item(:book="book" kidoku)
    portal(to="modalView")
      transition(name="modal-show")
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import PriceDisplay from '@/components/atoms/PriceDisplay.vue'
import BooksEmpty from '@/components/molecules/BooksEmpty.vue'
import BookListItem from '@/components/organs/BookListItem.vue'

@Component({
  components: {
    PriceDisplay,
    BooksEmpty,
    BookListItem
  }
})
export default class Kidoku extends Vue {
  public $store!: ExStore

  get books() {
    return this.$store.getters.kidokuBooks
  }
  get kidokuPrice() {
    return this.$store.getters.kidokuPrice
  }
}
</script>

<style lang="sass" scoped>
.kidoku
  width: 100%

.view
  width: 100%
  padding:
    left: 5%
    right: 5%

.list-item-container
  margin:
    top: 1.5rem
  padding:
    bottom: 1.5rem
  width: 100%
</style>
