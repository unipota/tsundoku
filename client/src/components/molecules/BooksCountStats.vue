<template lang="pug">
  .books-count-stats
    .count-number
      | {{booksCount}}
    .chart-wrapper
      books-count-chart(:chartData="booksCountChartData" :chartOptions="chartOptions" :styles="styles")
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { BookStats } from '@/types/Book'
import { ChartData, ChartOptions } from 'chart.js'
import dayjs from 'dayjs'
import 'dayjs/locale/ja'
import minMax from 'dayjs/plugin/minMax'
import '@/utils/roundedBarChart.js'

import BooksCountChart from '@/components/atoms/BooksCountChart.vue'

dayjs.extend(minMax)

// extend ChartOptions
interface RoundedChartOptions extends ChartOptions {
  cornerRadius: number
}

@Component({
  components: { BooksCountChart }
})
export default class BooksCountStats extends Vue {
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
      dayjs(stats.readHistories[stats.readHistories.length - 1].createdAt)
    )
    let day = dayjs.min([...booksRegisteredDateArray, dayjs()])
    const arrayOfDayjsToToday = []
    while (!day.isAfter(dayjs())) {
      arrayOfDayjsToToday.push(day)
      day = day.add(1, 'day')
    }
    const booksRegisteredCountsPerDay: number[] = arrayOfDayjsToToday.map(
      day => {
        return booksRegisteredDateArray.reduce(
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
      maintainAspectRatio: false,
      showLines: false,
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
  background: var(--toukei-black-bg)
.chart-wrapper
    height: 200px
</style>
