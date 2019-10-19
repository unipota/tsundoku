<template lang="pug">
  .user-icon
    img(v-if="imageSrc" :src="imageSrc")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component
export default class UserIcon extends Vue {
  @Prop({ type: String, default: '' }) readonly src!: string
}
</script>

<style lang="sass" scoped>
.user-icon
  width: 100%
  height: 100%
  border:
    radius: 100%
  background: var(--text-gray)
  overflow: hidden
</style>
