<template lang="pug">
  .toukei
    .stats-container
      .stats-wrapper
        books-count-stats(v-if="loaded" :allBookStats="allBookStats")
      .stats-wrapper
        tsundoku-stats(v-if="loaded" :allBookStats="allBookStats")
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

<style lang="sass" scoped>
.stats-container
  display: flex
  flex-wrap: wrap
  align-items: center
  padding:
    top: 64px
    left: 5%
    right: 5%

.stats-wrapper
  margin: 12px 12px
</style>
