<template lang="pug">
  .toukei
    div
      //- books-count-stats(v-if="loaded" :allBookStats="allBookStats")
      //- tsundoku-stats(v-if="loaded" :allBookStats="allBookStats")
    portal(to="modalView")
      transition(name="modal-show")  
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { BookStats } from '@/types/Book'

import BooksCountStats from '@/components/molecules/BooksCountStats.vue'
import TsundokuStats from '@/components/molecules/TsundokuStats.vue'

@Component({
  components: { BooksCountStats, TsundokuStats }
})
export default class Toukei extends Vue {
  public $store!: ExStore

  loaded = false

  async mounted() {
    await this.$store.dispatch('getAllBookStats')
    this.loaded = true
  }

  get allBookStats(): BookStats[] {
    return this.$store.state.bookStatsArray
  }
}
</script>

<style lang="sass" scoped></style>
