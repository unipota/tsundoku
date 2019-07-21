<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 14 14"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M8.55554 1C8.55554 0.447715 8.10783 0 7.55554 0H6.44443C5.89215 0 5.44443 0.447715 5.44443 1V5.44446L1 5.44446C0.447715 5.44446 0 5.89217 0 6.44446V7.55557C0 8.10785 0.447715 8.55557 1 8.55557L5.44443 8.55557V13C5.44443 13.5523 5.89215 14 6.44443 14H7.55554C8.10783 14 8.55554 13.5523 8.55554 13V8.55557L13 8.55557C13.5523 8.55557 14 8.10785 14 7.55557V6.44446C14 5.89218 13.5523 5.44446 13 5.44446L8.55554 5.44446V1Z"
      :fill="props.color"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconKidoku extends Vue {
  @Prop({ type: String, default: 'var(--text-black)' })
  private color!: string

  @Prop({ type: Number, default: 14 })
  private height!: number

  @Prop({ type: Number, default: 14 })
  private width!: number
}
</script>

<style lang="sass"></style>
