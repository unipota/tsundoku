<template>
  <component
    :is="`icon-${name}`"
    :color="color"
    :width="width"
    :height="height"
  />
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import IconClose from './IconClose.vue'
import IconDownArrow from './IconDownArrow.vue'
import IconKidoku from './IconKidoku.vue'
import IconLogo from './IconLogo.vue'
import IconPlus from './IconPlus.vue'
import IconScanner from './IconScanner.vue'
import IconSearch from './IconSearch.vue'
import IconToukei from './IconToukei.vue'
import IconTsundoku from './IconTsundoku.vue'

@Component({
  components: {
    IconClose,
    IconDownArrow,
    IconKidoku,
    IconLogo,
    IconPlus,
    IconScanner,
    IconSearch,
    IconToukei,
    IconTsundoku
  }
})
export default class Icon extends Vue {
  @Prop({ type: String, default: 'close' })
  private readonly name!: string

  @Prop({ type: String })
  private readonly color?: string

  @Prop({ type: Number })
  private readonly height?: number

  @Prop({ type: Number })
  private readonly width?: number
}
</script>
