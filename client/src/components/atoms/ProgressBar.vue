<template lang="pug">
  .progress-bar
    .progress(:style="progressStyle")
    .edited-progress(v-if="edit" :style="editedProgressStyle")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class ProgressBar extends Vue {
  @Prop({ type: Number, required: true })
  private progress!: number

  @Prop({ type: Boolean, default: false })
  private edit!: boolean

  @Prop({ type: Number, default: 0 })
  private editedProgress!: number

  mounted() {
    this.$emit('mounted', { width: this.$el.getBoundingClientRect().width })
  }

  get validatedProgress() {
    return Math.min(1, Math.max(0, this.progress))
  }

  get progressStyle() {
    return {
      width: `${this.validatedProgress * 100}%`,
      boxShadow:
        !this.edit && this.validatedProgress > 0 ? '0 0 0 4px white' : ''
    }
  }

  get editedProgressStyle() {
    return {
      width: `${this.validatedProgress * 100}%`
    }
  }
}
</script>

<style lang="sass" scoped>
$height: 8px
.progress-bar
  position: relative
  width: 100%
  height: $height
  background-color: var(--border-gray)
  border-radius: 100vw
  overflow: hidden

.progress, .edited-progress
  height: $height
  background-color: var(--kidoku-blue)
  border-radius: 100vw
  position: absolute
  top: 0
  left: 0
  z-index: 3
  transition: width .5s $easeInOutQuint

.edited-progress
  background-color: var(--kidoku-blue-bg)
  z-index: 5
  box-shadow: 0 0 0 4px white
</style>
