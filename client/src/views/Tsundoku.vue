<template lang="pug">
  .tsundoku
    portal(to="priceDisplay")
      price-display(key="price-display" tsundoku :price="tsundokuPrice")
    .view-header-container
    .view
      .list-item-container(v-for="book in books")
        book-list-item(:book="book")
    portal(to="modalView")
      transition(:name="`modal-show-${$store.state.viewType}`")
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import PriceDisplay from '@/components/atoms/PriceDisplay.vue'
import BookListItem from '@/components/organs/BookListItem.vue'

@Component({
  components: {
    PriceDisplay,
    BookListItem
  }
})
export default class Tsundoku extends Vue {
  public $store!: ExStore

  get books() {
    return this.$store.getters.tsundokuBooks
  }
  get tsundokuPrice() {
    return this.$store.getters.tsundokuPrice
  }
}
</script>

<style lang="sass" scoped>
.tsundoku
  width: 100%

.view
  width: 100%
  padding:
    left: 5%
    right: 5%

.list-item-container
  margin:
    top: 1rem
  padding:
    bottom: 1rem
  width: 100%
  border-bottom: 2px solid var(--border-gray)
</style>
