<script lang="ts">
import { mixins } from 'vue-class-component'
import { Component, Prop } from 'vue-property-decorator'
import { Bar } from 'vue-chartjs'
import { ChartData, ChartOptions } from 'chart.js'

@Component({})
export default class BooksCountChart extends mixins(Bar) {
  @Prop({ type: Object, required: true })
  private chartData!: ChartData

  @Prop({ type: Object, default: {} })
  private chartOptions!: ChartOptions

  mounted() {
    this.renderChart(this.chartData, this.chartOptions)
  }
}
</script>
