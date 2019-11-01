<script lang="ts">
import { mixins } from 'vue-class-component'
import { Component, Prop, Watch } from 'vue-property-decorator'
import { Line, mixins as chartMixins } from 'vue-chartjs'
import { ChartData, ChartOptions } from 'chart.js'

const { reactiveProp } = chartMixins

@Component({})
export default class LineChart extends mixins(Line, reactiveProp) {
  @Prop({ type: Object, required: true })
  private chartData!: ChartData

  @Prop({ type: Object, default: {} })
  private chartOptions!: ChartOptions

  mounted() {
    this.renderChart(this.chartData, this.chartOptions)
  }

  @Watch('chartOptions')
  onChartOptionsChanged() {
    this.renderChart(this.chartData, this.chartOptions)
  }
}
</script>
