<template>
  <bar-chart :chartData="datacollection" :options="options"/>
</template>

<script>
import BarChart from '@/components/Chart/BarChart.vue'
import {getUnionScoreList} from '@/api/user.js'
// import {getDict} from "@/utils/dictionary";
import infoList from "@/mixins/infoList";

export default {
  name: 'LineChartContainer',
  components: {BarChart},
  listApi: getUnionScoreList,
  mixins: [infoList],
  data() {
    return {
      scores: [],
      names: [],
      unionOptions: [],
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
  },
  async created() {
    await this.getDict('union')
    // console.log(dict)
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
      const scoreList = await getUnionScoreList({"page": 1, "pageSize": 20});
      const dataArray = scoreList.data.list;
      if (!dataArray) return;
      this.names.length = 0;
      this.scores.length = 0;
      dataArray.forEach((score) => {

        let unionName = this.filterDict(score.reqUnion, "union")
        this.names.push(unionName)

      });
      dataArray.forEach((score) => this.scores.push(score.score));

    }
  }
}
</script>
