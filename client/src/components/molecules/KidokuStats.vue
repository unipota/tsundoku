<template lang="pug">
.books-count-stats(@mouseenter="onMouseOver" @mouseleave="onMouseLeave")
  .info-wrapper
    .count-wrapper
      .count-number(v-tooltip="'現在のキドク総額'")
        span.count-unit ¥
        tweened-number(:num="priceCount" formatLocal)
      .icon-wrapper
        icon(name="kidoku" :width="36" :height="36")
    .count-label
      | キドク額
  .chart-scroller-wrapper
    .chart-scroller(ref="scroller")
      .chart-gradation-left
      .chart-gradation-right
      .chart-wrapper(:style="chartWrapperStyle")
        line-chart(:chartData="booksCountChartData" :chartOptions="chartOptions" :styles="styles")
  .chart-range-label(v-if="!hovered")
    | {{dayRange}}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { BookStats, ReadHistory } from '@/types/Book'
import { ChartData, ChartOptions } from 'chart.js'
import dayjs, { Dayjs } from 'dayjs'
import 'dayjs/locale/ja'
import minMax from 'dayjs/plugin/minMax'
import { tsundokuPrice } from '@/utils/tsundoku'

import Icon from '@/components/assets/Icon.vue'
import LineChart from '@/components/atoms/LineChart.vue'
import TweenedNumber from '@/components/atoms/TweenedNumber.vue'

dayjs.extend(minMax)

// extend ChartOptions
interface RoundedChartOptions extends ChartOptions {
  cornerRadius: number
}

interface ReadHistoryEx extends ReadHistory {
  price: number
}

@Component({
  components: { Icon, LineChart, TweenedNumber }
})
export default class KidokuStats extends Vue {
  @Prop({ type: Array, required: true })
  private allBookStats!: BookStats[]
  $store!: ExStore

  rendered: boolean = false
  hovered: boolean = false

  created() {
    dayjs.locale(this.$store.state.locale)
  }

  mounted() {
    const el = this.$refs.scroller as Element
    el.scrollLeft = el.scrollWidth
  }

  get priceCount(): number {
    return this.$store.getters.kidokuPrice
  }

  get booksRegisteredDateArray() {
    return this.allBookStats.map((stats: BookStats) =>
      dayjs(stats.readHistories[stats.readHistories.length - 1].createdAt)
    )
  }

  convertReadHistories(
    readHistories: ReadHistory[],
    price: number,
    totalPages: number
  ): ReadHistoryEx[] {
    const array: ReadHistoryEx[] = []
    array.push({
      ...readHistories.slice(-1)[0],
      price: 0
    })
    readHistories
      .slice(0, readHistories.length - 1)
      .reverse()
      .forEach(
        (readHistroy: ReadHistory, index: number, original: ReadHistory[]) => {
          array.unshift({
            price:
              readHistroy.readPages === totalPages
                ? price
                : array[0].readPages === totalPages &&
                  readHistroy.readPages !== totalPages
                ? price * -1
                : 0,
            createdAt: readHistroy.createdAt,
            readPages: readHistroy.readPages
          })
        }
      )
    return array
  }

  get allReadHistories(): ReadHistoryEx[] {
    return this.allBookStats.reduce(
      (acc: ReadHistoryEx[], stats: BookStats) => {
        return [
          ...acc,
          ...this.convertReadHistories(
            stats.readHistories,
            stats.price,
            stats.totalPages
          )
        ]
      },
      []
    )
  }

  get priceDiffPerDay() {
    return this.allReadHistories.reduce(
      (
        acc: { [key: string]: number },
        readHistory: ReadHistoryEx
      ): {
        [key: string]: number
      } => {
        const d = dayjs(readHistory.createdAt).format('YYYY-MM-DD')
        return {
          ...acc,
          [d]: acc[d] ? acc[d] + readHistory.price : readHistory.price
        }
      },
      {}
    )
  }

