<template lang="pug">
  .books-empty(:class="$store.getters.viewTypeClass")
    template(v-if="showLanding")
      .landing-wrap
        .tsundoku-logo
          icon(name="logo" :width="38" :height="38")
          .tsundoku-logo-text
            | ツンドク
        .copy-wrap
          .tsundoku-icon
            icon(name="tsundoku" :width="42" :height="42")
          | を
          .kidoku-icon
            icon(name="kidoku" :width="42" :height="42")
          | に
        .login-button-wrap
          rounded-button(@click="openLoginModal")
            | 新規登録/ログイン 
          .login-button-description
            | 後からも登録できます
    template(v-else)
      .icon
        icon(
          :name="name"
          :color="iconColor"
          :width="100"
          :height="100"
        )
    .message-container
      .message(v-if="name === 'tsundoku'")
        | {{ $t('noTsundoku') }}
      .message(v-else)
        | {{ $t('noKidoku') }}
      .message(v-if="name === 'tsundoku' && $store.state.viewType === 'mobile'")
        | {{ $t('addBooksFromHere') }}
      .message(v-if="name === 'tsundoku' && $store.state.viewType === 'desktop'")
        | {{ $t('addBooksFromSidebar') }}
    .arrow-container--mobile.arrow-container--mobile--portrait
      add-from-arrow(
        arrow-type="mobile-portrait"
        v-if="name === 'tsundoku' && $store.state.viewType === 'mobile'"
      )
    .arrow-container--mobile.arrow-container--mobile--landscape
      add-from-arrow(
        arrow-type="mobile-landscape"
        v-if="name === 'tsundoku' && $store.state.viewType === 'mobile'"
      )
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'
import AddFromArrow from '@/components/assets/AddFromArrow.vue'
import RoundedButton from '@/components/atoms/RoundedButton.vue'

@Component({
  components: {
    Icon,
    AddFromArrow,
    RoundedButton
  }
})
export default class BooksEmpty extends Vue {
  @Prop({ type: String, default: 'tsundoku' })
  name!: 'tsundoku' | 'kidoku'

  $store!: ExStore

  get iconColor() {
    return this.name === 'tsundoku'
      ? 'var(--tsundoku-red-bg)'
      : 'var(--kidoku-blue-bg)'
  }

  get showLanding() {
    return (
      !this.$store.state.userLogined && this.$store.getters.books.length === 0
    )
  }

  openLoginModal() {
    this.$store.commit('modal/push', { name: 'login' })
  }
}
</script>

<style lang="sass" scoped>
.landing-wrap
  display: flex
  flex-flow: column
  align-items: center
  padding:
    top: 48px
    bottom: 32px

.tsundoku-logo
  display: flex
  align-items: center

.tsundoku-logo-text
  font:
    weight: bold
    size: 32px
  padding:
    left: 16px
  color: var(--text-black)

.copy-wrap
  display: inline-flex
  align-items: center
  margin:
    top: 64px
  font:
    size: 18px
    weight: bold

.tsundoku-icon
  position: relative
  margin:
    right: 24px

  &::before, &::after
    position: absolute
    left: 50%
    white-space: nowrap
    color: var(--tsundoku-red)
    font:
      size: 18px

  &::before
    top: 0
    content: 'ツンドク'
    transform: translate(-50%,-100%)

  &::after
    bottom: 0
    content: '負債'
    transform: translate(-50%,100%)

.kidoku-icon
  position: relative
  margin:
    left: 24px
    right: 24px

  &::before, &::after
    position: absolute
    left: 50%
    white-space: nowrap
    color: var(--kidoku-blue)
    font:
      size: 18px

  &::before
    top: 0
    content: 'キドク'
    transform: translate(-50%,-100%)

  &::after
    bottom: 0
    content: '財産'
    transform: translate(-50%,100%)

.login-button-wrap
  margin:
    top: 54px

.login-button-description
  text-align: center
  margin:
    top: 8px
  font:
    size: 14px
    weight: bold
  color: var(--text-gray)

.books-empty
  display: flex
  align-items: center
  justify-content: center
  flex-direction: column
  position: relative

  width: 100%

  // それぞれのviewの高さに合わせる
  // [TODO] コンテナ側で高さを100%にしたい
  // &.is-desktop
  //   height: calc(100vh - 48px)
  // &.is-mobile
  //   height: calc(100vh - 64px - 90px)

.icon
  margin:
    top: 20vh
    left: auto
    right: auto
  width: 100px
  height: 100px

.message-container
  margin:
    top: 30px
  width: 100%
  display: flex
  align-items: center
  justify-content: center
  flex-direction: column

.message
  margin: 0.125rem
  font:
    weight: bold
    size: 1rem
  color: $text-gray

.arrow-container--mobile
  position: fixed
  display: none
  .is-mobile &
    display: block

  right: 40px
  bottom: 120px
  bottom: calc(env(safe-area-inset-bottom) + 120px)
  height: 80px
  width: 80px

  &.arrow-container--mobile--portrait
    display: none
    @media(orientation: portrait)
      display: block

  &.arrow-container--mobile--landscape
    display: none
    @media(orientation: landscape)
      display: block
      bottom: 120px
      bottom: calc(env(safe-area-inset-bottom) + 110px)
      right: 60px
      height: 60px
      width: 120px

.arrow-container--desktop
  display: none
  .is-desktop &
    display: block
</style>
