<template lang="pug">
  modal-frame.add-books-search(
    :title="$t('addBooksSearchTitle')"
    :is-content-non-scrollable="!isScrollable"
  )
    .content-wrap
      .search-field
        text-input(
          v-model="searchQuery"
          :placeholder="$t('addBooksSearchPlaceholder')"
          @keyup-enter="submitSearchQuery"
          withClearButton
          focus
        )
        .search-button(@click="submitSearchQuery" :class="{'is-disabled': !searchQuery}")
          icon(name="search" color="var(--text-black)" :width="18" :height="18")
      div.searching-animation(v-show="isSearching")
        .anim-container(ref="anim-container")
      div(v-if="isFirstView")
        book-cover(:hasShadow="true")
      div(v-else-if="!isSearching")
        router-link.edit-yourself(
          v-show="showEditBar"
          :to="firstRouteName + '/add-books-edit'"
        )
          icon.icon(
            name="logo"
            color="var(--border-gray)"
            :width="34"
            :height="27"
          )
          span.text
            | {{ $t('editYourself') }}
          icon.icon(name="right-arrow")
        add-book-card(
          v-for="(book, index) in searchResults"
          :key="index"
          :book="book"
          type="search"
        )
</template>

<script lang="ts">
import { Vue, Component, Watch } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { BookSimpleRecord } from '../types/Book'
import lottie from 'lottie-web/build/player/lottie_light'
import { AnimationItem } from 'lottie-web'

import Icon from '@/components/assets/Icon.vue'
import EditButton from '@/components/atoms/EditButton.vue'
import TextInput from '@/components/atoms/TextInput.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import AddBookCard from '@/components/molecules/AddBookCard.vue'
import BookCover from '@/components/atoms/BookCover.vue'

import tsundokuLoading from '@/animation/tsundoku.json'

@Component({
  components: {
    Icon,
    EditButton,
    TextInput,
    ModalFrame,
    AddBookCard,
    BookCover
  }
})
export default class AddBooksSearch extends Vue {
  public $store!: ExStore
  searchQuery = ''
  searchResults: BookSimpleRecord[] = []
  hasSubmittedSearchQuery = false
  isFirstView = true
  isScrollable = false
  isSearching = false
  loadingAnimation: AnimationItem | null = null

  mounted() {
    this.loadingAnimation = lottie.loadAnimation({
      container: this.$refs['anim-container'] as Element,
      renderer: 'svg',
      loop: true,
      autoplay: false,
      animationData: tsundokuLoading
    })
  }

  get showEditBar() {
    return !this.isSearching
  }

  get firstRouteName() {
    return this.$route.matched[0].path
  }

  submitSearchQuery() {
    if (!this.searchQuery) return
    this.isFirstView = false
    this.isSearching = true
    this.loadingAnimation && this.loadingAnimation.play()
    this.$store
      .dispatch('searchBooks', { search: this.searchQuery })
      .then((res: BookSimpleRecord[]) => {
        this.searchResults = res
        this.isSearching = false
        this.loadingAnimation && this.loadingAnimation.stop()
        this.hasSubmittedSearchQuery = true
        this.isScrollable = res.length >= 3 // [TODO] マジックナンバーではなくす
      })
  }

  @Watch('searchQuery')
  resetHasSubmittedSearchQuery(val: string) {
    // searchQuery がリセットされるたびに hasSubmittedSearchQuery もリセットする
    if (val.length === 0) {
      this.hasSubmittedSearchQuery = false
    }
  }
}
</script>

<style lang="sass" scoped>
.add-books-search
  -webkit-overflow-scrolling: touch

.book-cover
  margin: 20vh auto 40px auto

.edit-yourself
  width: calc(100% - 20px)
  margin: auto
  display: flex
  padding: 15px 30px
  border-radius: 8px
  box-shadow: 2px 4px 4px rgba(0, 0, 0, 0.1)
  .icon
    margin: auto 0
  .text
    margin: auto
    font-weight: bold

.add-book-card
  margin: 32px 0
  position: relative
  &:not(:last-child)
    &:before
      content : ""
      position: absolute
      left: calc((100% - 85%) / 2) // 中心
      bottom: -24px
      height: 1px
      width: 85%
      border-bottom: 2px solid var(--bg-gray)

.searching-animation
  display: flex
  justify-content: center
  align-items: center
  padding:
    top: 48px

.anim-container
  width: 100px
  height: 100px

.search-field
  display: flex
  width: 85%
  margin: 0 auto 20px auto

.search-button
  display: flex
  align-items: center
  justify-content: center
  flex-shrink: 0
  width: 40px
  height: 40px
  background: var(--border-gray)
  border-radius: 8px
  margin:
    left: 8px
  transition: background .3s
  cursor: pointer

  &:hover:not(.is-disabled)
    background: var(--border-gray-hovered)

  &.is-disabled
    opacity: 0.5
    cursor: auto
</style>
