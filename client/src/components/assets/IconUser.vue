<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 36 36"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      :fill="props.color"
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M33 18C33 26.2843 26.2843 33 18 33C9.71573 33 3 26.2843 3 18C3 9.71573 9.71573 3 18 3C26.2843 3 33 9.71573 33 18ZM27.6429 25.6339C22.8214 31.9286 12.6429 31.9286 8.35714 25.6339C8.35714 19.6071 27.6429 19.6071 27.6429 25.6339ZM18 18C20.3669 18 22.2857 16.0812 22.2857 13.7143C22.2857 11.3474 20.3669 9.42857 18 9.42857C15.6331 9.42857 13.7143 11.3474 13.7143 13.7143C13.7143 16.0812 15.6331 18 18 18Z"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconUser extends Vue {
  @Prop({ type: String, default: 'var(--text-gray)' })
  private color!: string

  @Prop({ type: Number, default: 36 })
  private height!: number

  @Prop({ type: Number, default: 36 })
  private width!: number
}
</script>

<style lang="sass"></style>
