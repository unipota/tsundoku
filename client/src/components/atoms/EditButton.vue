<template lang="pug">
  router-link.edit-button(:to="to")
    icon.icon-pen(name="pen" color="var(--text-gray)" :width="26" :height="26")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

import Icon from '../assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class EditButton extends Vue {
  @Prop({ type: String, required: true })
  private readonly to!: string
}
</script>

<style lang="sass" scoped>
.edit-button
  border: solid var(--text-gray)
  width: 33px
  height: 33px
  border-radius: 100%
  display: flex
  padding:
    right: 1px // そのままだとアイコンがやや右寄りに見えるので
  .icon-pen
    margin: auto
</style>
