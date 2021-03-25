<template>
  <div class="container">
    <bar-chart
        v-if="loaded"
        :chartdata="chartdata"
        :options="chartoptions"/>
    <el-button @click="getData">getData</el-button>
  </div>

</template>

<script>
import BarChart from '@/components/Chart/BarChart.vue'
import {getStudentUserScoreList} from '@/api/user.js'

export default {
  name: 'LineChartContainer',
  components: { BarChart },
  data: () => ({
    listApi: getStudentUserScoreList,
    scores: [],
    names: [],
    loaded: false,
    chartdata: null,
    chartoptions: null,
  }),
  async mounted () {
    this.loaded = false
      this.chartdata ={
        labels: this.names,
            datasets: [
          {
            label: '学生总成绩',
            backgroundColor: '#f87979',
            barPercentage: 0.5,
            barThickness: 30,
            maxBarThickness: 40,
            minBarLength: 2,
            data: this.scores
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
    try {
      await this.getData();
      this.loaded = true;
    } catch (e) {
      console.error(e)
    }
  },
  methods: {
    getData: async function () {
      const scoreList = await getStudentUserScoreList({"page":1,"pageSize":10});
      const dataArray = scoreList.data.list;
      if (!dataArray) return;
      dataArray.forEach((score) => this.names.push(score.nickName));
      dataArray.forEach((score) => this.scores.push(score.score));
    },
  }

}
</script>
