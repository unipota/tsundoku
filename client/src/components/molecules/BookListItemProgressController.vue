<template lang="pug">
  .book-list-item-progress-controller
    .book-list-item-progress-container
      .progress-wrap
        book-list-item-progress(
          :readPages="wrappedReadPages"
          :totalPages="book.totalPages"
          :edit="recordActive" 
          :editedReadPages.sync="editedReadPages")
      .record-read-pages-button-wrap
        record-read-pages-button(@click.stop="handleClickRecord" :active="recordActive")
      .check-button-wrap
        kidoku-button(@click="handleClickCheck")
    .progress-input-wrap(v-show="recordActive")
      progress-input(:totalPages="book.totalPages" v-model="editedReadPages" @cancel="handleCancel")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import BookListItemProgress from '@/components/molecules/BookListItemProgress.vue'
import RecordReadPagesButton from '@/components/atoms/RecordReadPagesButton.vue'
import KidokuButton from '@/components/atoms/KidokuButton.vue'
import ProgressInput from '@/components/atoms/ProgressInput.vue'
import { BookRecord } from '../../types/Book'

@Component({
  components: {
    BookListItemProgress,
    RecordReadPagesButton,
    KidokuButton,
    ProgressInput
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

  get wrappedReadPages(): number {
    return this.isAnimated ? this.fakedReadPages : this.book.readPages
  }

  handleClickRecord() {
    if (!this.recordActive) {
      this.recordActive = true
      this.editedReadPages = this.book.readPages
    } else {
      this.$store.commit('updateBookReadPages', {
        id: this.book.id,
        readPages: this.editedReadPages
      })
      this.recordActive = false
    }
  }

  async handleClickCheck() {
    if (this.book.readPages === this.book.totalPages) return
    this.fakedReadPages = this.book.totalPages
    this.isAnimated = true
    await new Promise(r => window.setTimeout(r, 500))
    this.$store.commit('updateBookReadPages', {
      id: this.book.id,
      readPages: this.book.totalPages
    })
    this.isAnimated = false
  }

  handleCancel() {
    this.recordActive = false
    this.editedReadPages = this.book.readPages
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
  position: relative
  z-index: 100
  margin-left: auto
</style>
