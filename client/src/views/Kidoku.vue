<template lang="pug">
  .kidoku
    .view-header-container
    .view
      .list-item-container(v-for="book in books")
        BookListItem(:book="book" kidoku)
    portal(to="modalView")
      modal-frame(v-show="routeDepth > 1")
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import ModalFrame from '@/components/atoms/ModalFrame.vue'
import BookListItem from '@/components/organs/BookListItem.vue'

@Component({
  components: {
    ModalFrame,
    BookListItem
  }
})
export default class Kidoku extends Vue {
  public $store!: ExStore
  get routeDepth() {
    return this.$route.matched.length
  }
  get books() {
    return this.$store.getters.getKidoku
  }
}
</script>

<style lang="sass" scoped>
.kidoku
  width: 100%
.view
  width: 100%
.list-item-container
  margin: 1rem 0
  width: 100%
</style>
