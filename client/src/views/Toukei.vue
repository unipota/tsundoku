<template lang="pug">
  .toukei
    .stats-tabs-wrapper
      .stats-tab
        | 累計
    .stats-container
      .stats-wrapper
        books-count-stats(v-if="loaded" :allBookStats="allBookStats")
      .stats-wrapper
        tsundoku-stats(v-if="loaded" :allBookStats="allBookStats")
      .stats-wrapper
        kidoku-stats(v-if="loaded" :allBookStats="allBookStats")
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
import KidokuStats from '@/components/molecules/KidokuStats.vue'

@Component({
  components: { BooksCountStats, TsundokuStats, KidokuStats }
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
.stats-tabs-wrapper
  display: flex
  align-items: center
  padding:
    top: 24px
    left: 5%

.stats-tab
  font:
    size: 20px
    weight: bold
  color: var(--text-black)
  border:
    radius: 9999vw
  background: var(--bg-gray)
  padding: 4px 16px
  margin:
    left: 12px
  cursor: pointer

.stats-container
  display: flex
  flex-wrap: wrap
  align-items: center
  padding:
    top: 12px
    left: 5%
    right: 5%

  .view-mobile &
    flex-flow: column

.stats-wrapper
  margin: 12px 12px
</style>
