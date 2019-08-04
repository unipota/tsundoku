<template lang="pug">
  .tsundoku
    portal(to="priceDisplay")
      price-display(key="price-display" tsundoku :price="tsundokuPrice")
    .view-header-container
      list-controller
    .view
      .list-item-container(v-for="book in books")
        book-list-item(:book="book")
    portal(to="modalView")
      transition(name="modal-show")
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import PriceDisplay from '@/components/atoms/PriceDisplay.vue'
import ListController from '@/components/molecules/ListController.vue'
import BookListItem from '@/components/organs/BookListItem.vue'

@Component({
  components: {
    PriceDisplay,
    ListController,
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
    top: 1.5rem
  padding:
    bottom: 1.5rem
  width: 100%
  position: relative

  &:not(:last-child)::after
    content: ''
    display: block
    position: absolute
    bottom: 0
    left: 0
    right: 0
    margin: auto
    width: calc(100% - 36px)
    height: 2px
    border:
      radius: 9999vw
    background: var(--border-gray)
</style>
