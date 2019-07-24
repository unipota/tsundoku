<template lang="pug">
  .tsundoku
    .view-header-container
    .view
      .list-item-container(v-for="book in books")
        book-list-item(:book="book")
    portal(to="modalView")
      transition(name="modal-show")
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
export default class Tsundoku extends Vue {
  public $store!: ExStore

  get routeDepth() {
    return this.$route.matched.length
  }

  get books() {
    return this.$store.getters.getTsundoku
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

.modal-show
  &-enter, &-leave-to
    transform: translateY(100%)

  &-enter-active, &-leave-active
    transition: transform .5s
</style>
