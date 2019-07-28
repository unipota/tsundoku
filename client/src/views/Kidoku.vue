<template lang="pug">
  .kidoku
    portal(to="priceDisplay")
      price-display(key="price-display" kidoku :price="kidokuPrice")
    .view-header-container
    .view
      .list-item-container(v-for="book in books")
        book-list-item(:book="book" kidoku)
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
  margin: 1rem 0
  width: 100%
</style>
