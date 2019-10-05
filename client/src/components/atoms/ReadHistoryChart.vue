<script lang="ts">
import { mixins } from 'vue-class-component'
import { Component, Prop } from 'vue-property-decorator'
import { Bar, mixins as chartMixins } from 'vue-chartjs'
import { ChartData } from 'chart.js'
import { ReadHistory } from '@/types/Book'

const { reactiveProp } = chartMixins

@Component({})
export default class ReadHistoryChart extends mixins(Bar, reactiveProp) {
  @Prop({ type: Array, required: true })
  private readHistories!: ReadHistory[]

  mounted() {
    this.renderChart(this.chartData)
  }

  get chartData(): ChartData {
    return {
      datasets: [
        {
          data: this.readHistories.map((h: ReadHistory) => h.readPages)
        }
      ]
    }
  }
}
</script>
