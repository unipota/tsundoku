<template lang="pug">
  .book-major-info
    .book-list-item__title(:class="`is-${size}`")
      | {{ book.title }}
    .book-list-item__author(:class="`is-${size}`")
      icon.icon(name="author")
      | {{ book.author }}
</template>

<script lang="ts">
type Size = 'normal' | 'small'

import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'
import Icon from '../assets/Icon.vue'

@Component({ components: { Icon } })
export default class BookListItem extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord

  @Prop({ type: String, default: 'normal' })
  private size!: Size
}
</script>

<style lang="sass">
.book-list-item__title
  font:
    size: 1.2rem
    weight: bold
  width: 100%
  margin-bottom: 8px
  overflow: hidden
  &.is-small
    font:
      size: 1rem

  display: -webkit-box
  -webkit-box-orient: vertical
  .is-mobile &
    -webkit-line-clamp: 1
  .is-desktop &
    -webkit-line-clamp: 2

  max-height: 3.6rem

.book-list-item__author
  height: 1.5rem
  width: 100%
  overflow: hidden
  white-space: nowrap
  text-overflow: ellipsis
  .icon
    margin:
      right: 7px
  &.is-small
    font:
      size: 0.8rem
</style>
