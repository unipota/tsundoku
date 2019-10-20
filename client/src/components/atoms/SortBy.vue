<template lang="pug">
  .sort-by
    transition(name="transition-sort-item")
      .sort-item(@click="changeSortBy" v-for="(item,index) in items" v-if="index ===  selectedId")
        | {{item}}
    .icon-sort-direction(:class="sortDirClass" @click="toggleSortDir")
      icon(name="down-arrow-2" color="var(--text-gray)" :width="16" :height="16")
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: {
    Icon
  }
})
export default class SortBy extends Vue {
  @Prop({ type: Array, required: true })
  private items!: string[]

  selectedId: number = 0
  sortDesc: boolean = true

  get selectedItem() {
    return this.items[this.selectedId]
  }

  get sortDirClass() {
    return {
      'is-desc': this.sortDesc
    }
  }

  changeSortBy() {
    this.selectedId =
      this.selectedId < this.items.length - 1 ? this.selectedId + 1 : 0
    this.$emit('change', this.selectedId)
  }

  toggleSortDir() {
    this.sortDesc = !this.sortDesc
    this.$emit('toggle', this.sortDesc)
  }
}
</script>

<style lang="sass" scoped>
.sort-by
  display: inline-flex
  align-items: center
  user-select: none

.sort-item
  position: relative
  font:
    weight: bold
    size: 16px
  color: var(--text-gray)
  cursor: pointer

  &::before
    content: ''
    display: block
    position: absolute
    z-index: -1
    width: calc(100% + 16px)
    height: calc(100% + 8px)
    left: -8px
    top: -4px
    background: var(--bg-gray)
    border-radius: 8px
    opacity: 0
    transition: opacity .3s

  &:hover::before
    opacity: 1

.icon-sort-direction
  position: relative
  display: flex
  width: 18px
  height: 18px
  justify-content: center
  align-items: center
  margin:
    left: 8px
    bottom: 2px
  cursor: pointer
  transform: rotate(-180deg)
  transition: transform .3s

  &::before
    content: ''
    display: block
    position: absolute
    z-index: -1
    width: 32px
    height: 32px
    background: var(--bg-gray)
    border-radius: 100%
    opacity: 0
    transition: opacity .3s

  &:hover::before
    opacity: 1

  &.is-desc
    transform: rotate(0)
</style>
