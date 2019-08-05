<template lang="pug">
  .book-progress-popover
    .modal-frame
      .modal-title
        span
          | 読書状況の更新
      .modal-body
        .modal-description
          span
            | スライダーを読み終わったページまで移動させましょう
      .modal-footer
        .button-cancel
        .button-apply
    .book-progress-controller
      .book-mock-body
        .book-progress-bar
      //- img.book-cover(:src="book.coverImageUrl")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookRecord } from '../../types/Book'

@Component({
  components: {}
})
export default class BookProgressPopover extends Vue {
  @Prop({ type: Object, required: true })
  private book!: BookRecord
}
</script>

<style lang="sass" scoped>
.book-progress-popover
  position: absolute
  left: 0
  top: 0
  width: 100vw
  height: 100vh
  background: rgba(0, 0, 0, 0.2)
  perspective: 1000
  transform-style: preserve-3d

.modal-frame
  width: 80%
  max-width: 300px
  margin: 10vh auto 0
  padding: 16px
  background-color: white
  border:
    radius: 8px

.modal-title
  text-align: center
  font:
    size: 1.4rem
    weight: bold

.modal-body
  margin:
    top: 12px

.book-progress-controller
  position: absolute
  left: 0
  right: 0
  bottom: 0

.book-mock-body
  width: 60%
  max-width: 300px
  height: 120px
  margin: auto
  overflow: hidden
  border:
    radius: 4px 4px 0 0
  background:
    color: white
    size: auto auto
    image: repeating-linear-gradient(90deg, transparent, transparent 10px, var(--border-gray) 10px, var(--border-gray) 12px )

.book-progress-bar
  width: 100%
  height: 8px
  background:
    color: var(--kidoku-blue)

.book-cover
  animation: slide-rotate-ver-r-fwd 2s cubic-bezier(0.455, 0.030, 0.515, 0.955) both

@keyframes slide-rotate-ver-r-fwd
  0%
    transform: translateZ(0) rotateY(0)
    transform-origin: center left
  100%
    transform: translateZ(130px) rotateY(-60deg)
    transform-origin: center left
    box-shadow: 1px 0 1px white, 2px 0 1px white, 3px 0 1px white
</style>
