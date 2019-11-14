<template lang="pug">
  .stats-share
    p.stats-share-title シェア
    .share-button-wrapper
      .sns-icon(@click="shareVia(0)")
        icon(name="twitter" :width="32" :height="32")
      .sns-icon(@click="shareVia(1)")
        img(src="@/img/line_logo.png")
      .sns-icon(@click="shareVia(2)")
        img(src="@/img/fb_logo.svg")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'

enum SNS {
  Twitter,
  Line,
  Facebook
}

const hashTag = '#ツンドク'
const twitterId = 'tsundokuApp'

@Component({ components: { Icon } })
export default class StatsShare extends Vue {
  $store!: ExStore

  get shareText(): string {
    return ``
  }

  async shareVia(snsName: SNS) {
    const shareURL = await this.$store.dispatch('getShareURL')
    let targetURL = ''

    switch (snsName) {
      case SNS.Twitter:
        targetURL = `https://twitter.com/intent/tweet?url=${encodeURI(
          shareURL
        )}&text=${encodeURI(
          this.shareText
        )}&via=${twitterId}&hashtags=${encodeURI(hashTag)}`
        break
      case SNS.Line:
        targetURL = `https://timeline.line.me/social-plugin/share?url=${encodeURI(
          `${this.shareText}\n${encodeURI(shareURL)}`
        )}`
        break
      case SNS.Facebook:
        targetURL = `https://www.facebook.com/sharer/sharer.php?u=${encodeURI(
          shareURL
        )}`
        break
    }
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
