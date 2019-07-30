<template lang="pug">
  div.wrapper
    span.icon-back(@click="$router.go(-1)")
      icon(name="left-arrow")
    span.title
      | {{ $t('title_' + name) }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

@Component({
  components: {
    Icon
  }
})
export default class PageTopBar extends Vue {
  @Prop({ type: String, required: true })
  private readonly name!: string
}
</script>

<style lang="sass" scoped>
.wrapper
  position: absolute
  left: 37px
  top: 38px
  display: flex
  .icon-back
    cursor: pointer
    margin: auto 0
    display: flex
  .title
    color: var(--text-black)
    font-weight: bold
    font-size: 22px
    margin:
      left: 27.5px
</style>
