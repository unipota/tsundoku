<template lang="pug">
  .book-cover(:class="{'has-shadow': hasShadow}")
    img.cover-image(
      v-if="url.length > 0"
      :key="url"
      :src="url"
      loading="lazy"
      alt="cover image"
    )
    dummy-book-cover(v-else)
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import DummyBookCover from '../assets/DummyBookCover.vue'

@Component({
  components: {
    DummyBookCover
  }
})
export default class BookCover extends Vue {
  @Prop({ type: String, default: '' })
  private readonly url!: string

  @Prop({ type: Boolean, default: false })
  private readonly hasShadow!: boolean
}
</script>

<style lang="sass" scoped>
.book-cover
  width: 100px
  height: 140px
  border-radius: 8px
  background-color: var(--text-white)
  overflow: hidden
  user-select: none
  &.has-shadow
    box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25)

.cover-image
  border-radius: 8px //for safari
  object-fit: cover
  width: 100%
  height: 100%
  transition: filter .3s .1s, opacity .1s
  will-change: filter, opacity
  opacity: 1

  // &.v-lazy-image
  //   filter: blur(1px)

  // &.v-lazy-image-loaded
  //   opacity: 1
  //   filter: blur(0)
</style>
