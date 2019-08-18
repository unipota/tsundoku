<template lang="pug">
  .books-empty(:class="BooksEmptyClasses")
    div(v-if="$store.getters.isFirstLanding")
      .title-logo
        title-logo(:width="205" :height="50")
      .icons
        span.icon-wrapper(v-for="icon in tsundokuToKidokuIcons")
          .icon
            .label(:class="icon.name")
              | {{ $t(icon.name) }}
            icon(
              :name="icon.name"
              :width="40"
              :height="40"
            )
            .label(:class="icon.name")
              | {{ $t(icon.description) }}
          span.icon-after
            | {{ $t(icon.after) }}
      login-button
    div(v-else)
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
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'
import AddFromArrow from '@/components/assets/AddFromArrow.vue'
import LoginButton from '@/components/atoms/LoginButton.vue'
import TitleLogo from '@/components/atoms/TitleLogo.vue'

@Component({
  components: {
    Icon,
    AddFromArrow,
    LoginButton,
    TitleLogo
  }
})
export default class BooksEmpty extends Vue {
  public $store!: ExStore

  @Prop({ type: String, default: 'tsundoku' })
  name!: 'tsundoku' | 'kidoku'

  get BooksEmptyClasses() {
    let ret: Map<string, boolean> = {}
    ret[this.$store.getters.viewTypeClass] = true
    if (this.$store.getters.isFirstLanding) {
      ret['first-landing'] = true
    } else {
      ret['not-first-landing'] = true
    }
    return ret
  }

  get iconColor() {
    return this.name === 'tsundoku'
      ? 'var(--tsundoku-red-bg)'
      : 'var(--kidoku-blue-bg)'
  }

  get tsundokuToKidokuIcons() {
    return [
      {
        name: 'tsundoku',
        width: 40,
        height: 40,
        description: 'debt',
        after: 'tsundoku-wo'
      },
      {
        name: 'kidoku',
        width: 40,
        height: 40,
        description: 'property',
        after: 'kidoku-ni'
      }
    ]
  }
}
</script>

<style lang="sass" scoped>
.books-empty
  display: flex
  align-items: center
  flex-direction: column
  position: relative

  width: 100%

  // それぞれのviewの高さに合わせる
  // [TODO] コンテナ側で高さを100%にしたい
  &.is-desktop
    height: calc(100vh - 48px)
  &.is-mobile
    height: calc(100vh - 64px - 90px)

.first-landing
  text-align: center

  .title-logo
    margin:
      bottom: 51px

  .icons
    display: flex
    margin:
      bottom: 35px

  .icon-wrapper
    display: flex
    &:not(:first-child)
      margin-left: 12px

  .icon
    text-align: center
    width: 64px // ツンドクとキドクで幅を揃えるため
    margin: auto 12px auto 0

  .label
    font-weight: bold
    margin: 5px auto
    &.tsundoku
      color: var(--tsundoku-red)
    &.kidoku
      color: var(--kidoku-blue)

  .icon-after
    margin: auto 0
    font-weight: bold
    color: var(--text-gray)

.not-first-landing
  justify-content: center

  .icon
    position: absolute
    top: calc(50% - 50px)
    left: 50%

    width: 100px
    height: 100px

    transform: translate(-50%, -50%)

.message-container
  position: absolute
  top: calc(50% + 18vh)
  left: 50%
  width: 100%
  transform: translate(-50%, -50%)

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
  position: absolute
  display: none
  .is-mobile &
    display: block

  bottom: 60px
  max-width: 120px
  right: 0
  height: calc(50vw - 110px)
  width: calc(50vw - 110px)

  @media(min-height: 600px)
    bottom: 0
    top: calc(50% + 21vh)
    transform: translateY(-50%)

  &.arrow-container--mobile--portrait
    display: none
    @media(orientation: portrait)
      display: block

  &.arrow-container--mobile--landscape
    display: none
    @media(orientation: landscape)
      display: block
      bottom: 10px
      max-width: 100vw
      right: 60px
      height: calc(50% - 40px)
      width: calc(50% - 160px)

.arrow-container--desktop
  position: absolute
  display: none
  .is-desktop &
    display: block
</style>
