<template lang="pug">
  .record-read-pages-button(@click="handleClick")
    icon(name="bookmark" color="white" :width="18" :height="18")
    span.label
      | {{$t('record')}}
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class RecordReadPagesButton extends Vue {
  handleClick(e: MouseEvent) {
    this.$emit('click', e)
  }
}
</script>

<style lang="sass" scoped>
.record-read-pages-button
  cursor: pointer
  display: flex
  align-items: center
  height: 36px
  background: var(--kidoku-blue)
  border:
    radius: 9999vw
  padding:
    left: 8px
    right: 12px

.label
  margin:
    left: 2px
  font:
    weight: bold
    size: 1rem
  color: white
  white-space: nowrap
</style>
