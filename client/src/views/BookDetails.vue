<template lang="pug">
  .book-details
    .header
      .header-bg(:style="headerBgStyle")
        .info-bg
          book-major-info(:book="book")
      .header-fade
      .close
        router-link.icon-close(:to="parentPath")
          icon-plus(color='white' :height="16")
      .info
        book-major-info(:book="book")
    .cover-wrap
      book-cover(:book="book")
    .body
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import BookCover from '@/components/atoms/BookCover.vue'
import BookMajorInfo from '@/components/atoms/BookMajorInfo.vue'
import IconPlus from '@/components/assets/IconPlus.vue'
import Vibrant from 'node-vibrant'
import { BookRecord } from '../types/Book'
import { RouteRecord } from 'vue-router'

@Component({
  components: {
    BookCover,
    BookMajorInfo,
    IconPlus
  }
})
export default class BookDetails extends Vue {
  public $store!: ExStore
  private vibrant!: Vibrant

  public async mounted() {}

  get book(): BookRecord {
    const bookId = this.$route.params['id']
    return this.$store.getters.getBookById(bookId)
  }

  get parentPath(): RouteRecord {
    const matched = this.$route.matched
    return matched[0]
  }

  get headerBgStyle() {
    return {
      backgroundImage: `url(${this.book.coverImageUrl})`
    }
  }
}
</script>

<style lang="sass" scoped>
// FIXME
.book-details
  position: fixed
  top: 0
  right: 0
  height: 100vh
  width: 400px
  background: white

.header
  position: relative

  width: 100%
  height: 240px
  padding: 24px

  overflow: hidden

.header-bg
  $blur-radius: 20px

  position: absolute
  top: -$blur-radius
  left: -$blur-radius
  width: calc(100% + #{$blur-radius * 2})
  height: calc(100% + #{$blur-radius * 2})

  background:
    size: cover
    repeat: no-repeat
    position: center center

  filter: blur(20px)

.header-fade
  position: absolute
  top: 0
  left: 0
  width: 100%
  height: 100%

  background: linear-gradient(to bottom, rgba(0, 0, 0, 1) 0%, rgba(0, 0, 0, 0.7) 30%, rgba(0, 0, 0, 0) 70%)
  opacity: 0.5

.close
  display: flex
  justify-content: flex-end

  width: 100%

.icon-close
  transform: rotate(45deg)

.info
  margin: 12px
  position: relative

  color: white

.cover-wrap
  position: relative
  top: -64px
  left: 24px
</style>
