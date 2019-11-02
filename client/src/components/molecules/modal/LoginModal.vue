<template lang="pug">
  popup-modal-frame(name="login" title="登録/ログイン")
    div.page-wrapper
      //- div.title-logo
      //-   title-logo
      div.columns
        div.column.icon-cloud-tsundoku
          icon(name="cloud")
          icon(name="tsundoku" color="var(--text-gray)" :width="28" :height="24")
        div.column.descriptions-wrapper
          div.description
            icon(name="check" color="var(--kidoku-blue)" :width="checkWidth" :height="checkHeight").icon-check
            span.text
              | {{ $t('multi_device') }}
          div.description
            icon(name="check" color="var(--kidoku-blue)" :width="checkWidth" :height="checkHeight").icon-check
            span.text
              | {{ $t('backup') }}
      div.buttons
        span.button(v-for="(service, key) in services" @click="moveToLink(service.link)")
          icon.icon(:name="key" :width="service.iconWidth" :height="service.iconHeight")
          span.text
            | {{ $t( 'login_with_' + key) }}
      div.will-not-post
        icon.icon(name="exclamation-mark")
        div.text
          | {{ $t('will_not_post') }}
      div.terms-and-privacy-policy
        span.link(@click="openAboutModal")
          | ツンドクについて
        br
        span.link(@click="openPrivacyPolicyModal")
          | {{ $t('privacy_policy') }}
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import PopupModalFrame from '@/components/atoms/PopupModalFrame.vue'
import Icon from '@/components/assets/Icon.vue'
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
  components: { PopupModalFrame, Icon, TitleLogo }
})
export default class LoginModal extends Vue {
  private checkWidth = 18
  private checkHeight = 14.14
  private services: Services = {
    // TODO: google, github の認証URLが違う場合は書き換える
    google: {
      name: 'Google',
      link: '/auth/google/oauth',
      iconWidth: 21,
      iconHeight: 21
    },
    twitter: {
      name: 'Twitter',
      link: '/auth/twitter/oauth',
      iconHeight: 17.44,
      iconWidth: 21.53
    },
    github: {
      name: 'GitHub',
      link: '/auth/github/oauth',
      iconHeight: 19.56,
      iconWidth: 20
    }
  }
  public moveToLink(link: string) {
    location.href = link
  }

  openAboutModal() {
    this.$store.commit('modal/push', { name: 'about' })
  }

  openPrivacyPolicyModal() {
    this.$store.commit('modal/push', { name: 'privacy' })
  }
}
</script>

<style lang="sass" scoped>
.page-wrapper
  text-align: center
  display: block
  padding:
    bottom: 32px
  .title-logo
    margin:
      top: 24px
  .columns
    display: flex
    width: max-content
    margin: 36px auto 0 auto
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
          color: var(--text-black)
          font-weight: bold
          margin:
            left: 8px
  .buttons
    margin:
      top: 32px
    .button
      display: flex
      background: var(--border-gray)
      border-radius: 29px
      height: 38px
      width: 270px
      margin: 0 auto
      font-weight: bold
      cursor: pointer
      transition: background .3s
      &:not(:first-child)
        margin:
          top: 12px
      .icon
        margin: auto 0 auto 14px
      .text
        color: var(--text-black)
        margin: auto 0 auto 24px

      &:hover
        background: var(--border-gray-hovered)

  .will-not-post
    display: flex
    background: var(--danger-red-bg)
    color: var(--danger-red)
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
      top: 36px
    color: var(--text-black-fade60)
    font-weight: bold

  .link
    cursor: pointer
</style>
