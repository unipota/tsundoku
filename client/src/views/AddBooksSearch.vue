<template lang="pug">
  modal-frame
    div
      | add books search
      input(v-model="searchQuery" @keyup.enter="submitSearchQuery")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { BookRecord } from '@/types/Book'

import ModalFrame from '@/components/atoms/ModalFrame.vue'

@Component({
  components: { ModalFrame }
})
export default class AddBooksSearch extends Vue {
  public $store!: ExStore
  searchQuery: string = ''
  searchResults: BookRecord[] = []

  submitSearchQuery() {
    this.$store
      .dispatch('searchBooksByISBN', { isbn: this.searchQuery })
      .then((res: BookRecord[]) => (this.searchResults = res))
  }
}
</script>

<style lang="sass"></style>
