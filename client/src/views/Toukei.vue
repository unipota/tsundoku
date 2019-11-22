<template lang="pug">
  .toukei
    portal(to="priceDisplay")
      .share-button-wrapper
        .sns-icon(@click="shareVia('twitter')")
          icon(name="twitter" :width="24" :height="24")
        .sns-icon(@click="shareVia('line')")
          img(src="@/img/line_logo.png")
        .sns-icon(@click="shareVia('facebook')")
          img(src="@/img/fb_logo.svg")
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
      .stats-wrapper
        stats-share(v-if="loaded" :tsundokuPrice="tsundokuPrice" :booksCount="booksCount")
    portal(to="modalView")
      transition(name="modal-show")  
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { BookStats } from '@/types/Book'
import { SNS, getShareText, getTargetURL } from '@/utils/share'

import Icon from '@/components/assets/Icon.vue'
import BooksCountStats from '@/components/molecules/BooksCountStats.vue'
import TsundokuStats from '@/components/molecules/TsundokuStats.vue'
import KidokuStats from '@/components/molecules/KidokuStats.vue'
import StatsShare from '@/components/molecules/StatsShare.vue'

@Component({
  components: { Icon, BooksCountStats, TsundokuStats, KidokuStats, StatsShare }
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

  get tsundokuPrice() {
    return this.$store.getters.tsundokuPrice
  }

  get booksCount(): number {
    return this.$store.getters.tsundokuBooks.length
  }

  async shareVia(snsName: SNS) {
    const shareURL = await this.$store.dispatch('getShareURL')
    let targetURL = getTargetURL(
      snsName,
      getShareText(this.tsundokuPrice, this.booksCount),
      shareURL
    )
    // iOS Safari not allow window.open
    window.open(targetURL) || (location.href = targetURL)
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
  cursor: pointer

  .view-desktop &
    margin:
      left: 12px

.stats-container
  display: flex
  flex-wrap: wrap
  padding:
    top: 12px
    left: 5%
    right: 5%
    bottom: 32px

  .view-mobile &
    flex-flow: column
    align-items: center

.stats-wrapper
  margin: 12px 12px

.share-button-wrapper
  display: flex
  align-items: center
  height: 48px
  padding: 0 18px
  border:
    radius: 999vw
  background: var(--bg-gray)

.sns-icon
  width: 24px
  height: 24px
  margin:
    left: 8px
    right: 8px
  cursor: pointer
</style>
