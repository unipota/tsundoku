<template lang="pug">
  modal-frame.add-books-search(
    :title="$t('addBooksSearchTitle')"
    is-content-non-scrollable
  )
    text-input(
      v-model="searchQuery"
      :placeholder="$t('addBooksSearchPlaceholder')"
      @keyup-enter="submitSearchQuery"
      withClearButton
      focus
    )
      icon(name="search" color="var(--text-gray)" :width="18" :height="18")
    template(v-if="isFirstView")
      book-cover(:hasShadow="true")
      .anim-container(ref="anim-container")
    template(v-else)
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

import Icon from '@/components/assets/Icon.vue'
import EditButton from '@/components/atoms/EditButton.vue'
import TextInput from '@/components/atoms/TextInput.vue'
import ModalFrame from '@/components/atoms/ModalFrame.vue'
import AddBookCard from '@/components/molecules/AddBookCard.vue'
import BookCover from '@/components/atoms/BookCover.vue'

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

  get goodSearchResult() {
    //  検索結果が1件以上 && クエリがタイトルに部分一致するような検索結果が存在する
    return (
      this.searchResults.length > 0 &&
      this.searchResults.find(result => result.title.includes(this.searchQuery))
    )
  }

  get showEditBar() {
    return this.hasSubmittedSearchQuery && !this.goodSearchResult
  }

  get firstRouteName() {
    return this.$route.matched[0].path
  }

  mounted() {
    const anim = lottie.loadAnimation({
      container: this.$refs['anim-container'] as Element,
      renderer: 'svg',
      loop: true,
      autoplay: false,
      path: 'json/tsundoku.json'
    })
    anim.setSpeed(0.8)
    // anim.play()
  }

  submitSearchQuery() {
    this.$store
      .dispatch('searchBooks', { search: this.searchQuery })
      .then((res: BookSimpleRecord[]) => {
        this.searchResults = res
        this.hasSubmittedSearchQuery = true
        this.isFirstView = false
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

.text-input
  width: 85%
  margin: 0 auto 20px auto

.edit-yourself
  width: 95%
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

.anim-container
  width: 100px
  height: 100px
</style>
