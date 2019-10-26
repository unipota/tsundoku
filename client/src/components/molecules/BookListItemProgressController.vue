<template lang="pug">
  .book-list-item-progress-controller
    template(v-if="isValidPages")
      .book-list-item-progress-container
        .progress-wrap
          book-list-item-progress(
            :readPages="wrappedReadPages"
            :totalPages="book.totalPages"
            :edit="recordActive" 
            :editedReadPages.sync="editedReadPages")
        .record-read-pages-button-wrap
          record-read-pages-button(@click.stop="handleClickRecord" :active="recordActive" v-tooltip="'読書状況を記録する'")
        .check-button-wrap
          kidoku-button(@click="handleClickCheck" v-tooltip="'キドクにする'" :disable="isKidoku")
      .progress-input-wrap(:style="progressInputStyle")
        .progress-input-body(ref="progressInputBody")
          transition(name="dummy-transition" @after-enter="afterEnter" @before-leave="beforeLeave")
            progress-input(v-show="recordActive" :totalPages="book.totalPages" v-model="editedReadPages" @cancel="handleCancel")
    template(v-else)
      .warning-invalid-page-number
        .icon-warning
          icon(name="exclamation-mark" :width="12" :height="12" color="white")
        span.warning-text
          | ページ数が未設定のため読書状況を記録できません
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import BookListItemProgress from '@/components/molecules/BookListItemProgress.vue'
import RecordReadPagesButton from '@/components/atoms/RecordReadPagesButton.vue'
import KidokuButton from '@/components/atoms/KidokuButton.vue'
import ProgressInput from '@/components/atoms/ProgressInput.vue'
import Icon from '@/components/assets/Icon.vue'
import { BookRecord } from '../../types/Book'

@Component({
  components: {
    BookListItemProgress,
    RecordReadPagesButton,
    KidokuButton,
    ProgressInput,
    Icon
  }
})
export default class BookListItemProgressController extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord

  public $store!: ExStore

  recordActive: boolean = false
  editedReadPages: number = this.book.readPages
  isAnimated: boolean = false
  fakedReadPages: number = 0
  progressInputHeight: number = 0

  get wrappedReadPages(): number {
    return this.isAnimated ? this.fakedReadPages : this.book.readPages
  }

  handleClickRecord() {
    if (!this.recordActive) {
      this.recordActive = true
      this.editedReadPages = this.book.readPages
    } else {
      if (this.book.readPages === this.editedReadPages) {
        this.recordActive = false
        return
      }
      const book = this.book
      book.readPages = this.editedReadPages
      this.$store.dispatch('updateBook', { book })
      this.recordActive = false
    }
  }

  async handleClickCheck() {
    if (this.book.readPages === this.book.totalPages) return
    this.recordActive = false
    this.fakedReadPages = this.book.totalPages
    this.isAnimated = true
    await new Promise(r => window.setTimeout(r, 500))
    const book = this.book
    book.readPages = this.book.totalPages
    this.$store.dispatch('updateBook', { book })
    this.isAnimated = false
  }

  handleCancel() {
    this.recordActive = false
    this.editedReadPages = this.book.readPages
  }

  get progressInputStyle() {
    return {
      height: `${this.progressInputHeight}px`
    }
  }

  get isKidoku() {
    return this.book.readPages === this.book.totalPages
  }

  get isValidPages() {
    return this.book.totalPages > 0
  }

  afterEnter() {
    this.$nextTick(() => {
      const ref = this.$refs.progressInputBody as HTMLElement
      this.progressInputHeight = ref.getBoundingClientRect().height
    })
  }

  beforeLeave() {
    this.progressInputHeight = 0
  }
}
</script>

<style lang="sass" scoped>
.book-list-item-progress-controller
  display: flex
  flex-flow: column

.book-list-item-progress-container
  display: flex
  align-items: center
  flex-wrap: nowrap
  margin:
    top: 8px
  padding:
    left: 4px
    right: 4px

.progress-wrap
  margin-right: 8px
  flex-grow: 1
  max-width: 250px
  margin:
    left: auto

.record-read-pages-button-wrap
  margin-right: 8px
  flex-shrink: 0

.check-button-wrap
  flex-shrink: 0

.progress-input-wrap
  transition: height .3s $easeInOutQuint
  margin-left: auto
  overflow: hidden

// .progress-input-body

.dummy-transition
  &-leave-active
    transition: all .3s

.warning-invalid-page-number
  margin:
    top: 8px
  padding:
    right: 4px
  display: inline-flex
  justify-content: flex-end
  align-items: center

.icon-warning
  flex-shrink: 0
  width: 18px
  height: 18px
  border-radius: 100%
  background: var(--danger-red)
  display: flex
  justify-content: center
  align-items: center

.warning-text
  padding-left: 6px
  font:
    size: 14px
    weight: bold
  color: var(--danger-red-fade60)
</style>
