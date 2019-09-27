<template lang="pug">
  .books-count-stats
    .count-number
      | {{booksCount}}
    .chart-wrapper
      line-chart(:chartData="booksCountChartData" :chartOptions="chartOptions" :styles="styles")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookStats, ReadHistory } from '@/types/Book'
import { ChartData, ChartOptions } from 'chart.js'
import dayjs from 'dayjs'
import 'dayjs/locale/ja'
import minMax from 'dayjs/plugin/minMax'

import LineChart from '@/components/atoms/LineChart.vue'

dayjs.extend(minMax)

@Component({
  components: { LineChart }
})
export default class TsundokuStats extends Vue {
  @Prop({ type: Array, required: true })
  private allBookStats!: BookStats[]

  created() {
    dayjs.locale(this.$store.state.locale)
  }

  get booksCount(): number {
    return this.allBookStats.length
  }

  get booksCountChartData(): ChartData {
    const booksRegisteredDateArray = this.allBookStats.map((stats: BookStats) =>
      dayjs(stats.readHistories[0].createdAt)
    )
    let day = dayjs.min([...booksRegisteredDateArray, dayjs()])
    const arrayOfDayjsToToday = []
    while (!day.isAfter(dayjs())) {
      arrayOfDayjsToToday.push(day)
      day = day.add(1, 'day')
    }
    const tsundokuPricePerDay: number[] = arrayOfDayjsToToday.map(day => {
      return this.allBookStats.reduce((acc: number, stats: BookStats) => {
        return (
          acc +
          stats.readHistories.find((readHistory: ReadHistory) =>
            day.isSame(dayjs(readHistory.createdAt), 'day')
          ).readPages
        )
      }, 0)
    })

    return {
      labels: arrayOfDayjsToToday.map(day => day.format('MM/DD')),
      datasets: [
        {
          label: '',
          data: booksRegisteredCountsPerDay,
          backgroundColor: '#E3402A'
        }
      ]
    }
  }

  get chartOptions(): ChartOptions {
    return {
      responsive: false,
      maintainAspectRatio: false,
      legend: {
        display: false
      },
      scales: {
        yAxes: [
          {
            gridLines: {
              display: false,
              drawBorder: false
            },
            ticks: {
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
}
</script>

<style lang="sass" scoped>
.books-count-stats
  max-width: 400px
  padding: 12px
  border:
    radius: 16px
  background: var(--tsundoku-red-bg)
.chart-wrapper
    height: 200px
</style>
