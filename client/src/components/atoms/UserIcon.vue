<template lang="pug">
  .user-icon
    img(v-if="imageSrc" :src="imageSrc")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class UserIcon extends Vue {
  @Prop({ type: String, default: '' }) readonly imageSrc!: string
}
</script>

<style lang="sass" scoped>
.user-icon
  width: 48px
  height: 48px
  border:
    radius: 100%
  background: var(--text-gray)
  overflow: hidden
</style>
