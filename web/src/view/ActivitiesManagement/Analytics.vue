
<!--<script>-->
<!--import { Bar } from 'vue-chartjs'-->

<!--export default {-->
<!--  extends: Bar,-->
<!--  data: () => ({-->
<!--    chartdata: {-->
<!--      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],-->
<!--      datasets: [-->
<!--        {-->
<!--          backgroundColor: '#f87979',-->
<!--          barPercentage: 0.5,-->
<!--          barThickness: 30,-->
<!--          maxBarThickness: 40,-->
<!--          minBarLength: 2,-->
<!--          data: [10, 20, 30, 40, 50, 60, 70, 120]-->
<!--        }-->
<!--      ]-->
<!--    },-->
<!--    options: {-->
<!--      responsive: true,-->
<!--      maintainAspectRatio: false,-->
<!--      scales: {-->
<!--        yAxes: [{-->
<!--          ticks: {-->
<!--            beginAtZero: true-->
<!--          }-->
<!--        }]-->
<!--      }-->
<!--    }-->

<!--  }),-->

<!--  mounted () {-->
<!--    this.renderChart(this.chartdata, this.options)-->
<!--  }-->
<!--}-->

<!--</script>-->

<!--<style scoped>-->

<!--</style>-->


<template>
<!--  <div class="container">-->
    <bar-chart
        v-if="loaded"
        :chartdata="chartdata"
        :options="chartoptions"/>
<!--  </div>-->
</template>

<script>
import BarChart from '@/components/Chart/BarChart.vue'
import {getStudentUserScoreList} from '@/api/user.js'

export default {
  name: 'LineChartContainer',
  components: { BarChart },
  data: () => ({
    listApi: getStudentUserScoreList,
    loaded: false,
    chartdata: null,
    chartoptions: null,
  }),
  async mounted () {
    this.loaded = false
    try {
      this.chartdata ={
        labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
            datasets: [
          {
            label: 'Data One',
            backgroundColor: '#f87979',
            barPercentage: 0.5,
            barThickness: 30,
            maxBarThickness: 40,
            minBarLength: 2,
            data: [10, 20, 30, 40, 50, 60, 70, 120, 220, 320, 500, 1024]
          },
              {
                label: 'Data Two',
                backgroundColor: '#f87979',
                barPercentage: 0.5,
                barThickness: 30,
                maxBarThickness: 40,
                minBarLength: 2,
                data: [10, 20, 30, 40, 50, 60, 70, 120, 220, 320, 500, 1024]
              }
        ]
      }
      this.chartoptions = {
        responsive: true,
            maintainAspectRatio: false,
            scales: {
          yAxes: [{
            ticks: {
              beginAtZero: true
            }
          }]
        }
      }
      // const { userlist } = await fetch('/api/userlist')
      // this.chartdata = userlist
      this.loaded = true
    } catch (e) {
      console.error(e)
    }
  },
  async created() {
    console.log(await getStudentUserScoreList({"page":1,"pageSize":10}))
  }
}
</script>
