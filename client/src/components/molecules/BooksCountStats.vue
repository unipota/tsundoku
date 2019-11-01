<template lang="pug">
.books-count-stats(@mouseenter="onMouseOver" @mouseleave="onMouseLeave")
  .info-wrapper
    .count-wrapper
      .count-number(v-tooltip="'今までに積んだ累計冊数'")
        | {{booksCount}}
        span.count-unit 冊
      .icon-wrapper
        icon(name="logo")
    .count-label
      | 積んだ冊数
  .chart-scroller-wrapper
    .chart-scroller(ref="scroller")
      .chart-gradation-left
      .chart-gradation-right
      .chart-wrapper(:style="chartWrapperStyle")
        bar-chart(:chartData="booksCountChartData" :chartOptions="chartOptions" :styles="styles")
  .chart-range-label(v-if="!hovered")
    | {{dayRange}}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookStats } from '@/types/Book'
import { ChartData, ChartOptions } from 'chart.js'
import dayjs from 'dayjs'
import 'dayjs/locale/ja'
import minMax from 'dayjs/plugin/minMax'
import '@/utils/roundedBarChart.js'
import '@/utils/chartjs'

import Icon from '@/components/assets/Icon.vue'
import BarChart from '@/components/atoms/BarChart.vue'

dayjs.extend(minMax)

// extend ChartOptions
interface RoundedChartOptions extends ChartOptions {
  cornerRadius: number
}

@Component({
  components: { BarChart, Icon }
})
export default class BooksCountStats extends Vue {
  @Prop({ type: Array, required: true })
  private allBookStats!: BookStats[]

  rendered: boolean = false
  hovered: boolean = false

  created() {
    dayjs.locale(this.$store.state.locale)
  }

  mounted() {
    const el = this.$refs.scroller as Element
    el.scrollLeft = el.scrollWidth
  }

  get booksCount(): number {
    return this.allBookStats.length
  }

  get booksRegisteredDateArray() {
    return this.allBookStats.map((stats: BookStats) =>
      dayjs(stats.readHistories[stats.readHistories.length - 1].createdAt)
    )
  }

  get booksCountChartData(): ChartData {
    let day = dayjs.min([...this.booksRegisteredDateArray, dayjs()])
    const arrayOfDayjsToToday = []
    while (!day.isAfter(dayjs())) {
      arrayOfDayjsToToday.push(day)
      day = day.add(1, 'day')
    }
    const booksRegisteredCountsPerDay: number[] = arrayOfDayjsToToday.map(
      day => {
        return this.booksRegisteredDateArray.reduce(
          (acc, regDay) => (regDay.isSame(day, 'day') ? acc + 1 : acc),
          0
        )
      }
    )

    return {
      labels: arrayOfDayjsToToday.map(day => day.format('MM/DD')),
      datasets: [
        {
          label: '',
          data: booksRegisteredCountsPerDay,
          backgroundColor: '#4B4B4B'
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
      showLines: false,
      title: {
        display: false
      },
      legend: {
        display: false
      },
      tooltips: {
        mode: 'x-axis',
        displayColors: false,
        callbacks: {
          label: (tooltipItem, data): string => {
            return tooltipItem.value + '冊'
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
              fontColor: '#939393',
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
  background: var(--border-gray)

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
  color: var(--text-black)
  font:
    size: 38px
    weight: bold

.count-unit
  padding:
    left: 4px
  font:
    size: 18px

.count-label
  color: var(--text-gray)
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
  background: linear-gradient(to right, var(--border-gray) 0%, transparent)

.chart-gradation-right
  right: 0
  background: linear-gradient(to left, var(--border-gray) 0%, transparent)

.chart-wrapper
  width: 200%
  // height: 120px
  padding: 0 8px 4px

.chart-range-label
  font:
    weight: bold
  color: var(--text-gray)
  height: 24px
  padding:
    left: 16px
</style>
