<template>
  <!-- eslint-disable-next-line vue/require-component-is -->
  <component
    :is="`icon-${name}`"
    :color="color"
    :width="width"
    :height="height"
  />
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

// あえて遅延読み込みをしない
import IconAuthor from './IconAuthor.vue'
import IconBookmark from './IconBookmark.vue'
import IconCheck from './IconCheck.vue'
import IconCloud from './IconCloud.vue'
import IconClose from './IconClose.vue'
import IconDownArrow from './IconDownArrow.vue'
import IconExclamationMark from './IconExclamationMark.vue'
import IconGithub from './IconGithub.vue'
import IconGoogle from './IconGoogle.vue'
import IconKidoku from './IconKidoku.vue'
import IconLeftArrow from './IconLeftArrow.vue'
import IconLogo from './IconLogo.vue'
import IconPen from './IconPen.vue'
import IconPlus from './IconPlus.vue'
import IconRemove from './IconRemove.vue'
import IconRightArrow from './IconRightArrow.vue'
import IconScanner from './IconScanner.vue'
import IconSearch from './IconSearch.vue'
import IconToukei from './IconToukei.vue'
import IconTsundoku from './IconTsundoku.vue'
import IconTwitter from './IconTwitter.vue'
import IconSetting from './IconSetting.vue'
import IconUser from './IconUser.vue'
import IconDownArrow2 from './IconDownArrow2.vue'
import IconUpArrow2 from './IconUpArrow2.vue'
import IconList from './IconList.vue'
import IconGrid from './IconGrid.vue'

@Component({
  components: {
    IconAuthor,
    IconBookmark,
    IconCheck,
    IconCloud,
    IconClose,
    IconDownArrow,
    IconExclamationMark,
    IconGithub,
    IconGoogle,
    IconKidoku,
    IconLeftArrow,
    IconLogo,
    IconPen,
    IconPlus,
    IconRemove,
    IconRightArrow,
    IconScanner,
    IconSearch,
    IconToukei,
    IconTsundoku,
    IconTwitter,
    IconSetting,
    IconUser,
    IconDownArrow2,
    IconUpArrow2,
    IconList,
    IconGrid
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
