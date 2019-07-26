<template lang="pug">
  .tsundoku
    .view-header-container
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

import BookListItem from '@/components/organs/BookListItem.vue'

@Component({
  components: {
    BookListItem
  }
})
export default class Tsundoku extends Vue {
  public $store!: ExStore

  get books() {
    return this.$store.getters.tsundokuBooks
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
  margin: 1rem 0
  width: 100%
</style>
