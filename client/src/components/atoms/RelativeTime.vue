<template lang="pug">
  span
    | {{relativeTime}}
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'vue-property-decorator'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/ja'

dayjs.extend(relativeTime)

@Component({
  components: {}
})
export default class RelativeTime extends Vue {
  @Prop({ type: String, required: true })
  private from!: string

  @Prop({ type: String, default: 'ja' })
  private locale!: string

  created() {
    dayjs.locale(this.locale)
    // setInterval(() => {
    //   this.now = dayjs()
    // }, 60 * 1000)
  }

  @Watch('locale')
  onLocaleChanged() {
    dayjs.locale(this.locale)
  }

  get relativeTime(): string {
    return dayjs(this.from).fromNow()
  }
}
</script>

<style lang="sass" scoped></style>
