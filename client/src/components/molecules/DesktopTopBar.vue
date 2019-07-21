<template lang="pug">
.desktop-top-bar
  .left
    .text-input-container
      TextInput(:placeholder="$t('search_books')")
  .logo
    IconLogo
  .right
    | right

</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import TextInput from '@/components/atoms/TextInput.vue'
import IconLogo from '@/components/assets/IconLogo.vue'

@Component({
  components: {
    TextInput,
    IconLogo
  }
})
export default class DesktopTopBar extends Vue {}
</script>

<style lang="sass" scoped>

$header-height: 80px

.desktop-top-bar
  height: $header-height
  width: 100vw
  display: flex
  padding: 0 20px
  align-items: center
  justify-content: space-between
  background-color: white
  border-bottom: 1px solid $border-gray

.logo
  position: absolute
  top: $header-height / 2
  left: 50vw
  transform: translate(-50%, -50%)

.text-input-container
  width: 320px
</style>

<i18n>
{
  "ja": {
    "search_books": "本のタイトル・著者名で検索"
  },
  "en": {
    "search_books": "Search by title / author"
  }
}
</i18n>
