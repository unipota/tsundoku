<template lang="pug">
  .check-button(@click="handleClick")
    icon(name="check" color="white" :width="19" :height="14")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class CheckButton extends Vue {
  handleClick(e: MouseEvent) {
    this.$emit('click', e)
  }
}
</script>

<style lang="sass" scoped>
.check-button
  cursor: pointer
  display: flex
  align-items: center
  justify-content: center
  width: 36px
  height: 36px
  border:
    radius: 100%
  background: var(--kidoku-blue)
</style>
