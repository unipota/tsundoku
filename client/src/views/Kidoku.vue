<template lang="pug">
  .kidoku
    .view-header-container
    .view
      .list-item-container(v-for="book in books")
        BookListItem(:book="book")
    router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import BookListItem from '@/components/molecules/BookListItem.vue'
import { ExStore } from 'vuex'

@Component({
  components: {
    BookListItem
  }
})
export default class Kidoku extends Vue {
  public $store!: ExStore
  get books() {
    return this.$store.getters.getKidoku
  }
}
</script>

<style lang="sass"></style>
