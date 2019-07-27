<template lang="pug">
  div.page-wrapper
    page-top-bar(name="register")
    div.title-logo
      title-logo
    div.columns
      div.column.icon-cloud-tsundoku
        icon(name="cloud")
        icon(name="tsundoku" color="var(--text-gray)" :width="28" :height="24")
      div.column.descriptions-wrapper
        div.description
          icon(name="check" :color="checkColor" :width="checkWidth" :height="checkHeight").icon-check
          span.text
            | {{ $t('multi_device') }}
        div.description
          icon(name="check" :color="checkColor" :width="checkWidth" :height="checkHeight").icon-check
          span.text
            | {{ $t('backup') }}
    div.register-buttons
      a.register-button(v-for="(service, key) in services" :href="service.link")
        icon.icon(:name="key" :width="service.iconWidth" :height="service.iconHeight")
        span.text
          | {{ $t('register_with_' + key) }}
    div.will-not-post
      icon.icon(name="exclamation-mark")
      div.text
        | {{ $t('will_not_post') }}
    div.terms-and-privacy-policy
      router-link.terms(to="/terms")
        | {{ $t('terms') }}
      span.dot
        | ･
      router-link.privacy-policy(to="/privacy-policy")
        | {{ $t('privacy_policy') }}
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import PageTopBar from '@/components/molecules/PageTopBar.vue'
import TitleLogo from '@/components/atoms/TitleLogo.vue'

type serviceNames = 'google' | 'twitter' | 'github'
type Services = {
  [P in serviceNames]: {
    name: string
    link: string
    iconWidth: number
    iconHeight: number
  }
}

@Component({
  components: {
    Icon,
    TitleLogo,
    PageTopBar,
    ModalFrame
  }
})
export default class Register extends Vue {
  private checkColor = '#085092'
  private checkWidth = 18
  private checkHeight = 14.14
  private services: Services = {
    google: {
      name: 'Google',
      link: 'https://google.com',
      iconWidth: 21,
      iconHeight: 21
    },
    twitter: {
      name: 'Twitter',
      link: 'https://twitter.com',
      iconHeight: 17.44,
      iconWidth: 21.53
    },
    github: {
      name: 'GitHub',
      link: 'https://github.com',
      iconHeight: 19.56,
      iconWidth: 20
    }
  }
}
</script>

<style lang="sass">
.page-wrapper
  text-align: center
  position: absolute
  top: 0
  left: 0
  right: 0
  display: block
  .title-logo
    margin:
      top: 102px
  .columns
    display: flex
    width: max-content
    margin: 48px auto 0 auto
    .icon-cloud-tsundoku
      svg
        display: block
        margin: 0 auto
    .descriptions-wrapper
      margin: auto 0 auto 22.83px
      .description
        &:not(:first-child)
          margin:
            top: 5px
        .text
          color: $text-gray
          font-weight: bold
          margin:
            left: 8px
  .register-buttons
    margin:
      top: 49px
    .register-button
      display: flex
      background: $border-gray
      border-radius: 29px
      height: 38px
      width: 255px
      margin: 0 auto
      font-weight: bold
      &:not(:first-child)
        margin:
          top: 12px
      .icon
        margin: auto 0 auto 14px
      .text
        margin: auto 0 auto 36px
  .will-not-post
    display: flex
    background: $danger-red-bg
    color: $danger-red
    margin: 39px auto 0 auto
    width: 301px
    height: 73px
    border-radius: 16px
    .icon
      margin: auto 0 auto 28px
    .text
      margin: auto 0 auto 31px
      width: 190px
      font-weight: bold
      opacity: 0.6
  .terms-and-privacy-policy
    margin:
      top: 44px
    color: rgba(79, 79, 79, 0.6)
    font-weight: bold
    .dot
      margin: 0 4px
</style>