  get booksCountChartData(): ChartData {
    let day = dayjs.min([...this.booksRegisteredDateArray, dayjs()])
    const arrayOfDayjsToToday = []
    while (!day.isAfter(dayjs())) {
      arrayOfDayjsToToday.push(day)
      day = day.add(1, 'day')
    }

    const tsundokuPriceArrayPerDay = arrayOfDayjsToToday
      .map((day: Dayjs): number => {
        return this.priceDiffPerDay[day.format('YYYY-MM-DD')]
          ? this.priceDiffPerDay[day.format('YYYY-MM-DD')]
          : 0
      })
      .reduce((acc: number[], priceDiff: number) => {
        return acc[0] ? [...acc, acc.slice(-1)[0] + priceDiff] : [priceDiff]
      }, [])

    return {
      labels: arrayOfDayjsToToday.map(day => day.format('MM/DD')),
      datasets: [
        {
          label: '',
          data: tsundokuPriceArrayPerDay,
          backgroundColor: '#2488D0',
          borderWidth: 0,
          pointBorderColor: 'transparent',
          pointBackgroundColor: 'transparent'
        }
      ]
    }
  }

  get chartOptions(): RoundedChartOptions {
    return {
      cornerRadius: 8,
      responsive: true,
      responsiveAnimationDuration: 0,
      animation: {
        duration: 600,
        easing: 'easeOutExpo',
        onComplete: () => {
          this.rendered = true
        }
      },
      maintainAspectRatio: false,
      showLines: true,
      title: {
        display: false
      },
      legend: {
        display: false
      },
      elements: {
        line: {
          cubicInterpolationMode: 'monotone',
          borderWidth: 0
        }
      },
      tooltips: {
        mode: 'x-axis',
        displayColors: false,
        callbacks: {
          label: (tooltipItem, data): string => {
            return '¥' + tooltipItem.value
          }
        }
      },
      scales: {
        display: false,
        yAxes: [
          {
            gridLines: {
              display: false,
              drawBorder: false
            },
            ticks: {
              display: false,
              beginAtZero: true,
              stepSize: 1
            }
          }
        ],
        xAxes: [
          {
            gridLines: {
              display: false,
              drawBorder: false
            },
            ticks: {
              display: this.hovered,
              fontColor: 'rgba(35, 136, 208, 0.6)',
              fontStyle: 'bold'
            }
          }
        ]
      }
    }
  }

  get styles() {
    return {
      position: 'relative',
      width: '100%',
      height: '100%'
    }
  }

  get chartWrapperStyle() {
    return {
      height: this.hovered ? `${120 + 24}px` : `${120}px`
    }
  }

  get dayRange(): string {
    return `${this.booksRegisteredDateArray[0].format(
      'MM/DD'
    )}~${dayjs().format('MM/DD')}`
  }

  onMouseOver() {
    this.hovered = true
  }

  onMouseLeave() {
    this.hovered = false
  }
}
</script>

<style lang="sass" scoped>
.books-count-stats
  max-width: 320px
  padding: 16px
  border:
    radius: 24px
  background: var(--kidoku-blue-bg)

.info-wrapper
  padding:
    top: 8px
    right: 8px
    left: 16px
    bottom: 32px

.count-wrapper
  display: flex
  flex-flow: row
  justify-content: space-between
  align-items: center

.count-number
  color: var(--kidoku-blue)
  font:
    size: 38px
    weight: bold

.count-unit
  padding:
    right: 4px
  font:
    size: 24px

.count-label
  color: var(--kidoku-blue-fade60)
  font:
    size: 18px
    weight: bold

.chart-scroller-wrapper
  position: relative

.chart-scroller
  width: 100%
  overflow:
    x: scroll
    y: hidden

.chart-gradation-left, .chart-gradation-right
  position: absolute
  z-index: 10
  display: block
  top: 0
  height: 100%
  width: 16px

.chart-gradation-left
  left: 0
  background: linear-gradient(to right, #c3d4f0 0%, transparent)

.chart-gradation-right
  right: 0
  background: linear-gradient(to left, #c3d4f0 0%, transparent)

.chart-wrapper
  width: 200%
  // height: 120px
  padding: 0 8px 4px

.chart-range-label
  font:
    weight: bold
  color: var(--kidoku-blue-fade60)
  height: 24px
  padding:
    left: 16px
</style>
