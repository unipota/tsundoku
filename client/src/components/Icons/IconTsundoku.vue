<template functional>
  <svg width="46" height="40" viewBox="0 0 46 40" fill="none" xmlns="http://www.w3.org/2000/svg">
    <rect x="2.83081" y="30.1384" width="35.3846" height="8.49231" rx="1" :fill="props.color" />
    <rect y="1" width="35.3846" height="8.49231" rx="1" :fill="props.color" />
    <rect x="14.1538" y="15.5693" width="31.8462" height="8.49231" rx="1" :fill="props.color" />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {},
})
export default class IconTsundoku extends Vue {
  @Prop({ type: String, default: 'var(--tsundoku-red)' })
  public color!: string
}
</script>

<style lang="sass">
</style>
