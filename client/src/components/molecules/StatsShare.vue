<template lang="pug">
  .stats-share
    p.stats-share-title シェア
    .share-button-wrapper
      .sns-icon(@click="shareVia('twitter')")
        icon(name="twitter" :width="32" :height="32")
      .sns-icon(@click="shareVia('line')")
        img(src="@/img/line_logo.png")
      .sns-icon(@click="shareVia('facebook')")
        img(src="@/img/fb_logo.svg")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { SNS, getShareText, getTargetURL } from '@/utils/share'

import Icon from '@/components/assets/Icon.vue'

@Component({ components: { Icon } })
export default class StatsShare extends Vue {
  $store!: ExStore

  @Prop({ type: Number, required: true })
  tsundokuPrice!: number

  @Prop({ type: Number, required: true })
  booksCount!: number

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
.stats-share
  display: flex
  align-items: center
  justify-content: space-between
  width: 320px
  padding: 16px 24px
  border:
    radius: 24px
  background: var(--bg-gray)

.share-button-wrapper
  display: flex
  align-items: center
  height: 100%

.stats-share-title
  font:
    size: 18px
    weight: bold

.sns-icon
  width: 32px
  height: 32px
  margin:
    right: 16px
  cursor: pointer
</style>
