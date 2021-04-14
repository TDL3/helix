<template>
  <line-chart style="height: 700px" :chartData="datacollection" :options="options"/>
</template>

<script>
import LineChart from '@/components/chart/LineChart.vue'
import {getAnalytics} from '@/api/items'
export default {
  name: 'LineChartContainer',
  components: {LineChart},
  data() {
    return {
      yLabels: [],
      newItems: [],
      recoveredItems: [],
      datacollection: null,
      options: null,
      page: 0
    }
  },
  async created() {
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
    // await setInterval(this.fillData, 5000);
  },
  methods: {
    getHeight() {
      let a = this.$refs.dope.clientHeight
      let b = this.$refs.bruh.clientHeight
      console.log(a, b)
    },
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
        labels: this.yLabels,
        datasets: [
          {
            label: '新增物品',
            data: this.newItems,
            backgroundColor: '#ff638466',
            borderWidth: 2,
            borderColor: '#ff6384',
            fill: false
          },
          {
            label: '找回物品',
            data: this.recoveredItems,
            backgroundColor: '#36a2eb66',
            borderWidth: 2,
            borderColor: '#36a2eb',
            fill: false
          }
        ]
      }
    },
    getData: async function () {
      const res = await getAnalytics();
      const dataArray = res.data.list;
      if (!dataArray) return;
      let newItems = res.data.list.newItems
      let recoveredItems = res.data.list.recoveredItems
      // console.log(newItems, recoveredItems);
      this.newItems = newItems.itemsCount
      this.recoveredItems = recoveredItems.itemsCount
      this.yLabels = newItems.monthsCount
    }
  },
  destroyed() {
    clearInterval();
  }
}
</script>

<style>
.small {
  /*max-width: 600px;*/
  /*max-height: 1080px;*/
  /*margin:  150px auto;*/
  height: max-content;
  position: relative;
}
</style>
