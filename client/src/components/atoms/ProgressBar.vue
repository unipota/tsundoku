<template lang="pug">
  .progress-bar
    .progress(:style="progressStyle")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class ProgressBar extends Vue {
  @Prop({ type: Number, required: true }) private progress!: number

  get validatedProgress() {
    return Math.min(1, Math.max(0, this.progress))
  }

  get progressStyle() {
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
  background-color: $border-gray
  border-radius: 100vw
  overflow: hidden

.progress
  height: $height
  background-color: $kidoku-blue
  border-radius: 100vw
  position: absolute
  top: 0
  left: 0
  z-index: 2
  box-shadow: 0 0 0 4px white
  transition: width .5s $easeInOutQuint
</style>
