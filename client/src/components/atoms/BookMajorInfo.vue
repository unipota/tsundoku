<template lang="pug">
  .book-major-info
    .book-list-item__title(:class="`is-${size}`")
      | {{ title }}
    .book-list-item__author(v-if="!isAuthorEmpty" :class="`is-${size}`")
      icon.icon(name="author" color="currentColor")
      span(v-for="(author, index) in authors")
        | {{ author + (index < authors.length-1 ? ' / ' : '')}}
</template>

<script lang="ts">
type Size = 'normal' | 'small'

import { Vue, Component, Prop } from 'vue-property-decorator'
import Icon from '../assets/Icon.vue'

@Component({ components: { Icon } })
export default class BookMajorInfo extends Vue {
  @Prop({ type: String, required: true })
  private title!: string

  @Prop({ type: Array, required: true })
  private authors!: string[]

  @Prop({ type: String, default: 'normal' })
  private size!: Size

  get isAuthorEmpty() {
    return !this.authors.some(author => author !== '')
  }
}
</script>

<style lang="sass">
.book-major-info
  width: 100%
  overflow: hidden
.book-list-item__title
  font:
    size: 1.2rem
    weight: bold
  width: 100%
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
  padding:
    left: 4px
  .icon
    margin:
      right: 7px
  &.is-small
    font:
      size: 0.8rem
</style>
