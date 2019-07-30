<template lang="pug">
  router-link.edit-button-wrapper(:to="to")
    div.edit-button
      icon.icon-pen(name="pen" color="var(--text-gray)" :width="18" :height="18")
      span.text
        | {{ $t('edit') }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import Icon from '../assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class EditButtonLarge extends Vue {
  @Prop({ type: String, required: true })
  private readonly to!: string
}
</script>

<style lang="sass" scoped>
.edit-button-wrapper
  border: solid var(--text-gray)
  width: 99px
  height: 33px
  border-radius: 33px
  display: flex
  padding:
    top: 2px
  .edit-button
    margin: 0 auto
    display: flex
    .icon-pen
      margin: auto 4px auto 0
    .text
      color: var(--text-gray)
      font-weight: bold
      font-size: 13px
      margin: auto 0
</style>
