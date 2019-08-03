<template lang="pug">
  .books-empty
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
      .message(v-if="name === 'tsundoku'")
        | {{ $t('addBooksFromHere') }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: {
    Icon
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

  width: 100%
  height: calc(100vh - 48px)

.icon
  width: 100px
  height: 100px

.message-container
  display: flex
  align-items: center
  justify-content: center
  flex-direction: column

  margin-top: 2rem

.message
  margin: 0.125rem
  font:
    weight: bold
    size: 1rem
  color: $text-gray
</style>
