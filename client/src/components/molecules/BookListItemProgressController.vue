<template lang="pug">
  .book-list-item-progress-container
    .progress-wrap
      book-list-item-progress(:book="book")
    .record-read-pages-button-wrap
      record-read-pages-button(@click="handleClickRecord")
    .check-button-wrap
      check-button(@click="handleClickCheck")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'

import BookListItemProgress from '@/components/molecules/BookListItemProgress.vue'
import RecordReadPagesButton from '@/components/atoms/RecordReadPagesButton.vue'
import CheckButton from '@/components/atoms/CheckButton.vue'

@Component({
  components: { BookListItemProgress, RecordReadPagesButton, CheckButton }
})
export default class BookListItemProgressController extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord

  handleClickRecord(e: MouseEvent) {
    this.$emit('click-record', e)
  }
  handleClickCheck(e: MouseEvent) {
    this.$emit('click-check', e)
  }
}
</script>

<style lang="sass" scoped>
.book-list-item-progress-container
  display: flex
  align-items: center
  flex-wrap: nowrap
  margin:
    top: 4px
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
</style>
