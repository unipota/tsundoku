<template lang="pug">
  .books-empty(:class="$store.getters.viewTypeClass")
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
import Icon from '@/components/assets/Icon.vue'
import AddFromArrow from '@/components/assets/AddFromArrow.vue'

@Component({
  components: {
    Icon,
    AddFromArrow
  }
})
export default class BooksEmpty extends Vue {
  @Prop({ type: String, default: 'tsundoku' })
  name!: 'tsundoku' | 'kidoku'

  get iconColor() {
    return this.name === 'tsundoku'
      ? 'var(--tsundoku-red-bg)'
      : 'var(--kidoku-blue-bg)'
  }
}
</script>

<style lang="sass" scoped>
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
