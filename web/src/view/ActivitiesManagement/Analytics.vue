<template>
  <bar-chart :chartData="datacollection" :options="options"/>
</template>

<script>
import BarChart from '@/components/Chart/BarChart.vue'
import {getStudentUserScoreList} from '@/api/user.js'

export default {
  name: 'LineChartContainer',
  components: {BarChart},
  data() {
    return {
      listApi: getStudentUserScoreList,
      scores: [],
      names: [],
      datacollection: null,
      options: null,
      page: 0
    }
  },
  async mounted() {
    this.options = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          ticks: {
            beginAtZero: true
          }
        }]
      }
    };
    await this.fillData();
    await setInterval(this.fillData, 5000);
  },
  methods: {
    async fillData() {
      try {
        await this.getData();
      } catch (e) {
        this.$message({
          type: 'warning',
          message: '获取数据失败'
        });
        console.error(e);
      }
      this.datacollection = {
        labels: this.names,
        datasets: [
          {
            label: '学生总成绩',
            backgroundColor: ['#f8797966', '#ff638466', '#36a2eb66', '#cc65fe66', '#ffce5666'],
            borderWidth: 2,
            borderColor: ['#f87979', '#ff6384', '#36a2eb', '#cc65fe', '#ffce56'],
            barPercentage: 0.5,
            barThickness: 50,
            maxBarThickness: 60,
            minBarLength: 2,
            data: this.scores
          }
        ]
      }
    },
    getData: async function () {
      let num = this.countTo(3)
      console.log(num);
      const scoreList = await getStudentUserScoreList({"page": num, "pageSize": 5});
      const dataArray = scoreList.data.list;
      if (!dataArray) return;
      this.names.length = 0;
      this.scores.length = 0;
      dataArray.forEach((score) => this.names.push(score.nickName));
      dataArray.forEach((score) => this.scores.push(score.score));
    },
    getRandomInt(min, max) {
      return Math.floor(Math.random() * (max - min + 1) + min);
    },
    countTo(max) {
      if (this.page === max) {
        this.page = 0;
      }
      this.page += 1;
      return this.page
    }
  },
  destroyed() {
    clearInterval();
  }

}
</script>
